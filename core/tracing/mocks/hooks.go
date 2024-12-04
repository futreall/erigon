// Code generated by MockGen. DO NOT EDIT.
// Source: hooks.go
//
// Generated by this command:
//
//	mockgen -source hooks.go -destination mocks/hooks.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	chain "github.com/erigontech/erigon-lib/chain"
	common "github.com/erigontech/erigon-lib/common"
	tracing "github.com/erigontech/erigon/core/tracing"
	types "github.com/erigontech/erigon/core/types"
	uint256 "github.com/holiman/uint
	gomock "go.uber.org/mock/gomock"
)

// MockOpContext is a mock of OpContext interface.
type MockOpContext struct {
	ctrl     *gomock.Controller
	recorder *MockOpContextMockRecorder
}

// MockOpContextMockRecorder is the mock recorder for MockOpContext.
type MockOpContextMockRecorder struct {
	mock *MockOpContext
}

// NewMockOpContext creates a new mock instance.
func NewMockOpContext(ctrl *gomock.Controller) *MockOpContext {
	mock := &MockOpContext{ctrl: ctrl}
	mock.recorder = &MockOpContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpContext) EXPECT() *MockOpContextMockRecorder {
	return m.recorder
}

// Address mocks base method.
func (m *MockOpContext) Address() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// Address indicates an expected call of Address.
func (mr *MockOpContextMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockOpContext)(nil).Address))
}

// CallInput mocks base method.
func (m *MockOpContext) CallInput() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallInput")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// CallInput indicates an expected call of CallInput.
func (mr *MockOpContextMockRecorder) CallInput() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallInput", reflect.TypeOf((*MockOpContext)(nil).CallInput))
}

// CallValue mocks base method.
func (m *MockOpContext) CallValue() *uint256.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallValue")
	ret0, _ := ret[0].(*uint256.Int)
	return ret0
}

// CallValue indicates an expected call of CallValue.
func (mr *MockOpContextMockRecorder) CallValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallValue", reflect.TypeOf((*MockOpContext)(nil).CallValue))
}

// Caller mocks base method.
func (m *MockOpContext) Caller() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Caller")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// Caller indicates an expected call of Caller.
func (mr *MockOpContextMockRecorder) Caller() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Caller", reflect.TypeOf((*MockOpContext)(nil).Caller))
}

// Code mocks base method.
func (m *MockOpContext) Code() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Code")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Code indicates an expected call of Code.
func (mr *MockOpContextMockRecorder) Code() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Code", reflect.TypeOf((*MockOpContext)(nil).Code))
}

// CodeHash mocks base method.
func (m *MockOpContext) CodeHash() common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeHash")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// CodeHash indicates an expected call of CodeHash.
func (mr *MockOpContextMockRecorder) CodeHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeHash", reflect.TypeOf((*MockOpContext)(nil).CodeHash))
}

// MemoryData mocks base method.
func (m *MockOpContext) MemoryData() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemoryData")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// MemoryData indicates an expected call of MemoryData.
func (mr *MockOpContextMockRecorder) MemoryData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemoryData", reflect.TypeOf((*MockOpContext)(nil).MemoryData))
}

// StackData mocks base method.
func (m *MockOpContext) StackData() []uint256.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StackData")
	ret0, _ := ret[0].([]uint256.Int)
	return ret0
}

// StackData indicates an expected call of StackData.
func (mr *MockOpContextMockRecorder) StackData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StackData", reflect.TypeOf((*MockOpContext)(nil).StackData))
}

// MockIntraBlockState is a mock of IntraBlockState interface.
type MockIntraBlockState struct {
	ctrl     *gomock.Controller
	recorder *MockIntraBlockStateMockRecorder
}

// MockIntraBlockStateMockRecorder is the mock recorder for MockIntraBlockState.
type MockIntraBlockStateMockRecorder struct {
	mock *MockIntraBlockState
}

// NewMockIntraBlockState creates a new mock instance.
func NewMockIntraBlockState(ctrl *gomock.Controller) *MockIntraBlockState {
	mock := &MockIntraBlockState{ctrl: ctrl}
	mock.recorder = &MockIntraBlockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntraBlockState) EXPECT() *MockIntraBlockStateMockRecorder {
	return m.recorder
}

