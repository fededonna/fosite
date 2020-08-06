// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fededonna/fosite (interfaces: Requester)

// Package internal is a generated GoMock package.
package internal

import (
	url "net/url"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"

	fosite "github.com/fededonna/fosite"
)

// MockRequester is a mock of Requester interface
type MockRequester struct {
	ctrl     *gomock.Controller
	recorder *MockRequesterMockRecorder
}

// MockRequesterMockRecorder is the mock recorder for MockRequester
type MockRequesterMockRecorder struct {
	mock *MockRequester
}

// NewMockRequester creates a new mock instance
func NewMockRequester(ctrl *gomock.Controller) *MockRequester {
	mock := &MockRequester{ctrl: ctrl}
	mock.recorder = &MockRequesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRequester) EXPECT() *MockRequesterMockRecorder {
	return m.recorder
}

// AppendRequestedScope mocks base method
func (m *MockRequester) AppendRequestedScope(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AppendRequestedScope", arg0)
}

// AppendRequestedScope indicates an expected call of AppendRequestedScope
func (mr *MockRequesterMockRecorder) AppendRequestedScope(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendRequestedScope", reflect.TypeOf((*MockRequester)(nil).AppendRequestedScope), arg0)
}

// GetClient mocks base method
func (m *MockRequester) GetClient() fosite.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(fosite.Client)
	return ret0
}

// GetClient indicates an expected call of GetClient
func (mr *MockRequesterMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockRequester)(nil).GetClient))
}

// GetGrantedAudience mocks base method
func (m *MockRequester) GetGrantedAudience() fosite.Arguments {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGrantedAudience")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetGrantedAudience indicates an expected call of GetGrantedAudience
func (mr *MockRequesterMockRecorder) GetGrantedAudience() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGrantedAudience", reflect.TypeOf((*MockRequester)(nil).GetGrantedAudience))
}

// GetGrantedScopes mocks base method
func (m *MockRequester) GetGrantedScopes() fosite.Arguments {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGrantedScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetGrantedScopes indicates an expected call of GetGrantedScopes
func (mr *MockRequesterMockRecorder) GetGrantedScopes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGrantedScopes", reflect.TypeOf((*MockRequester)(nil).GetGrantedScopes))
}

// GetID mocks base method
func (m *MockRequester) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockRequesterMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockRequester)(nil).GetID))
}

// GetRequestForm mocks base method
func (m *MockRequester) GetRequestForm() url.Values {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestForm")
	ret0, _ := ret[0].(url.Values)
	return ret0
}

// GetRequestForm indicates an expected call of GetRequestForm
func (mr *MockRequesterMockRecorder) GetRequestForm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestForm", reflect.TypeOf((*MockRequester)(nil).GetRequestForm))
}

// GetRequestedAt mocks base method
func (m *MockRequester) GetRequestedAt() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetRequestedAt indicates an expected call of GetRequestedAt
func (mr *MockRequesterMockRecorder) GetRequestedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestedAt", reflect.TypeOf((*MockRequester)(nil).GetRequestedAt))
}

// GetRequestedAudience mocks base method
func (m *MockRequester) GetRequestedAudience() fosite.Arguments {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestedAudience")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetRequestedAudience indicates an expected call of GetRequestedAudience
func (mr *MockRequesterMockRecorder) GetRequestedAudience() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestedAudience", reflect.TypeOf((*MockRequester)(nil).GetRequestedAudience))
}

// GetRequestedScopes mocks base method
func (m *MockRequester) GetRequestedScopes() fosite.Arguments {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestedScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetRequestedScopes indicates an expected call of GetRequestedScopes
func (mr *MockRequesterMockRecorder) GetRequestedScopes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestedScopes", reflect.TypeOf((*MockRequester)(nil).GetRequestedScopes))
}

// GetSession mocks base method
func (m *MockRequester) GetSession() fosite.Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession")
	ret0, _ := ret[0].(fosite.Session)
	return ret0
}

// GetSession indicates an expected call of GetSession
func (mr *MockRequesterMockRecorder) GetSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockRequester)(nil).GetSession))
}

// GrantAudience mocks base method
func (m *MockRequester) GrantAudience(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GrantAudience", arg0)
}

// GrantAudience indicates an expected call of GrantAudience
func (mr *MockRequesterMockRecorder) GrantAudience(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantAudience", reflect.TypeOf((*MockRequester)(nil).GrantAudience), arg0)
}

// GrantScope mocks base method
func (m *MockRequester) GrantScope(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GrantScope", arg0)
}

// GrantScope indicates an expected call of GrantScope
func (mr *MockRequesterMockRecorder) GrantScope(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantScope", reflect.TypeOf((*MockRequester)(nil).GrantScope), arg0)
}

// Merge mocks base method
func (m *MockRequester) Merge(arg0 fosite.Requester) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Merge", arg0)
}

// Merge indicates an expected call of Merge
func (mr *MockRequesterMockRecorder) Merge(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Merge", reflect.TypeOf((*MockRequester)(nil).Merge), arg0)
}

// Sanitize mocks base method
func (m *MockRequester) Sanitize(arg0 []string) fosite.Requester {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sanitize", arg0)
	ret0, _ := ret[0].(fosite.Requester)
	return ret0
}

// Sanitize indicates an expected call of Sanitize
func (mr *MockRequesterMockRecorder) Sanitize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sanitize", reflect.TypeOf((*MockRequester)(nil).Sanitize), arg0)
}

// SetID mocks base method
func (m *MockRequester) SetID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetID", arg0)
}

// SetID indicates an expected call of SetID
func (mr *MockRequesterMockRecorder) SetID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetID", reflect.TypeOf((*MockRequester)(nil).SetID), arg0)
}

// SetRequestedAudience mocks base method
func (m *MockRequester) SetRequestedAudience(arg0 fosite.Arguments) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRequestedAudience", arg0)
}

// SetRequestedAudience indicates an expected call of SetRequestedAudience
func (mr *MockRequesterMockRecorder) SetRequestedAudience(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRequestedAudience", reflect.TypeOf((*MockRequester)(nil).SetRequestedAudience), arg0)
}

// SetRequestedScopes mocks base method
func (m *MockRequester) SetRequestedScopes(arg0 fosite.Arguments) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRequestedScopes", arg0)
}

// SetRequestedScopes indicates an expected call of SetRequestedScopes
func (mr *MockRequesterMockRecorder) SetRequestedScopes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRequestedScopes", reflect.TypeOf((*MockRequester)(nil).SetRequestedScopes), arg0)
}

// SetSession mocks base method
func (m *MockRequester) SetSession(arg0 fosite.Session) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSession", arg0)
}

// SetSession indicates an expected call of SetSession
func (mr *MockRequesterMockRecorder) SetSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSession", reflect.TypeOf((*MockRequester)(nil).SetSession), arg0)
}
