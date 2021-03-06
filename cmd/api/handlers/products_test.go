package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/handlers"
	products "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	URL_PRODUCTS = "/api/v1/products/"
)

type responseProductArray struct {
	Code  int
	Data  []products.Product
	Error string
}

type responseProduct struct {
	Code  int
	Data  products.Product
	Error string
}

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
		ProductTypeId:                  1,
		SellerId:                       1,
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
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	ps = append(ps, prod1, prod2)
	return ps
}

func createProductRequestTest(method string, url string, body string) (
	*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestProductStore(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		ps := createProductsArray()
		expected := `{"id": 1,
			"product_code": "01",
			"description": "leite",
			"width": 0.1,
			"height": 0.1,
			"length": 0.1,
			"net_weight": 0.1,
			"expiration_rate": 0.1,
			"recommended_freezing_temperature": 1.1,
			"freezing_rate": 1.1,
			"product_type_id": 1,
			"seller_id": 1}`
		req, rr := createProductRequestTest(
			http.MethodPost,
			URL_PRODUCTS,
			expected)
		mockService.On("Store", context.Background(), ps[0]).Return(ps[0], nil)
		productRouterGroup.POST("/", handlerProduct.Store())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusCreated, rr.Code, resp.Code)
		assert.Equal(t, ps[0], resp.Data)
		assert.Equal(t, resp.Error, "")
	})
	t.Run("create_fail_to_save", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		ps := createProductsArray()
		expected := `{"id": 1,
			"product_code": "01",
			"description": "leite",
			"width": 0.1,
			"height": 0.1,
			"length": 0.1,
			"net_weight": 0.1,
			"expiration_rate": 0.1,
			"recommended_freezing_temperature": 1.1,
			"freezing_rate": 1.1,
			"product_type_id": 1,
			"seller_id": 1}`
		req, rr := createProductRequestTest(
			http.MethodPost,
			URL_PRODUCTS,
			expected)
		mockService.On("Store", mock.AnythingOfType("*context.emptyCtx"),
			ps[0]).Return(products.Product{},
			fmt.Errorf("fail to save"))
		productRouterGroup.POST("/", handlerProduct.Store())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusNotFound, rr.Code, resp.Code)
		assert.Equal(t, products.Product{}, resp.Data)
		assert.Equal(t, resp.Error, "fail to save")
	})
	t.Run("create_conflict", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		ps := createProductsArray()
		expected := `{"id": 1,
			"product_code": "01",
			"description": "leite",
			"width": 0.1,
			"height": 0.1,
			"length": 0.1,
			"net_weight": 0.1,
			"expiration_rate": 0.1,
			"recommended_freezing_temperature": 1.1,
			"freezing_rate": 1.1,
			"product_type_id": 1,
			"seller_id": 1}`
		req, rr := createProductRequestTest(
			http.MethodPost,
			URL_PRODUCTS,
			expected)
		mockService.On("Store", context.Background(), ps[0]).Return(
			products.Product{}, fmt.Errorf(products.ERROR_UNIQUE_PRODUCT_CODE))
		productRouterGroup.POST("/", handlerProduct.Store())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusConflict, rr.Code, resp.Code)
		assert.Equal(t, products.Product{}, resp.Data)
		assert.Equal(t, resp.Error, products.ERROR_UNIQUE_PRODUCT_CODE)
	})
	t.Run("create_wrong_body", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		expected := `{"id": 0,
			"product_code": "",
			"description": "",
			"width": 0,
			"height": 0,
			"length": 0,
			"net_weight": 0,
			"expiration_rate": 0,
			"recommended_freezing_temperature": 0,
			"freezing_rate": 0,
			"product_type_id": 0,
			"seller_id": 0}`
		req, rr := createProductRequestTest(
			http.MethodPost,
			URL_PRODUCTS,
			expected)
		productRouterGroup.POST("/", handlerProduct.Store())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code, resp.Code)
		assert.Equal(t, products.Product{}, resp.Data)
		assert.Equal(t, resp.Error, handler.ERROR_PRODUCT_CODE)
	})
	t.Run("create_error_bind", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		expected := `{"id": 1,
			"product_code": "01",
			"description": "leite",
			"width": "string",
			"height": 0.1,
			"length": 0.1,
			"net_weight": 0.1,
			"expiration_rate": 0.1,
			"recommended_freezing_temperature": 1.1,
			"freezing_rate": 1.1,
			"product_type_id": 1,
			"seller_id": 1}`
		req, rr := createProductRequestTest(
			http.MethodPost,
			URL_PRODUCTS,
			expected)
		bindError := "json: cannot unmarshal string into Go struct field Product.width of type float64"
		productRouterGroup.POST("/", handlerProduct.Store())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusBadRequest, rr.Code, resp.Code)
		assert.Equal(t, products.Product{}, resp.Data)
		assert.Equal(t, resp.Error, bindError)
	})
}

func TestProductGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		ps := createProductsArray()
		req, rr := createProductRequestTest(http.MethodGet, URL_PRODUCTS, "")
		mockService.On("GetAll", context.Background()).Return(ps, nil)
		productRouterGroup.GET("/", handlerProduct.GetAll())
		server.ServeHTTP(rr, req)
		resp := responseProductArray{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusOK, rr.Code, resp.Code)
		assert.Equal(t, ps, resp.Data)
		assert.Equal(t, resp.Error, "")
	})
}

func TestProductGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		ps := createProductsArray()
		req, rr := createProductRequestTest(
			http.MethodGet,
			URL_PRODUCTS+"1",
			"")
		mockService.On("GetById", context.Background(), 1).Return(ps[0], nil)
		productRouterGroup.GET("/:id", handlerProduct.GetById())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusOK, rr.Code, resp.Code)
		assert.Equal(t, ps[0], resp.Data)
		assert.Equal(t, resp.Error, "")
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		ps := products.Product{}
		req, rr := createProductRequestTest(
			http.MethodGet,
			URL_PRODUCTS+"3",
			"")
		mockService.On("GetById", context.Background(), 3).Return(
			ps, fmt.Errorf("produto 3 não encontrado"))
		productRouterGroup.GET("/:id", handlerProduct.GetById())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusNotFound, rr.Code, resp.Code)
		assert.Equal(t, ps, resp.Data)
		assert.Equal(t, resp.Error, "produto 3 não encontrado")
	})
	t.Run("find_by_id_id_non_number", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		ps := products.Product{}
		req, rr := createProductRequestTest(
			http.MethodGet,
			URL_PRODUCTS+"A",
			"")
		productRouterGroup.GET("/:id", handlerProduct.GetById())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, ps, resp.Data)
		assert.Equal(t, resp.Error, handler.ERROR_ID)
	})
}

