// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	port "github.com/rpuggian/ports-service/proto"
)

// PortService is an autogenerated mock type for the PortService type
type PortService struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx, id
func (_m *PortService) Find(ctx context.Context, id string) (*port.Port, error) {
	ret := _m.Called(ctx, id)

	var r0 *port.Port
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*port.Port, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *port.Port); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*port.Port)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, _a1
func (_m *PortService) Store(ctx context.Context, _a1 *port.Port) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *port.Port) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPortService creates a new instance of PortService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPortService(t interface {
	mock.TestingT
	Cleanup(func())
}) *PortService {
	mock := &PortService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
