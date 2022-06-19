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
		mockRepository.On("CreateWarehouse", 1, data.WarehouseCode, data.Address, data.Telephone, data.MinCapacity, data.MinTemperature).Return(expected, nil)
		mockRepository.On("IncrementID").Return(1)

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
