package employee_test

import (
	"fmt"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee/mocks"
	"github.com/stretchr/testify/assert"
)

func createEmployeeArray() []employee.Employee {
	var emps []employee.Employee
	employee1 := employee.Employee{
		ID:          1,
		CardNumber:  117899,
		FirstName:   "Jose",
		LastName:    "Neves",
		WareHouseID: 456521,
	}
	employee2 := employee.Employee{
		ID:          2,
		CardNumber:  7878447,
		FirstName:   "Antonio",
		LastName:    "Moraes",
		WareHouseID: 11224411,
	}
	emps = append(emps, employee1, employee2)
	return emps
}

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
		mockRepository.On("Delete", 80).Return(e)
		err := service.Delete(80)
		assert.Equal(t, e, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := employee.NewService(mockRepository)
		employees := createEmployeeArray()
		mockRepository.On("GetAll").Return(employees, nil)
		employee, err := service.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, employee, employees)

	})
}

func TestGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := employee.NewService(mockRepository)
		employees := createEmployeeArray()
		id := 2
		mockRepository.On("GetById", id).Return(employees[id-1], nil)
		employee, err := service.GetById(id)
		assert.Nil(t, err)
		assert.Equal(t, employee, employees[id-1])
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := employee.NewService(mockRepository)
		expectedError := fmt.Errorf("o funcionário não existe")
		id := 9
		mockRepository.On("GetById", id).Return(employee.Employee{}, expectedError)
		_, err := service.GetById(id)
		assert.Equal(t, expectedError, err)
	})
}
