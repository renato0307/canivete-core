// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	"github.com/renato0307/canivete-core/pkg/programming"
	mock "github.com/stretchr/testify/mock"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// NewUuid provides a mock function with given fields:
func (_m *Interface) NewUuid() programming.UuidOutput {
	ret := _m.Called()

	var r0 programming.UuidOutput
	if rf, ok := ret.Get(0).(func() programming.UuidOutput); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(programming.UuidOutput)
	}

	return r0
}
