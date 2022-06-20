package warehouse_test

import (
	"fmt"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/mock/mock_repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeValidDBWarehouse() warehouse.Warehouse {
	return warehouse.Warehouse{
		ID:             1,
		WarehouseCode:  "j753",
		Address:        "Rua das Margaridas",
		Telephone:      "4833334444",
		MinCapacity:    100,
		MinTemperature: 10,
	}
}

func Test_CreateWarehouse(t *testing.T) {
	t.Run("Deve conter os campos necessários para ser criado um Warehouse.", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := warehouse.NewService(mockRepository)

		data := warehouse.Warehouse{
			WarehouseCode:  "j753",
			Address:        "Rua das Margaridas",
			Telephone:      "4833334444",
			MinCapacity:    100,
			MinTemperature: 10,
		}

		expected := makeValidDBWarehouse()

		mockRepository.On("FindByWarehouseCode", mock.AnythingOfType("string")).Return(warehouse.Warehouse{},
			fmt.Errorf("o warehouse com esse `warehouse_code`: %s não foi encontrado", data.WarehouseCode))

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

		w := makeValidDBWarehouse()

		expected := warehouse.Warehouse{}

		mockRepository.On("FindByWarehouseCode", mock.AnythingOfType("string")).Return(w, nil)

		result, err := service.CreateWarehouse(data.WarehouseCode, data.Address, data.Telephone, data.MinCapacity, data.MinTemperature)

		assert.Equal(t, result, expected)
		assert.Equal(t, err, fmt.Errorf("o `warehouse_code` já está em uso"))
		assert.Error(t, err)
	})

	t.Run("Deve retornar um erro caso CreateWarehouse, retorne um error", func(t *testing.T) {
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

		mockRepository.On("FindByWarehouseCode", mock.AnythingOfType("string")).Return(expected, fmt.Errorf("o warehouse com esse `warehouse_code`: %s não foi encontrado", data.WarehouseCode))

		mockRepository.On("IncrementID").Return(1)
		mockRepository.On("CreateWarehouse", 1, data.WarehouseCode, data.Address, data.Telephone, data.MinCapacity, data.MinTemperature).Return(expected, fmt.Errorf("não foi possível ler o arquivo"))

		result, err := service.CreateWarehouse(data.WarehouseCode, data.Address, data.Telephone, data.MinCapacity, data.MinTemperature)

		assert.Equal(t, result, expected)
		assert.Equal(t, err, fmt.Errorf("não foi possível ler o arquivo"))
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

		expected := makeValidDBWarehouse()

		mockRepository.On("GetByID", 1).Return(expected, nil)

		result, err := service.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, result, expected)
		assert.NotEmpty(t, result)
	})
}

func Test_UpdateWarehouseID(t *testing.T) {
	t.Run("Deve atualizar com sucesso o campo `warehouse_code` do Warehouse com o ID informado.", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := warehouse.NewService(mockRepository)

		expected := makeValidDBWarehouse()

		mockRepository.On("FindByWarehouseCode", mock.AnythingOfType("string")).Return(warehouse.Warehouse{}, fmt.Errorf("o warehouse com esse `warehouse_code`: %s não foi encontrado", expected.WarehouseCode))

		mockRepository.On("UpdatedWarehouseID", 1, "j753").Return(expected, nil)

		result, err := service.UpdatedWarehouseID(1, "j753")

		assert.Nil(t, err)
		assert.Equal(t, result, expected)
		assert.NotEmpty(t, result)
	})

	t.Run("Deve retornar um erro se já exister um Warehouse com o mesmo `warehouse_code`.", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := warehouse.NewService(mockRepository)

		expected := makeValidDBWarehouse()

		mockRepository.On("FindByWarehouseCode", mock.AnythingOfType("string")).Return(expected, nil)

		result, err := service.UpdatedWarehouseID(1, "j753")

		assert.NotNil(t, err)
		assert.Equal(t, err, fmt.Errorf("o `warehouse_code` já está em uso"))
		assert.Equal(t, result, warehouse.Warehouse{})
	})
}

func Test_DeleteWarehouse(t *testing.T) {
	t.Run("Deve deletar um Warehouse com sucesso passando um id válido.", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := warehouse.NewService(mockRepository)

		expected := makeValidDBWarehouse()

		mockRepository.On("DeleteWarehouse", expected.ID).Return(nil)

		err := service.DeleteWarehouse(expected.ID)

		assert.Nil(t, err)
		assert.Equal(t, err, nil)
	})
}
