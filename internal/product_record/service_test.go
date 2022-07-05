package productrecord_test
/*
import (
	"context"
	"testing"

	productrecord "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_record"
	products "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_record/mocks"
	mockProducts "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product/mocks"
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
	ps = append(ps, prod1, prod2)
	return ps
}

func createProductRecordArray() []productrecord.ProductRecord {
	var prs []productrecord.ProductRecord
	prod1 := productrecord.ProductRecord{
		ID:             1,
		LastUpdateDate: "2090-28-03",
		PurchasePrice:  10.10,
		SalePrice:      100.10,
		ProductId:      1,
	}
	prod2 := productrecord.ProductRecord{
		ID:             2,
		LastUpdateDate: "2086-19-04",
		PurchasePrice:  20.20,
		SalePrice:      200.20,
		ProductId:      2,
	}
	prs = append(prs, prod1, prod2)
	return prs
}
/*
func TestStore(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		mockProductRepository := mockProducts.NewRepository(t)
		mockProductService := products.NewService(mockProductRepository)
		ps := createProductArray()
		mockRepository := mocks.NewRepository(t)
		service := productrecord.NewService(mockRepository, mockProductService)
		expected := createProductRecordArray()[1]
		mockProductRepository.On("GetById", 2).Return(ps[1], nil)
		mockRepository.On("checkIfProductExists", expected).Return(true)
		mockRepository.On("checkDatetime", expected.LastUpdateDate).Return(true)
		prod, err := service.Store(context.Background(), expected)
		assert.Nil(t, err)
		assert.Equal(t, expected, prod)
	})
}
*/