// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	section "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section"
	mock "github.com/stretchr/testify/mock"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// Create provides a mock function with given fields: secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID
func (_m *Services) Create(secNum int, curTemp int, minTemp int, curCap int, minCap int, maxCap int, wareID int, typeID int) (section.Section, error) {
	ret := _m.Called(secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID)

	var r0 section.Section
	if rf, ok := ret.Get(0).(func(int, int, int, int, int, int, int, int) section.Section); ok {
		r0 = rf(secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID)
	} else {
		r0 = ret.Get(0).(section.Section)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, int, int, int, int, int, int) error); ok {
		r1 = rf(secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSection provides a mock function with given fields: id
func (_m *Services) DeleteSection(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Services) GetAll() ([]section.Section, error) {
	ret := _m.Called()

	var r0 []section.Section
	if rf, ok := ret.Get(0).(func() []section.Section); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]section.Section)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *Services) GetByID(id int) (section.Section, error) {
	ret := _m.Called(id)

	var r0 section.Section
	if rf, ok := ret.Get(0).(func(int) section.Section); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(section.Section)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSecID provides a mock function with given fields: id, secNum
func (_m *Services) UpdateSecID(id int, secNum int) (section.Section, section.CodeError) {
	ret := _m.Called(id, secNum)

	var r0 section.Section
	if rf, ok := ret.Get(0).(func(int, int) section.Section); ok {
		r0 = rf(id, secNum)
	} else {
		r0 = ret.Get(0).(section.Section)
	}

	var r1 section.CodeError
	if rf, ok := ret.Get(1).(func(int, int) section.CodeError); ok {
		r1 = rf(id, secNum)
	} else {
		r1 = ret.Get(1).(section.CodeError)
	}

	return r0, r1
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
