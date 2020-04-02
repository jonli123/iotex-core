// Code generated by MockGen. DO NOT EDIT.
// Source: ./blockchain/blockchain.go

// Package mock_blockchain is a generated GoMock package.
package mock_blockchain

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	hash "github.com/iotexproject/go-pkgs/hash"
	action "github.com/iotexproject/iotex-core/action"
	blockchain "github.com/iotexproject/iotex-core/blockchain"
	block "github.com/iotexproject/iotex-core/blockchain/block"
	genesis "github.com/iotexproject/iotex-core/blockchain/genesis"
	reflect "reflect"
	time "time"
)

// MockBlockchain is a mock of Blockchain interface
type MockBlockchain struct {
	ctrl     *gomock.Controller
	recorder *MockBlockchainMockRecorder
}

// MockBlockchainMockRecorder is the mock recorder for MockBlockchain
type MockBlockchainMockRecorder struct {
	mock *MockBlockchain
}

// NewMockBlockchain creates a new mock instance
func NewMockBlockchain(ctrl *gomock.Controller) *MockBlockchain {
	mock := &MockBlockchain{ctrl: ctrl}
	mock.recorder = &MockBlockchainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockchain) EXPECT() *MockBlockchainMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockBlockchain) Start(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockBlockchainMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBlockchain)(nil).Start), arg0)
}

// Stop mocks base method
func (m *MockBlockchain) Stop(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockBlockchainMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBlockchain)(nil).Stop), arg0)
}

// BlockHeaderByHeight mocks base method
func (m *MockBlockchain) BlockHeaderByHeight(height uint64) (*block.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockHeaderByHeight", height)
	ret0, _ := ret[0].(*block.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockHeaderByHeight indicates an expected call of BlockHeaderByHeight
func (mr *MockBlockchainMockRecorder) BlockHeaderByHeight(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockHeaderByHeight", reflect.TypeOf((*MockBlockchain)(nil).BlockHeaderByHeight), height)
}

// BlockFooterByHeight mocks base method
func (m *MockBlockchain) BlockFooterByHeight(height uint64) (*block.Footer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockFooterByHeight", height)
	ret0, _ := ret[0].(*block.Footer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockFooterByHeight indicates an expected call of BlockFooterByHeight
func (mr *MockBlockchainMockRecorder) BlockFooterByHeight(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockFooterByHeight", reflect.TypeOf((*MockBlockchain)(nil).BlockFooterByHeight), height)
}

// ChainID mocks base method
func (m *MockBlockchain) ChainID() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainID")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// ChainID indicates an expected call of ChainID
func (mr *MockBlockchainMockRecorder) ChainID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainID", reflect.TypeOf((*MockBlockchain)(nil).ChainID))
}

// ChainAddress mocks base method
func (m *MockBlockchain) ChainAddress() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainAddress")
	ret0, _ := ret[0].(string)
	return ret0
}

// ChainAddress indicates an expected call of ChainAddress
func (mr *MockBlockchainMockRecorder) ChainAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainAddress", reflect.TypeOf((*MockBlockchain)(nil).ChainAddress))
}

// TipHash mocks base method
func (m *MockBlockchain) TipHash() hash.Hash256 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TipHash")
	ret0, _ := ret[0].(hash.Hash256)
	return ret0
}

// TipHash indicates an expected call of TipHash
func (mr *MockBlockchainMockRecorder) TipHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TipHash", reflect.TypeOf((*MockBlockchain)(nil).TipHash))
}

// TipHeight mocks base method
func (m *MockBlockchain) TipHeight() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TipHeight")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// TipHeight indicates an expected call of TipHeight
func (mr *MockBlockchainMockRecorder) TipHeight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TipHeight", reflect.TypeOf((*MockBlockchain)(nil).TipHeight))
}

// Genesis mocks base method
func (m *MockBlockchain) Genesis() genesis.Genesis {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Genesis")
	ret0, _ := ret[0].(genesis.Genesis)
	return ret0
}

// Genesis indicates an expected call of Genesis
func (mr *MockBlockchainMockRecorder) Genesis() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Genesis", reflect.TypeOf((*MockBlockchain)(nil).Genesis))
}

// Context mocks base method
func (m *MockBlockchain) Context() (context.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Context indicates an expected call of Context
func (mr *MockBlockchainMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBlockchain)(nil).Context))
}

