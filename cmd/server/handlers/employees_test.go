package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type responseOne struct {
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

		resp := responseOne{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusCreated, rr.Code, resp.Code)
		assert.Equal(t, emps[0], resp.Data)
		assert.Equal(t, resp.Error, "")
	})

	t.Run("create_conflict", func(t *testing.T) {
		mockService := mocks.NewServices(t)
		handlerEmployee := handler.NewEmployee(mockService)

		server := gin.Default()
		employeesRouterGroup := server.Group(URL_EMPLOYEES)

		emps := createEmployeesArray()
		expected := `{"id": 1, "card_number_id": 117899, "first_name": "Jose", "last_name": "Neves", "warehouse_id": 456521}`
		req, rr := createEmployeeRequestTest(http.MethodPost, URL_EMPLOYEES, expected)
		mockService.On("GetAll").Return([]employee.Employee{})
		mockService.On("Create", emps[0].CardNumber, emps[0].FirstName, emps[0].LastName, emps[0].WareHouseID).Return(employee.Employee{}, fmt.Errorf("funcionário com cartão nº: 117899 já existe no banco de dados"))
		employeesRouterGroup.POST("/", handlerEmployee.Create())
		server.ServeHTTP(rr, req)

		resp := responseOne{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusConflict, rr.Code, resp.Code)
		assert.Equal(t, employee.Employee{}, resp.Data)
		assert.Equal(t, resp.Error, "funcionário com cartão nº: 117899 já existe no banco de dados")
	})
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
	t.Run("delete_non_existent", func(t *testing.T) {
		mockService := mocks.NewServices(t)
		handlerEmployee := handler.NewEmployee(mockService)

		server := gin.Default()
		employeeRouterGroup := server.Group(URL_EMPLOYEES)

		req, rr := createEmployeeRequestTest(http.MethodDelete, URL_EMPLOYEES+"8765", "")

		mockService.On("Delete", 8765).Return(fmt.Errorf(""))
		employeeRouterGroup.DELETE("/:id", handlerEmployee.Delete())
		server.ServeHTTP(rr, req)

		resp := response{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusNotFound, rr.Code, resp.Code)
		assert.Equal(t, resp.Data, []employee.Employee([]employee.Employee(nil)))
		assert.Equal(t, resp.Error, "")
	})
}

func TestEmployeesGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		mockService := mocks.NewServices(t)
		handlerEmployee := handler.NewEmployee(mockService)

		server := gin.Default()
		employeeRouterGroup := server.Group(URL_EMPLOYEES)

		emps := createEmployeesArray()
		req, rr := createEmployeeRequestTest(http.MethodGet, URL_EMPLOYEES+"1", "")

		mockService.On("GetById", 1).Return(emps[0], nil)
		employeeRouterGroup.GET("/:id", handlerEmployee.GetById())
		server.ServeHTTP(rr, req)

		resp := responseOne{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, rr.Code, resp.Code)
		assert.Equal(t, emps[0], resp.Data)
		assert.Equal(t, resp.Error, "")
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		mockService := mocks.NewServices(t)
		handlerEmployee := handler.NewEmployee(mockService)

		server := gin.Default()
		employeeRouterGroup := server.Group(URL_EMPLOYEES)

		emps := employee.Employee{}
		req, rr := createEmployeeRequestTest(http.MethodGet, URL_EMPLOYEES+"8765", "")

		mockService.On("GetById", 8765).Return(emps, fmt.Errorf(""))
		employeeRouterGroup.GET("/:id", handlerEmployee.GetById())
		server.ServeHTTP(rr, req)

		resp := responseOne{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusNotFound, rr.Code, resp.Code)
		assert.Equal(t, emps, resp.Data)
		assert.Equal(t, resp.Error, "")
	})
}

func TestEmployeesUpdate(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		mockService := mocks.NewServices(t)
		handlerEmployee := handler.NewEmployee(mockService)

		server := gin.Default()
		employeeRouterGroup := server.Group(URL_EMPLOYEES)

		expected := `{"id": 1, "card_number_id": 111133, "first_name": "Novo", "last_name": "Nome", "warehouse_id": 7711}`
		req, rr := createEmployeeRequestTest(http.MethodPatch, URL_EMPLOYEES+"1", expected)

		newEmployee1 := employee.Employee{
			ID:          1,
			CardNumber:  111133,
			FirstName:   "Novo",
			LastName:    "Nome",
			WareHouseID: 7711,
		}

		mockService.On("Update", newEmployee1, 1).Return(newEmployee1, nil)
		employeeRouterGroup.PATCH("/:id", handlerEmployee.Update())
		server.ServeHTTP(rr, req)

		resp := responseOne{}
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, rr.Code, resp.Code)
		assert.Equal(t, newEmployee1, resp.Data)

	})
}
