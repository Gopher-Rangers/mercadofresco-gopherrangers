package products_test

import (
	"context"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality"
	mocksLocality "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality/mocks"
	"testing"

	products "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	mocks "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product/mocks"
	seller "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"
	mockSeller "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller/mocks"

	"github.com/stretchr/testify/assert"
)

func createProductsArray() []products.Product {
	var ps []products.Product
	prod1 := products.Product{
		ID:                             1,
		ProductCode:                    "01",
		Description:                    "leite",
		Width:                          0.1,
		Height:                         0.1,
		Length:                         0.1,
		NetWeight:                      0.1,
		ExpirationRate:                 0.1,
		RecommendedFreezingTemperature: 1.1,
		FreezingRate:                   1.1,
		ProductTypeId:                  01,
		SellerId:                       01,
	}
	prod2 := products.Product{
		ID:                             2,
		ProductCode:                    "02",
		Description:                    "café",
		Width:                          0.2,
		Height:                         0.2,
		Length:                         0.2,
		NetWeight:                      0.2,
		ExpirationRate:                 0.2,
		RecommendedFreezingTemperature: 2.2,
		FreezingRate:                   2.2,
		ProductTypeId:                  02,
		SellerId:                       02,
	}
	ps = append(ps, prod1, prod2)
	return ps
}

func TestStore(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		expected := products.Product{
			ID:                             3,
			ProductCode:                    "03",
			Description:                    "queijo",
			Width:                          0.3,
			Height:                         0.3,
			Length:                         0.3,
			NetWeight:                      0.3,
			ExpirationRate:                 0.3,
			RecommendedFreezingTemperature: 3.3,
			FreezingRate:                   3.3,
			ProductTypeId:                  03,
			SellerId:                       03,
		}
		mockRepository.On("CheckProductType", context.Background(),
			expected.ProductTypeId).Return(true)
		mockSellerRepository.On("GetOne", context.Background(), 3).Return(
			seller.Seller{}, nil)
		mockRepository.On("CheckProductCode", context.Background(),
			expected.ID, expected.ProductCode).Return(true)
		mockRepository.On("Store", context.Background(), expected).Return(
			expected, nil)
		prod, err := service.Store(context.Background(), expected)
		assert.Nil(t, err)
		assert.Equal(t, expected, prod)
	})
	t.Run("create_inexistent_seller", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		expected := products.Product{
			ID:                             3,
			ProductCode:                    "03",
			Description:                    "queijo",
			Width:                          0.3,
			Height:                         0.3,
			Length:                         0.3,
			NetWeight:                      0.3,
			ExpirationRate:                 0.3,
			RecommendedFreezingTemperature: 3.3,
			FreezingRate:                   3.3,
			ProductTypeId:                  03,
			SellerId:                       03,
		}
		mockRepository.On("CheckProductType", context.Background(),
			expected.ProductTypeId).Return(true)
		mockSellerRepository.On("GetOne", context.Background(), 3).Return(
			seller.Seller{}, fmt.Errorf("id does not exist"))
		prod, err := service.Store(context.Background(), expected)
		assert.Equal(t, err, fmt.Errorf(products.ERROR_INEXISTENT_SELLER))
		assert.Equal(t, prod, products.Product{})
	})
	t.Run("create_inexistent_product_type", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		expected := products.Product{
			ID:                             3,
			ProductCode:                    "03",
			Description:                    "queijo",
			Width:                          0.3,
			Height:                         0.3,
			Length:                         0.3,
			NetWeight:                      0.3,
			ExpirationRate:                 0.3,
			RecommendedFreezingTemperature: 3.3,
			FreezingRate:                   3.3,
			ProductTypeId:                  03,
			SellerId:                       03,
		}
		mockRepository.On("CheckProductType", context.Background(),
			expected.ProductTypeId).Return(false)
		prod, err := service.Store(context.Background(), expected)
		assert.Equal(t, err, fmt.Errorf(products.ERROR_INEXISTENT_PRODUCT_TYPE))
		assert.Equal(t, prod, products.Product{})
	})
	t.Run("create_conflict", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		expected := products.Product{
			ID:                             3,
			ProductCode:                    "02",
			Description:                    "queijo",
			Width:                          0.3,
			Height:                         0.3,
			Length:                         0.3,
			NetWeight:                      0.3,
			ExpirationRate:                 0.3,
			RecommendedFreezingTemperature: 3.3,
			FreezingRate:                   3.3,
			ProductTypeId:                  03,
			SellerId:                       03,
		}
		mockRepository.On("CheckProductType", context.Background(),
			expected.ProductTypeId).Return(true)
		mockSellerRepository.On("GetOne", context.Background(), 3).Return(
			seller.Seller{}, nil)
		mockRepository.On("CheckProductCode", context.Background(),
			expected.ID, expected.ProductCode).Return(false)
		prod, err := service.Store(context.Background(), expected)
		fmt.Println(err)
		assert.Equal(t, err, fmt.Errorf("the product code must be unique"))
		assert.Equal(t, products.Product{}, prod)
	})
	t.Run("create_error", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		expected := products.Product{
			ID:                             3,
			ProductCode:                    "03",
			Description:                    "queijo",
			Width:                          0.3,
			Height:                         0.3,
			Length:                         0.3,
			NetWeight:                      0.3,
			ExpirationRate:                 0.3,
			RecommendedFreezingTemperature: 3.3,
			FreezingRate:                   3.3,
			ProductTypeId:                  03,
			SellerId:                       03,
		}
		mockRepository.On("CheckProductType", context.Background(),
			expected.ProductTypeId).Return(true)
		mockSellerRepository.On("GetOne", context.Background(), 3).Return(
			seller.Seller{}, nil)
		mockRepository.On("CheckProductCode", context.Background(),
			expected.ID, expected.ProductCode).Return(true)
		mockRepository.On("Store", context.Background(), expected).Return(
			products.Product{}, fmt.Errorf("fail to save"))
		prod, err := service.Store(context.Background(), expected)
		assert.Equal(t, err, fmt.Errorf("fail to save"))
		assert.Equal(t, prod, products.Product{})
	})
}

func TestGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		ps := createProductsArray()
		mockRepository.On("GetAll", context.Background()).Return(ps, nil)
		prod, err := service.GetAll(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, prod, ps)
	})
}

func TestGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		ps := createProductsArray()
		mockRepository.On("GetById", context.Background(), 1).Return(ps[0], nil)
		prod, err := service.GetById(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, prod, ps[0])
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		e := fmt.Errorf("produto 3 não encontrado")
		mockRepository.On("GetById", context.Background(), 3).Return(
			products.Product{}, e)
		prod, err := service.GetById(context.Background(), 3)
		assert.Equal(t, e, err)
		assert.Equal(t, prod, products.Product{})
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update_existent", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		expected := products.Product{
			ID:                             1,
			ProductCode:                    "01",
			Description:                    "requeijao",
			Width:                          0.1,
			Height:                         0.1,
			Length:                         0.1,
			NetWeight:                      0.1,
			ExpirationRate:                 0.1,
			RecommendedFreezingTemperature: 1.1,
			FreezingRate:                   1.1,
			ProductTypeId:                  01,
			SellerId:                       01,
		}
		mockRepository.On("CheckProductType", context.Background(),
			expected.ProductTypeId).Return(true)
		mockSellerRepository.On("GetOne", context.Background(), 1).Return(
			seller.Seller{}, nil)
		mockRepository.On("CheckProductCode", context.Background(),
			expected.ID, expected.ProductCode).Return(true)
		mockRepository.On("Update", context.Background(), expected, 1).Return(
			expected, nil)
		prod, err := service.Update(context.Background(), expected, 1)
		assert.Nil(t, err)
		assert.Equal(t, prod, expected)
	})
	t.Run("update_inexistent_product_type", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		expected := products.Product{
			ID:                             1,
			ProductCode:                    "01",
			Description:                    "requeijao",
			Width:                          0.1,
			Height:                         0.1,
			Length:                         0.1,
			NetWeight:                      0.1,
			ExpirationRate:                 0.1,
			RecommendedFreezingTemperature: 1.1,
			FreezingRate:                   1.1,
			ProductTypeId:                  01,
			SellerId:                       01,
		}
		mockRepository.On("CheckProductType", context.Background(),
			expected.ProductTypeId).Return(false)
		prod, err := service.Update(context.Background(), expected, 1)
		assert.Equal(t, err, fmt.Errorf(products.ERROR_INEXISTENT_PRODUCT_TYPE))
		assert.Equal(t, prod, products.Product{})
	})
	t.Run("update_inexistent_seller", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		expected := products.Product{
			ID:                             1,
			ProductCode:                    "01",
			Description:                    "requeijao",
			Width:                          0.1,
			Height:                         0.1,
			Length:                         0.1,
			NetWeight:                      0.1,
			ExpirationRate:                 0.1,
			RecommendedFreezingTemperature: 1.1,
			FreezingRate:                   1.1,
			ProductTypeId:                  01,
			SellerId:                       01,
		}
		mockRepository.On("CheckProductType", context.Background(),
			expected.ProductTypeId).Return(true)
		mockSellerRepository.On("GetOne", context.Background(), 1).Return(
			seller.Seller{}, fmt.Errorf("id does not exist"))
		prod, err := service.Update(context.Background(), expected, 1)
		assert.Equal(t, err, fmt.Errorf(products.ERROR_INEXISTENT_SELLER))
		assert.Equal(t, prod, products.Product{})
	})
	t.Run("update_non_existent", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		expected := products.Product{
			ID:                             3,
			ProductCode:                    "03",
			Description:                    "queijo",
			Width:                          0.3,
			Height:                         0.3,
			Length:                         0.3,
			NetWeight:                      0.3,
			ExpirationRate:                 0.3,
			RecommendedFreezingTemperature: 3.3,
			FreezingRate:                   3.3,
			ProductTypeId:                  03,
			SellerId:                       03,
		}
		mockRepository.On("CheckProductType", context.Background(),
			expected.ProductTypeId).Return(true)
		mockSellerRepository.On("GetOne", context.Background(), 3).Return(
			seller.Seller{}, nil)
		e := fmt.Errorf("produto 3 não encontrado")
		mockRepository.On("CheckProductCode", context.Background(),
			expected.ID, expected.ProductCode).Return(true)
		mockRepository.On("Update", context.Background(), expected, 3).Return(
			products.Product{}, e)
		prod, err := service.Update(context.Background(), expected, 3)
		assert.Equal(t, e, err)
		assert.Equal(t, prod, products.Product{})
	})
	t.Run("update_conflict", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		expected := products.Product{
			ID:                             1,
			ProductCode:                    "02",
			Description:                    "requeijao",
			Width:                          0.1,
			Height:                         0.1,
			Length:                         0.1,
			NetWeight:                      0.1,
			ExpirationRate:                 0.1,
			RecommendedFreezingTemperature: 1.1,
			FreezingRate:                   1.1,
			ProductTypeId:                  01,
			SellerId:                       01,
		}
		mockRepository.On("CheckProductType", context.Background(),
			expected.ProductTypeId).Return(true)
		mockSellerRepository.On("GetOne", context.Background(), 1).Return(
			seller.Seller{}, nil)
		mockRepository.On("CheckProductCode", context.Background(),
			expected.ID, expected.ProductCode).Return(false)
		prod, err := service.Update(context.Background(), expected, 1)
		assert.Equal(t, err, fmt.Errorf("the product code must be unique"))
		assert.Equal(t, products.Product{}, prod)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		mockRepository.On("Delete", context.Background(), 1).Return(nil)
		err := service.Delete(context.Background(), 1)
		assert.Nil(t, err)
	})
	t.Run("delete_non_existent", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		mockLocalityRepository := mocksLocality.NewRepository(t)
		localityService := locality.NewService(mockLocalityRepository)
		sellerService := seller.NewService(mockSellerRepository, localityService)
		mockRepository := mocks.NewRepository(t)
		service := products.NewService(mockRepository, sellerService)
		e := fmt.Errorf("produto 3 não encontrado")
		mockRepository.On("Delete", context.Background(), 3).Return(e)
		err := service.Delete(context.Background(), 3)
		assert.Equal(t, e, err)
	})
}
