// Code generated by MockGen. DO NOT EDIT.
// Source: C:\dev\terraform-provider-azuredevops\terraform-provider-azuredevops-ni-pre\vendor\github.com\microsoft\azure-devops-go-api\azuredevops\pipelineschecks\client.go

// Package mock_pipelineschecks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pipelineschecks "github.com/microsoft/azure-devops-go-api/azuredevops/pipelineschecks"
)

// PipelinesChecksClientExtrasV5 is a mock of Client interface.
type PipelinesChecksClientExtrasV5 struct {
	ctrl     *gomock.Controller
	recorder *PipelinesChecksClientExtrasV5MockRecorder
}

// PipelinesChecksClientExtrasV5MockRecorder is the mock recorder for PipelinesChecksClientExtrasV5.
type PipelinesChecksClientExtrasV5MockRecorder struct {
	mock *PipelinesChecksClientExtrasV5
}

// NewPipelinesChecksClientExtrasV5 creates a new mock instance.
func NewPipelinesChecksClientExtrasV5(ctrl *gomock.Controller) *PipelinesChecksClientExtrasV5 {
	mock := &PipelinesChecksClientExtrasV5{ctrl: ctrl}
	mock.recorder = &PipelinesChecksClientExtrasV5MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *PipelinesChecksClientExtrasV5) EXPECT() *PipelinesChecksClientExtrasV5MockRecorder {
	return m.recorder
}

// GetCheckConfiguration mocks base method.
func (m *PipelinesChecksClientExtrasV5) GetCheckConfiguration(arg0 context.Context, arg1 pipelineschecks.GetCheckConfigurationArgs) (*pipelineschecks.CheckConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*pipelineschecks.CheckConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckConfiguration indicates an expected call of GetCheckConfiguration.
func (mr *PipelinesChecksClientExtrasV5MockRecorder) GetCheckConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckConfiguration", reflect.TypeOf((*PipelinesChecksClientExtrasV5)(nil).GetCheckConfiguration), arg0, arg1)
}
