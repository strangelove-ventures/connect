// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RequestHandler is an autogenerated mock type for the RequestHandler type
type RequestHandler struct {
	mock.Mock
}

type RequestHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *RequestHandler) EXPECT() *RequestHandler_Expecter {
	return &RequestHandler_Expecter{mock: &_m.Mock}
}

// Do provides a mock function with given fields: ctx, url
func (_m *RequestHandler) Do(ctx context.Context, url string) (*http.Response, error) {
	ret := _m.Called(ctx, url)

	if len(ret) == 0 {
		panic("no return value specified for Do")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*http.Response, error)); ok {
		return rf(ctx, url)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *http.Response); ok {
		r0 = rf(ctx, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestHandler_Do_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Do'
type RequestHandler_Do_Call struct {
	*mock.Call
}

// Do is a helper method to define mock.On call
//   - ctx context.Context
//   - url string
func (_e *RequestHandler_Expecter) Do(ctx interface{}, url interface{}) *RequestHandler_Do_Call {
	return &RequestHandler_Do_Call{Call: _e.mock.On("Do", ctx, url)}
}

func (_c *RequestHandler_Do_Call) Run(run func(ctx context.Context, url string)) *RequestHandler_Do_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *RequestHandler_Do_Call) Return(_a0 *http.Response, _a1 error) *RequestHandler_Do_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RequestHandler_Do_Call) RunAndReturn(run func(context.Context, string) (*http.Response, error)) *RequestHandler_Do_Call {
	_c.Call.Return(run)
	return _c
}

// Type provides a mock function with no fields
func (_m *RequestHandler) Type() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RequestHandler_Type_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Type'
type RequestHandler_Type_Call struct {
	*mock.Call
}

// Type is a helper method to define mock.On call
func (_e *RequestHandler_Expecter) Type() *RequestHandler_Type_Call {
	return &RequestHandler_Type_Call{Call: _e.mock.On("Type")}
}

func (_c *RequestHandler_Type_Call) Run(run func()) *RequestHandler_Type_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RequestHandler_Type_Call) Return(_a0 string) *RequestHandler_Type_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequestHandler_Type_Call) RunAndReturn(run func() string) *RequestHandler_Type_Call {
	_c.Call.Return(run)
	return _c
}

// NewRequestHandler creates a new instance of RequestHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequestHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequestHandler {
	mock := &RequestHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