// Exist mocks base method.
func (m *MockIntraBlockState) Exist(arg0 common.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exist indicates an expected call of Exist.
func (mr *MockIntraBlockStateMockRecorder) Exist(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockIntraBlockState)(nil).Exist), arg0)
}

// GetBalance mocks base method.
func (m *MockIntraBlockState) GetBalance(arg0 common.Address) *uint256.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", arg0)
	ret0, _ := ret[0].(*uint256.Int)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockIntraBlockStateMockRecorder) GetBalance(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockIntraBlockState)(nil).GetBalance), arg0)
}

// GetCode mocks base method.
func (m *MockIntraBlockState) GetCode(arg0 common.Address) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCode", arg0)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetCode indicates an expected call of GetCode.
func (mr *MockIntraBlockStateMockRecorder) GetCode(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCode", reflect.TypeOf((*MockIntraBlockState)(nil).GetCode), arg0)
}

// GetNonce mocks base method.
func (m *MockIntraBlockState) GetNonce(arg0 common.Address) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNonce", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetNonce indicates an expected call of GetNonce.
func (mr *MockIntraBlockStateMockRecorder) GetNonce(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNonce", reflect.TypeOf((*MockIntraBlockState)(nil).GetNonce), arg0)
}

// GetRefund mocks base method.
func (m *MockIntraBlockState) GetRefund() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefund")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetRefund indicates an expected call of GetRefund.
func (mr *MockIntraBlockStateMockRecorder) GetRefund() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefund", reflect.TypeOf((*MockIntraBlockState)(nil).GetRefund))
}

// GetState mocks base method.
func (m *MockIntraBlockState) GetState(addr common.Address, key *common.Hash, value *uint256.Int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetState", addr, key, value)
}

// GetState indicates an expected call of GetState.
func (mr *MockIntraBlockStateMockRecorder) GetState(addr, key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockIntraBlockState)(nil).GetState), addr, key, value)
}

// Mocktracer is a mock of tracer interface.
type Mocktracer struct {
	ctrl     *gomock.Controller
	recorder *MocktracerMockRecorder
}

// MocktracerMockRecorder is the mock recorder for Mocktracer.
type MocktracerMockRecorder struct {
	mock *Mocktracer
}

// NewMocktracer creates a new mock instance.
func NewMocktracer(ctrl *gomock.Controller) *Mocktracer {
	mock := &Mocktracer{ctrl: ctrl}
	mock.recorder = &MocktracerMockRecorder{mock}
	return mock
}

func (m *Mocktracer) Hooks() *tracing.Hooks {
	return &tracing.Hooks{
		OnTxStart:   m.TxStartHook,
		OnTxEnd:     m.TxEndHook,
		OnEnter:     m.EnterHook,
		OnExit:      m.ExitHook,
		OnOpcode:    m.OpcodeHook,
		OnFault:     m.FaultHook,
		OnGasChange: m.GasChangeHook,
		// Chain events
		OnBlockchainInit:  m.BlockchainInitHook,
		OnBlockStart:      m.BlockStartHook,
		OnBlockEnd:        m.BlockEndHook,
		OnGenesisBlock:    m.GenesisBlockHook,
		OnSystemCallStart: m.OnSystemCallStartHook,
		OnSystemCallEnd:   m.OnSystemCallEndHook,
		// State events
		OnBalanceChange: m.BalanceChangeHook,
		OnNonceChange:   m.NonceChangeHook,
		OnCodeChange:    m.CodeChangeHook,
		OnStorageChange: m.StorageChangeHook,
		OnLog:           m.LogHook,
	}
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocktracer) EXPECT() *MocktracerMockRecorder {
	return m.recorder
}

// BalanceChangeHook mocks base method.
func (m *Mocktracer) BalanceChangeHook(addr common.Address, prev, new *uint256.Int, reason tracing.BalanceChangeReason) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BalanceChangeHook", addr, prev, new, reason)
}

