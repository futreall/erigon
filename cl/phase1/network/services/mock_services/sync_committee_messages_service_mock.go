// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erigontech/erigon/cl/phase1/network/services (interfaces: SyncCommitteeMessagesService)
//
// Generated by this command:
//
//	mockgen -typed=true -destination=./mock_services/sync_committee_messages_service_mock.go -package=mock_services . SyncCommitteeMessagesService
//

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	cltypes "github.com/erigontech/erigon/cl/cltypes"
	gomock "go.uber.org/mock/gomock"
)

// MockSyncCommitteeMessagesService is a mock of SyncCommitteeMessagesService interface.
type MockSyncCommitteeMessagesService struct {
	ctrl     *gomock.Controller
	recorder *MockSyncCommitteeMessagesServiceMockRecorder
	isgomock struct{}
}

// MockSyncCommitteeMessagesServiceMockRecorder is the mock recorder for MockSyncCommitteeMessagesService.
type MockSyncCommitteeMessagesServiceMockRecorder struct {
	mock *MockSyncCommitteeMessagesService
}

// NewMockSyncCommitteeMessagesService creates a new mock instance.
func NewMockSyncCommitteeMessagesService(ctrl *gomock.Controller) *MockSyncCommitteeMessagesService {
	mock := &MockSyncCommitteeMessagesService{ctrl: ctrl}
	mock.recorder = &MockSyncCommitteeMessagesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncCommitteeMessagesService) EXPECT() *MockSyncCommitteeMessagesServiceMockRecorder {
	return m.recorder
}

// ProcessMessage mocks base method.
func (m *MockSyncCommitteeMessagesService) ProcessMessage(ctx context.Context, subnet *uint64, msg *cltypes.SyncCommitteeMessageWithGossipData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessMessage", ctx, subnet, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessMessage indicates an expected call of ProcessMessage.
func (mr *MockSyncCommitteeMessagesServiceMockRecorder) ProcessMessage(ctx, subnet, msg any) *MockSyncCommitteeMessagesServiceProcessMessageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessMessage", reflect.TypeOf((*MockSyncCommitteeMessagesService)(nil).ProcessMessage), ctx, subnet, msg)
	return &MockSyncCommitteeMessagesServiceProcessMessageCall{Call: call}
}

// MockSyncCommitteeMessagesServiceProcessMessageCall wrap *gomock.Call
type MockSyncCommitteeMessagesServiceProcessMessageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSyncCommitteeMessagesServiceProcessMessageCall) Return(arg0 error) *MockSyncCommitteeMessagesServiceProcessMessageCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSyncCommitteeMessagesServiceProcessMessageCall) Do(f func(context.Context, *uint64, *cltypes.SyncCommitteeMessageWithGossipData) error) *MockSyncCommitteeMessagesServiceProcessMessageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSyncCommitteeMessagesServiceProcessMessageCall) DoAndReturn(f func(context.Context, *uint64, *cltypes.SyncCommitteeMessageWithGossipData) error) *MockSyncCommitteeMessagesServiceProcessMessageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