func TestProductUpdate(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		ps := products.Product{
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
			ProductTypeId:                  1,
			SellerId:                       1,
		}
		expected := `{"id": 1,
			"product_code": "01",
			"description": "requeijao",
			"width": 0.1,
			"height": 0.1,
			"length": 0.1,
			"net_weight": 0.1,
			"expiration_rate": 0.1,
			"recommended_freezing_temperature": 1.1,
			"freezing_rate": 1.1,
			"product_type_id": 1,
			"seller_id": 1}`
		req, rr := createProductRequestTest(
			http.MethodPatch, URL_PRODUCTS+"1", expected)
		mockService.On("Update", context.Background(), ps, 1).Return(ps, nil)
		productRouterGroup.PATCH("/:id", handlerProduct.Update())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusOK, rr.Code, resp.Code)
		assert.Equal(t, ps, resp.Data)
		assert.Equal(t, resp.Error, "")
	})
	t.Run("update_fail_to_save", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		ps := products.Product{
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
			ProductTypeId:                  1,
			SellerId:                       1,
		}
		expected := `{"id": 1,
			"product_code": "01",
			"description": "requeijao",
			"width": 0.1,
			"height": 0.1,
			"length": 0.1,
			"net_weight": 0.1,
			"expiration_rate": 0.1,
			"recommended_freezing_temperature": 1.1,
			"freezing_rate": 1.1,
			"product_type_id": 1,
			"seller_id": 1}`
		req, rr := createProductRequestTest(
			http.MethodPatch, URL_PRODUCTS+"1", expected)
		mockService.On("Update", context.Background(), ps, 1).Return(
			products.Product{}, fmt.Errorf("fail to save"))
		productRouterGroup.PATCH("/:id", handlerProduct.Update())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusNotFound, rr.Code, resp.Code)
		assert.Equal(t, products.Product{}, resp.Data)
		assert.Equal(t, resp.Error, "fail to save")
	})
	t.Run("update_conflict", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		ps := products.Product{
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
			ProductTypeId:                  1,
			SellerId:                       1,
		}
		expected := `{"id": 1,
			"product_code": "01",
			"description": "requeijao",
			"width": 0.1,
			"height": 0.1,
			"length": 0.1,
			"net_weight": 0.1,
			"expiration_rate": 0.1,
			"recommended_freezing_temperature": 1.1,
			"freezing_rate": 1.1,
			"product_type_id": 1,
			"seller_id": 1}`
		req, rr := createProductRequestTest(
			http.MethodPatch, URL_PRODUCTS+"1", expected)
		mockService.On("Update", context.Background(), ps, 1).Return(
			products.Product{}, fmt.Errorf(products.ERROR_UNIQUE_PRODUCT_CODE))
		productRouterGroup.PATCH("/:id", handlerProduct.Update())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusConflict, rr.Code, resp.Code)
		assert.Equal(t, products.Product{}, resp.Data)
		assert.Equal(t, resp.Error, products.ERROR_UNIQUE_PRODUCT_CODE)
	})
	t.Run("update_error_bind", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		expected := `{"id": 1,
			"product_code": "01",
			"description": "requeijao",
			"width": "string",
			"height": 0.1,
			"length": 0.1,
			"net_weight": 0.1,
			"expiration_rate": 0.1,
			"recommended_freezing_temperature": 1.1,
			"freezing_rate": 1.1,
			"product_type_id": 1,
			"seller_id": 1}`
		req, rr := createProductRequestTest(
			http.MethodPatch,
			URL_PRODUCTS+"1",
			expected)
		bindError := "json: cannot unmarshal string into Go struct field Product.width of type float64"
		productRouterGroup.PATCH("/:id", handlerProduct.Update())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusBadRequest, rr.Code, resp.Code)
		assert.Equal(t, resp.Data, products.Product{})
		assert.Equal(t, resp.Error, bindError)
	})
	t.Run("update_id_non_number", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		expected := `{"id": 1,
			"product_code": "01",
			"description": "requeijao",
			"width": 0.1,
			"height": 0.1,
			"length": 0.1,
			"net_weight": 0.1,
			"expiration_rate": 0.1,
			"recommended_freezing_temperature": 1.1,
			"freezing_rate": 1.1,
			"product_type_id": 1,
			"seller_id": 1}`
		req, rr := createProductRequestTest(http.MethodPatch, URL_PRODUCTS+"non_number", expected)
		productRouterGroup.PATCH("/:id", handlerProduct.Update())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusBadRequest, rr.Code, resp.Code)
		assert.Equal(t, products.Product{}, resp.Data)
		assert.Equal(t, resp.Error, handler.ERROR_ID)
	})
	t.Run("update_wrong_body", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		expected := `{"id": 0,
			"product_code": "",
			"description": "",
			"width": 0,
			"height": 0,
			"length": 0,
			"net_weight": 0,
			"expiration_rate": 0,
			"recommended_freezing_temperature": 0,
			"freezing_rate": 0,
			"product_type_id": 0,
			"seller_id": 0}`
		req, rr := createProductRequestTest(http.MethodPatch, URL_PRODUCTS+"1", expected)
		productRouterGroup.PATCH("/:id", handlerProduct.Update())
		server.ServeHTTP(rr, req)
		resp := responseProduct{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code, resp.Code)
		assert.Equal(t, products.Product{}, resp.Data)
		assert.Equal(t, resp.Error, handler.ERROR_PRODUCT_CODE)
	})
}

func TestProductDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		req, rr := createProductRequestTest(
			http.MethodDelete,
			URL_PRODUCTS+"1",
			"")
		mockService.On("Delete", context.Background(), 1).Return(nil)
		productRouterGroup.DELETE("/:id", handlerProduct.Delete())
		server.ServeHTTP(rr, req)
		resp := responseProductArray{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusNoContent, rr.Code, resp.Code)
		assert.Equal(t, resp.Data, []products.Product([]products.Product(nil)))
		assert.Equal(t, resp.Error, "")
	})
	t.Run("delete_non_existent", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		req, rr := createProductRequestTest(
			http.MethodDelete,
			URL_PRODUCTS+"1",
			"")
		mockService.On("Delete", context.Background(), 1).Return(fmt.Errorf("produto 1 não encontrado"))
		productRouterGroup.DELETE("/:id", handlerProduct.Delete())
		server.ServeHTTP(rr, req)
		resp := responseProductArray{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, resp.Data, []products.Product([]products.Product(nil)))
		assert.Equal(t, resp.Error, "produto 1 não encontrado")
	})
	t.Run("delete_id_non_number", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProduct(mockService)
		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCTS)
		req, rr := createProductRequestTest(
			http.MethodDelete,
			URL_PRODUCTS+"non_number",
			"")
		productRouterGroup.DELETE("/:id", handlerProduct.Delete())
		server.ServeHTTP(rr, req)
		resp := responseProductArray{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, resp.Data, []products.Product([]products.Product(nil)))
		assert.Equal(t, resp.Error, handler.ERROR_ID)
	})
}

func TestNewRequestProduct(t *testing.T) {
	t.Run("fake_test_new_request_product_for_swag", func(t *testing.T) {
		handlerProduct := handler.NewRequestProduct()
		assert.Equal(t, handlerProduct, handler.NewRequestProduct())
	})
}
