package products_test

import (
	"fmt"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product/mocks"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	t.Run("create_ok", func (t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository)

		expected := products.Product {
			ProductCode: "01",
			Description: "leite",
			Width: 0.1,
			Height: 0.1,
			Length: 0.1,
			NetWeight: 0.1,
			ExpirationRate: "10/10/2022",
			RecommendedFreezingTemperature: 1.1,
			FreezingRate: 1.1,
			ProductTypeId: 01,
			SellerId: 01,
		}
		var ps []products.Product
		ps = append(ps, expected)
		mockRepository.On("GetAll").Return(ps, nil)
		mockRepository.On("LastID").Return(0, nil)
		mockRepository.On("Store", expected, 1).Return(expected, nil)
		prod, err := service.Store(expected)

		assert.Nil(t, err)
		assert.Equal(t, expected, prod)
	})

	t.Run("create_conflict", func (t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository)

		expected := products.Product {
			ProductCode: "01",
			Description: "leite",
			Width: 0.1,
			Height: 0.1,
			Length: 0.1,
			NetWeight: 0.1,
			ExpirationRate: "10/10/2022",
			RecommendedFreezingTemperature: 1.1,
			FreezingRate: 1.1,
			ProductTypeId: 01,
			SellerId: 01,
		}
		var ps []products.Product
		ps = append(ps, expected)

		mockRepository.On("GetAll").Return(ps, nil)
		mockRepository.On("LastID").Return(0, nil)
		mockRepository.On("Store", expected, 1).Return(products.Product{}, fmt.Errorf("the product code must be unique"))
		prod, err := service.Store(expected)

		assert.Equal(t, err, fmt.Errorf("the product code must be unique"))
		assert.Equal(t, products.Product{}, prod)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("find_all", func (t *testing.T) {

	})
}

func TestGetById(t *testing.T) {
	t.Run("find_by_id_existent", func (t *testing.T) {

	})

	t.Run("find_by_id_non_existent", func (t *testing.T) {

	})
}

func TestUpdate(t *testing.T) {
	t.Run("update_existent", func (t *testing.T) {

	})

	t.Run("update_non_existent", func (t *testing.T) {

	})
}

func TestDelete(t *testing.T) {
	t.Run("delete_ok", func (t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository)

		mockRepository.On("Delete", 1).Return(nil)
		err := service.Delete(1)

		assert.Nil(t, err)
	})

	t.Run("delete_non_existent", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository)

		mockRepository.On("Delete", 1).Return(fmt.Errorf("produto 1 não encontrado"))
		err := service.Delete(1)

		assert.Equal(t, fmt.Errorf("produto 1 não encontrado"), err)
	})
}
