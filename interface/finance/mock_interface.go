// Code generated by mockery v1.0.0. DO NOT EDIT.

package finance

import mock "github.com/stretchr/testify/mock"

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

// CalculateCompoundInterests provides a mock function with given fields: p, n, t, m, y, rInt
func (_m *MockInterface) CalculateCompoundInterests(p float64, n float64, t float64, m float64, y float64, rInt float64) CompoundInterestsOutput {
	ret := _m.Called(p, n, t, m, y, rInt)

	var r0 CompoundInterestsOutput
	if rf, ok := ret.Get(0).(func(float64, float64, float64, float64, float64, float64) CompoundInterestsOutput); ok {
		r0 = rf(p, n, t, m, y, rInt)
	} else {
		r0 = ret.Get(0).(CompoundInterestsOutput)
	}

	return r0
}