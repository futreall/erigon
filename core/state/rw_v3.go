// Copyright 2024 The Erigon Authors
// This file is part of Erigon.
//
// Erigon is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Erigon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Erigon. If not, see <http://www.gnu.org/licenses/>.

package state

import (
	"context"
	"encoding/binary"
	"fmt"
	"sync"
	"time"

	"github.com/erigontech/erigon-lib/log/v3"
	"github.com/holiman/uint256"

	"github.com/erigontech/erigon-lib/common"
	"github.com/erigontech/erigon-lib/common/dbg"
	"github.com/erigontech/erigon-lib/common/length"
	"github.com/erigontech/erigon-lib/etl"
	"github.com/erigontech/erigon-lib/kv"
	"github.com/erigontech/erigon-lib/metrics"
	"github.com/erigontech/erigon-lib/state"
	libstate "github.com/erigontech/erigon-lib/state"
	"github.com/erigontech/erigon/core/types/accounts"
	"github.com/erigontech/erigon/turbo/shards"
)

var execTxsDone = metrics.NewCounter(`exec_txs_done`)

type StateV3 struct {
	domains      *libstate.SharedDomains
	triggerLock  sync.Mutex
	triggers     map[uint64]*TxTask
	senderTxNums map[common.Address]uint64

	applyPrevAccountBuf []byte // buffer for ApplyState. Doesn't need mutex because Apply is single-threaded
	addrIncBuf          []byte // buffer for ApplyState. Doesn't need mutex because Apply is single-threaded
	logger              log.Logger

	trace bool
}

func NewStateV3(domains *libstate.SharedDomains, logger log.Logger) *StateV3 {
	return &StateV3{
		domains:             domains,
		triggers:            map[uint64]*TxTask{},
		senderTxNums:        map[common.Address]uint64{},
		applyPrevAccountBuf: make([]byte, 256),
		logger:              logger,
		//trace: true,
	}
}

func (rs *StateV3) ReTry(txTask *TxTask, in *QueueWithRetry) {
	txTask.Reset()
	in.ReTry(txTask)
}
func (rs *StateV3) AddWork(ctx context.Context, txTask *TxTask, in *QueueWithRetry) {
	txTask.Reset()
	in.Add(ctx, txTask)
}

func (rs *StateV3) RegisterSender(txTask *TxTask) bool {
	//TODO: it deadlocks on panic, fix it
	defer func() {
		rec := recover()
		if rec != nil {
			fmt.Printf("panic?: %s,%s\n", rec, dbg.Stack())
		}
	}()
	rs.triggerLock.Lock()
	defer rs.triggerLock.Unlock()
	lastTxNum, deferral := rs.senderTxNums[*txTask.Sender]
	if deferral {
		// Transactions with the same sender have obvious data dependency, no point running it before lastTxNum
		// So we add this data dependency as a trigger
		//fmt.Printf("trigger[%d] sender [%x]<=%x\n", lastTxNum, *txTask.Sender, txTask.Tx.Hash())
		rs.triggers[lastTxNum] = txTask
	}
	//fmt.Printf("senderTxNums[%x]=%d\n", *txTask.Sender, txTask.TxNum)
	rs.senderTxNums[*txTask.Sender] = txTask.TxNum
	return !deferral
}

func (rs *StateV3) CommitTxNum(sender *common.Address, txNum uint64, in *QueueWithRetry) (count int) {
	execTxsDone.Inc()

	rs.triggerLock.Lock()
	defer rs.triggerLock.Unlock()
	if triggered, ok := rs.triggers[txNum]; ok {
		in.ReTry(triggered)
		count++
		delete(rs.triggers, txNum)
	}
	if sender != nil {
		if lastTxNum, ok := rs.senderTxNums[*sender]; ok && lastTxNum == txNum {
			// This is the last transaction so far with this sender, remove
			delete(rs.senderTxNums, *sender)
		}
	}
	return count
}