// BalanceChangeHook indicates an expected call of BalanceChangeHook.
func (mr *MocktracerMockRecorder) BalanceChangeHook(addr, prev, new, reason any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceChangeHook", reflect.TypeOf((*Mocktracer)(nil).BalanceChangeHook), addr, prev, new, reason)
}

// BlockEndHook mocks base method.
func (m *Mocktracer) BlockEndHook(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlockEndHook", err)
}

// BlockEndHook indicates an expected call of BlockEndHook.
func (mr *MocktracerMockRecorder) BlockEndHook(err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockEndHook", reflect.TypeOf((*Mocktracer)(nil).BlockEndHook), err)
}

// BlockStartHook mocks base method.
func (m *Mocktracer) BlockStartHook(event tracing.BlockEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlockStartHook", event)
}

// BlockStartHook indicates an expected call of BlockStartHook.
func (mr *MocktracerMockRecorder) BlockStartHook(event any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockStartHook", reflect.TypeOf((*Mocktracer)(nil).BlockStartHook), event)
}

// BlockchainInitHook mocks base method.
func (m *Mocktracer) BlockchainInitHook(chainConfig *chain.Config) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlockchainInitHook", chainConfig)
}

// BlockchainInitHook indicates an expected call of BlockchainInitHook.
func (mr *MocktracerMockRecorder) BlockchainInitHook(chainConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockchainInitHook", reflect.TypeOf((*Mocktracer)(nil).BlockchainInitHook), chainConfig)
}

// CodeChangeHook mocks base method.
func (m *Mocktracer) CodeChangeHook(addr common.Address, prevCodeHash common.Hash, prevCode []byte, codeHash common.Hash, code []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CodeChangeHook", addr, prevCodeHash, prevCode, codeHash, code)
}

// CodeChangeHook indicates an expected call of CodeChangeHook.
func (mr *MocktracerMockRecorder) CodeChangeHook(addr, prevCodeHash, prevCode, codeHash, code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeChangeHook", reflect.TypeOf((*Mocktracer)(nil).CodeChangeHook), addr, prevCodeHash, prevCode, codeHash, code)
}

// EnterHook mocks base method.
func (m *Mocktracer) EnterHook(depth int, typ byte, from, to common.Address, precompile bool, input []byte, gas uint64, value *uint256.Int, code []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnterHook", depth, typ, from, to, precompile, input, gas, value, code)
}

// EnterHook indicates an expected call of EnterHook.
func (mr *MocktracerMockRecorder) EnterHook(depth, typ, from, to, precompile, input, gas, value, code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnterHook", reflect.TypeOf((*Mocktracer)(nil).EnterHook), depth, typ, from, to, precompile, input, gas, value, code)
}

// ExitHook mocks base method.
func (m *Mocktracer) ExitHook(depth int, output []byte, gasUsed uint64, err error, reverted bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExitHook", depth, output, gasUsed, err, reverted)
}

// ExitHook indicates an expected call of ExitHook.
func (mr *MocktracerMockRecorder) ExitHook(depth, output, gasUsed, err, reverted any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExitHook", reflect.TypeOf((*Mocktracer)(nil).ExitHook), depth, output, gasUsed, err, reverted)
}

// FaultHook mocks base method.
func (m *Mocktracer) FaultHook(pc uint64, op byte, gas, cost uint64, scope tracing.OpContext, depth int, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FaultHook", pc, op, gas, cost, scope, depth, err)
}

// FaultHook indicates an expected call of FaultHook.
func (mr *MocktracerMockRecorder) FaultHook(pc, op, gas, cost, scope, depth, err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FaultHook", reflect.TypeOf((*Mocktracer)(nil).FaultHook), pc, op, gas, cost, scope, depth, err)
}

// GasChangeHook mocks base method.
func (m *Mocktracer) GasChangeHook(old, new uint64, reason tracing.GasChangeReason) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GasChangeHook", old, new, reason)
}

// GasChangeHook indicates an expected call of GasChangeHook.
func (mr *MocktracerMockRecorder) GasChangeHook(old, new, reason any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GasChangeHook", reflect.TypeOf((*Mocktracer)(nil).GasChangeHook), old, new, reason)
}

