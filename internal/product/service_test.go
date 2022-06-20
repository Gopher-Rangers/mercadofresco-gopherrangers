package products_test

import (
	"fmt"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product/mocks"

	"github.com/stretchr/testify/assert"
)

func createProductsArray() []products.Product {
	var ps []products.Product
	prod1 := products.Product {
		ID: 1,
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
	prod2 := products.Product {
		ID: 2,
		ProductCode: "02",
		Description: "café",
		Width: 0.2,
		Height: 0.2,
		Length: 0.2,
		NetWeight: 0.2,
		ExpirationRate: "10/10/2022",
		RecommendedFreezingTemperature: 2.2,
		FreezingRate: 2.2,
		ProductTypeId: 02,
		SellerId: 02,
	}
	ps = append(ps, prod1, prod2)
	return ps
}

func TestStore(t *testing.T) {
	t.Run("create_ok", func (t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository)
		ps := createProductsArray()
		expected := ps[0]
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
		ps := createProductsArray()
		expected := products.Product {
			ID: 3,
			ProductCode: "02",
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
		mockRepository.On("GetAll").Return(ps, nil)
		prod, err := service.Store(expected)
		fmt.Println(err)
		assert.Equal(t, err, fmt.Errorf("the product code must be unique"))
		assert.Equal(t, products.Product{}, prod)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("find_all", func (t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository)
		ps := createProductsArray()
		mockRepository.On("GetAll").Return(ps, nil)
		prod, err := service.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, prod, ps)
	})
}

func TestGetById(t *testing.T) {
	t.Run("find_by_id_existent", func (t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository)
		ps := createProductsArray()
		mockRepository.On("GetById", 1).Return(ps[0], nil)
		prod, err := service.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, prod, ps[0])
	})
	t.Run("find_by_id_non_existent", func (t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository)
		mockRepository.On("GetById", 3).Return(products.Product{}, fmt.Errorf("produto 3 não encontrado"))
		prod, err := service.GetById(3)
		assert.Equal(t, fmt.Errorf("produto 3 não encontrado"), err)
		assert.Equal(t, prod, products.Product{})
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update_existent", func (t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository)
		ps := createProductsArray()
		expected := products.Product {
			ID: 1,
			ProductCode: "01",
			Description: "queijo",
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
		mockRepository.On("GetAll").Return(ps, nil)
		mockRepository.On("Update", expected, 1).Return(expected, nil)
		prod, err := service.Update(expected, 1)
		assert.Nil(t, err)
		assert.Equal(t, prod, expected)
	})
	t.Run("update_non_existent", func (t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository)
		ps := createProductsArray()
		expected := products.Product {
			ID: 1,
			ProductCode: "01",
			Description: "queijo",
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
		mockRepository.On("GetAll").Return(ps, nil)
		mockRepository.On("Update", expected, 3).Return( products.Product{}, fmt.Errorf("produto 1 não encontrado"))
		prod, err := service.Update(expected, 3)
		assert.Equal(t, fmt.Errorf("produto 1 não encontrado"), err)
		assert.Equal(t, prod, products.Product{})
	})
	t.Run("update_conflict", func (t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository)
		ps := createProductsArray()
		expected := products.Product {
			ID: 1,
			ProductCode: "02",
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
		mockRepository.On("GetAll").Return(ps, nil)
		prod, err := service.Update(expected, 1)
		assert.Equal(t, err, fmt.Errorf("the product code must be unique"))
		assert.Equal(t, products.Product{}, prod)
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
