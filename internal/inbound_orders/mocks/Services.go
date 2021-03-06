// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	inboundorders "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/inbound_orders"
	mock "github.com/stretchr/testify/mock"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// Create provides a mock function with given fields: orderDate, orderNumber, employeeId, productBatchId, warehouseId
func (_m *Services) Create(orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (inboundorders.InboundOrder, error) {
	ret := _m.Called(orderDate, orderNumber, employeeId, productBatchId, warehouseId)

	var r0 inboundorders.InboundOrder
	if rf, ok := ret.Get(0).(func(string, string, int, int, int) inboundorders.InboundOrder); ok {
		r0 = rf(orderDate, orderNumber, employeeId, productBatchId, warehouseId)
	} else {
		r0 = ret.Get(0).(inboundorders.InboundOrder)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int, int, int) error); ok {
		r1 = rf(orderDate, orderNumber, employeeId, productBatchId, warehouseId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCounterByEmployee provides a mock function with given fields: id
func (_m *Services) GetCounterByEmployee(id int) int {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewServices interface {
	mock.TestingT
	Cleanup(func())
}

// NewServices creates a new instance of Services. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServices(t mockConstructorTestingTNewServices) *Services {
	mock := &Services{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
