// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// OracleKeeper is an autogenerated mock type for the OracleKeeper type
type OracleKeeper struct {
	mock.Mock
}

// GetAllCurrencyPairs provides a mock function with given fields: ctx
func (_m *OracleKeeper) GetAllCurrencyPairs(ctx types.Context) []oracletypes.CurrencyPair {
	ret := _m.Called(ctx)

	var r0 []oracletypes.CurrencyPair
	if rf, ok := ret.Get(0).(func(types.Context) []oracletypes.CurrencyPair); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]oracletypes.CurrencyPair)
		}
	}

	return r0
}

// SetPriceForCurrencyPair provides a mock function with given fields: ctx, cp, qp
func (_m *OracleKeeper) SetPriceForCurrencyPair(ctx types.Context, cp oracletypes.CurrencyPair, qp oracletypes.QuotePrice) error {
	ret := _m.Called(ctx, cp, qp)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, oracletypes.CurrencyPair, oracletypes.QuotePrice) error); ok {
		r0 = rf(ctx, cp, qp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOracleKeeper creates a new instance of OracleKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOracleKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *OracleKeeper {
	mock := &OracleKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