// MintNewBlock mocks base method
func (m *MockBlockchain) MintNewBlock(actionMap map[string][]action.SealedEnvelope, timestamp time.Time) (*block.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintNewBlock", actionMap, timestamp)
	ret0, _ := ret[0].(*block.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MintNewBlock indicates an expected call of MintNewBlock
func (mr *MockBlockchainMockRecorder) MintNewBlock(actionMap, timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintNewBlock", reflect.TypeOf((*MockBlockchain)(nil).MintNewBlock), actionMap, timestamp)
}

// CommitBlock mocks base method
func (m *MockBlockchain) CommitBlock(blk *block.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitBlock", blk)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitBlock indicates an expected call of CommitBlock
func (mr *MockBlockchainMockRecorder) CommitBlock(blk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitBlock", reflect.TypeOf((*MockBlockchain)(nil).CommitBlock), blk)
}

// ValidateBlock mocks base method
func (m *MockBlockchain) ValidateBlock(blk *block.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateBlock", blk)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateBlock indicates an expected call of ValidateBlock
func (mr *MockBlockchainMockRecorder) ValidateBlock(blk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateBlock", reflect.TypeOf((*MockBlockchain)(nil).ValidateBlock), blk)
}

// AddSubscriber mocks base method
func (m *MockBlockchain) AddSubscriber(arg0 blockchain.BlockCreationSubscriber) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSubscriber", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSubscriber indicates an expected call of AddSubscriber
func (mr *MockBlockchainMockRecorder) AddSubscriber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubscriber", reflect.TypeOf((*MockBlockchain)(nil).AddSubscriber), arg0)
}

// RemoveSubscriber mocks base method
func (m *MockBlockchain) RemoveSubscriber(arg0 blockchain.BlockCreationSubscriber) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSubscriber", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSubscriber indicates an expected call of RemoveSubscriber
func (mr *MockBlockchainMockRecorder) RemoveSubscriber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSubscriber", reflect.TypeOf((*MockBlockchain)(nil).RemoveSubscriber), arg0)
}

// MockBlockBuilderFactory is a mock of BlockBuilderFactory interface
type MockBlockBuilderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockBlockBuilderFactoryMockRecorder
}

// MockBlockBuilderFactoryMockRecorder is the mock recorder for MockBlockBuilderFactory
type MockBlockBuilderFactoryMockRecorder struct {
	mock *MockBlockBuilderFactory
}

// NewMockBlockBuilderFactory creates a new mock instance
func NewMockBlockBuilderFactory(ctrl *gomock.Controller) *MockBlockBuilderFactory {
	mock := &MockBlockBuilderFactory{ctrl: ctrl}
	mock.recorder = &MockBlockBuilderFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockBuilderFactory) EXPECT() *MockBlockBuilderFactoryMockRecorder {
	return m.recorder
}

// NewBlockBuilder mocks base method
func (m *MockBlockBuilderFactory) NewBlockBuilder(arg0 context.Context, arg1 map[string][]action.SealedEnvelope, arg2 func(action.Envelope) (action.SealedEnvelope, error)) (*block.Builder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewBlockBuilder", arg0, arg1, arg2)
	ret0, _ := ret[0].(*block.Builder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewBlockBuilder indicates an expected call of NewBlockBuilder
func (mr *MockBlockBuilderFactoryMockRecorder) NewBlockBuilder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBlockBuilder", reflect.TypeOf((*MockBlockBuilderFactory)(nil).NewBlockBuilder), arg0, arg1, arg2)
}

// MockActPoolManager is a mock of ActPoolManager interface
type MockActPoolManager struct {
	ctrl     *gomock.Controller
	recorder *MockActPoolManagerMockRecorder
}

// MockActPoolManagerMockRecorder is the mock recorder for MockActPoolManager
type MockActPoolManagerMockRecorder struct {
	mock *MockActPoolManager
}

// NewMockActPoolManager creates a new mock instance
func NewMockActPoolManager(ctrl *gomock.Controller) *MockActPoolManager {
	mock := &MockActPoolManager{ctrl: ctrl}
	mock.recorder = &MockActPoolManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActPoolManager) EXPECT() *MockActPoolManagerMockRecorder {
	return m.recorder
}

// GetActionByHash mocks base method
func (m *MockActPoolManager) GetActionByHash(hash hash.Hash256) (action.SealedEnvelope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActionByHash", hash)
	ret0, _ := ret[0].(action.SealedEnvelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActionByHash indicates an expected call of GetActionByHash
func (mr *MockActPoolManagerMockRecorder) GetActionByHash(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionByHash", reflect.TypeOf((*MockActPoolManager)(nil).GetActionByHash), hash)
}