func (rs *StateV3) applyState(txTask *TxTask, domains *libstate.SharedDomains) error {
	var acc accounts.Account

	//maps are unordered in Go! don't iterate over it. SharedDomains.deleteAccount will call GetLatest(Code) and expecting it not been delete yet
	if txTask.WriteLists != nil {
		for _, domain := range []kv.Domain{kv.AccountsDomain, kv.CodeDomain, kv.StorageDomain} {
			list, ok := txTask.WriteLists[domain.String()]
			if !ok {
				continue
			}

			for i, key := range list.Keys {
				if list.Vals[i] == nil {
					if err := domains.DomainDel(domain, []byte(key), nil, nil, 0); err != nil {
						return err
					}
				} else {
					if err := domains.DomainPut(domain, []byte(key), nil, list.Vals[i], nil, 0); err != nil {
						return err
					}
				}
			}
		}
	}

	emptyRemoval := txTask.Rules.IsSpuriousDragon
	for addr, increase := range txTask.BalanceIncreaseSet {
		increase := increase
		addrBytes := addr.Bytes()
		enc0, step0, err := domains.GetLatest(kv.AccountsDomain, addrBytes, nil)
		if err != nil {
			return err
		}
		acc.Reset()
		if len(enc0) > 0 {
			if err := accounts.DeserialiseV3(&acc, enc0); err != nil {
				return err
			}
		}
		acc.Balance.Add(&acc.Balance, &increase)
		if emptyRemoval && acc.Nonce == 0 && acc.Balance.IsZero() && acc.IsEmptyCodeHash() {
			if err := domains.DomainDel(kv.AccountsDomain, addrBytes, nil, enc0, step0); err != nil {
				return err
			}
		} else {
			enc1 := accounts.SerialiseV3(&acc)
			if err := domains.DomainPut(kv.AccountsDomain, addrBytes, nil, enc1, enc0, step0); err != nil {
				return err
			}
		}
	}
	return nil
}

func (rs *StateV3) Domains() *libstate.SharedDomains {
	return rs.domains
}

func (rs *StateV3) SetTxNum(txNum, blockNum uint64) {
	rs.domains.SetTxNum(txNum)
	rs.domains.SetBlockNum(blockNum)
}

func (rs *StateV3) ApplyState4(ctx context.Context, txTask *TxTask) error {
	if txTask.HistoryExecution {
		return nil
	}
	//defer rs.domains.BatchHistoryWriteStart().BatchHistoryWriteEnd()

	if err := rs.applyState(txTask, rs.domains); err != nil {
		return fmt.Errorf("StateV3.ApplyState: %w", err)
	}
	returnReadList(txTask.ReadLists)
	returnWriteList(txTask.WriteLists)

	if err := rs.ApplyLogsAndTraces4(txTask, rs.domains); err != nil {
		return fmt.Errorf("StateV3.ApplyLogsAndTraces: %w", err)
	}

	if (txTask.TxNum+1)%rs.domains.StepSize() == 0 /*&& txTask.TxNum > 0 */ {
		// We do not update txNum before commitment cuz otherwise committed state will be in the beginning of next file, not in the latest.
		// That's why we need to make txnum++ on SeekCommitment to get exact txNum for the latest committed state.
		//fmt.Printf("[commitment] running due to txNum reached aggregation step %d\n", txNum/rs.domains.StepSize())
		_, err := rs.domains.ComputeCommitment(ctx, true, txTask.BlockNum,
			fmt.Sprintf("applying step %d", txTask.TxNum/rs.domains.StepSize()))
		if err != nil {
			return fmt.Errorf("StateV3.ComputeCommitment: %w", err)
		}
	}

	txTask.ReadLists, txTask.WriteLists = nil, nil
	return nil
}

func (rs *StateV3) ApplyLogsAndTraces4(txTask *TxTask, domains *libstate.SharedDomains) error {
	if dbg.DiscardHistory() {
		return nil
	}
	shouldPruneNonEssentials := txTask.PruneNonEssentials && txTask.Config != nil

	for addr := range txTask.TraceFroms {
		if shouldPruneNonEssentials && addr != txTask.Config.DepositContract {
			continue
		}
		if err := domains.IndexAdd(kv.TblTracesFromIdx, addr[:]); err != nil {
			return err
		}
	}

	for addr := range txTask.TraceTos {
		if shouldPruneNonEssentials && addr != txTask.Config.DepositContract {
			continue
		}
		if err := domains.IndexAdd(kv.TblTracesToIdx, addr[:]); err != nil {
			return err
		}
	}

	for _, lg := range txTask.Logs {
		if shouldPruneNonEssentials && lg.Address != txTask.Config.DepositContract {
			continue
		}
		if err := domains.IndexAdd(kv.TblLogAddressIdx, lg.Address[:]); err != nil {
			return err
		}
		for _, topic := range lg.Topics {
			if err := domains.IndexAdd(kv.TblLogTopicsIdx, topic[:]); err != nil {
				return err
			}
		}
	}
	return nil
}

