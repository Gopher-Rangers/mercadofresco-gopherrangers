package usecases_test

import (
	"fmt"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/usecases"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/usecases/mock/mock_repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeValidDBWarehouse() domain.Warehouse {
	return domain.Warehouse{
		ID:            1,
		WarehouseCode: "j753",
		Address:       "Rua das Margaridas",
		Telephone:     "4833334444",
		LocalityID:    1,
	}
}

func Test_CreateWarehouse(t *testing.T) {
	t.Run("Deve conter os campos necessários para ser criado um Warehouse.", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := usecases.NewService(mockRepository)

		data := domain.Warehouse{
			WarehouseCode: "j753",
			Address:       "Rua das Margaridas",
			Telephone:     "4833334444",
			LocalityID:    1,
		}

		expected := makeValidDBWarehouse()

		mockRepository.On("FindByWarehouseCode", mock.AnythingOfType("string")).Return(domain.Warehouse{},
			fmt.Errorf("o warehouse com esse `warehouse_code`: %s não foi encontrado", data.WarehouseCode))

		mockRepository.On("CreateWarehouse", data.WarehouseCode, data.Address, data.Telephone, data.LocalityID).Return(expected, nil)

		result, err := service.CreateWarehouse(data.WarehouseCode, data.Address, data.Telephone, data.LocalityID)

		assert.Nil(t, err)
		assert.Equal(t, result, expected)

	})

	t.Run("Deve retornar um warehouse vazio se já existir um `warehouse_code`. ", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := usecases.NewService(mockRepository)

		data := domain.Warehouse{
			WarehouseCode: "j753",
			Address:       "Rua das Margaridas",
			Telephone:     "4833334444",
			LocalityID:    1,
		}

		w := makeValidDBWarehouse()

		expected := domain.Warehouse{}

		mockRepository.On("FindByWarehouseCode", mock.AnythingOfType("string")).Return(w, nil)

		result, err := service.CreateWarehouse(data.WarehouseCode, data.Address, data.Telephone, data.LocalityID)

		assert.Equal(t, result, expected)
		assert.Equal(t, err, fmt.Errorf("o `warehouse_code` já está em uso"))
		assert.Error(t, err)
	})

	t.Run("Deve retornar um erro caso CreateWarehouse, retorne um error", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := usecases.NewService(mockRepository)

		data := domain.Warehouse{
			WarehouseCode: "j753",
			Address:       "Rua das Margaridas",
			Telephone:     "4833334444",
			LocalityID:    1,
		}

		expected := domain.Warehouse{}

		mockRepository.On("FindByWarehouseCode", mock.AnythingOfType("string")).Return(expected, fmt.Errorf("o warehouse com esse `warehouse_code`: %s não foi encontrado", data.WarehouseCode))

		mockRepository.On("CreateWarehouse", data.WarehouseCode, data.Address, data.Telephone, data.LocalityID).Return(expected, fmt.Errorf("não foi possível ler o arquivo"))

		result, err := service.CreateWarehouse(data.WarehouseCode, data.Address, data.Telephone, data.LocalityID)

		assert.Equal(t, result, expected)
		assert.Equal(t, err, fmt.Errorf("não foi possível ler o arquivo"))
		assert.Error(t, err)
	})
}

func Test_GetAll(t *testing.T) {
	t.Run("Deve retornar todos os elementos que estão na lista de warehouses", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := usecases.NewService(mockRepository)

		w := makeValidDBWarehouse()
		expected := []domain.Warehouse{w}

		mockRepository.On("GetAll").Return(expected)

		result := service.GetAll()

		assert.Equal(t, result, expected)
		assert.NotEmpty(t, result)
	})

}

func Test_GetById(t *testing.T) {
	t.Run("Deve retornar warehouse vazio e um erro, se um elemento com o id especifíco não existir.", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := usecases.NewService(mockRepository)

		mockRepository.On("GetByID", 1).Return(domain.Warehouse{}, fmt.Errorf("o id: %d não foi encontrado", 1))

		result, err := service.GetByID(1)

		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, result, domain.Warehouse{})
		assert.Empty(t, result)
	})

	t.Run("Deve retornar um Warehouse, com o id solicitado.", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := usecases.NewService(mockRepository)

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
		service := usecases.NewService(mockRepository)

		expected := makeValidDBWarehouse()

		mockRepository.On("FindByWarehouseCode", mock.AnythingOfType("string")).Return(domain.Warehouse{}, fmt.Errorf("o warehouse com esse `warehouse_code`: %s não foi encontrado", expected.WarehouseCode))

		mockRepository.On("UpdatedWarehouseID", 1, "j753").Return(expected, nil)

		result, err := service.UpdatedWarehouseID(1, "j753")

		assert.Nil(t, err)
		assert.Equal(t, result, expected)
		assert.NotEmpty(t, result)
	})

	t.Run("Deve retornar um erro se já exister um Warehouse com o mesmo `warehouse_code`.", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := usecases.NewService(mockRepository)

		expected := makeValidDBWarehouse()

		mockRepository.On("FindByWarehouseCode", mock.AnythingOfType("string")).Return(expected, nil)

		result, err := service.UpdatedWarehouseID(1, "j753")

		assert.NotNil(t, err)
		assert.Equal(t, err, fmt.Errorf("o `warehouse_code` já está em uso"))
		assert.Equal(t, result, domain.Warehouse{})
	})

	t.Run("Deve retornar um erro caso UpdatedWarehouseID, retorne um error", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := usecases.NewService(mockRepository)

		expected := makeValidDBWarehouse()

		mockRepository.On("FindByWarehouseCode", mock.AnythingOfType("string")).Return(domain.Warehouse{}, fmt.Errorf("o warehouse com esse `warehouse_code`: %s não foi encontrado", expected.WarehouseCode))

		mockRepository.On("UpdatedWarehouseID", 1, "j753").Return(domain.Warehouse{}, fmt.Errorf("o id: %d informado não existe", expected.ID))

		result, err := service.UpdatedWarehouseID(1, expected.WarehouseCode)

		assert.NotNil(t, err)
		assert.Equal(t, result, domain.Warehouse{})
		assert.Equal(t, err, fmt.Errorf("o id: %d informado não existe", expected.ID))
		assert.Error(t, err)
	})
}

func Test_DeleteWarehouse(t *testing.T) {
	t.Run("Deve deletar um Warehouse com sucesso passando um id válido.", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := usecases.NewService(mockRepository)

		expected := makeValidDBWarehouse()

		mockRepository.On("DeleteWarehouse", expected.ID).Return(nil)

		err := service.DeleteWarehouse(expected.ID)

		assert.Nil(t, err)
		assert.Equal(t, err, nil)
	})

	t.Run("Deve retornar um erro se não achar um Warehouse com o id passado.", func(t *testing.T) {
		mockRepository := mock_repository.NewRepository(t)
		service := usecases.NewService(mockRepository)

		expected := makeValidDBWarehouse()

		mockRepository.On("DeleteWarehouse", expected.ID).Return(fmt.Errorf("não foi achado warehouse com esse id: %d", expected.ID))

		err := service.DeleteWarehouse(expected.ID)

		assert.NotNil(t, err)
		assert.Equal(t, err, fmt.Errorf("não foi achado warehouse com esse id: %d", expected.ID))
	})
}
