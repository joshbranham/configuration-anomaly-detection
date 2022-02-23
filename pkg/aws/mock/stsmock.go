// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/sts/stsiface (interfaces: STSAPI)

// Package aws is a generated GoMock package.
package aws

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	sts "github.com/aws/aws-sdk-go/service/sts"
	gomock "github.com/golang/mock/gomock"
)

// MockSTSAPI is a mock of STSAPI interface.
type MockSTSAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSTSAPIMockRecorder
}

// MockSTSAPIMockRecorder is the mock recorder for MockSTSAPI.
type MockSTSAPIMockRecorder struct {
	mock *MockSTSAPI
}

// NewMockSTSAPI creates a new mock instance.
func NewMockSTSAPI(ctrl *gomock.Controller) *MockSTSAPI {
	mock := &MockSTSAPI{ctrl: ctrl}
	mock.recorder = &MockSTSAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSTSAPI) EXPECT() *MockSTSAPIMockRecorder {
	return m.recorder
}

// AssumeRole mocks base method.
func (m *MockSTSAPI) AssumeRole(arg0 *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRole", arg0)
	ret0, _ := ret[0].(*sts.AssumeRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRole indicates an expected call of AssumeRole.
func (mr *MockSTSAPIMockRecorder) AssumeRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRole", reflect.TypeOf((*MockSTSAPI)(nil).AssumeRole), arg0)
}

// AssumeRoleRequest mocks base method.
func (m *MockSTSAPI) AssumeRoleRequest(arg0 *sts.AssumeRoleInput) (*request.Request, *sts.AssumeRoleOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRoleRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.AssumeRoleOutput)
	return ret0, ret1
}

// AssumeRoleRequest indicates an expected call of AssumeRoleRequest.
func (mr *MockSTSAPIMockRecorder) AssumeRoleRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleRequest", reflect.TypeOf((*MockSTSAPI)(nil).AssumeRoleRequest), arg0)
}