var (
	mxState3UnwindRunning = metrics.GetOrCreateGauge("state3_unwind_running")
	mxState3Unwind        = metrics.GetOrCreateSummary("state3_unwind")
)

func (rs *StateV3) Unwind(ctx context.Context, tx kv.RwTx, blockUnwindTo, txUnwindTo uint64, accumulator *shards.Accumulator, changeset *[kv.DomainLen][]state.DomainEntryDiff) error {
	mxState3UnwindRunning.Inc()
	defer mxState3UnwindRunning.Dec()
	st := time.Now()
	defer mxState3Unwind.ObserveDuration(st)
	var currentInc uint64

	//TODO: why we don't call accumulator.ChangeCode???
	handle := func(k, v []byte, table etl.CurrentTableReader, next etl.LoadNextFunc) error {
		if len(k) == length.Addr {
			if len(v) > 0 {
				var acc accounts.Account
				if err := accounts.DeserialiseV3(&acc, v); err != nil {
					return fmt.Errorf("%w, %x", err, v)
				}
				var address common.Address
				copy(address[:], k)

				newV := make([]byte, acc.EncodingLengthForStorage())
				acc.EncodeForStorage(newV)
				if accumulator != nil {
					accumulator.ChangeAccount(address, acc.Incarnation, newV)
				}
			} else {
				var address common.Address
				copy(address[:], k)
				if accumulator != nil {
					accumulator.DeleteAccount(address)
				}
			}
			return nil
		}

		var address common.Address
		var location common.Hash
		copy(address[:], k[:length.Addr])
		copy(location[:], k[length.Addr:])
		if accumulator != nil {
			accumulator.ChangeStorage(address, currentInc, location, common.Copy(v))
		}
		return nil
	}

	stateChanges := etl.NewCollector("", "", etl.NewOldestEntryBuffer(etl.BufferOptimalSize), rs.logger)
	defer stateChanges.Close()
	stateChanges.SortAndFlushInBackground(true)

	accountDiffs := changeset[kv.AccountsDomain]
	for _, kv := range accountDiffs {
		if err := stateChanges.Collect(kv.Key[:length.Addr], kv.Value); err != nil {
			return err
		}
	}
	storageDiffs := changeset[kv.StorageDomain]
	for _, kv := range storageDiffs {
		if err := stateChanges.Collect(kv.Key, kv.Value); err != nil {
			return err
		}
	}

	if err := stateChanges.Load(tx, "", handle, etl.TransformArgs{Quit: ctx.Done()}); err != nil {
		return err
	}
	if err := rs.domains.Unwind(ctx, tx, blockUnwindTo, txUnwindTo, changeset); err != nil {
		return err
	}

	return nil
}

func (rs *StateV3) DoneCount() uint64 {
	return execTxsDone.GetValueUint64()
}

func (rs *StateV3) SizeEstimate() (r uint64) {
	if rs.domains != nil {
		r += rs.domains.SizeEstimate()
	}
	return r
}

func (rs *StateV3) ReadsValid(readLists map[string]*libstate.KvList) bool {
	return rs.domains.ReadsValid(readLists)
}

// StateWriterBufferedV3 - used by parallel workers to accumulate updates and then send them to conflict-resolution.
type StateWriterBufferedV3 struct {
	rs           *StateV3
	trace        bool
	writeLists   map[string]*libstate.KvList
	accountPrevs map[string][]byte
	accountDels  map[string]*accounts.Account
	storagePrevs map[string][]byte
	codePrevs    map[string]uint64
	accumulator  *shards.Accumulator
}

func NewStateWriterBufferedV3(rs *StateV3, accumulator *shards.Accumulator) *StateWriterBufferedV3 {
	return &StateWriterBufferedV3{
		rs:          rs,
		writeLists:  newWriteList(),
		accumulator: accumulator,
		//trace:      true,
	}
}

func (w *StateWriterBufferedV3) SetTxNum(ctx context.Context, txNum uint64) {
	w.rs.domains.SetTxNum(txNum)
}
func (w *StateWriterBufferedV3) SetTx(tx kv.Tx) {}