// GenesisBlockHook mocks base method.
func (m *Mocktracer) GenesisBlockHook(genesis *types.Block, alloc types.GenesisAlloc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GenesisBlockHook", genesis, alloc)
}

// GenesisBlockHook indicates an expected call of GenesisBlockHook.
func (mr *MocktracerMockRecorder) GenesisBlockHook(genesis, alloc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenesisBlockHook", reflect.TypeOf((*Mocktracer)(nil).GenesisBlockHook), genesis, alloc)
}

// LogHook mocks base method.
func (m *Mocktracer) LogHook(log *types.Log) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogHook", log)
}

// LogHook indicates an expected call of LogHook.
func (mr *MocktracerMockRecorder) LogHook(log any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogHook", reflect.TypeOf((*Mocktracer)(nil).LogHook), log)
}

// NonceChangeHook mocks base method.
func (m *Mocktracer) NonceChangeHook(addr common.Address, prev, new uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NonceChangeHook", addr, prev, new)
}

// NonceChangeHook indicates an expected call of NonceChangeHook.
func (mr *MocktracerMockRecorder) NonceChangeHook(addr, prev, new any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NonceChangeHook", reflect.TypeOf((*Mocktracer)(nil).NonceChangeHook), addr, prev, new)
}

// OnSystemCallEndHook mocks base method.
func (m *Mocktracer) OnSystemCallEndHook() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnSystemCallEndHook")
}

// OnSystemCallEndHook indicates an expected call of OnSystemCallEndHook.
func (mr *MocktracerMockRecorder) OnSystemCallEndHook() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnSystemCallEndHook", reflect.TypeOf((*Mocktracer)(nil).OnSystemCallEndHook))
}

// OnSystemCallStartHook mocks base method.
func (m *Mocktracer) OnSystemCallStartHook() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnSystemCallStartHook")
}

// OnSystemCallStartHook indicates an expected call of OnSystemCallStartHook.
func (mr *MocktracerMockRecorder) OnSystemCallStartHook() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnSystemCallStartHook", reflect.TypeOf((*Mocktracer)(nil).OnSystemCallStartHook))
}

// OpcodeHook mocks base method.
func (m *Mocktracer) OpcodeHook(pc uint64, op byte, gas, cost uint64, scope tracing.OpContext, rData []byte, depth int, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OpcodeHook", pc, op, gas, cost, scope, rData, depth, err)
}

// OpcodeHook indicates an expected call of OpcodeHook.
func (mr *MocktracerMockRecorder) OpcodeHook(pc, op, gas, cost, scope, rData, depth, err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpcodeHook", reflect.TypeOf((*Mocktracer)(nil).OpcodeHook), pc, op, gas, cost, scope, rData, depth, err)
}

// StorageChangeHook mocks base method.
func (m *Mocktracer) StorageChangeHook(addr common.Address, slot *common.Hash, prev, new uint256.Int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StorageChangeHook", addr, slot, prev, new)
}

// StorageChangeHook indicates an expected call of StorageChangeHook.
func (mr *MocktracerMockRecorder) StorageChangeHook(addr, slot, prev, new any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageChangeHook", reflect.TypeOf((*Mocktracer)(nil).StorageChangeHook), addr, slot, prev, new)
}

// TxEndHook mocks base method.
func (m *Mocktracer) TxEndHook(receipt *types.Receipt, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TxEndHook", receipt, err)
}

// TxEndHook indicates an expected call of TxEndHook.
func (mr *MocktracerMockRecorder) TxEndHook(receipt, err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxEndHook", reflect.TypeOf((*Mocktracer)(nil).TxEndHook), receipt, err)
}

// TxStartHook mocks base method.
func (m *Mocktracer) TxStartHook(vm *tracing.VMContext, tx types.Transaction, from common.Address) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TxStartHook", vm, tx, from)
}

// TxStartHook indicates an expected call of TxStartHook.
func (mr *MocktracerMockRecorder) TxStartHook(vm, tx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxStartHook", reflect.TypeOf((*Mocktracer)(nil).TxStartHook), vm, tx, from)
}
