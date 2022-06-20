package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	URL = "/api/v1/products/"
)

type responseArray struct {
	Code int
	Data []products.Product
	Error string
}

type responseId struct {
	Code int
	Data products.Product
	Error string
}

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
}

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

func TestGetAll(t *testing.T) {
	t.Run("find_all", func (t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL)

		ps := createProductsArray()
		req, rr := createRequestTest(http.MethodGet, URL, "")
	
		mockService.On("GetAll").Return(ps, nil)
		productRouterGroup.GET("/", handlerProduct.GetAll())
		server.ServeHTTP(rr, req)

		resp := responseArray{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, rr.Code, resp.Code)
		assert.Equal(t, ps, resp.Data)
		assert.Equal(t, resp.Error, "")
	})
	t.Run("find_all_invalid_token", func (t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL)

		req, rr := createRequestTestIvalidToken(http.MethodGet, URL, "")

		productRouterGroup.GET("/", handlerProduct.GetAll())
		server.ServeHTTP(rr, req)

		resp := responseArray{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusUnauthorized, rr.Code, resp.Code)
		assert.Equal(t, resp.Data, []products.Product([]products.Product(nil)))
		assert.Equal(t, resp.Error, handler.ERROR_TOKEN)
	})
}

func TestGetById(t *testing.T) {
	t.Run("find_by_id_existent", func (t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL)

		ps := createProductsArray()
		req, rr := createRequestTest(http.MethodGet, URL + "1", "")
	
		mockService.On("GetById", 1).Return(ps[0], nil)
		productRouterGroup.GET("/:id", handlerProduct.GetById())
		server.ServeHTTP(rr, req)

		resp := responseId{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, rr.Code, resp.Code)
		assert.Equal(t, ps[0], resp.Data)
		assert.Equal(t, resp.Error, "")
	})
	t.Run("find_by_id_non_existent", func (t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL)

		ps := products.Product{}
		req, rr := createRequestTest(http.MethodGet, URL + "3", "")
	
		mockService.On("GetById", 3).Return(ps, fmt.Errorf("produto 3 não encontrado"))
		productRouterGroup.GET("/:id", handlerProduct.GetById())
		server.ServeHTTP(rr, req)

		resp := responseId{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusNotFound, rr.Code, resp.Code)
		assert.Equal(t, ps, resp.Data)
		assert.Equal(t, resp.Error, "produto 3 não encontrado")
	})
	t.Run("find_by_id_invalid_token", func (t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL)

		ps := products.Product{}
		req, rr := createRequestTestIvalidToken(http.MethodGet, URL + "1", "")

		productRouterGroup.GET("/:id", handlerProduct.GetById())
		server.ServeHTTP(rr, req)

		resp := responseId{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Equal(t, ps, resp.Data)
		assert.Equal(t, resp.Error, handler.ERROR_TOKEN)
	})
	t.Run("find_by_id_id_non_number", func (t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL)

		ps := products.Product{}
		req, rr := createRequestTest(http.MethodGet, URL + "A", "")

		productRouterGroup.GET("/:id", handlerProduct.GetById())
		server.ServeHTTP(rr, req)

		resp := responseId{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, ps, resp.Data)
		assert.Equal(t, resp.Error, handler.ERROR_ID)
	})
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

		resp := responseArray{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusNoContent, rr.Code, resp.Code)
		assert.Equal(t, resp.Data, []products.Product([]products.Product(nil)))
		assert.Equal(t, resp.Error, "")
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

		resp := responseArray{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, resp.Data, []products.Product([]products.Product(nil)))
		assert.Equal(t, resp.Error, "produto 1 não encontrado")
	})
	t.Run("delete_id_non_number", func (t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL)

		req, rr := createRequestTest(http.MethodDelete, URL + "non_number", "")

		productRouterGroup.DELETE("/:id", handlerProduct.Delete())
		server.ServeHTTP(rr, req)

		resp := responseArray{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, resp.Data, []products.Product([]products.Product(nil)))
		assert.Equal(t, resp.Error, handler.ERROR_ID)
	})
	t.Run("delete_invalid_token", func (t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL)

		req, rr := createRequestTestIvalidToken(http.MethodDelete, URL + "1", "")

		productRouterGroup.DELETE("/:id", handlerProduct.Delete())
		server.ServeHTTP(rr, req)

		resp := responseArray{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Equal(t, resp.Data, []products.Product([]products.Product(nil)))
		assert.Equal(t, resp.Error, handler.ERROR_TOKEN)
	})
}
