// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	buyer "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *Repository) Create(ctx context.Context, _a1 buyer.Buyer) (buyer.Buyer, error) {
	ret := _m.Called(ctx, _a1)

	var r0 buyer.Buyer
	if rf, ok := ret.Get(0).(func(context.Context, buyer.Buyer) buyer.Buyer); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(buyer.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, buyer.Buyer) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]buyer.Buyer, error) {
	ret := _m.Called(ctx)

	var r0 []buyer.Buyer
	if rf, ok := ret.Get(0).(func(context.Context) []buyer.Buyer); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]buyer.Buyer)
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
func (_m *Repository) GetById(ctx context.Context, id int) (buyer.Buyer, error) {
	ret := _m.Called(ctx, id)

	var r0 buyer.Buyer
	if rf, ok := ret.Get(0).(func(context.Context, int) buyer.Buyer); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(buyer.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *Repository) Update(ctx context.Context, _a1 buyer.Buyer) (buyer.Buyer, error) {
	ret := _m.Called(ctx, _a1)

	var r0 buyer.Buyer
	if rf, ok := ret.Get(0).(func(context.Context, buyer.Buyer) buyer.Buyer); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(buyer.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, buyer.Buyer) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t NewRepositoryT) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}