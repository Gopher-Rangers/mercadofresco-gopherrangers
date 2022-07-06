// Code generated by mockery v2.13.1. DO NOT EDIT.

package mock_repository

import (
	domain "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/domain"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateWarehouse provides a mock function with given fields: code, address, tel, minCap, minTemp
func (_m *Repository) CreateWarehouse(code string, address string, tel string, minCap int, minTemp int) (domain.Warehouse, error) {
	ret := _m.Called(code, address, tel, minCap, minTemp)

	var r0 domain.Warehouse
	if rf, ok := ret.Get(0).(func(string, string, string, int, int) domain.Warehouse); ok {
		r0 = rf(code, address, tel, minCap, minTemp)
	} else {
		r0 = ret.Get(0).(domain.Warehouse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, int, int) error); ok {
		r1 = rf(code, address, tel, minCap, minTemp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWarehouse provides a mock function with given fields: id
func (_m *Repository) DeleteWarehouse(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByWarehouseCode provides a mock function with given fields: code
func (_m *Repository) FindByWarehouseCode(code string) (domain.Warehouse, error) {
	ret := _m.Called(code)

	var r0 domain.Warehouse
	if rf, ok := ret.Get(0).(func(string) domain.Warehouse); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Get(0).(domain.Warehouse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() []domain.Warehouse {
	ret := _m.Called()

	var r0 []domain.Warehouse
	if rf, ok := ret.Get(0).(func() []domain.Warehouse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Warehouse)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *Repository) GetByID(id int) (domain.Warehouse, error) {
	ret := _m.Called(id)

	var r0 domain.Warehouse
	if rf, ok := ret.Get(0).(func(int) domain.Warehouse); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Warehouse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatedWarehouseID provides a mock function with given fields: id, code
func (_m *Repository) UpdatedWarehouseID(id int, code string) (domain.Warehouse, error) {
	ret := _m.Called(id, code)

	var r0 domain.Warehouse
	if rf, ok := ret.Get(0).(func(int, string) domain.Warehouse); ok {
		r0 = rf(id, code)
	} else {
		r0 = ret.Get(0).(domain.Warehouse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(id, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