func (w *StateWriterBufferedV3) ResetWriteSet() {
	w.writeLists = newWriteList()
	w.accountPrevs = nil
	w.accountDels = nil
	w.storagePrevs = nil
	w.codePrevs = nil
}

func (w *StateWriterBufferedV3) WriteSet() map[string]*libstate.KvList {
	return w.writeLists
}

func (w *StateWriterBufferedV3) PrevAndDels() (map[string][]byte, map[string]*accounts.Account, map[string][]byte, map[string]uint64) {
	return w.accountPrevs, w.accountDels, w.storagePrevs, w.codePrevs
}

func (w *StateWriterBufferedV3) UpdateAccountData(address common.Address, original, account *accounts.Account) error {
	if w.trace {
		fmt.Printf("acc %x: {Balance: %d, Nonce: %d, Inc: %d, CodeHash: %x}\n", address, &account.Balance, account.Nonce, account.Incarnation, account.CodeHash)
	}
	if original.Incarnation > account.Incarnation {
		//del, before create: to clanup code/storage
		if err := w.rs.domains.DomainDel(kv.CodeDomain, address[:], nil, nil, 0); err != nil {
			return err
		}
		if err := w.rs.domains.IterateStoragePrefix(address[:], func(k, v []byte, step uint64) error {
			w.writeLists[kv.StorageDomain.String()].Push(string(k), nil)
			return nil
		}); err != nil {
			return err
		}
	}
	value := accounts.SerialiseV3(account)
	if w.accumulator != nil {
		w.accumulator.ChangeAccount(address, account.Incarnation, value)
	}
	w.writeLists[kv.AccountsDomain.String()].Push(string(address[:]), value)

	return nil
}

func (w *StateWriterBufferedV3) UpdateAccountCode(address common.Address, incarnation uint64, codeHash common.Hash, code []byte) error {
	if w.trace {
		fmt.Printf("code: %x, %x, valLen: %d\n", address.Bytes(), codeHash, len(code))
	}
	if w.accumulator != nil {
		w.accumulator.ChangeCode(address, incarnation, code)
	}
	w.writeLists[kv.CodeDomain.String()].Push(string(address[:]), code)
	return nil
}

func (w *StateWriterBufferedV3) DeleteAccount(address common.Address, original *accounts.Account) error {
	if w.trace {
		fmt.Printf("del acc: %x\n", address)
	}
	if w.accumulator != nil {
		w.accumulator.DeleteAccount(address)
	}
	w.writeLists[kv.AccountsDomain.String()].Push(string(address.Bytes()), nil)
	return nil
}

func (w *StateWriterBufferedV3) WriteAccountStorage(address common.Address, incarnation uint64, key *common.Hash, original, value *uint256.Int) error {
	if *original == *value {
		return nil
	}
	compositeS := string(append(address.Bytes(), key.Bytes()...))
	w.writeLists[kv.StorageDomain.String()].Push(compositeS, value.Bytes())
	if w.trace {
		fmt.Printf("storage: %x,%x,%x\n", address, *key, value.Bytes())
	}
	if w.accumulator != nil && key != nil && value != nil {
		k := *key
		v := value.Bytes()
		w.accumulator.ChangeStorage(address, incarnation, k, v)
	}
	return nil
}

func (w *StateWriterBufferedV3) CreateContract(address common.Address) error {
	if w.trace {
		fmt.Printf("create contract: %x\n", address)
	}

	//seems don't need delete code here - tests starting fail
	//err := w.rs.domains.IterateStoragePrefix(address[:], func(k, v []byte) error {
	//	w.writeLists[string(kv.StorageDomain)].Push(string(k), nil)
	//	return nil
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}

// StateWriterV3 - used by parallel workers to accumulate updates and then send them to conflict-resolution.
type StateWriterV3 struct {
	rs          *StateV3
	trace       bool
	accumulator *shards.Accumulator
}

func NewStateWriterV3(rs *StateV3, accumulator *shards.Accumulator) *StateWriterV3 {
	return &StateWriterV3{
		rs:          rs,
		accumulator: accumulator,
		//trace: true,
	}
}

func (w *StateWriterV3) ResetWriteSet() {}

func (w *StateWriterV3) WriteSet() map[string]*libstate.KvList {
	return nil
}

func (w *StateWriterV3) PrevAndDels() (map[string][]byte, map[string]*accounts.Account, map[string][]byte, map[string]uint64) {
	return nil, nil, nil, nil
}

