package warehouses_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/handlers/warehouses"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/usecases/mock/mock_service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeValidDBWarehouse() domain.Warehouse {
	return domain.Warehouse{
		WarehouseCode: "j753",
		Address:       "Rua das Margaridas",
		Telephone:     "4833334444",
		LocalityID:    1,
	}
}

const (
	URLwarehouses = "/api/v1/warehouse"
)

type warehouseResponseBody struct {
	Code  int              `json:"code"`
	Data  domain.Warehouse `json:"data"`
	Error string           `json:"error"`
}

type warehouseResponseBodyArray struct {
	Code  int                `json:"code"`
	Data  []domain.Warehouse `json:"data"`
	Error string             `json:"error"`
}

func Test_CreateWarehouse(t *testing.T) {

	service := mock_service.NewService(t)
	controller := warehouses.NewWarehouse(service)
	server := gin.Default()

	gin.SetMode(gin.TestMode) // Pra deixar o framework do gin em modo de test

	server.POST(URLwarehouses, controller.CreateWarehouse)

	t.Run("Deve retornar um status code 201, quando a entrada de dados for bem-sucedida e retornará um warehouse.", func(t *testing.T) {

		data := makeValidDBWarehouse()

		service.On("CreateWarehouse", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(data, nil).Once()

		dataJSON, _ := json.Marshal(data) // Retorna Array de Bytes.

		body := strings.NewReader(string(dataJSON)) // NewReader -> leitor para consumir esses Bytes.

		rr := httptest.NewRecorder() // Monitorar e armazenar requisições http.(Response)

		req, _ := http.NewRequest(http.MethodPost, URLwarehouses, body)

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

		req, _ := http.NewRequest(http.MethodPost, URLwarehouses, invalidBody)

		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), "\"error\":")
	})

	t.Run("Deve retornar um status code 409, se `warehouse_code` já estiver em uso.", func(t *testing.T) {

		service.On("CreateWarehouse", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(domain.Warehouse{}, errors.New("o `warehouse_code` já está em uso")).Once()

		data := makeValidDBWarehouse()

		dataJSON, _ := json.Marshal(data)

		body := strings.NewReader(string(dataJSON))

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, URLwarehouses, body)

		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
		assert.Contains(t, rr.Body.String(), "o `warehouse_code` já está em uso")
	})
}

func Test_GetAll(t *testing.T) {

	service := mock_service.NewService(t)
	controller := warehouses.NewWarehouse(service)
	server := gin.Default()

	gin.SetMode(gin.TestMode)

	server.GET(URLwarehouses, controller.GetAll)

	t.Run("Deve retornar uma lista de warehouses se a solicitação for bem sucedida.", func(t *testing.T) {

		data := makeValidDBWarehouse()

		service.On("GetAll").Return([]domain.Warehouse{data})

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, URLwarehouses, nil)

		server.ServeHTTP(rr, req)

		respBody := warehouseResponseBodyArray{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, data, respBody.Data[0])
	})
}

func Test_GetByID(t *testing.T) {

	service := mock_service.NewService(t)
	controller := warehouses.NewWarehouse(service)
	server := gin.Default()

	server.GET(URLwarehouses+"/:id", controller.GetByID)

	t.Run("Deve retornar um código 404, quando o Warehouse não existir.", func(t *testing.T) {

		service.On("GetByID", 1).Return(domain.Warehouse{}, errors.New("O warehouse não foi encontrado!")).Once()

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, URLwarehouses+"/1", nil)

		server.ServeHTTP(rr, req)

		respBody := warehouseResponseBody{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Contains(t, "O warehouse não foi encontrado!", respBody.Error)
	})

	t.Run("Deve retornar um código 400, e uma mensagem de erro, quando o id passado não for um número.", func(t *testing.T) {

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, URLwarehouses+"/casa", nil)

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

		req, _ := http.NewRequest(http.MethodGet, URLwarehouses+"/1", nil)

		server.ServeHTTP(rr, req)

		respBody := warehouseResponseBody{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusOK, respBody.Code)
		assert.Equal(t, data, respBody.Data)
	})
}

func Test_UpdatedWarehouseID(t *testing.T) {

	service := mock_service.NewService(t)
	controller := warehouses.NewWarehouse(service)
	server := gin.Default()

	gin.SetMode(gin.TestMode)

	server.PATCH(URLwarehouses+"/:id", controller.UpdatedWarehouseID)

	t.Run("Deve retornar um status code 200, e o Warehouse atualizado, quando a solicitação for bem sucedida.", func(t *testing.T) {

		data := makeValidDBWarehouse()

		service.On("UpdatedWarehouseID", mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(data, nil).Once()

		dataJSON, _ := json.Marshal(data)

		body := strings.NewReader(string(dataJSON))

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPatch, URLwarehouses+"/1", body)

		server.ServeHTTP(rr, req)

		respBody := warehouseResponseBody{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, data, respBody.Data)
		assert.Empty(t, respBody.Error)
	})

	t.Run("Deve retornar um status code 422, se o objeto JSON não contiver os campos necessários", func(t *testing.T) {

		invalidBody := bytes.NewBuffer([]byte(`
		{
			"minimum_capacity": 10,
			"minimum_temperature": 8
		}
		`))

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPatch, URLwarehouses+"/1", invalidBody)

		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), "\"error\":")
	})

	t.Run("Deve retornar um código 400, e uma mensagem de erro, quando o id passado não for um número.", func(t *testing.T) {

		data := makeValidDBWarehouse()

		dataJSON, _ := json.Marshal(data)

		body := strings.NewReader(string(dataJSON))

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPatch, URLwarehouses+"/casa", body)

		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "O id passado não é um número!")
	})

	t.Run("Deve retornar um código 404, se o Warehouse a ser atualizado não existir.", func(t *testing.T) {

		data := makeValidDBWarehouse()

		service.On("UpdatedWarehouseID", mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(domain.Warehouse{}, fmt.Errorf("o id: %d informado não existe", 1)).Once()

		dataJSON, _ := json.Marshal(data)

		body := strings.NewReader(string(dataJSON))

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPatch, URLwarehouses+"/1", body)

		respBody := warehouseResponseBody{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Contains(t, "o id: 1 informado não existe", respBody.Error)
	})
}

func Test_DeleteWarehouse(t *testing.T) {

	service := mock_service.NewService(t)
	controller := warehouses.NewWarehouse(service)
	server := gin.Default()

	gin.SetMode(gin.TestMode)

	server.DELETE(URLwarehouses+"/:id", controller.DeleteWarehouse)

	t.Run("Deve retornar um código 404, se o Warehouse não existir.", func(t *testing.T) {

		service.On("DeleteWarehouse", 1).Return(fmt.Errorf("não foi achado warehouse com esse id: %d", 1)).Once()

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodDelete, URLwarehouses+"/1", nil)

		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Contains(t, rr.Body.String(), "não foi achado warehouse com esse id")
	})

	t.Run("Deve retornar um código 400, e uma mensagem de erro, quando o id passado não for um número.", func(t *testing.T) {

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodDelete, URLwarehouses+"/casa", nil)

		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "O id passado não é um número!")
	})

	t.Run("Deve retornar um código 204, se o Warehouse for deletado com sucesso.", func(t *testing.T) {

		service.On("DeleteWarehouse", 1).Return(nil).Once()

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodDelete, URLwarehouses+"/1", nil)

		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)
		assert.Empty(t, rr.Body.String())
	})
}
