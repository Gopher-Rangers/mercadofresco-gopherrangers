package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/handlers"
	productrecord "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_record"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_record/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	URL_PRODUCT_RECORD_POST = "/api/v1/productRecords/"
	URL_PRODUCT_RECORD_GET  = "/api/v1/reportRecords/"
)

type responseProductRecordGetArray struct {
	Code  int
	Data  []productrecord.ProductRecordGet
	Error string
}

type responseProductRecordGet struct {
	Code  int
	Data  productrecord.ProductRecordGet
	Error string
}

type responseProductRecord struct {
	Code  int
	Data  productrecord.ProductRecord
	Error string
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

func createProductRecordRequestTest(method string, url string, body string) (
	*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestProductRecordStore(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProductRecord(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCT_RECORD_POST)

		ps := createProductRecordArray()
		expected := `{"id": 1,
			"last_update_date": "2025-07-06 13:30:00",
			"purchase_price": 10.10,
			"sale_price": 100.10,
			"product_id": 1}`
		req, rr := createProductRecordRequestTest(
			http.MethodPost,
			URL_PRODUCT_RECORD_POST,
			expected)
		mockService.On("Store",
			mock.AnythingOfType("*context.emptyCtx"),
			ps[0]).Return(ps[0], nil)
		productRouterGroup.POST("/", handlerProduct.Store())
		server.ServeHTTP(rr, req)

		resp := responseProductRecord{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusCreated, rr.Code, resp.Code)
		assert.Equal(t, resp.Error, "")
		assert.Equal(t, ps[0], resp.Data)
	})
	t.Run("create_conflict", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProductRecord(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCT_RECORD_POST)

		ps := createProductRecordArray()
		expected := `{"id": 1,
			"last_update_date": "2025-07-06 13:30:00",
			"purchase_price": 10.10,
			"sale_price": 100.10,
			"product_id": 1}`
		req, rr := createProductRequestTest(
			http.MethodPost,
			URL_PRODUCT_RECORD_POST,
			expected)
		mockService.On("Store", mock.AnythingOfType("*context.emptyCtx"),
			ps[0]).Return(productrecord.ProductRecord{},
			fmt.Errorf(productrecord.ERROR_INEXISTENT_PRODUCT))
		productRouterGroup.POST("/", handlerProduct.Store())
		server.ServeHTTP(rr, req)

		resp := responseProductRecord{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusConflict, rr.Code, resp.Code)
		assert.Equal(t, productrecord.ProductRecord{}, resp.Data)
		assert.Equal(t, resp.Error, productrecord.ERROR_INEXISTENT_PRODUCT)
	})
	t.Run("create_fail_to_save", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProductRecord(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCT_RECORD_POST)

		ps := createProductRecordArray()
		expected := `{"id": 1,
			"last_update_date": "2025-07-06 13:30:00",
			"purchase_price": 10.10,
			"sale_price": 100.10,
			"product_id": 1}`
		req, rr := createProductRequestTest(
			http.MethodPost,
			URL_PRODUCT_RECORD_POST,
			expected)
		mockService.On("Store", mock.AnythingOfType("*context.emptyCtx"),
			ps[0]).Return(productrecord.ProductRecord{},
			fmt.Errorf("fail to save"))
		productRouterGroup.POST("/", handlerProduct.Store())
		server.ServeHTTP(rr, req)

		resp := responseProductRecord{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusNotFound, rr.Code, resp.Code)
		assert.Equal(t, productrecord.ProductRecord{}, resp.Data)
		assert.Equal(t, resp.Error, "fail to save")
	})
	t.Run("create_error_bind", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProductRecord(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCT_RECORD_POST)

		expected := `{"id": 1,
			"last_update_date": "2025-07-06 13:30:00",
			"purchase_price": "string",
			"sale_price": 100.10,
			"product_id": 1}`
		req, rr := createProductRequestTest(
			http.MethodPost,
			URL_PRODUCT_RECORD_POST,
			expected)
		bindError := "json: cannot unmarshal string into Go struct field ProductRecord.purchase_price of type float64"

		productRouterGroup.POST("/", handlerProduct.Store())
		server.ServeHTTP(rr, req)

		resp := responseProductRecord{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusBadRequest, rr.Code, resp.Code)
		assert.Equal(t, productrecord.ProductRecord{}, resp.Data)
		assert.Equal(t, resp.Error, bindError)
	})
	t.Run("create_wrong_body", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProductRecord(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCT_RECORD_POST)

		expected := `{"id": 0,
			"last_update_date": "",
			"purchase_price": 0,
			"sale_price": 0,
			"product_id": 0}`
		req, rr := createProductRequestTest(
			http.MethodPost,
			URL_PRODUCT_RECORD_POST,
			expected)
		productRouterGroup.POST("/", handlerProduct.Store())
		server.ServeHTTP(rr, req)

		resp := responseProductRecord{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code, resp.Code)
		assert.Equal(t, productrecord.ProductRecord{}, resp.Data)
		assert.Equal(t, resp.Error, handler.ERROR_LAST_UPDATE_DATE)
	})
}

