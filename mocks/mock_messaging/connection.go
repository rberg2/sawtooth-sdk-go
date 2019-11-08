// Code generated by MockGen. DO NOT EDIT.
// Source: connection.go

// Package mock_messaging is a generated GoMock package.
package mock_messaging

import (
	gomock "github.com/golang/mock/gomock"
	validator_pb2 "protobuf/validator_pb2"
	zmq4 "github.com/pebbe/zmq4"
	reflect "reflect"
)

// MockConnection is a mock of Connection interface
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// SendData mocks base method
func (m *MockConnection) SendData(id string, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendData", id, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendData indicates an expected call of SendData
func (mr *MockConnectionMockRecorder) SendData(id, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendData", reflect.TypeOf((*MockConnection)(nil).SendData), id, data)
}

// SendNewMsg mocks base method
func (m *MockConnection) SendNewMsg(t validator_pb2.Message_MessageType, c []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendNewMsg", t, c)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendNewMsg indicates an expected call of SendNewMsg
func (mr *MockConnectionMockRecorder) SendNewMsg(t, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendNewMsg", reflect.TypeOf((*MockConnection)(nil).SendNewMsg), t, c)
}

// SendNewMsgTo mocks base method
func (m *MockConnection) SendNewMsgTo(id string, t validator_pb2.Message_MessageType, c []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendNewMsgTo", id, t, c)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendNewMsgTo indicates an expected call of SendNewMsgTo
func (mr *MockConnectionMockRecorder) SendNewMsgTo(id, t, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendNewMsgTo", reflect.TypeOf((*MockConnection)(nil).SendNewMsgTo), id, t, c)
}

// SendMsg mocks base method
func (m *MockConnection) SendMsg(t validator_pb2.Message_MessageType, c []byte, corrId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", t, c, corrId)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockConnectionMockRecorder) SendMsg(t, c, corrId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockConnection)(nil).SendMsg), t, c, corrId)
}

// SendMsgTo mocks base method
func (m *MockConnection) SendMsgTo(id string, t validator_pb2.Message_MessageType, c []byte, corrId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsgTo", id, t, c, corrId)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsgTo indicates an expected call of SendMsgTo
func (mr *MockConnectionMockRecorder) SendMsgTo(id, t, c, corrId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsgTo", reflect.TypeOf((*MockConnection)(nil).SendMsgTo), id, t, c, corrId)
}

// RecvData mocks base method
func (m *MockConnection) RecvData() (string, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvData")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RecvData indicates an expected call of RecvData
func (mr *MockConnectionMockRecorder) RecvData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvData", reflect.TypeOf((*MockConnection)(nil).RecvData))
}

// RecvMsg mocks base method
func (m *MockConnection) RecvMsg() (string, *validator_pb2.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*validator_pb2.Message)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockConnectionMockRecorder) RecvMsg() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockConnection)(nil).RecvMsg))
}

// RecvMsgWithId mocks base method
func (m *MockConnection) RecvMsgWithId(corrId string) (string, *validator_pb2.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsgWithId", corrId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*validator_pb2.Message)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RecvMsgWithId indicates an expected call of RecvMsgWithId
func (mr *MockConnectionMockRecorder) RecvMsgWithId(corrId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsgWithId", reflect.TypeOf((*MockConnection)(nil).RecvMsgWithId), corrId)
}

// Close mocks base method
func (m *MockConnection) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
}

// Socket mocks base method
func (m *MockConnection) Socket() *zmq4.Socket {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Socket")
	ret0, _ := ret[0].(*zmq4.Socket)
	return ret0
}

// Socket indicates an expected call of Socket
func (mr *MockConnectionMockRecorder) Socket() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Socket", reflect.TypeOf((*MockConnection)(nil).Socket))
}

// Monitor mocks base method
func (m *MockConnection) Monitor(arg0 zmq4.Event) (*zmq4.Socket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Monitor", arg0)
	ret0, _ := ret[0].(*zmq4.Socket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Monitor indicates an expected call of Monitor
func (mr *MockConnectionMockRecorder) Monitor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Monitor", reflect.TypeOf((*MockConnection)(nil).Monitor), arg0)
}

// Identity mocks base method
func (m *MockConnection) Identity() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identity")
	ret0, _ := ret[0].(string)
	return ret0
}

// Identity indicates an expected call of Identity
func (mr *MockConnectionMockRecorder) Identity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identity", reflect.TypeOf((*MockConnection)(nil).Identity))
}
