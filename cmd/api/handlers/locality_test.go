package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

const (
	URL_LOCALITY = "/api/v1/localities/"
)

func TestLocality_ReportSellers(t *testing.T) {
	t.Run("Deve retornar status 200", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerLocality := NewLocality(mockService)

		reportSeller := locality.ReportSeller{LocalityID: 1, LocalityName: "Gru", SellersCount: 3}
		dataJson, _ := json.Marshal(reportSeller)

		mockService.On("ReportSellers", mock.Anything, 1).Return(reportSeller, nil)

		server := gin.Default()
		serverLocalityGroup := server.Group(URL_LOCALITY)
		serverLocalityGroup.GET("/reportSellers", handlerLocality.ReportSellers)

		req, rr := createRequestTest(http.MethodGet, URL_LOCALITY+"reportSellers?id=1", string(dataJson))
		server.ServeHTTP(rr, req)

		assert.Equal(t, 200, rr.Code)
	})

	t.Run("Deve retornar status 400", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerLocality := NewLocality(mockService)

		mockService.On("ReportSellers", mock.Anything, 1).Return(locality.ReportSeller{}, fmt.Errorf("error"))

		server := gin.Default()
		serverLocalityGroup := server.Group(URL_LOCALITY)
		serverLocalityGroup.GET("/reportSellers", handlerLocality.ReportSellers)

		req, rr := createRequestTest(http.MethodGet, URL_LOCALITY+"reportSellers?id=1", "")
		server.ServeHTTP(rr, req)

		assert.Equal(t, 400, rr.Code)
	})

	t.Run("Deve retornar status 400 quando erro do parametro da url", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerLocality := NewLocality(mockService)

		server := gin.Default()
		serverLocalityGroup := server.Group(URL_LOCALITY)
		serverLocalityGroup.GET("/reportSellers", handlerLocality.ReportSellers)
		req, rr := createRequestTest(http.MethodGet, URL_LOCALITY+"reportSellers?err=1", "")
		server.ServeHTTP(rr, req)

		assert.Equal(t, 400, rr.Code)
	})

	t.Run("Deve retornar status 500 quando erro do conversao id", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerLocality := NewLocality(mockService)

		server := gin.Default()
		serverLocalityGroup := server.Group(URL_LOCALITY)
		serverLocalityGroup.GET("/reportSellers", handlerLocality.ReportSellers)
		req, rr := createRequestTest(http.MethodGet, URL_LOCALITY+"reportSellers?id=a", "")
		server.ServeHTTP(rr, req)

		assert.Equal(t, 500, rr.Code)
	})
}

func TestLocality_Create(t *testing.T) {
	t.Run("Deve retornar status 409 quando o zipcode já existir", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerLocality := NewLocality(mockService)

		inputLocality := locality.Locality{Id: 1, ZipCode: "6700", LocalityName: "Gru", ProvinceName: "SP", CountryName: "BRA"}
		dataJson, _ := json.Marshal(inputLocality)

		mockService.On("Create", mock.Anything, inputLocality.ZipCode, inputLocality.LocalityName, inputLocality.ProvinceName, inputLocality.CountryName).
			Return(locality.Locality{}, fmt.Errorf("zip_code already exists"))

		server := gin.Default()
		serverLocalityGroup := server.Group(URL_LOCALITY)
		serverLocalityGroup.POST("/", handlerLocality.Create)

		req, rr := createRequestTest(http.MethodPost, URL_LOCALITY, string(dataJson))
		server.ServeHTTP(rr, req)

		assert.Equal(t, 409, rr.Code)
	})

	t.Run("Deve retornar status 400 quando o erro", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerLocality := NewLocality(mockService)

		inputLocality := locality.Locality{Id: 1, ZipCode: "6700", LocalityName: "Gru", ProvinceName: "SP", CountryName: "BRA"}
		dataJson, _ := json.Marshal(inputLocality)

		mockService.On("Create", mock.Anything, inputLocality.ZipCode, inputLocality.LocalityName, inputLocality.ProvinceName, inputLocality.CountryName).
			Return(locality.Locality{}, fmt.Errorf("error"))

		server := gin.Default()
		serverLocalityGroup := server.Group(URL_LOCALITY)
		serverLocalityGroup.POST("/", handlerLocality.Create)

		req, rr := createRequestTest(http.MethodPost, URL_LOCALITY, string(dataJson))
		server.ServeHTTP(rr, req)

		assert.Equal(t, 400, rr.Code)
	})

	t.Run("Deve retornar status 422 quando erro", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerLocality := NewLocality(mockService)

		inputLocality := locality.Locality{Id: 1, LocalityName: "Gru", ProvinceName: "SP", CountryName: "BRA"}
		dataJson, _ := json.Marshal(inputLocality)

		server := gin.Default()
		serverLocalityGroup := server.Group(URL_LOCALITY)
		serverLocalityGroup.POST("/", handlerLocality.Create)

		req, rr := createRequestTest(http.MethodPost, URL_LOCALITY, string(dataJson))
		server.ServeHTTP(rr, req)

		assert.Equal(t, 422, rr.Code)
	})

	t.Run("Deve retornar status 201 quando sucesso", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerLocality := NewLocality(mockService)

		inputLocality := locality.Locality{Id: 1, ZipCode: "6700", LocalityName: "Gru", ProvinceName: "SP", CountryName: "BRA"}
		dataJson, _ := json.Marshal(inputLocality)

		mockService.On("Create", mock.Anything, inputLocality.ZipCode, inputLocality.LocalityName, inputLocality.ProvinceName, inputLocality.CountryName).
			Return(inputLocality, nil)

		server := gin.Default()
		serverLocalityGroup := server.Group(URL_LOCALITY)
		serverLocalityGroup.POST("/", handlerLocality.Create)

		req, rr := createRequestTest(http.MethodPost, URL_LOCALITY, string(dataJson))
		server.ServeHTTP(rr, req)

		assert.Equal(t, 201, rr.Code)
	})
}

