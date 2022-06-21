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
		ID:             1,
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

func Test_CreateWarehouse(t *testing.T) {

	service := mock_service.NewService(t)
	controller := handlers.NewWarehouse(service)
	server := gin.Default()

	gin.SetMode(gin.TestMode) // Pra deixar o framework do gin em modo de test

	server.POST(URL, controller.CreateWarehouse)

	t.Run("Deve retornar um status code 201, quando a entra de dados for bem-sucedida e retornará um warehouse.", func(t *testing.T) {

		data := makeValidDBWarehouse()

		service.On("CreateWarehouse", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(data, nil).Once()

		dataJSON, _ := json.Marshal(data) // Retorna Array de Bytes.

		body := strings.NewReader(string(dataJSON)) // NewReader -> leitor para consumir esses Bytes.

		rr := httptest.NewRecorder() // Monitorar e armazenar requisições http.(Response)

		req, _ := http.NewRequest(http.MethodPost, URL, body)

		server.ServeHTTP(rr, req) // gerenciador de requisições

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, "{\"code\":201,\"data\":{\"id\":1,\"warehouse_code\":\"j753\",\"address\":\"Rua das Margaridas\",\"telephone\":\"4833334444\",\"minimun_capacity\":100,\"minimun_temperature\":10}}", rr.Body.String())
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

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"code\":200,\"data\":[{\"id\":1,\"warehouse_code\":\"j753\",\"address\":\"Rua das Margaridas\",\"telephone\":\"4833334444\",\"minimun_capacity\":100,\"minimun_temperature\":10}]}", rr.Body.String())
	})
}
