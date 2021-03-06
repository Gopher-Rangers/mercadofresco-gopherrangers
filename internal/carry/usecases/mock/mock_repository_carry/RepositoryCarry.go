// Code generated by mockery v2.13.1. DO NOT EDIT.

package mock_repository_carry

import (
	domain "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryCarry is an autogenerated mock type for the RepositoryCarry type
type RepositoryCarry struct {
	mock.Mock
}

// CreateCarry provides a mock function with given fields: carry
func (_m *RepositoryCarry) CreateCarry(carry domain.Carry) (domain.Carry, error) {
	ret := _m.Called(carry)

	var r0 domain.Carry
	if rf, ok := ret.Get(0).(func(domain.Carry) domain.Carry); ok {
		r0 = rf(carry)
	} else {
		r0 = ret.Get(0).(domain.Carry)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Carry) error); ok {
		r1 = rf(carry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCarryByCid provides a mock function with given fields: cid
func (_m *RepositoryCarry) GetCarryByCid(cid string) (domain.Carry, error) {
	ret := _m.Called(cid)

	var r0 domain.Carry
	if rf, ok := ret.Get(0).(func(string) domain.Carry); ok {
		r0 = rf(cid)
	} else {
		r0 = ret.Get(0).(domain.Carry)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepositoryCarry interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositoryCarry creates a new instance of RepositoryCarry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositoryCarry(t mockConstructorTestingTNewRepositoryCarry) *RepositoryCarry {
	mock := &RepositoryCarry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
