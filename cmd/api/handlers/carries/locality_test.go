package carries_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/handlers/carries"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases/mock/mock_repository_locality"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type localityResponseBody struct {
	Code  int             `json:"code"`
	Data  domain.Locality `json:"data"`
	Error string          `json:"error"`
}

type localityResponseBodyArray struct {
	Code  int               `json:"code"`
	Data  []domain.Locality `json:"data"`
	Error string            `json:"error"`
}

const (
	URLlocalityCarry = "/api/v1/localities/reportCarries"
)

func Test_GetCarryLocality(t *testing.T) {

	repository := mock_repository_locality.NewRepositoryLocality(t)
	service := usecases.NewServiceLocality(repository)
	controller := carries.NewLocality(service)

	server := gin.Default()

	gin.SetMode(gin.TestMode)

	server.GET(URLlocalityCarry, controller.GetCarryLocality)

	t.Run("Deve retornar um código 404, quando a locality da carry não existir.", func(t *testing.T) {

		repository.On("GetCarryLocalityByID", 1).Return(domain.Locality{}, errors.New("a localidade não foi encontrada!")).Once()

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, URLlocalityCarry+"?id=1", nil)

		server.ServeHTTP(rr, req)

		respBody := localityResponseBody{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Contains(t, "a localidade não foi encontrada!", respBody.Error)
	})

	t.Run("Deve retornar um código 200, e um locality da carry é encontrada, quando o id existir no BD", func(t *testing.T) {

		repository.On("GetCarryLocalityByID", 1).Return(domain.Locality{
			ID:    1,
			Name:  "Florianopolis",
			Count: 3,
		}, nil).Once()

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, URLlocalityCarry+"?id=1", nil)

		server.ServeHTTP(rr, req)

		respBody := localityResponseBodyArray{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, []domain.Locality{
			{
				ID:    1,
				Name:  "Florianopolis",
				Count: 3,
			},
		}, respBody.Data)
	})

	t.Run("Deve retornar um código 400, quando não conseguir acessar o BD.", func(t *testing.T) {

		repository.On("GetAllCarriesLocality").Return([]domain.Locality{}, errors.New("erro ao acessar o banco de dados")).Once()

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, URLlocalityCarry, nil)

		server.ServeHTTP(rr, req)

		respBody := localityResponseBody{}

		json.Unmarshal(rr.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, "erro ao acessar o banco de dados", respBody.Error)
	})

}