func TestLocality_GetAll(t *testing.T) {

	t.Run("Deve retornar status 200", func(t *testing.T) {
		mockService := mocks.NewRepository(t)
		handlerLocality := NewLocality(mockService)

		localityList := []locality.Locality{{1, "6700", "Gru", "SP", "BRA"}, {2, "6701", "Rio", "RJ", "BRA"}}
		dataJson, _ := json.Marshal(localityList)

		req, rr := createRequestTest(http.MethodGet, URL_LOCALITY, string(dataJson))
		mockService.On("GetAll", mock.Anything).Return(localityList, nil)

		server := gin.Default()
		localityServerGroup := server.Group(URL_LOCALITY)
		localityServerGroup.GET("/", handlerLocality.GetAll)

		server.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)
	})

	t.Run("Deve retornar status 404", func(t *testing.T) {
		mockService := mocks.NewRepository(t)
		handlerLocality := NewLocality(mockService)

		localityList := []locality.Locality{{1, "6700", "Gru", "SP", "BRA"}, {2, "6701", "Rio", "RJ", "BRA"}}
		dataJson, _ := json.Marshal(localityList)

		req, rr := createRequestTest(http.MethodGet, URL_LOCALITY, string(dataJson))
		mockService.On("GetAll", mock.Anything).Return([]locality.Locality{}, fmt.Errorf("error"))

		server := gin.Default()
		localityServerGroup := server.Group(URL_LOCALITY)
		localityServerGroup.GET("/", handlerLocality.GetAll)

		server.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)
	})

}

func TestLocality_ValidateLocalityFields(t *testing.T) {
	t.Run("Deve retornar a mensagem de erro do campo inválido", func(t *testing.T) {
		localityZipCode := requestLocality{ZipCode: "6700", LocalityName: "Gru", ProvinceName: "SP", CountryName: "BRA"}
		localityLocalityName := requestLocality{ZipCode: "6700", LocalityName: "", ProvinceName: "SP", CountryName: "BRA"}
		localityProvinceName := requestLocality{ZipCode: "6700", LocalityName: "Gru", ProvinceName: "", CountryName: "BRA"}
		localityCountryName := requestLocality{ZipCode: "6700", LocalityName: "Gru", ProvinceName: "SP", CountryName: ""}

		errLocalityZipCode := validateLocalityFields(localityZipCode)
		errLocalityLocalityName := validateLocalityFields(localityLocalityName)
		errLocalityProvinceName := validateLocalityFields(localityProvinceName)
		errLocalityCountryName := validateLocalityFields(localityCountryName)

		assert.NotEqual(t, "invalid input in field zip_code", errLocalityZipCode)
		assert.NotEqual(t, "invalid input in field locality_name", errLocalityLocalityName)
		assert.NotEqual(t, "invalid input in field province_name", errLocalityProvinceName)
		assert.NotEqual(t, "invalid input in field country_name", errLocalityCountryName)
	})

	t.Run("Deve retornar nil quando não houver erro", func(t *testing.T) {
		localityOk := requestLocality{ZipCode: "6700", LocalityName: "Gru", ProvinceName: "SP", CountryName: "BRA"}
		errLocalityOk := validateLocalityFields(localityOk)

		assert.Nil(t, errLocalityOk)
	})
}
