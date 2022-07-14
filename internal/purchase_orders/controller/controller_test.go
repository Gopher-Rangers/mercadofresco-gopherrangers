package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/controller"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/domain/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	URL = "/api/v1/purchase-orders/"
)

type responseDataOrdersArray struct {
	Code int
	Data []domain.PurchaseOrders
}

type responseData struct {
	Code  int
	Data  domain.PurchaseOrders
	Error string
}

func createBaseData() []domain.PurchaseOrders {
	var purchases []domain.PurchaseOrders
	purchaseOne := domain.PurchaseOrders{
		ID:              1,
		OrderNumber:     "Order1",
		OrderDate:       "2008-11-11",
		TrackingCode:    "1",
		BuyerId:         1,
		ProductRecordId: 1,
		OrderStatusId:   1,
	}
	purchaseTwo := domain.PurchaseOrders{
		ID:              1,
		OrderNumber:     "Order1",
		OrderDate:       "2008-11-11",
		TrackingCode:    "1",
		BuyerId:         1,
		ProductRecordId: 1,
		OrderStatusId:   1,
	}
	purchases = append(purchases, purchaseOne, purchaseTwo)
	return purchases
}

func createRequestTestIvalidToken(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "invalid_token")
	return req, httptest.NewRecorder()
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", os.Getenv("TOKEN"))
	return req, httptest.NewRecorder()
}

func TestCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerPurchase := controller.NewPurchaseOrder(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		buyerData := createBaseData()
		expected := `{"order_number": "order1",
        "order_date": "2008-11-11T13:23:44Z",
        "tracking_code": "1521",
        "buyer_id": 1,
        "product_record_id": 1,
        "order_status_id": 1}`

		req, response := createRequestTest(http.MethodPost, URL, expected)
		mockService.On("Create", context.Background(), domain.PurchaseOrders{OrderNumber: "order1",
			OrderDate:       "2008-11-11T13:23:44Z",
			TrackingCode:    "1521",
			BuyerId:         1,
			ProductRecordId: 1,
			OrderStatusId:   1}).Return(buyerData[0], nil)
		buyerRouterGroup.POST("/", handlerPurchase.Create)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)
		resp.Data.ID = 1

		assert.Equal(t, http.StatusCreated, response.Code, resp.Code)
		assert.Equal(t, buyerData[0], resp.Data)
		assert.Equal(t, resp.Error, "")
	})
	t.Run("create_wrong_body", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerPurchase := controller.NewPurchaseOrder(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		expected := `{"order_number": "order1",
        "order_date": "2008-11-11T13:23:44Z",
        "buyer_id": 1,
        "product_record_id": 1,
        "order_status_id": 1}`

		req, response := createRequestTest(http.MethodPost, URL, expected)

		buyerRouterGroup.POST("/", handlerPurchase.Create)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
		assert.Equal(t, domain.PurchaseOrders{}, resp.Data)
		assert.Equal(t, resp.Error, "invalid body")
	})
	t.Run("create_conflict", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := controller.NewPurchaseOrder(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		expected := `{"order_number": "Order1",
       "order_date": "2008-11-11",
       "tracking_code": "1521",
       "buyer_id": 1,
       "product_record_id": 1,
       "order_status_id": 1}`

		req, response := createRequestTest(http.MethodPost, URL, expected)
		mockService.On("Create", context.Background(), domain.PurchaseOrders{
			OrderNumber:     "Order1",
			OrderDate:       "2008-11-11",
			TrackingCode:    "1521",
			BuyerId:         1,
			ProductRecordId: 1,
			OrderStatusId:   1,
		}).Return(domain.PurchaseOrders{},
			fmt.Errorf("the order number must be unique"))
		buyerRouterGroup.POST("/", buyerHandler.Create)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusConflict, response.Code, resp.Code)
		assert.Equal(t, domain.PurchaseOrders{}, resp.Data)
		assert.Equal(t, resp.Error, "the order number must be unique")
	})
}

func TestGetById(t *testing.T) {
	t.Run("get_by_id", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerPurchase := controller.NewPurchaseOrder(mockService)

		server := gin.Default()
		routerGroup := server.Group(URL)

		data := createBaseData()

		req, response := createRequestTest(http.MethodGet, URL+"1", "")
		mockService.On("GetById", context.Background(), 1).Return(data[0], nil)
		routerGroup.GET("/:id", handlerPurchase.GetPurchaseOrderById)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, response.Code, resp.Code)
		assert.Equal(t, data[0], resp.Data)
		assert.Equal(t, resp.Error, "")
	})
	t.Run("get_by_id_not_found", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerPurchase := controller.NewPurchaseOrder(mockService)

		server := gin.Default()
		routerGroup := server.Group(URL)

		req, response := createRequestTest(http.MethodGet, URL+"1123", "")
		mockService.On("GetById", context.Background(), 1123).Return(domain.PurchaseOrders{}, fmt.Errorf("purchase order with id 1123 not founded"))
		routerGroup.GET("/:id", handlerPurchase.GetPurchaseOrderById)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusNotFound, response.Code, resp.Code)
		assert.Equal(t, domain.PurchaseOrders{}, resp.Data)
		assert.Equal(t, resp.Error, "purchase order with id 1123 not founded")
	})
}
