package handlers_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	//"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	URL = "/api/v1/products/"
)

/*func createProductsArray() []products.Product {
	var ps []products.Product
	prod1 := products.Product {
		ID: 1,
		ProductCode: "01",
		Description: "leite",
		Width: 0.1,
		Height: 0.1,
		Length: 0.1,
		NetWeight: 0.1,
		ExpirationRate: "01/01/2022",
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
		ExpirationRate: "02/02/2022",
		RecommendedFreezingTemperature: 2.2,
		FreezingRate: 2.2,
		ProductTypeId: 02,
		SellerId: 02,
	}
	ps = append(ps, prod1, prod2)
	return ps
}*/

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", os.Getenv("TOKEN"))
	return req, httptest.NewRecorder()
}

func createRequestTestIvalidToken(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "invalid_token")
	return req, httptest.NewRecorder()
}


func TestDelete(t *testing.T) {
	t.Run("delete_ok", func (t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL)

		req, rr := createRequestTest(http.MethodDelete, URL + "1", "")

		mockService.On("Delete", 1).Return(nil)
		productRouterGroup.DELETE("/:id", handlerProduct.Delete())
		server.ServeHTTP(rr, req)

		assert.Equal(t, 204, rr.Code)
	})
	t.Run("delete_non_existent", func (t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL)

		req, rr := createRequestTest(http.MethodDelete, URL + "1", "")

		mockService.On("Delete", 1).Return(fmt.Errorf("produto 1 não encontrado"))
		productRouterGroup.DELETE("/:id", handlerProduct.Delete())
		server.ServeHTTP(rr, req)

		assert.Equal(t, 404, rr.Code)
	})
	t.Run("delete_id_non_number", func (t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL)

		req, rr := createRequestTest(http.MethodDelete, URL + "non_number", "")

		productRouterGroup.DELETE("/:id", handlerProduct.Delete())
		server.ServeHTTP(rr, req)

		assert.Equal(t, 400, rr.Code)
	})
	t.Run("delete_invalid_token", func (t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL)

		req, rr := createRequestTestIvalidToken(http.MethodDelete, URL + "1", "")

		productRouterGroup.DELETE("/:id", handlerProduct.Delete())
		server.ServeHTTP(rr, req)

		assert.Equal(t, 401, rr.Code)
	})
}
