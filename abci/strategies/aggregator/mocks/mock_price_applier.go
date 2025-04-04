// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"

	big "math/big"

	mock "github.com/stretchr/testify/mock"

	pkgtypes "github.com/skip-mev/connect/v2/pkg/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// PriceApplier is an autogenerated mock type for the PriceApplier type
type PriceApplier struct {
	mock.Mock
}

type PriceApplier_Expecter struct {
	mock *mock.Mock
}

func (_m *PriceApplier) EXPECT() *PriceApplier_Expecter {
	return &PriceApplier_Expecter{mock: &_m.Mock}
}

// ApplyPricesFromVoteExtensions provides a mock function with given fields: ctx, req
func (_m *PriceApplier) ApplyPricesFromVoteExtensions(ctx types.Context, req *abcitypes.RequestFinalizeBlock) (map[pkgtypes.CurrencyPair]*big.Int, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for ApplyPricesFromVoteExtensions")
	}

	var r0 map[pkgtypes.CurrencyPair]*big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, *abcitypes.RequestFinalizeBlock) (map[pkgtypes.CurrencyPair]*big.Int, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *abcitypes.RequestFinalizeBlock) map[pkgtypes.CurrencyPair]*big.Int); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[pkgtypes.CurrencyPair]*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, *abcitypes.RequestFinalizeBlock) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PriceApplier_ApplyPricesFromVoteExtensions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyPricesFromVoteExtensions'
type PriceApplier_ApplyPricesFromVoteExtensions_Call struct {
	*mock.Call
}

// ApplyPricesFromVoteExtensions is a helper method to define mock.On call
//   - ctx types.Context
//   - req *abcitypes.RequestFinalizeBlock
func (_e *PriceApplier_Expecter) ApplyPricesFromVoteExtensions(ctx interface{}, req interface{}) *PriceApplier_ApplyPricesFromVoteExtensions_Call {
	return &PriceApplier_ApplyPricesFromVoteExtensions_Call{Call: _e.mock.On("ApplyPricesFromVoteExtensions", ctx, req)}
}

func (_c *PriceApplier_ApplyPricesFromVoteExtensions_Call) Run(run func(ctx types.Context, req *abcitypes.RequestFinalizeBlock)) *PriceApplier_ApplyPricesFromVoteExtensions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(*abcitypes.RequestFinalizeBlock))
	})
	return _c
}

func (_c *PriceApplier_ApplyPricesFromVoteExtensions_Call) Return(_a0 map[pkgtypes.CurrencyPair]*big.Int, _a1 error) *PriceApplier_ApplyPricesFromVoteExtensions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PriceApplier_ApplyPricesFromVoteExtensions_Call) RunAndReturn(run func(types.Context, *abcitypes.RequestFinalizeBlock) (map[pkgtypes.CurrencyPair]*big.Int, error)) *PriceApplier_ApplyPricesFromVoteExtensions_Call {
	_c.Call.Return(run)
	return _c
}

// GetPricesForValidator provides a mock function with given fields: validator
func (_m *PriceApplier) GetPricesForValidator(validator types.ConsAddress) map[pkgtypes.CurrencyPair]*big.Int {
	ret := _m.Called(validator)

	if len(ret) == 0 {
		panic("no return value specified for GetPricesForValidator")
	}

	var r0 map[pkgtypes.CurrencyPair]*big.Int
	if rf, ok := ret.Get(0).(func(types.ConsAddress) map[pkgtypes.CurrencyPair]*big.Int); ok {
		r0 = rf(validator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[pkgtypes.CurrencyPair]*big.Int)
		}
	}

	return r0
}

// PriceApplier_GetPricesForValidator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPricesForValidator'
type PriceApplier_GetPricesForValidator_Call struct {
	*mock.Call
}

// GetPricesForValidator is a helper method to define mock.On call
//   - validator types.ConsAddress
func (_e *PriceApplier_Expecter) GetPricesForValidator(validator interface{}) *PriceApplier_GetPricesForValidator_Call {
	return &PriceApplier_GetPricesForValidator_Call{Call: _e.mock.On("GetPricesForValidator", validator)}
}

func (_c *PriceApplier_GetPricesForValidator_Call) Run(run func(validator types.ConsAddress)) *PriceApplier_GetPricesForValidator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.ConsAddress))
	})
	return _c
}

func (_c *PriceApplier_GetPricesForValidator_Call) Return(_a0 map[pkgtypes.CurrencyPair]*big.Int) *PriceApplier_GetPricesForValidator_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PriceApplier_GetPricesForValidator_Call) RunAndReturn(run func(types.ConsAddress) map[pkgtypes.CurrencyPair]*big.Int) *PriceApplier_GetPricesForValidator_Call {
	_c.Call.Return(run)
	return _c
}

// NewPriceApplier creates a new instance of PriceApplier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPriceApplier(t interface {
	mock.TestingT
	Cleanup(func())
}) *PriceApplier {
	mock := &PriceApplier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