func (w *StateWriterV3) UpdateAccountData(address common.Address, original, account *accounts.Account) error {
	if w.trace {
		fmt.Printf("acc %x: {Balance: %d, Nonce: %d, Inc: %d, CodeHash: %x}\n", address, &account.Balance, account.Nonce, account.Incarnation, account.CodeHash)
	}
	if original.Incarnation > account.Incarnation {
		//del, before create: to clanup code/storage
		if err := w.rs.domains.DomainDel(kv.CodeDomain, address[:], nil, nil, 0); err != nil {
			return err
		}
		if err := w.rs.domains.DomainDelPrefix(kv.StorageDomain, address[:]); err != nil {
			return err
		}
	}
	value := accounts.SerialiseV3(account)
	if w.accumulator != nil {
		w.accumulator.ChangeAccount(address, account.Incarnation, value)
	}

	if err := w.rs.domains.DomainPut(kv.AccountsDomain, address[:], nil, value, nil, 0); err != nil {
		return err
	}
	return nil
}

func (w *StateWriterV3) UpdateAccountCode(address common.Address, incarnation uint64, codeHash common.Hash, code []byte) error {
	if w.trace {
		fmt.Printf("code: %x, %x, valLen: %d\n", address.Bytes(), codeHash, len(code))
	}
	if err := w.rs.domains.DomainPut(kv.CodeDomain, address[:], nil, code, nil, 0); err != nil {
		return err
	}
	if w.accumulator != nil {
		w.accumulator.ChangeCode(address, incarnation, code)
	}
	return nil
}

func (w *StateWriterV3) DeleteAccount(address common.Address, original *accounts.Account) error {
	if w.trace {
		fmt.Printf("del acc: %x\n", address)
	}
	if err := w.rs.domains.DomainDel(kv.AccountsDomain, address[:], nil, nil, 0); err != nil {
		return err
	}
	// if w.accumulator != nil { TODO: investigate later. basically this will always panic. keeping this out should be fine anyway.
	// 	w.accumulator.DeleteAccount(address)
	// }
	return nil
}

func (w *StateWriterV3) WriteAccountStorage(address common.Address, incarnation uint64, key *common.Hash, original, value *uint256.Int) error {
	if *original == *value {
		return nil
	}
	composite := append(address.Bytes(), key.Bytes()...)
	v := value.Bytes()
	if w.trace {
		fmt.Printf("storage: %x,%x,%x\n", address, *key, v)
	}
	if len(v) == 0 {
		return w.rs.domains.DomainDel(kv.StorageDomain, composite, nil, nil, 0)
	}
	if w.accumulator != nil && key != nil && value != nil {
		k := *key
		w.accumulator.ChangeStorage(address, incarnation, k, v)
	}

	return w.rs.domains.DomainPut(kv.StorageDomain, composite, nil, v, nil, 0)
}

func (w *StateWriterV3) CreateContract(address common.Address) error {
	if w.trace {
		fmt.Printf("create contract: %x\n", address)
	}

	//seems don't need delete code here. IntraBlockState take care of it.
	//if err := w.rs.domains.DomainDelPrefix(kv.StorageDomain, address[:]); err != nil {
	//	return err
	//}
	return nil
}

type ReaderV3 struct {
	txNum     uint64
	trace     bool
	tx        kv.TemporalGetter
	composite []byte
}

func NewReaderV3(tx kv.TemporalGetter) *ReaderV3 {
	return &ReaderV3{
		//trace:     true,
		tx:        tx,
		composite: make([]byte, 20+32),
	}
}

func (r *ReaderV3) DiscardReadList()                     {}
func (r *ReaderV3) SetTxNum(txNum uint64)                { r.txNum = txNum }
func (r *ReaderV3) SetTx(tx kv.Tx)                       {}
func (r *ReaderV3) ReadSet() map[string]*libstate.KvList { return nil }
func (r *ReaderV3) SetTrace(trace bool)                  { r.trace = trace }
func (r *ReaderV3) ResetReadSet()                        {}

