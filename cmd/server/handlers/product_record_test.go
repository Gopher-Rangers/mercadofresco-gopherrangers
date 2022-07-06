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
	productrecord "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_record"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_record/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	URL_PRODUCT_RECORD = "/api/v1/productRecords/"
)

/*type responseProductRecordGetArray struct {
	Code  int
	Data  []productrecord.ProductRecordGet
	Error string
}

type responseProductRecordGet struct {
	Code  int
	Data  productrecord.ProductRecordGet
	Error string
}*/

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

func createProductRecordRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", os.Getenv("TOKEN"))
	return req, httptest.NewRecorder()
}

func createProductRecordRequestTestIvalidToken(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "invalid_token")
	return req, httptest.NewRecorder()
}

func TestProductRecordStore(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProductRecord(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCT_RECORD)

		ps := createProductRecordArray()
		expected := `{"id": 1,
			"last_update_date": "2025-07-06 13:30:00",
			"purchase_price": 10.10,
			"sale_price": 100.10,
			"product_id": 1}`
		req, rr := createProductRecordRequestTest(
			http.MethodPost,
			URL_PRODUCT_RECORD,
			expected)
		mockService.On("Store", mock.AnythingOfType("*context.emptyCtx"), ps[0]).Return(ps[0], nil)
		productRouterGroup.POST("/", handlerProduct.Store())
		server.ServeHTTP(rr, req)

		resp := responseProductRecord{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusCreated, rr.Code, resp.Code)
		assert.Equal(t, resp.Error, "")
		assert.Equal(t, ps[0], resp.Data)
	})
	t.Run("create_invalid_token", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProductRecord(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCT_RECORD)

		expected := `{"id": 1,
			"last_update_date": "2025-07-06 13:30:00",
			"purchase_price": 10.10,
			"sale_price": 100.10,
			"product_id": 1}`
		req, rr := createProductRecordRequestTestIvalidToken(
			http.MethodPost,
			URL_PRODUCT_RECORD,
			expected)
		productRouterGroup.POST("/", handlerProduct.Store())
		server.ServeHTTP(rr, req)

		resp := responseProductRecord{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusUnauthorized, rr.Code, resp.Code)
		assert.Equal(t, productrecord.ProductRecord{}, resp.Data)
		assert.Equal(t, resp.Error, handler.ERROR_TOKEN)
	})
	t.Run("create_conflict", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProductRecord(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCT_RECORD)

		ps := createProductRecordArray()
		expected := `{"id": 1,
			"last_update_date": "2025-07-06 13:30:00",
			"purchase_price": 10.10,
			"sale_price": 100.10,
			"product_id": 1}`
		req, rr := createProductRequestTest(
			http.MethodPost,
			URL_PRODUCT_RECORD,
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
	t.Run("create_wrong_body", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerProduct := handler.NewProductRecord(mockService)

		server := gin.Default()
		productRouterGroup := server.Group(URL_PRODUCT_RECORD)

		expected := `{"id": 0,
			"last_update_date": "",
			"purchase_price": 0,
			"sale_price": 0,
			"product_id": 0}`
		req, rr := createProductRequestTest(
			http.MethodPost,
			URL_PRODUCT_RECORD,
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
