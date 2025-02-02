// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	io "io"

	mock "github.com/stretchr/testify/mock"

	treatment "github.com/caraml-dev/xp/clients/treatment"
)

// ClientInterface is an autogenerated mock type for the ClientInterface type
type ClientInterface struct {
	mock.Mock
}

// FetchTreatment provides a mock function with given fields: ctx, projectId, params, body, reqEditors
func (_m *ClientInterface) FetchTreatment(ctx context.Context, projectId int64, params *treatment.FetchTreatmentParams, body treatment.FetchTreatmentJSONRequestBody, reqEditors ...treatment.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, projectId, params, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, int64, *treatment.FetchTreatmentParams, treatment.FetchTreatmentJSONRequestBody, ...treatment.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, projectId, params, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, *treatment.FetchTreatmentParams, treatment.FetchTreatmentJSONRequestBody, ...treatment.RequestEditorFn) error); ok {
		r1 = rf(ctx, projectId, params, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchTreatmentWithBody provides a mock function with given fields: ctx, projectId, params, contentType, body, reqEditors
func (_m *ClientInterface) FetchTreatmentWithBody(ctx context.Context, projectId int64, params *treatment.FetchTreatmentParams, contentType string, body io.Reader, reqEditors ...treatment.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, projectId, params, contentType, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, int64, *treatment.FetchTreatmentParams, string, io.Reader, ...treatment.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, projectId, params, contentType, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, *treatment.FetchTreatmentParams, string, io.Reader, ...treatment.RequestEditorFn) error); ok {
		r1 = rf(ctx, projectId, params, contentType, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClientInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewClientInterface creates a new instance of ClientInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientInterface(t mockConstructorTestingTNewClientInterface) *ClientInterface {
	mock := &ClientInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
