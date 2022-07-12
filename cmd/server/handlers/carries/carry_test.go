package carries_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers/carries"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases/mock/mock_repository_carry"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeValidDBCarry() domain.Carry {
	return domain.Carry{
		ID:         1,
		Cid:        "CID#5",
		Name:       "mercado-livre",
		Address:    "Criciuma, 666",
		Telephone:  "99999999",
		LocalityID: 2,
	}
}

const (
	URLcarry = "/api/v1/carries"
)

type carryResponseBody struct {
	Code  int          `json:"code"`
	Data  domain.Carry `json:"data"`
	Error string       `json:"error"`
}

func Test_CreateCarry(t *testing.T) {

	repository := mock_repository_carry.NewRepositoryCarry(t)
	service := usecases.NewServiceCarry(repository)
	controller := carries.NewCarry(service)
	server := gin.Default()

	gin.SetMode(gin.TestMode)

	server.POST(URLcarry, controller.CreateCarry)

	t.Run("Deve retornar um status code 201, quando a entrada de dados for bem-sucedida e retornará uma Carry.", func(t *testing.T) {

		data := makeValidDBCarry()

		repository.On("GetCarryByCid", mock.AnythingOfType("string")).Return(domain.Carry{}, fmt.Errorf("o carry com esse `cid`: %s não foi encontrado", data.Cid))

		repository.On("CreateCarry", mock.Anything).Return(data, nil).Once()

		dataJSON, _ := json.Marshal(data)

		body := strings.NewReader(string(dataJSON))

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, URLcarry, body)

		server.ServeHTTP(rr, req)

		respBody := carryResponseBody{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, data, respBody.Data)
		assert.Empty(t, respBody.Error)
	})

	t.Run("Deve retornar um status code 422, se o objeto JSON não contiver os campos necessários", func(t *testing.T) {

		invalidBody := bytes.NewBuffer([]byte(`
		{
			"company_name": "mercado-livre",
			"address": "Criciuma"
		}
		`))

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, URLcarry, invalidBody)

		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), "\"error\":")
	})

	t.Run("Deve retornar um status code 409, se o `cid` já estiver em uso.", func(t *testing.T) {

		data := makeValidDBCarry()

		repository.On("GetCarryByCid", mock.AnythingOfType("string")).Return(data, nil)

		repository.On("CreateCarry", mock.Anything).Return(domain.Carry{}, errors.New("o `cid` já está em uso")).Once()

		dataJSON, _ := json.Marshal(data)

		body := strings.NewReader(string(dataJSON))

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, URLcarry, body)

		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
		assert.Contains(t, rr.Body.String(), "o `cid` já está em uso")
	})

}
