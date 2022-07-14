package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/handlers"
	// employee "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"

	inboundorders "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/inbound_orders"
	mock "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/inbound_orders/mocks"

	// mockEmp "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	URL_INBOUND_ORDER = "/api/v1/inboundOrders/"
)

type responseInboundOrder struct {
	Code  int
	Data  []inboundorders.InboundOrder
	Error string
}

func createInboundOrdersArray() []inboundorders.InboundOrder {
	var iosArray []inboundorders.InboundOrder
	io1 := inboundorders.InboundOrder{
		ID:             1,
		OrderDate:      "2022-04-04",
		OrderNumber:    "order#1",
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}
	io2 := inboundorders.InboundOrder{
		ID:             2,
		OrderDate:      "2022-04-04",
		OrderNumber:    "order#2",
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}

	iosArray = append(iosArray, io1, io2)
	return iosArray
}

func createInboundOrdersRequestTest(method string, url string, body string) (
	*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestInboundOrdersCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		mockService := mock.NewServices(t)
		handlerInboundOrder := handler.NewInboundOrder(mockService)
		server := gin.Default()
		inboundOrderRouterGroup := server.Group(URL_INBOUND_ORDER)
		ios := createInboundOrdersArray()
		expected := `{
			"id": 1,
			"order_date": "2022-04-04",
			"order_number": "order#1",
			"employee_id": 1,
			"product_batch_id": 1,
			"warehouse_id": 1}`
		req, rr := createInboundOrdersRequestTest(
			http.MethodPost,
			URL_INBOUND_ORDER,
			expected)

		mockService.On("Create", ios[0].OrderDate, ios[0].OrderNumber, ios[0].EmployeeId, ios[0].ProductBatchId, ios[0].WarehouseId).Return(ios[0], nil)
		inboundOrderRouterGroup.POST("/", handlerInboundOrder.Create())
		server.ServeHTTP(rr, req)
		resp := responseInboundOrder{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusCreated, rr.Code, resp.Code)
		assert.Equal(t, ios[0], resp.Data)
		assert.Equal(t, resp.Error, "")
	})
}