// AssumeRoleWithContext mocks base method.
func (m *MockSTSAPI) AssumeRoleWithContext(arg0 context.Context, arg1 *sts.AssumeRoleInput, arg2 ...request.Option) (*sts.AssumeRoleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssumeRoleWithContext", varargs...)
	ret0, _ := ret[0].(*sts.AssumeRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleWithContext indicates an expected call of AssumeRoleWithContext.
func (mr *MockSTSAPIMockRecorder) AssumeRoleWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithContext", reflect.TypeOf((*MockSTSAPI)(nil).AssumeRoleWithContext), varargs...)
}

// AssumeRoleWithSAML mocks base method.
func (m *MockSTSAPI) AssumeRoleWithSAML(arg0 *sts.AssumeRoleWithSAMLInput) (*sts.AssumeRoleWithSAMLOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRoleWithSAML", arg0)
	ret0, _ := ret[0].(*sts.AssumeRoleWithSAMLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleWithSAML indicates an expected call of AssumeRoleWithSAML.
func (mr *MockSTSAPIMockRecorder) AssumeRoleWithSAML(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithSAML", reflect.TypeOf((*MockSTSAPI)(nil).AssumeRoleWithSAML), arg0)
}

// AssumeRoleWithSAMLRequest mocks base method.
func (m *MockSTSAPI) AssumeRoleWithSAMLRequest(arg0 *sts.AssumeRoleWithSAMLInput) (*request.Request, *sts.AssumeRoleWithSAMLOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRoleWithSAMLRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.AssumeRoleWithSAMLOutput)
	return ret0, ret1
}

// AssumeRoleWithSAMLRequest indicates an expected call of AssumeRoleWithSAMLRequest.
func (mr *MockSTSAPIMockRecorder) AssumeRoleWithSAMLRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithSAMLRequest", reflect.TypeOf((*MockSTSAPI)(nil).AssumeRoleWithSAMLRequest), arg0)
}

// AssumeRoleWithSAMLWithContext mocks base method.
func (m *MockSTSAPI) AssumeRoleWithSAMLWithContext(arg0 context.Context, arg1 *sts.AssumeRoleWithSAMLInput, arg2 ...request.Option) (*sts.AssumeRoleWithSAMLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssumeRoleWithSAMLWithContext", varargs...)
	ret0, _ := ret[0].(*sts.AssumeRoleWithSAMLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleWithSAMLWithContext indicates an expected call of AssumeRoleWithSAMLWithContext.
func (mr *MockSTSAPIMockRecorder) AssumeRoleWithSAMLWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithSAMLWithContext", reflect.TypeOf((*MockSTSAPI)(nil).AssumeRoleWithSAMLWithContext), varargs...)
}

// AssumeRoleWithWebIdentity mocks base method.
func (m *MockSTSAPI) AssumeRoleWithWebIdentity(arg0 *sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRoleWithWebIdentity", arg0)
	ret0, _ := ret[0].(*sts.AssumeRoleWithWebIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleWithWebIdentity indicates an expected call of AssumeRoleWithWebIdentity.
func (mr *MockSTSAPIMockRecorder) AssumeRoleWithWebIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithWebIdentity", reflect.TypeOf((*MockSTSAPI)(nil).AssumeRoleWithWebIdentity), arg0)
}

// AssumeRoleWithWebIdentityRequest mocks base method.
func (m *MockSTSAPI) AssumeRoleWithWebIdentityRequest(arg0 *sts.AssumeRoleWithWebIdentityInput) (*request.Request, *sts.AssumeRoleWithWebIdentityOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRoleWithWebIdentityRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.AssumeRoleWithWebIdentityOutput)
	return ret0, ret1
}

// AssumeRoleWithWebIdentityRequest indicates an expected call of AssumeRoleWithWebIdentityRequest.
func (mr *MockSTSAPIMockRecorder) AssumeRoleWithWebIdentityRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithWebIdentityRequest", reflect.TypeOf((*MockSTSAPI)(nil).AssumeRoleWithWebIdentityRequest), arg0)
}

// AssumeRoleWithWebIdentityWithContext mocks base method.
func (m *MockSTSAPI) AssumeRoleWithWebIdentityWithContext(arg0 context.Context, arg1 *sts.AssumeRoleWithWebIdentityInput, arg2 ...request.Option) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssumeRoleWithWebIdentityWithContext", varargs...)
	ret0, _ := ret[0].(*sts.AssumeRoleWithWebIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleWithWebIdentityWithContext indicates an expected call of AssumeRoleWithWebIdentityWithContext.
func (mr *MockSTSAPIMockRecorder) AssumeRoleWithWebIdentityWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithWebIdentityWithContext", reflect.TypeOf((*MockSTSAPI)(nil).AssumeRoleWithWebIdentityWithContext), varargs...)
}

// DecodeAuthorizationMessage mocks base method.
func (m *MockSTSAPI) DecodeAuthorizationMessage(arg0 *sts.DecodeAuthorizationMessageInput) (*sts.DecodeAuthorizationMessageOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeAuthorizationMessage", arg0)
	ret0, _ := ret[0].(*sts.DecodeAuthorizationMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeAuthorizationMessage indicates an expected call of DecodeAuthorizationMessage.
func (mr *MockSTSAPIMockRecorder) DecodeAuthorizationMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeAuthorizationMessage", reflect.TypeOf((*MockSTSAPI)(nil).DecodeAuthorizationMessage), arg0)
}

// DecodeAuthorizationMessageRequest mocks base method.
func (m *MockSTSAPI) DecodeAuthorizationMessageRequest(arg0 *sts.DecodeAuthorizationMessageInput) (*request.Request, *sts.DecodeAuthorizationMessageOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeAuthorizationMessageRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.DecodeAuthorizationMessageOutput)
	return ret0, ret1
}

// DecodeAuthorizationMessageRequest indicates an expected call of DecodeAuthorizationMessageRequest.
func (mr *MockSTSAPIMockRecorder) DecodeAuthorizationMessageRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeAuthorizationMessageRequest", reflect.TypeOf((*MockSTSAPI)(nil).DecodeAuthorizationMessageRequest), arg0)
}

