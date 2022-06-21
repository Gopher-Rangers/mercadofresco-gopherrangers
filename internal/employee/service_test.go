package employee_test

import (
	"fmt"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := employee.NewService(mockRepository)
		mockRepository.On("Delete", 1).Return(nil)
		err := service.Delete(1)
		assert.Nil(t, err)
	})
	t.Run("delete_non_existent", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := employee.NewService(mockRepository)
		e := fmt.Errorf("Funcionário 80 não existe")
		mockRepository.On("Delete", 1).Return(e)
		err := service.Delete(1)
		assert.Equal(t, e, err)
	})
}

type NewServiceT interface {
	mock.TestingT
	Cleanup(func())
}

func NewService(t NewServiceT) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