func TestProductRecordGet(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProductRecord(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCT_RECORD_GET)

		ps := createProductRecordGetArray()
		req, rr := createProductRequestTest(http.MethodGet,
			URL_PRODUCT_RECORD_GET, "")

		mockService.On("GetAll",
			mock.AnythingOfType("*context.emptyCtx")).Return(ps, nil)
		productRouterGroup.GET("/", handlerProduct.Get())
		server.ServeHTTP(rr, req)

		resp := responseProductRecordGetArray{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, rr.Code, resp.Code)
		assert.Equal(t, ps, resp.Data)
		assert.Equal(t, resp.Error, "")
	})
	t.Run("find_by_id_id_non_number", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProductRecord(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCT_RECORD_GET)

		ps := productrecord.ProductRecordGet{}
		req, rr := createProductRequestTest(http.MethodGet,
			URL_PRODUCT_RECORD_GET+"?id=A", "")

		productRouterGroup.GET("", handlerProduct.Get())
		server.ServeHTTP(rr, req)

		resp := responseProductRecordGet{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, ps, resp.Data)
		assert.Equal(t, resp.Error, handler.ERROR_ID)
	})
	t.Run("find_by_id_existent", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProductRecord(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCT_RECORD_GET)

		ps := createProductRecordGetArray()
		req, rr := createProductRequestTest(http.MethodGet,
			URL_PRODUCT_RECORD_GET+"?id=1", "")

		mockService.On("GetById",
			mock.AnythingOfType("*context.emptyCtx"),
			1).Return(ps[0], nil)
		productRouterGroup.GET("", handlerProduct.Get())
		server.ServeHTTP(rr, req)

		resp := responseProductRecordGet{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, rr.Code, resp.Code)
		assert.Equal(t, ps[0], resp.Data)
		assert.Equal(t, resp.Error, "")
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProductRecord(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCT_RECORD_GET)

		ps := productrecord.ProductRecordGet{}
		req, rr := createProductRequestTest(http.MethodGet,
			URL_PRODUCT_RECORD_GET+"?id=3", "")

		mockService.On("GetById",
			mock.AnythingOfType("*context.emptyCtx"),
			3).Return(ps, fmt.Errorf("produto 3 não encontrado"))
		productRouterGroup.GET("", handlerProduct.Get())
		server.ServeHTTP(rr, req)

		resp := responseProductRecordGet{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusNotFound, rr.Code, resp.Code)
		assert.Equal(t, ps, resp.Data)
		assert.Equal(t, resp.Error, "produto 3 não encontrado")
	})
}

func TestNewRequestProductRecord(t *testing.T) {
	t.Run("fake_test_new_request_product_for_swag", func(t *testing.T) {
		handlerProduct := handler.NewRequestProductRecord()

		assert.Equal(t, handlerProduct, handler.NewRequestProductRecord())
	})
}
