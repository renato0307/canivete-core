// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	datetime "github.com/renato0307/canivete-core/interface/datetime"
	mock "github.com/stretchr/testify/mock"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// FromUnitTimestamp provides a mock function with given fields: unixTime
func (_m *Interface) FromUnitTimestamp(unixTime int64) datetime.FromUnixTimestampOutput {
	ret := _m.Called(unixTime)

	var r0 datetime.FromUnixTimestampOutput
	if rf, ok := ret.Get(0).(func(int64) datetime.FromUnixTimestampOutput); ok {
		r0 = rf(unixTime)
	} else {
		r0 = ret.Get(0).(datetime.FromUnixTimestampOutput)
	}

	return r0
}