// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	locality "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, zipCode, localityName, provinceName, countryName
func (_m *Service) Create(ctx context.Context, zipCode string, localityName string, provinceName string, countryName string) (locality.Locality, error) {
	ret := _m.Called(ctx, zipCode, localityName, provinceName, countryName)

	var r0 locality.Locality
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) locality.Locality); ok {
		r0 = rf(ctx, zipCode, localityName, provinceName, countryName)
	} else {
		r0 = ret.Get(0).(locality.Locality)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, zipCode, localityName, provinceName, countryName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *Service) GetAll(ctx context.Context) ([]locality.Locality, error) {
	ret := _m.Called(ctx)

	var r0 []locality.Locality
	if rf, ok := ret.Get(0).(func(context.Context) []locality.Locality); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]locality.Locality)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *Service) GetById(ctx context.Context, id int) (locality.Locality, error) {
	ret := _m.Called(ctx, id)

	var r0 locality.Locality
	if rf, ok := ret.Get(0).(func(context.Context, int) locality.Locality); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(locality.Locality)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReportSellers provides a mock function with given fields: ctx, id
func (_m *Service) ReportSellers(ctx context.Context, id int) (locality.ReportSeller, error) {
	ret := _m.Called(ctx, id)

	var r0 locality.ReportSeller
	if rf, ok := ret.Get(0).(func(context.Context, int) locality.ReportSeller); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(locality.ReportSeller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
