// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erigontech/erigon/txnprovider/txpool (interfaces: Pool)
//
// Generated by this command:
//
//	mockgen -typed=true -destination=./pool_mock.go -package=txpool . Pool
//

// Package txpool is a generated GoMock package.
package txpool

import (
	context "context"
	reflect "reflect"

	remoteproto "github.com/erigontech/erigon-lib/gointerfaces/remoteproto"
	kv "github.com/erigontech/erigon-lib/kv"
	txpoolcfg "github.com/erigontech/erigon/txnprovider/txpool/txpoolcfg"
	gomock "go.uber.org/mock/gomock"
)

// MockPool is a mock of Pool interface.
type MockPool struct {
	ctrl     *gomock.Controller
	recorder *MockPoolMockRecorder
}

// MockPoolMockRecorder is the mock recorder for MockPool.
type MockPoolMockRecorder struct {
	mock *MockPool
}

// NewMockPool creates a new mock instance.
func NewMockPool(ctrl *gomock.Controller) *MockPool {
	mock := &MockPool{ctrl: ctrl}
	mock.recorder = &MockPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPool) EXPECT() *MockPoolMockRecorder {
	return m.recorder
}

// AddLocalTxns mocks base method.
func (m *MockPool) AddLocalTxns(arg0 context.Context, arg1 TxnSlots) ([]txpoolcfg.DiscardReason, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLocalTxns", arg0, arg1)
	ret0, _ := ret[0].([]txpoolcfg.DiscardReason)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLocalTxns indicates an expected call of AddLocalTxns.
func (mr *MockPoolMockRecorder) AddLocalTxns(arg0, arg1 any) *MockPoolAddLocalTxnsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLocalTxns", reflect.TypeOf((*MockPool)(nil).AddLocalTxns), arg0, arg1)
	return &MockPoolAddLocalTxnsCall{Call: call}
}

// MockPoolAddLocalTxnsCall wrap *gomock.Call
type MockPoolAddLocalTxnsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPoolAddLocalTxnsCall) Return(arg0 []txpoolcfg.DiscardReason, arg1 error) *MockPoolAddLocalTxnsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPoolAddLocalTxnsCall) Do(f func(context.Context, TxnSlots) ([]txpoolcfg.DiscardReason, error)) *MockPoolAddLocalTxnsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPoolAddLocalTxnsCall) DoAndReturn(f func(context.Context, TxnSlots) ([]txpoolcfg.DiscardReason, error)) *MockPoolAddLocalTxnsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AddNewGoodPeer mocks base method.
func (m *MockPool) AddNewGoodPeer(arg0 PeerID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddNewGoodPeer", arg0)
}

// AddNewGoodPeer indicates an expected call of AddNewGoodPeer.
func (mr *MockPoolMockRecorder) AddNewGoodPeer(arg0 any) *MockPoolAddNewGoodPeerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewGoodPeer", reflect.TypeOf((*MockPool)(nil).AddNewGoodPeer), arg0)
	return &MockPoolAddNewGoodPeerCall{Call: call}
}

// MockPoolAddNewGoodPeerCall wrap *gomock.Call
type MockPoolAddNewGoodPeerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPoolAddNewGoodPeerCall) Return() *MockPoolAddNewGoodPeerCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPoolAddNewGoodPeerCall) Do(f func(PeerID)) *MockPoolAddNewGoodPeerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPoolAddNewGoodPeerCall) DoAndReturn(f func(PeerID)) *MockPoolAddNewGoodPeerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AddRemoteTxns mocks base method.
func (m *MockPool) AddRemoteTxns(arg0 context.Context, arg1 TxnSlots) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddRemoteTxns", arg0, arg1)
}

// AddRemoteTxns indicates an expected call of AddRemoteTxns.
func (mr *MockPoolMockRecorder) AddRemoteTxns(arg0, arg1 any) *MockPoolAddRemoteTxnsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRemoteTxns", reflect.TypeOf((*MockPool)(nil).AddRemoteTxns), arg0, arg1)
	return &MockPoolAddRemoteTxnsCall{Call: call}
}

// MockPoolAddRemoteTxnsCall wrap *gomock.Call
type MockPoolAddRemoteTxnsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPoolAddRemoteTxnsCall) Return() *MockPoolAddRemoteTxnsCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPoolAddRemoteTxnsCall) Do(f func(context.Context, TxnSlots)) *MockPoolAddRemoteTxnsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPoolAddRemoteTxnsCall) DoAndReturn(f func(context.Context, TxnSlots)) *MockPoolAddRemoteTxnsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FilterKnownIdHashes mocks base method.
func (m *MockPool) FilterKnownIdHashes(arg0 kv.Tx, arg1 Hashes) (Hashes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterKnownIdHashes", arg0, arg1)
	ret0, _ := ret[0].(Hashes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterKnownIdHashes indicates an expected call of FilterKnownIdHashes.
func (mr *MockPoolMockRecorder) FilterKnownIdHashes(arg0, arg1 any) *MockPoolFilterKnownIdHashesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterKnownIdHashes", reflect.TypeOf((*MockPool)(nil).FilterKnownIdHashes), arg0, arg1)
	return &MockPoolFilterKnownIdHashesCall{Call: call}
}

// MockPoolFilterKnownIdHashesCall wrap *gomock.Call
type MockPoolFilterKnownIdHashesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPoolFilterKnownIdHashesCall) Return(arg0 Hashes, arg1 error) *MockPoolFilterKnownIdHashesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPoolFilterKnownIdHashesCall) Do(f func(kv.Tx, Hashes) (Hashes, error)) *MockPoolFilterKnownIdHashesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPoolFilterKnownIdHashesCall) DoAndReturn(f func(kv.Tx, Hashes) (Hashes, error)) *MockPoolFilterKnownIdHashesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetRlp mocks base method.
func (m *MockPool) GetRlp(arg0 kv.Tx, arg1 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRlp", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRlp indicates an expected call of GetRlp.
func (mr *MockPoolMockRecorder) GetRlp(arg0, arg1 any) *MockPoolGetRlpCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRlp", reflect.TypeOf((*MockPool)(nil).GetRlp), arg0, arg1)
	return &MockPoolGetRlpCall{Call: call}
}