// DecodeAuthorizationMessageWithContext mocks base method.
func (m *MockSTSAPI) DecodeAuthorizationMessageWithContext(arg0 context.Context, arg1 *sts.DecodeAuthorizationMessageInput, arg2 ...request.Option) (*sts.DecodeAuthorizationMessageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DecodeAuthorizationMessageWithContext", varargs...)
	ret0, _ := ret[0].(*sts.DecodeAuthorizationMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeAuthorizationMessageWithContext indicates an expected call of DecodeAuthorizationMessageWithContext.
func (mr *MockSTSAPIMockRecorder) DecodeAuthorizationMessageWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeAuthorizationMessageWithContext", reflect.TypeOf((*MockSTSAPI)(nil).DecodeAuthorizationMessageWithContext), varargs...)
}

// GetAccessKeyInfo mocks base method.
func (m *MockSTSAPI) GetAccessKeyInfo(arg0 *sts.GetAccessKeyInfoInput) (*sts.GetAccessKeyInfoOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessKeyInfo", arg0)
	ret0, _ := ret[0].(*sts.GetAccessKeyInfoOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessKeyInfo indicates an expected call of GetAccessKeyInfo.
func (mr *MockSTSAPIMockRecorder) GetAccessKeyInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessKeyInfo", reflect.TypeOf((*MockSTSAPI)(nil).GetAccessKeyInfo), arg0)
}

// GetAccessKeyInfoRequest mocks base method.
func (m *MockSTSAPI) GetAccessKeyInfoRequest(arg0 *sts.GetAccessKeyInfoInput) (*request.Request, *sts.GetAccessKeyInfoOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessKeyInfoRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.GetAccessKeyInfoOutput)
	return ret0, ret1
}

// GetAccessKeyInfoRequest indicates an expected call of GetAccessKeyInfoRequest.
func (mr *MockSTSAPIMockRecorder) GetAccessKeyInfoRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessKeyInfoRequest", reflect.TypeOf((*MockSTSAPI)(nil).GetAccessKeyInfoRequest), arg0)
}

// GetAccessKeyInfoWithContext mocks base method.
func (m *MockSTSAPI) GetAccessKeyInfoWithContext(arg0 context.Context, arg1 *sts.GetAccessKeyInfoInput, arg2 ...request.Option) (*sts.GetAccessKeyInfoOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessKeyInfoWithContext", varargs...)
	ret0, _ := ret[0].(*sts.GetAccessKeyInfoOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessKeyInfoWithContext indicates an expected call of GetAccessKeyInfoWithContext.
func (mr *MockSTSAPIMockRecorder) GetAccessKeyInfoWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessKeyInfoWithContext", reflect.TypeOf((*MockSTSAPI)(nil).GetAccessKeyInfoWithContext), varargs...)
}

// GetCallerIdentity mocks base method.
func (m *MockSTSAPI) GetCallerIdentity(arg0 *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCallerIdentity", arg0)
	ret0, _ := ret[0].(*sts.GetCallerIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCallerIdentity indicates an expected call of GetCallerIdentity.
func (mr *MockSTSAPIMockRecorder) GetCallerIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallerIdentity", reflect.TypeOf((*MockSTSAPI)(nil).GetCallerIdentity), arg0)
}

// GetCallerIdentityRequest mocks base method.
func (m *MockSTSAPI) GetCallerIdentityRequest(arg0 *sts.GetCallerIdentityInput) (*request.Request, *sts.GetCallerIdentityOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCallerIdentityRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.GetCallerIdentityOutput)
	return ret0, ret1
}

// GetCallerIdentityRequest indicates an expected call of GetCallerIdentityRequest.
func (mr *MockSTSAPIMockRecorder) GetCallerIdentityRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallerIdentityRequest", reflect.TypeOf((*MockSTSAPI)(nil).GetCallerIdentityRequest), arg0)
}