func (r *ReaderV3) ReadAccountData(address common.Address) (*accounts.Account, error) {
	enc, _, err := r.tx.GetLatest(kv.AccountsDomain, address[:], nil)
	if err != nil {
		return nil, err
	}
	if len(enc) == 0 {
		if r.trace {
			fmt.Printf("ReadAccountData [%x] => [empty], txNum: %d\n", address, r.txNum)
		}
		return nil, nil
	}

	var acc accounts.Account
	if err := accounts.DeserialiseV3(&acc, enc); err != nil {
		return nil, err
	}
	if r.trace {
		fmt.Printf("ReadAccountData [%x] => [nonce: %d, balance: %d, codeHash: %x], txNum: %d\n", address, acc.Nonce, &acc.Balance, acc.CodeHash, r.txNum)
	}
	return &acc, nil
}

func (r *ReaderV3) ReadAccountStorage(address common.Address, incarnation uint64, key *common.Hash) ([]byte, error) {
	r.composite = append(append(r.composite[:0], address[:]...), key.Bytes()...)
	enc, _, err := r.tx.GetLatest(kv.StorageDomain, r.composite, nil)
	if err != nil {
		return nil, err
	}
	if r.trace {
		if enc == nil {
			fmt.Printf("ReadAccountStorage [%x] => [empty], txNum: %d\n", r.composite, r.txNum)
		} else {
			fmt.Printf("ReadAccountStorage [%x] => [%x], txNum: %d\n", r.composite, enc, r.txNum)
		}
	}
	return enc, nil
}

func (r *ReaderV3) ReadAccountCode(address common.Address, incarnation uint64, codeHash common.Hash) ([]byte, error) {
	//if codeHash == emptyCodeHashH { // TODO: how often do we have this case on mainnet/bor-mainnet?
	//	return nil, nil
	//}
	enc, _, err := r.tx.GetLatest(kv.CodeDomain, address[:], nil)
	if err != nil {
		return nil, err
	}
	if r.trace {
		fmt.Printf("ReadAccountCode [%x] => [%x], txNum: %d\n", address, enc, r.txNum)
	}
	return enc, nil
}

func (r *ReaderV3) ReadAccountCodeSize(address common.Address, incarnation uint64, codeHash common.Hash) (int, error) {
	enc, _, err := r.tx.GetLatest(kv.CodeDomain, address[:], nil)
	if err != nil {
		return 0, err
	}
	size := len(enc)
	if r.trace {
		fmt.Printf("ReadAccountCodeSize [%x] => [%d], txNum: %d\n", address, size, r.txNum)
	}
	return size, nil
}

func (r *ReaderV3) ReadAccountIncarnation(address common.Address) (uint64, error) {
	return 0, nil
}

type ReaderParallelV3 struct {
	txNum     uint64
	trace     bool
	sd        *libstate.SharedDomains
	composite []byte

	discardReadList bool
	readLists       map[string]*libstate.KvList
}

func NewReaderParallelV3(sd *libstate.SharedDomains) *ReaderParallelV3 {
	return &ReaderParallelV3{
		//trace:     true,
		sd:        sd,
		readLists: newReadList(),
		composite: make([]byte, 20+32),
	}
}

func (r *ReaderParallelV3) DiscardReadList()                     { r.discardReadList = true }
func (r *ReaderParallelV3) SetTxNum(txNum uint64)                { r.txNum = txNum }
func (r *ReaderParallelV3) SetTx(tx kv.Tx)                       {}
func (r *ReaderParallelV3) ReadSet() map[string]*libstate.KvList { return r.readLists }
func (r *ReaderParallelV3) SetTrace(trace bool)                  { r.trace = trace }
func (r *ReaderParallelV3) ResetReadSet()                        { r.readLists = newReadList() }

func (r *ReaderParallelV3) ReadAccountData(address common.Address) (*accounts.Account, error) {
	enc, _, err := r.sd.GetLatest(kv.AccountsDomain, address[:], nil)
	if err != nil {
		return nil, err
	}
	if !r.discardReadList {
		// lifecycle of `r.readList` is less than lifecycle of `r.rs` and `r.tx`, also `r.rs` and `r.tx` do store data immutable way
		r.readLists[kv.AccountsDomain.String()].Push(string(address[:]), enc)
	}
	if len(enc) == 0 {
		if r.trace {
			fmt.Printf("ReadAccountData [%x] => [empty], txNum: %d\n", address, r.txNum)
		}
		return nil, nil
	}

	var acc accounts.Account
	if err := accounts.DeserialiseV3(&acc, enc); err != nil {
		return nil, err
	}
	if r.trace {
		fmt.Printf("ReadAccountData [%x] => [nonce: %d, balance: %d, codeHash: %x], txNum: %d\n", address, acc.Nonce, &acc.Balance, acc.CodeHash, r.txNum)
	}
	return &acc, nil
}

