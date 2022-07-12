package productrecord_test

import (
	"context"
	"fmt"
	"testing"

	seller "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"
	mockSeller "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller/mocks"
	products "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	mockProducts "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product/mocks"
	productrecord "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_record"
	mocks "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_record/mocks"

	"github.com/stretchr/testify/assert"
)

func createProductArray() []products.Product {
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
	prod3 := products.Product{
		ID:                             3,
		ProductCode:                    "03",
		Description:                    "doce de leite",
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
	ps = append(ps, prod1, prod2, prod3)
	return ps
}

func createProductRecordArray() []productrecord.ProductRecord {
	var prs []productrecord.ProductRecord
	prod1 := productrecord.ProductRecord{
		ID:             1,
		LastUpdateDate: "2025-07-06 13:30:00",
		PurchasePrice:  10.10,
		SalePrice:      100.10,
		ProductId:      1,
	}
	prod2 := productrecord.ProductRecord{
		ID:             2,
		LastUpdateDate: "2025-07-06 13:30:00",
		PurchasePrice:  20.20,
		SalePrice:      200.20,
		ProductId:      2,
	}
	prs = append(prs, prod1, prod2)
	return prs
}

func createProductRecordGetArray() []productrecord.ProductRecordGet {
	var prs []productrecord.ProductRecordGet
	prod1 := productrecord.ProductRecordGet{
		ProductId:    1,
		Description:  "leite",
		RecordsCount: 1,
	}
	prod2 := productrecord.ProductRecordGet{
		ProductId:    2,
		Description:  "café",
		RecordsCount: 1,
	}
	prod3 := productrecord.ProductRecordGet{
		ProductId:    3,
		Description:  "doce de leite",
		RecordsCount: 1,
	}
	prs = append(prs, prod1, prod2, prod3)
	return prs
}

func TestStore(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		sellerService := seller.NewService(mockSellerRepository)
		mockProductRepository := mockProducts.NewRepository(t)
		ProductService := products.NewService(
			mockProductRepository, sellerService)
		ps := createProductArray()
		mockRepository := mocks.NewRepository(t)
		service := productrecord.NewService(mockRepository, ProductService)
		expected := createProductRecordArray()[1]
		mockProductRepository.On("GetById", context.Background(), 2).Return(
			ps[1], nil)
		mockRepository.On("Store",
						context.Background(),
						expected).Return(expected, nil)
		prod, err := service.Store(context.Background(), expected)
		assert.Nil(t, err)
		assert.Equal(t, expected, prod)
	})
	t.Run("create_inexistent_product", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		sellerService := seller.NewService(mockSellerRepository)
		mockProductRepository := mockProducts.NewRepository(t)
		ProductService := products.NewService(
			mockProductRepository, sellerService)
		mockRepository := mocks.NewRepository(t)
		service := productrecord.NewService(mockRepository, ProductService)
		expected := createProductRecordArray()[1]
		mockProductRepository.On("GetById", context.Background(), 2).Return(
										products.Product{},
										fmt.Errorf("produt record 2 not found"))
		prod, err := service.Store(context.Background(), expected)
		assert.Equal(t, err, fmt.Errorf(productrecord.ERROR_INEXISTENT_PRODUCT))
		assert.Equal(t, prod, productrecord.ProductRecord{})
	})
	t.Run("create_lower_last_update_time", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		sellerService := seller.NewService(mockSellerRepository)
		mockProductRepository := mockProducts.NewRepository(t)
		ProductService := products.NewService(
			mockProductRepository, sellerService)
		ps := createProductArray()
		mockRepository := mocks.NewRepository(t)
		service := productrecord.NewService(mockRepository, ProductService)
		expected := productrecord.ProductRecord{
			ID:             3,
			LastUpdateDate: "2020-07-06 13:30:00",
			PurchasePrice:  30.30,
			SalePrice:      300.30,
			ProductId:      3,
		}
		mockProductRepository.On("GetById", context.Background(), 3).Return(
			ps[2], nil)
		prod, err := service.Store(context.Background(), expected)
		assert.Equal(t, err, fmt.Errorf(
			productrecord.ERROR_WRONG_LAST_UPDATE_DATE))
		assert.Equal(t, prod, productrecord.ProductRecord{})
	})
	t.Run("create_error_parsing_time", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		sellerService := seller.NewService(mockSellerRepository)
		mockProductRepository := mockProducts.NewRepository(t)
		ProductService := products.NewService(
			mockProductRepository, sellerService)
		ps := createProductArray()
		mockRepository := mocks.NewRepository(t)
		service := productrecord.NewService(mockRepository, ProductService)
		expected := productrecord.ProductRecord {
			ID:             3,
			LastUpdateDate: "errror_parsing_datetime",
			PurchasePrice:  30.30,
			SalePrice:      300.30,
			ProductId:      3,
		}
		mockProductRepository.On("GetById", context.Background(), 3).Return(
			ps[2], nil)
		prod, err := service.Store(context.Background(), expected)
		assert.NotNil(t, err)
		assert.Equal(t, prod, productrecord.ProductRecord{})
	})
	t.Run("create_fail_to_save", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		sellerService := seller.NewService(mockSellerRepository)
		mockProductRepository := mockProducts.NewRepository(t)
		ProductService := products.NewService(
			mockProductRepository, sellerService)
		ps := createProductArray()
		mockRepository := mocks.NewRepository(t)
		service := productrecord.NewService(mockRepository, ProductService)
		expected := createProductRecordArray()[1]
		errFail := fmt.Errorf("fail to save")
		mockProductRepository.On("GetById", context.Background(), 2).Return(
			ps[1], nil)
		mockRepository.On("Store",
						context.Background(),
						expected).Return(productrecord.ProductRecord{}, errFail)
		prod, err := service.Store(context.Background(), expected)
		assert.Equal(t, err, errFail)
		assert.Equal(t, prod, productrecord.ProductRecord{})
	})
}

func TestGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		sellerService := seller.NewService(mockSellerRepository)
		mockProductRepository := mockProducts.NewRepository(t)
		ProductService := products.NewService(mockProductRepository, sellerService)
		mockRepository := mocks.NewRepository(t)
		service := productrecord.NewService(mockRepository, ProductService)
		expectedGet := createProductRecordGetArray()[0]
		mockRepository.On("GetById",
						context.Background(), 1).Return(expectedGet, nil)
		prod, err := service.GetById(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, prod, expectedGet)
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		sellerService := seller.NewService(mockSellerRepository)
		mockProductRepository := mockProducts.NewRepository(t)
		ProductService := products.NewService(mockProductRepository, sellerService)
		mockRepository := mocks.NewRepository(t)
		service := productrecord.NewService(mockRepository, ProductService)
		errNotFound := fmt.Errorf("produt record 1 not found")
		mockRepository.On("GetById",
						context.Background(),
						1).Return(productrecord.ProductRecordGet{}, errNotFound)
		prod, err := service.GetById(context.Background(), 1)
		assert.Equal(t, err, errNotFound)
		assert.Equal(t, prod, productrecord.ProductRecordGet{})
	})
}

func TestGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		mockSellerRepository := mockSeller.NewRepository(t)
		sellerService := seller.NewService(mockSellerRepository)
		mockProductRepository := mockProducts.NewRepository(t)
		ProductService := products.NewService(mockProductRepository, sellerService)
		mockRepository := mocks.NewRepository(t)
		service := productrecord.NewService(mockRepository, ProductService)
		expectedGet := createProductRecordGetArray()
		mockRepository.On("GetAll",
						context.Background()).Return(expectedGet, nil)
		prod, err := service.GetAll(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, prod, expectedGet)
	})
}
