// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	time "time"

	types "github.com/skip-mev/connect/v2/x/marketmap/types"
)

// Oracle is an autogenerated mock type for the Oracle type
type Oracle struct {
	mock.Mock
}

type Oracle_Expecter struct {
	mock *mock.Mock
}

func (_m *Oracle) EXPECT() *Oracle_Expecter {
	return &Oracle_Expecter{mock: &_m.Mock}
}

// GetLastSyncTime provides a mock function with no fields
func (_m *Oracle) GetLastSyncTime() time.Time {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLastSyncTime")
	}

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// Oracle_GetLastSyncTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastSyncTime'
type Oracle_GetLastSyncTime_Call struct {
	*mock.Call
}

// GetLastSyncTime is a helper method to define mock.On call
func (_e *Oracle_Expecter) GetLastSyncTime() *Oracle_GetLastSyncTime_Call {
	return &Oracle_GetLastSyncTime_Call{Call: _e.mock.On("GetLastSyncTime")}
}

func (_c *Oracle_GetLastSyncTime_Call) Run(run func()) *Oracle_GetLastSyncTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Oracle_GetLastSyncTime_Call) Return(_a0 time.Time) *Oracle_GetLastSyncTime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Oracle_GetLastSyncTime_Call) RunAndReturn(run func() time.Time) *Oracle_GetLastSyncTime_Call {
	_c.Call.Return(run)
	return _c
}

// GetMarketMap provides a mock function with no fields
func (_m *Oracle) GetMarketMap() types.MarketMap {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMarketMap")
	}

	var r0 types.MarketMap
	if rf, ok := ret.Get(0).(func() types.MarketMap); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.MarketMap)
	}

	return r0
}

// Oracle_GetMarketMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMarketMap'
type Oracle_GetMarketMap_Call struct {
	*mock.Call
}

// GetMarketMap is a helper method to define mock.On call
func (_e *Oracle_Expecter) GetMarketMap() *Oracle_GetMarketMap_Call {
	return &Oracle_GetMarketMap_Call{Call: _e.mock.On("GetMarketMap")}
}

func (_c *Oracle_GetMarketMap_Call) Run(run func()) *Oracle_GetMarketMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Oracle_GetMarketMap_Call) Return(_a0 types.MarketMap) *Oracle_GetMarketMap_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Oracle_GetMarketMap_Call) RunAndReturn(run func() types.MarketMap) *Oracle_GetMarketMap_Call {
	_c.Call.Return(run)
	return _c
}

// GetPrices provides a mock function with no fields
func (_m *Oracle) GetPrices() map[string]*big.Float {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPrices")
	}

	var r0 map[string]*big.Float
	if rf, ok := ret.Get(0).(func() map[string]*big.Float); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*big.Float)
		}
	}

	return r0
}

// Oracle_GetPrices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPrices'
type Oracle_GetPrices_Call struct {
	*mock.Call
}

// GetPrices is a helper method to define mock.On call
func (_e *Oracle_Expecter) GetPrices() *Oracle_GetPrices_Call {
	return &Oracle_GetPrices_Call{Call: _e.mock.On("GetPrices")}
}

func (_c *Oracle_GetPrices_Call) Run(run func()) *Oracle_GetPrices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Oracle_GetPrices_Call) Return(_a0 map[string]*big.Float) *Oracle_GetPrices_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Oracle_GetPrices_Call) RunAndReturn(run func() map[string]*big.Float) *Oracle_GetPrices_Call {
	_c.Call.Return(run)
	return _c
}

// IsRunning provides a mock function with no fields
func (_m *Oracle) IsRunning() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsRunning")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Oracle_IsRunning_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsRunning'
type Oracle_IsRunning_Call struct {
	*mock.Call
}

// IsRunning is a helper method to define mock.On call
func (_e *Oracle_Expecter) IsRunning() *Oracle_IsRunning_Call {
	return &Oracle_IsRunning_Call{Call: _e.mock.On("IsRunning")}
}

func (_c *Oracle_IsRunning_Call) Run(run func()) *Oracle_IsRunning_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Oracle_IsRunning_Call) Return(_a0 bool) *Oracle_IsRunning_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Oracle_IsRunning_Call) RunAndReturn(run func() bool) *Oracle_IsRunning_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *Oracle) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Oracle_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Oracle_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Oracle_Expecter) Start(ctx interface{}) *Oracle_Start_Call {
	return &Oracle_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *Oracle_Start_Call) Run(run func(ctx context.Context)) *Oracle_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Oracle_Start_Call) Return(_a0 error) *Oracle_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Oracle_Start_Call) RunAndReturn(run func(context.Context) error) *Oracle_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with no fields
func (_m *Oracle) Stop() {
	_m.Called()
}

// Oracle_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type Oracle_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *Oracle_Expecter) Stop() *Oracle_Stop_Call {
	return &Oracle_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *Oracle_Stop_Call) Run(run func()) *Oracle_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Oracle_Stop_Call) Return() *Oracle_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *Oracle_Stop_Call) RunAndReturn(run func()) *Oracle_Stop_Call {
	_c.Run(run)
	return _c
}

// NewOracle creates a new instance of Oracle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOracle(t interface {
	mock.TestingT
	Cleanup(func())
}) *Oracle {
	mock := &Oracle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