func (r *ReaderParallelV3) ReadAccountStorage(address common.Address, incarnation uint64, key *common.Hash) ([]byte, error) {
	r.composite = append(append(r.composite[:0], address[:]...), key.Bytes()...)
	enc, _, err := r.sd.GetLatest(kv.StorageDomain, r.composite, nil)
	if err != nil {
		return nil, err
	}
	if !r.discardReadList {
		r.readLists[kv.StorageDomain.String()].Push(string(r.composite), enc)
	}
	if r.trace {
		if enc == nil {
			fmt.Printf("ReadAccountStorage [%x] => [empty], txNum: %d\n", r.composite, r.txNum)
		} else {
			fmt.Printf("ReadAccountStorage [%x] => [%x], txNum: %d\n", r.composite, enc, r.txNum)
		}
	}
	return enc, nil
}

func (r *ReaderParallelV3) ReadAccountCode(address common.Address, incarnation uint64, codeHash common.Hash) ([]byte, error) {
	enc, _, err := r.sd.GetLatest(kv.CodeDomain, address[:], nil)
	if err != nil {
		return nil, err
	}

	if !r.discardReadList {
		r.readLists[kv.CodeDomain.String()].Push(string(address[:]), enc)
	}
	if r.trace {
		fmt.Printf("ReadAccountCode [%x] => [%x], txNum: %d\n", address, enc, r.txNum)
	}
	return enc, nil
}

func (r *ReaderParallelV3) ReadAccountCodeSize(address common.Address, incarnation uint64, codeHash common.Hash) (int, error) {
	enc, _, err := r.sd.GetLatest(kv.CodeDomain, address[:], nil)
	if err != nil {
		return 0, err
	}
	if !r.discardReadList {
		var sizebuf [8]byte
		binary.BigEndian.PutUint64(sizebuf[:], uint64(len(enc)))
		r.readLists[libstate.CodeSizeTableFake].Push(string(address[:]), sizebuf[:])
	}
	size := len(enc)
	if r.trace {
		fmt.Printf("ReadAccountCodeSize [%x] => [%d], txNum: %d\n", address, size, r.txNum)
	}
	return size, nil
}

func (r *ReaderParallelV3) ReadAccountIncarnation(address common.Address) (uint64, error) {
	return 0, nil
}

var writeListPool = sync.Pool{
	New: func() any {
		return map[string]*libstate.KvList{
			kv.AccountsDomain.String(): {},
			kv.StorageDomain.String():  {},
			kv.CodeDomain.String():     {},
		}
	},
}

func newWriteList() map[string]*libstate.KvList {
	v := writeListPool.Get().(map[string]*libstate.KvList)
	for _, tbl := range v {
		tbl.Keys, tbl.Vals = tbl.Keys[:0], tbl.Vals[:0]
	}
	return v
	//return writeListPool.Get().(map[string]*libstate.KvList)
}
func returnWriteList(v map[string]*libstate.KvList) {
	if v == nil {
		return
	}
	//for _, tbl := range v {
	//	clear(tbl.Keys)
	//	clear(tbl.Vals)
	//	tbl.Keys, tbl.Vals = tbl.Keys[:0], tbl.Vals[:0]
	//}
	writeListPool.Put(v)
}

var readListPool = sync.Pool{
	New: func() any {
		return map[string]*libstate.KvList{
			kv.AccountsDomain.String(): {},
			kv.CodeDomain.String():     {},
			libstate.CodeSizeTableFake: {},
			kv.StorageDomain.String():  {},
		}
	},
}

func newReadList() map[string]*libstate.KvList {
	v := readListPool.Get().(map[string]*libstate.KvList)
	for _, tbl := range v {
		tbl.Keys, tbl.Vals = tbl.Keys[:0], tbl.Vals[:0]
	}
	return v
	//return readListPool.Get().(map[string]*libstate.KvList)
}
func returnReadList(v map[string]*libstate.KvList) {
	if v == nil {
		return
	}
	//for _, tbl := range v {
	//	clear(tbl.Keys)
	//	clear(tbl.Vals)
	//	tbl.Keys, tbl.Vals = tbl.Keys[:0], tbl.Vals[:0]
	//}
	readListPool.Put(v)
}
