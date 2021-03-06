package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/handlers"
	employee "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"

	// inboundorders "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/inbound_orders"

	mockEmp "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee/mocks"
	mockIo "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/inbound_orders/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	URL_EMPLOYEES = "/api/v1/employees/"
)

type responseEmployeeArray struct {
	Code  int
	Data  []employee.Employee
	Error string
}

type responseEmployee struct {
	Code  int
	Data  employee.Employee
	Error string
}

func createEmployeesArray() []employee.Employee {
	var emps []employee.Employee
	emp1 := employee.Employee{
		ID:          1,
		CardNumber:  123,
		FirstName:   "Sergio",
		LastName:    "Blabla",
		WareHouseID: 1,
	}
	emp2 := employee.Employee{
		ID:          2,
		CardNumber:  321,
		FirstName:   "Outro",
		LastName:    "Nome",
		WareHouseID: 1,
	}

	emps = append(emps, emp1, emp2)
	return emps
}

func createEmployee() []employee.Employee {
	var emps []employee.Employee
	empToCreate := employee.Employee{
		CardNumber:  123,
		FirstName:   "Sergio",
		LastName:    "Blabla",
		WareHouseID: 1,
	}

	emps = append(emps, empToCreate)
	return emps
}

func createEmployeeRequestTest(method string, url string, body string) (
	*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestEmployeeCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		mockEmpService := mockEmp.NewServices(t)
		mockIOService := mockIo.NewServices(t)
		handlerEmployee := handler.NewEmployee(mockEmpService, mockIOService)
		server := gin.Default()
		employeeRouterGroup := server.Group(URL_EMPLOYEES)
		empToCreate := createEmployee()
		emps := createEmployeesArray()
		expected := `{
			"card_number_id": 123,
			"first_name": "Sergio",
			"last_name": "Blabla",
			"warehouse_id": 1}`
		req, rr := createEmployeeRequestTest(
			http.MethodPost,
			URL_EMPLOYEES,
			expected)

		mockEmpService.On("Create", empToCreate[0].CardNumber, empToCreate[0].FirstName, empToCreate[0].LastName, empToCreate[0].WareHouseID).Return(emps[0], nil)
		employeeRouterGroup.POST("/", handlerEmployee.Create())
		server.ServeHTTP(rr, req)
		resp := responseEmployee{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusCreated, rr.Code, resp.Code)
		assert.Equal(t, emps[0], resp.Data)
		assert.Equal(t, resp.Error, "")
	})
}

func TestEmployeeGetAll(t *testing.T) {
	t.Run("get_all_ok", func(t *testing.T) {
		mockEmpService := mockEmp.NewServices(t)
		mockIOService := mockIo.NewServices(t)
		handlerEmployee := handler.NewEmployee(mockEmpService, mockIOService)
		server := gin.Default()
		employeeRouterGroup := server.Group(URL_EMPLOYEES)
		emps := createEmployeesArray()
		req, rr := createEmployeeRequestTest(http.MethodGet, URL_EMPLOYEES, "")

		mockEmpService.On("GetAll").Return(emps, nil)
		employeeRouterGroup.GET("/", handlerEmployee.GetAll())
		server.ServeHTTP(rr, req)
		resp := responseEmployeeArray{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusOK, rr.Code, resp.Code)
		assert.Equal(t, emps, resp.Data)
		assert.Equal(t, resp.Error, "")
	})
}

func TestEmployeeGetById(t *testing.T) {
	t.Run("get_by_id_existent", func(t *testing.T) {
		mockEmpService := mockEmp.NewServices(t)
		mockIOService := mockIo.NewServices(t)
		handlerEmployee := handler.NewEmployee(mockEmpService, mockIOService)
		server := gin.Default()
		employeeRouterGroup := server.Group(URL_EMPLOYEES)
		emps := createEmployeesArray()
		req, rr := createEmployeeRequestTest(http.MethodGet, URL_EMPLOYEES+"1", "")

		mockEmpService.On("GetById", 1).Return(emps[0], nil)
		employeeRouterGroup.GET("/:id", handlerEmployee.GetById())
		server.ServeHTTP(rr, req)
		resp := responseEmployee{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusOK, rr.Code, resp.Code)
		assert.Equal(t, emps[0], resp.Data)
		assert.Equal(t, resp.Error, "")
	})
}

func TestEmployeeDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		mockEmpService := mockEmp.NewServices(t)
		mockIOService := mockIo.NewServices(t)
		handlerEmployee := handler.NewEmployee(mockEmpService, mockIOService)
		server := gin.Default()
		employeeRouterGroup := server.Group(URL_EMPLOYEES)

		req, rr := createEmployeeRequestTest(http.MethodDelete, URL_EMPLOYEES+"1", "")

		mockEmpService.On("Delete", 1).Return(nil)
		employeeRouterGroup.DELETE("/:id", handlerEmployee.Delete())
		server.ServeHTTP(rr, req)
		resp := responseEmployee{}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Equal(t, http.StatusNoContent, rr.Code, resp.Code)
		assert.Equal(t, resp.Data, employee.Employee{})
		assert.Equal(t, resp.Error, "")
	})
}