// GetCallerIdentityWithContext mocks base method.
func (m *MockSTSAPI) GetCallerIdentityWithContext(arg0 context.Context, arg1 *sts.GetCallerIdentityInput, arg2 ...request.Option) (*sts.GetCallerIdentityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCallerIdentityWithContext", varargs...)
	ret0, _ := ret[0].(*sts.GetCallerIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCallerIdentityWithContext indicates an expected call of GetCallerIdentityWithContext.
func (mr *MockSTSAPIMockRecorder) GetCallerIdentityWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallerIdentityWithContext", reflect.TypeOf((*MockSTSAPI)(nil).GetCallerIdentityWithContext), varargs...)
}

// GetFederationToken mocks base method.
func (m *MockSTSAPI) GetFederationToken(arg0 *sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFederationToken", arg0)
	ret0, _ := ret[0].(*sts.GetFederationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFederationToken indicates an expected call of GetFederationToken.
func (mr *MockSTSAPIMockRecorder) GetFederationToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFederationToken", reflect.TypeOf((*MockSTSAPI)(nil).GetFederationToken), arg0)
}

// GetFederationTokenRequest mocks base method.
func (m *MockSTSAPI) GetFederationTokenRequest(arg0 *sts.GetFederationTokenInput) (*request.Request, *sts.GetFederationTokenOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFederationTokenRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.GetFederationTokenOutput)
	return ret0, ret1
}

// GetFederationTokenRequest indicates an expected call of GetFederationTokenRequest.
func (mr *MockSTSAPIMockRecorder) GetFederationTokenRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFederationTokenRequest", reflect.TypeOf((*MockSTSAPI)(nil).GetFederationTokenRequest), arg0)
}

// GetFederationTokenWithContext mocks base method.
func (m *MockSTSAPI) GetFederationTokenWithContext(arg0 context.Context, arg1 *sts.GetFederationTokenInput, arg2 ...request.Option) (*sts.GetFederationTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFederationTokenWithContext", varargs...)
	ret0, _ := ret[0].(*sts.GetFederationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFederationTokenWithContext indicates an expected call of GetFederationTokenWithContext.
func (mr *MockSTSAPIMockRecorder) GetFederationTokenWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFederationTokenWithContext", reflect.TypeOf((*MockSTSAPI)(nil).GetFederationTokenWithContext), varargs...)
}

// GetSessionToken mocks base method.
func (m *MockSTSAPI) GetSessionToken(arg0 *sts.GetSessionTokenInput) (*sts.GetSessionTokenOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionToken", arg0)
	ret0, _ := ret[0].(*sts.GetSessionTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionToken indicates an expected call of GetSessionToken.
func (mr *MockSTSAPIMockRecorder) GetSessionToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionToken", reflect.TypeOf((*MockSTSAPI)(nil).GetSessionToken), arg0)
}

// GetSessionTokenRequest mocks base method.
func (m *MockSTSAPI) GetSessionTokenRequest(arg0 *sts.GetSessionTokenInput) (*request.Request, *sts.GetSessionTokenOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionTokenRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.GetSessionTokenOutput)
	return ret0, ret1
}

// GetSessionTokenRequest indicates an expected call of GetSessionTokenRequest.
func (mr *MockSTSAPIMockRecorder) GetSessionTokenRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionTokenRequest", reflect.TypeOf((*MockSTSAPI)(nil).GetSessionTokenRequest), arg0)
}

// GetSessionTokenWithContext mocks base method.
func (m *MockSTSAPI) GetSessionTokenWithContext(arg0 context.Context, arg1 *sts.GetSessionTokenInput, arg2 ...request.Option) (*sts.GetSessionTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSessionTokenWithContext", varargs...)
	ret0, _ := ret[0].(*sts.GetSessionTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionTokenWithContext indicates an expected call of GetSessionTokenWithContext.
func (mr *MockSTSAPIMockRecorder) GetSessionTokenWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionTokenWithContext", reflect.TypeOf((*MockSTSAPI)(nil).GetSessionTokenWithContext), varargs...)
}
