package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/mock/mock_service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeValidDBWarehouse() warehouse.Warehouse {
	return warehouse.Warehouse{
		WarehouseCode:  "j753",
		Address:        "Rua das Margaridas",
		Telephone:      "4833334444",
		MinCapacity:    100,
		MinTemperature: 10,
	}
}

const (
	URL = "/api/v1/warehouse"
)

type warehouseResponseBody struct {
	Code  int                 `json:"code"`
	Data  warehouse.Warehouse `json:"data"`
	Error string              `json:"error"`
}

func Test_CreateWarehouse(t *testing.T) {

	service := mock_service.NewService(t)
	controller := handlers.NewWarehouse(service)
	server := gin.Default()

	gin.SetMode(gin.TestMode) // Pra deixar o framework do gin em modo de test

	server.POST(URL, controller.CreateWarehouse)

	t.Run("Deve retornar um status code 201, quando a entrada de dados for bem-sucedida e retornará um warehouse.", func(t *testing.T) {

		data := makeValidDBWarehouse()

		service.On("CreateWarehouse", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(data, nil).Once()

		dataJSON, _ := json.Marshal(data) // Retorna Array de Bytes.

		body := strings.NewReader(string(dataJSON)) // NewReader -> leitor para consumir esses Bytes.

		rr := httptest.NewRecorder() // Monitorar e armazenar requisições http.(Response)

		req, _ := http.NewRequest(http.MethodPost, URL, body)

		server.ServeHTTP(rr, req) // gerenciador de requisições

		respBody := warehouseResponseBody{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, data, respBody.Data)
		assert.Empty(t, respBody.Error)
	})

	t.Run("Deve retornar um status code 422, se o objeto JSON não contiver os campos necessários", func(t *testing.T) {

		invalidBody := bytes.NewBuffer([]byte(`
		{
			"warehouse_code": "valid_code",
			"minimum_capacity": 10,
			"minimum_temperature": 8
		}
		`))

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, URL, invalidBody)

		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), "\"error\":")
	})

	t.Run("Deve retornar um status code 409, se `warehouse_code` já estiver em uso.", func(t *testing.T) {

		service.On("CreateWarehouse", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(warehouse.Warehouse{}, errors.New("o `warehouse_code` já está em uso")).Once()

		data := makeValidDBWarehouse()

		dataJSON, _ := json.Marshal(data)

		body := strings.NewReader(string(dataJSON))

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, URL, body)

		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
		assert.Contains(t, rr.Body.String(), "o `warehouse_code` já está em uso")
	})
}

func Test_GetAll(t *testing.T) {

	service := mock_service.NewService(t)
	controller := handlers.NewWarehouse(service)
	server := gin.Default()

	gin.SetMode(gin.TestMode)

	server.GET(URL, controller.GetAll)

	t.Run("Deve retornar uma lista de warehouses se a solicitação for bem sucedida.", func(t *testing.T) {

		data := makeValidDBWarehouse()

		service.On("GetAll").Return([]warehouse.Warehouse{data})

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, URL, nil)

		server.ServeHTTP(rr, req)

		respBody := warehouseResponseBody{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, data, respBody)
	})
}

func Test_GetByID(t *testing.T) {

	service := mock_service.NewService(t)
	controller := handlers.NewWarehouse(service)
	server := gin.Default()

	server.GET(URL+"/:id", controller.GetByID)

	t.Run("Deve retornar um código 404, quando o Warehouse não existir.", func(t *testing.T) {

		service.On("GetByID", 1).Return(warehouse.Warehouse{}, errors.New("O warehouse não foi encontrado!")).Once()

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, URL+"/1", nil)

		server.ServeHTTP(rr, req)

		respBody := warehouseResponseBody{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Contains(t, "O warehouse não foi encontrado!", respBody.Error)
	})

	t.Run("Deve retornar um código 400, e uma mensagem de erro, quando o id passado não for um número.", func(t *testing.T) {

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, URL+"/casa", nil)

		server.ServeHTTP(rr, req)

		respBody := warehouseResponseBody{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, "O id passado não é um número!", respBody.Error)
	})

	t.Run("Deve retornar um código 200, e um Warehouse quando o id passado for o mesmo do Warehouse no BD.", func(t *testing.T) {

		data := makeValidDBWarehouse()

		service.On("GetByID", 1).Return(data, nil).Once()

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, URL+"/1", nil)

		server.ServeHTTP(rr, req)

		respBody := warehouseResponseBody{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusOK, respBody.Code)
		assert.Equal(t, data, respBody.Data)
	})
}

func Test_UpdatedWarehouseID(t *testing.T) {

	service := mock_service.NewService(t)
	controller := handlers.NewWarehouse(service)
	server := gin.Default()

	gin.SetMode(gin.TestMode)

	server.PATCH(URL+"/:id", controller.UpdatedWarehouseID)

	t.Run("Deve retornar um status code 200, e o Warehouse atualizado, quando a solicitação for bem sucedida.", func(t *testing.T) {

		data := makeValidDBWarehouse()

		service.On("UpdatedWarehouseID", mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(data, nil).Once()

		dataJSON, _ := json.Marshal(data)

		body := strings.NewReader(string(dataJSON))

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPatch, URL+"/1", body)

		server.ServeHTTP(rr, req)

		respBody := warehouseResponseBody{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, data, respBody.Data)
		// assert.Empty(t, respBody.Error)
	})

	// t.Run("", func(t *testing.T) {

	// 	invalidBody := bytes.NewBuffer([]byte(`
	// 	{
	// 		"warehouse_code": "valid_code",
	// 		"minimum_capacity": 10,
	// 		"minimum_temperature": 8
	// 	}
	// 	`))

	// 	rr := httptest.NewRecorder()

	// 	req, _ := http.NewRequest(http.MethodPost, URL, invalidBody)

	// 	server.ServeHTTP(rr, req)

	// 	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	// 	assert.Contains(t, rr.Body.String(), "\"error\":")
	// })

	// t.Run("", func(t *testing.T) {

	// 	service.On("UptadedWarehouseID", mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return().Once()

	// 	data := makeValidDBWarehouse()

	// 	dataJSON, _ := json.Marshal(data)

	// 	body := strings.NewReader(string(dataJSON))

	// 	rr := httptest.NewRecorder()

	// 	req, _ := http.NewRequest(http.MethodPost, URL, body)

	// 	server.ServeHTTP(rr, req)

	// 	assert.Equal(t, http.StatusConflict, rr.Code)
	// 	assert.Contains(t, rr.Body.String(), "o `warehouse_code` já está em uso")
	// })
}