// MockPoolGetRlpCall wrap *gomock.Call
type MockPoolGetRlpCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPoolGetRlpCall) Return(arg0 []byte, arg1 error) *MockPoolGetRlpCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPoolGetRlpCall) Do(f func(kv.Tx, []byte) ([]byte, error)) *MockPoolGetRlpCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPoolGetRlpCall) DoAndReturn(f func(kv.Tx, []byte) ([]byte, error)) *MockPoolGetRlpCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IdHashKnown mocks base method.
func (m *MockPool) IdHashKnown(arg0 kv.Tx, arg1 []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdHashKnown", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IdHashKnown indicates an expected call of IdHashKnown.
func (mr *MockPoolMockRecorder) IdHashKnown(arg0, arg1 any) *MockPoolIdHashKnownCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdHashKnown", reflect.TypeOf((*MockPool)(nil).IdHashKnown), arg0, arg1)
	return &MockPoolIdHashKnownCall{Call: call}
}

// MockPoolIdHashKnownCall wrap *gomock.Call
type MockPoolIdHashKnownCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPoolIdHashKnownCall) Return(arg0 bool, arg1 error) *MockPoolIdHashKnownCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPoolIdHashKnownCall) Do(f func(kv.Tx, []byte) (bool, error)) *MockPoolIdHashKnownCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPoolIdHashKnownCall) DoAndReturn(f func(kv.Tx, []byte) (bool, error)) *MockPoolIdHashKnownCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OnNewBlock mocks base method.
func (m *MockPool) OnNewBlock(arg0 context.Context, arg1 *remoteproto.StateChangeBatch, arg2, arg3, arg4 TxnSlots) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnNewBlock", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnNewBlock indicates an expected call of OnNewBlock.
func (mr *MockPoolMockRecorder) OnNewBlock(arg0, arg1, arg2, arg3, arg4 any) *MockPoolOnNewBlockCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNewBlock", reflect.TypeOf((*MockPool)(nil).OnNewBlock), arg0, arg1, arg2, arg3, arg4)
	return &MockPoolOnNewBlockCall{Call: call}
}

// MockPoolOnNewBlockCall wrap *gomock.Call
type MockPoolOnNewBlockCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPoolOnNewBlockCall) Return(arg0 error) *MockPoolOnNewBlockCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPoolOnNewBlockCall) Do(f func(context.Context, *remoteproto.StateChangeBatch, TxnSlots, TxnSlots, TxnSlots) error) *MockPoolOnNewBlockCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPoolOnNewBlockCall) DoAndReturn(f func(context.Context, *remoteproto.StateChangeBatch, TxnSlots, TxnSlots, TxnSlots) error) *MockPoolOnNewBlockCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Started mocks base method.
func (m *MockPool) Started() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Started")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Started indicates an expected call of Started.
func (mr *MockPoolMockRecorder) Started() *MockPoolStartedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Started", reflect.TypeOf((*MockPool)(nil).Started))
	return &MockPoolStartedCall{Call: call}
}

// MockPoolStartedCall wrap *gomock.Call
type MockPoolStartedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPoolStartedCall) Return(arg0 bool) *MockPoolStartedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPoolStartedCall) Do(f func() bool) *MockPoolStartedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPoolStartedCall) DoAndReturn(f func() bool) *MockPoolStartedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ValidateSerializedTxn mocks base method.
func (m *MockPool) ValidateSerializedTxn(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateSerializedTxn", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateSerializedTxn indicates an expected call of ValidateSerializedTxn.
func (mr *MockPoolMockRecorder) ValidateSerializedTxn(arg0 any) *MockPoolValidateSerializedTxnCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateSerializedTxn", reflect.TypeOf((*MockPool)(nil).ValidateSerializedTxn), arg0)
	return &MockPoolValidateSerializedTxnCall{Call: call}
}

// MockPoolValidateSerializedTxnCall wrap *gomock.Call
type MockPoolValidateSerializedTxnCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPoolValidateSerializedTxnCall) Return(arg0 error) *MockPoolValidateSerializedTxnCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPoolValidateSerializedTxnCall) Do(f func([]byte) error) *MockPoolValidateSerializedTxnCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPoolValidateSerializedTxnCall) DoAndReturn(f func([]byte) error) *MockPoolValidateSerializedTxnCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
