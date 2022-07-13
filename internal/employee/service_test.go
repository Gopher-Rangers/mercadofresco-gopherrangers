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
		e := fmt.Errorf("funcionario nao existe")
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
		employee, _ := service.GetAll()
		assert.Equal(t, employee, employees)

	})
}

func TestGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := employee.NewService(mockRepository)
		employees := createEmployeeArray()
		id := 2

		mockRepository.On("GetAll").Return(employees, nil)
		mockRepository.On("GetById", id).Return(employees[id-1], nil)
		employee, err := service.GetById(id)
		assert.Nil(t, err)
		assert.Equal(t, employee, employees[id-1])
	})
}

func TestCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := employee.NewService(mockRepository)
		employees := createEmployeeArray()
		expected := employee.Employee{
			CardNumber:  98765431,
			FirstName:   "Novo",
			LastName:    "Func",
			WareHouseID: 1174,
		}

		mockRepository.On("GetAll").Return(employees, nil)
		mockRepository.On("Create", expected.CardNumber, expected.FirstName, expected.LastName, expected.WareHouseID).Return(expected, nil)
		employee, err := service.Create(expected.CardNumber, expected.FirstName, expected.LastName, expected.WareHouseID)
		assert.Nil(t, err)
		assert.Equal(t, expected, employee)
	})
	t.Run("create_conflict", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := employee.NewService(mockRepository)
		employees := createEmployeeArray()
		expected := employee.Employee{
			CardNumber:  7878447,
			FirstName:   "Novo",
			LastName:    "Func",
			WareHouseID: 1174,
		}
		mockRepository.On("GetAll").Return(employees, nil)
		_, err := service.Create(expected.CardNumber, expected.FirstName, expected.LastName, expected.WareHouseID)
		assert.Equal(t, err, fmt.Errorf("funcionário com cartão nº: 7878447 já existe no banco de dados"))
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update_existent", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := employee.NewService(mockRepository)
		expected := employee.Employee{
			FirstName:   "Nome",
			LastName:    "Diferente",
			WareHouseID: 76657665445,
		}

		mockRepository.On("GetById", 2).Return(expected, nil)
		mockRepository.On("Update", 2, expected.FirstName, expected.LastName, expected.WareHouseID).Return(expected, nil)
		employee, err := service.Update(expected, 2)
		assert.Nil(t, err)
		assert.Equal(t, expected, employee)
	})
	t.Run("update_non_existent", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := employee.NewService(mockRepository)
		expected := employee.Employee{
			FirstName:   "Nome",
			LastName:    "Diferente",
			WareHouseID: 76657665445,
		}
		e := fmt.Errorf("funcionario nao existe")

		mockRepository.On("GetById", 15).Return(expected, nil)
		mockRepository.On("Update", 15, expected.FirstName, expected.LastName, expected.WareHouseID).Return(expected, e)
		_, err := service.Update(expected, 15)
		assert.Equal(t, e, err)
	})
}
