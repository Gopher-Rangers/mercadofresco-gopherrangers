package service_test

import (
	"fmt"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/mock/mock_repository"
	"github.com/stretchr/testify/assert"
)

func Test_CreateWarehouse(t *testing.T) {
	t.Run("Deve conter os campos necessários para ser criado", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := warehouse.NewService(mockRepository)

		data := warehouse.Warehouse{
			WarehouseCode:  "j753",
			Address:        "Rua das Margaridas",
			Telephone:      "4833334444",
			MinCapacity:    100,
			MinTemperature: 10,
		}

		expected := warehouse.Warehouse{
			ID:             1,
			WarehouseCode:  "j753",
			Address:        "Rua das Margaridas",
			Telephone:      "4833334444",
			MinCapacity:    100,
			MinTemperature: 10,
		}

		mockRepository.On("GetAll").Return([]warehouse.Warehouse{})
		mockRepository.On("IncrementID").Return(1)
		mockRepository.On("CreateWarehouse", 1, data.WarehouseCode, data.Address, data.Telephone, data.MinCapacity, data.MinTemperature).Return(expected, nil)

		result, err := service.CreateWarehouse(data.WarehouseCode, data.Address, data.Telephone, data.MinCapacity, data.MinTemperature)

		assert.Nil(t, err)
		assert.Equal(t, result, expected)

	})

	t.Run("Deve retornar um warehouse vazio se já existir um `warehouse_code`. ", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := warehouse.NewService(mockRepository)

		data := warehouse.Warehouse{
			WarehouseCode:  "j753",
			Address:        "Rua das Margaridas",
			Telephone:      "4833334444",
			MinCapacity:    100,
			MinTemperature: 10,
		}

		expected := warehouse.Warehouse{}

		mockRepository.On("GetAll").Return([]warehouse.Warehouse{data})

		result, err := service.CreateWarehouse(data.WarehouseCode, data.Address, data.Telephone, data.MinCapacity, data.MinTemperature)

		assert.Equal(t, result, expected)
		assert.Equal(t, err, fmt.Errorf("o warehouse_code: %s já existe no banco de dados", data.WarehouseCode))
		assert.Error(t, err)
	})
}

func Test_GetAll(t *testing.T) {
	t.Run("Deve retornar todos os elementos que estão na lista de warehouses", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := warehouse.NewService(mockRepository)

		expected := []warehouse.Warehouse{}

		mockRepository.On("GetAll").Return(expected)

		result := service.GetAll()

		assert.Equal(t, result, expected)
		assert.Empty(t, result)
	})

}

func Test_GetById(t *testing.T) {
	t.Run("Deve retornar warehouse vazio e um erro, se um elemento com o id especifíco não existir.", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := warehouse.NewService(mockRepository)

		mockRepository.On("GetByID", 1).Return(warehouse.Warehouse{}, fmt.Errorf("o id: %d não foi encontrado", 1))

		result, err := service.GetByID(1)

		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, result, warehouse.Warehouse{})
		assert.Empty(t, result)
	})

	t.Run("Deve retornar um Warehouse, com o id solicitado.", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := warehouse.NewService(mockRepository)

		expected := warehouse.Warehouse{
			ID:             1,
			WarehouseCode:  "j753",
			Address:        "Rua das Margaridas",
			Telephone:      "4833334444",
			MinCapacity:    100,
			MinTemperature: 10,
		}

		mockRepository.On("GetByID", 1).Return(expected, nil)

		result, err := service.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, result, expected)
		assert.NotEmpty(t, result)
	})
}

func Test_UpdateWarehouseID(t *testing.T) {
	t.Run("Deve atualizar com sucesso o campo `warehouse_code`, se já existir um warehouse com o ID informado.", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := warehouse.NewService(mockRepository)

		w := warehouse.Warehouse{
			ID:             1,
			WarehouseCode:  "j753",
			Address:        "Rua das Margaridas",
			Telephone:      "4833334444",
			MinCapacity:    100,
			MinTemperature: 10,
		}

		mockRepository.On("UpdateWarehouseID", 1, "j753").Return(w)

		result, err := service.UpdatedWarehouseID(1, "j753")

		assert.Nil(t, err)
		assert.Equal(t, result, w)
		assert.NotEmpty(t, result)
	})

	t.Run("", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := warehouse.NewService(mockRepository)

		expected := warehouse.Warehouse{
			ID:             1,
			WarehouseCode:  "j753",
			Address:        "Rua das Margaridas",
			Telephone:      "4833334444",
			MinCapacity:    100,
			MinTemperature: 10,
		}

		mockRepository.On("UpdateWarehouseID", 1).Return(expected, nil)

		result, err := service.UpdatedWarehouseID(1, "j753")

		assert.Nil(t, err)
		assert.Equal(t, result, expected)
		assert.NotEmpty(t, result)
	})
}
