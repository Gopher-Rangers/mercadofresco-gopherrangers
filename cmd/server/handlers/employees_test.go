package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	employee "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	URL_EMPLOYEES = "/api/v1/employees/"
)

type response struct {
	Code  int
	Data  []employee.Employee
	Error string
}

type responseCreate struct {
	Code  int
	Data  employee.Employee
	Error string
}

func createEmployeesArray() []employee.Employee {
	var emps []employee.Employee
	employee1 := employee.Employee{
		ID:          1,
		CardNumber:  117899,
		FirstName:   "Jose",
		LastName:    "Neves",
		WareHouseID: 456521,
	}
	employee2 := employee.Employee{
		ID:          2,
		CardNumber:  7878447,
		FirstName:   "Antonio",
		LastName:    "Moraes",
		WareHouseID: 11224411,
	}
	emps = append(emps, employee1, employee2)
	return emps
}

func createEmployeeRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestEmployeeCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		mockService := mocks.NewServices(t)
		handlerEmployee := handler.NewEmployee(mockService)

		server := gin.Default()
		employeesRouterGroup := server.Group(URL_EMPLOYEES)

		emps := createEmployeesArray()
		expected := `{"id": 1, "card_number_id": 117899, "first_name": "Jose", "last_name": "Neves", "warehouse_id": 456521}`
		req, rr := createEmployeeRequestTest(http.MethodPost, URL_EMPLOYEES, expected)
		mockService.On("GetAll").Return([]employee.Employee{})
		mockService.On("Create", emps[0].CardNumber, emps[0].FirstName, emps[0].LastName, emps[0].WareHouseID).Return(emps[0], nil)
		employeesRouterGroup.POST("/", handlerEmployee.Create())
		server.ServeHTTP(rr, req)

		resp := responseCreate{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusCreated, rr.Code, resp.Code)
		assert.Equal(t, emps[0], resp.Data)
		assert.Equal(t, resp.Error, "")
	})
	// 	t.Run("create_conflict", func(t *testing.T) {
	// 		mockService := mocks.NewService(t)
	// 		handlerEmployee := handler.NewEmployee(mockService)

	// 		server := gin.Default()
	// 		productRouterGroup := server.Group(URL_EMPLOYEES)

	// 		emps := createEmployeesArray()
	// 		expected := `{"id": 1,
	// 								"card_number": 117899,
	// "first_name": "Jose",
	// "last_name": "Neves",
	// "ware_house_id": 456521,}`
	// 		req, rr := createProductRequestTest(http.MethodPost, URL_PRODUCTS, expected)
	// 		mockService.On("Store", ps[0]).Return(products.Product{}, fmt.Errorf(products.ERROR_UNIQUE_PRODUCT_CODE))
	// 		productRouterGroup.POST("/", handlerProduct.Store())
	// 		server.ServeHTTP(rr, req)

	// 		resp := responseId{}
	// 		json.Unmarshal(rr.Body.Bytes(), &resp)

	// 		assert.Equal(t, http.StatusConflict, rr.Code, resp.Code)
	// 		assert.Equal(t, products.Product{}, resp.Data)
	// 		assert.Equal(t, resp.Error, products.ERROR_UNIQUE_PRODUCT_CODE)
	// 	})
	// 	t.Run("create_wrong_body", func(t *testing.T) {
	// 		mockService := mocks.NewService(t)
	// 		handlerProduct := handler.NewProduct(mockService)

	// 		server := gin.Default()
	// 		productRouterGroup := server.Group(URL_PRODUCTS)

	// 		expected := `{"id": 0,
	// 										"product_code": "",
	// 										"description": "",
	// 										"width": 0,
	// 										"height": 0,
	// 										"length": 0,
	// 										"net_weight": 0,
	// 										"expiration_rate": "",
	// 										"recommended_freezing_temperature": 0,
	// 										"freezing_rate": 0,
	// 										"product_type_id": 0,
	// 										"seller_id": 0}`
	// 		req, rr := createProductRequestTest(http.MethodPost, URL_PRODUCTS, expected)
	// 		productRouterGroup.POST("/", handlerProduct.Store())
	// 		server.ServeHTTP(rr, req)

	// 		resp := responseId{}
	// 		json.Unmarshal(rr.Body.Bytes(), &resp)

	// 		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code, resp.Code)
	// 		assert.Equal(t, products.Product{}, resp.Data)
	// 		assert.Equal(t, resp.Error, handler.ERROR_PRODUCT_CODE)
	// 	})
}

func TestEmployeesGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		mockService := mocks.NewServices(t)
		handlerEmployee := handler.NewEmployee(mockService)

		server := gin.Default()
		employeeRouterGroup := server.Group(URL_EMPLOYEES)

		emps := createEmployeesArray()
		req, rr := createEmployeeRequestTest(http.MethodGet, URL_EMPLOYEES, "")

		mockService.On("GetAll").Return(emps, nil)
		employeeRouterGroup.GET("/", handlerEmployee.GetAll())
		server.ServeHTTP(rr, req)

		resp := response{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, rr.Code, resp.Code)
		assert.Equal(t, emps, resp.Data)
		assert.Equal(t, resp.Error, "")
	})
}

func TestEmployeesDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		mockService := mocks.NewServices(t)
		handlerEmployee := handler.NewEmployee(mockService)

		server := gin.Default()
		employeeRouterGroup := server.Group(URL_EMPLOYEES)

		req, rr := createEmployeeRequestTest(http.MethodDelete, URL_EMPLOYEES+"1", "")

		mockService.On("Delete", 1).Return(nil)
		employeeRouterGroup.DELETE("/:id", handlerEmployee.Delete())
		server.ServeHTTP(rr, req)

		resp := response{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusNoContent, rr.Code, resp.Code)
		assert.Equal(t, resp.Data, []employee.Employee([]employee.Employee(nil)))
		assert.Equal(t, resp.Error, "")
	})
}
