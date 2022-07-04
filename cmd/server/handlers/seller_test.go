package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	URL = "/api/v1/sellers/"
)

type responseArray struct {
	Code  int
	Data  []seller.Seller
	Error string
}

type responseId struct {
	Code  int
	Data  seller.Seller
	Error string
}

func TestSeller_Update(t *testing.T) {
	t.Run("Se o tipo de vendedor a ser atualizado não existir, um código 404 será devolvido.", func(t *testing.T) {
		var id int = 2
		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		expected := seller.Seller{CompanyId: 2, CompanyName: "Expected", Address: "BR", Telephone: "5501154545454"}
		expectedError := fmt.Errorf("the id %d does not exists", id)

		dataJson, _ := json.Marshal(expected)

		mockService.On("GetOne", mock.Anything, 2).Return(seller.Seller{}, expectedError)

		server := gin.Default()
		serverSellerGroup := server.Group(URL)

		serverSellerGroup.PUT("/:id", handlerSeller.Update)

		req, rr := createRequestTest(http.MethodPut, URL+"2", string(dataJson))
		server.ServeHTTP(rr, req)

		assert.Equal(t, 404, rr.Code)
	})
	t.Run("Quando a atualização dos dados for bem sucedida, o vendedor será devolvido com as informações atualizadas juntamente com um código 200", func(t *testing.T) {

		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		data := seller.Seller{Id: 1, CompanyId: 2, CompanyName: "Data", Address: "ARG", Telephone: "9999999"}
		expected := seller.Seller{CompanyId: 2, CompanyName: "Expected", Address: "BR", Telephone: "5501154545454"}

		dataJson, _ := json.Marshal(expected)

		mockService.On("GetOne", mock.Anything, 1).Return(data, nil)
		mockService.On("Update", mock.Anything, 1, expected.CompanyId, expected.CompanyName, expected.Address, expected.Telephone).
			Return(expected, nil)

		server := gin.Default()
		serverSellerGroup := server.Group(URL)

		serverSellerGroup.PUT("/:id", handlerSeller.Update)

		req, rr := createRequestTest(http.MethodPut, URL+"1", string(dataJson))
		server.ServeHTTP(rr, req)

		assert.Equal(t, 200, rr.Code)
	})

	t.Run("Se o campo a ser atualizado estiver não conforme, um código 422 será devolvido.", func(t *testing.T) {

		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		expected := seller.Seller{Id: 1, CompanyName: "Expected", Address: "BR", Telephone: "5501154545454"}

		dataJson, _ := json.Marshal(expected)

		mockService.On("GetOne", mock.Anything, 1).Return(seller.Seller{}, nil)

		server := gin.Default()
		serverSellerGroup := server.Group(URL)

		serverSellerGroup.PUT("/:id", handlerSeller.Update)

		req, rr := createRequestTest(http.MethodPut, URL+"1", string(dataJson))
		server.ServeHTTP(rr, req)

		assert.Equal(t, 422, rr.Code)
	})

	t.Run("Se o id não for um número, o código 500 será devolvido.", func(t *testing.T) {

		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		expected := seller.Seller{Id: 1, CompanyName: "Expected", Address: "BR", Telephone: "5501154545454"}

		dataJson, _ := json.Marshal(expected)

		server := gin.Default()
		serverSellerGroup := server.Group(URL)

		serverSellerGroup.PUT("/:id", handlerSeller.Update)

		req, rr := createRequestTest(http.MethodPut, URL+"err", string(dataJson))
		server.ServeHTTP(rr, req)

		assert.Equal(t, 500, rr.Code)
	})
}

func TestSeller_Create(t *testing.T) {
	t.Run("Se o cid já existir, ele retornará um erro 409 Conflict.", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		input := seller.Seller{CompanyId: 5, CompanyName: "Meli", Address: "BR", Telephone: "5501154545454"}
		expectedError := errors.New("the cid already exists")

		dataJson, _ := json.Marshal(input)

		mockService.On("Create", mock.Anything, input.CompanyId, input.CompanyName, input.Address, input.Telephone).Return(seller.Seller{}, expectedError)

		server := gin.Default()

		serverSellerGroup := server.Group(URL)

		serverSellerGroup.POST("/", handlerSeller.Create)

		req, rr := createRequestTest(http.MethodPost, URL, string(dataJson))
		server.ServeHTTP(rr, req)

		assert.Equal(t, 409, rr.Code)
	})

	t.Run("Se o objeto JSON não contiver os campos necessários, um código 422 será retornado.", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		input := seller.Seller{CompanyName: "Meli", Address: "BR", Telephone: "5501154545454"}

		dataJson, _ := json.Marshal(input)

		server := gin.Default()

		serverSellerGroup := server.Group(URL)
		serverSellerGroup.POST("/", handlerSeller.Create)

		req, rr := createRequestTest(http.MethodPost, URL, string(dataJson))
		server.ServeHTTP(rr, req)

		assert.Equal(t, 422, rr.Code)
	})

	t.Run("Quando a entrada de dados for bem-sucedida, um código 201 será retornado junto com o objeto inserido.", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		expected := seller.Seller{Id: 1, CompanyId: 5, CompanyName: "Meli", Address: "BR", Telephone: "5501154545454"}
		input := seller.Seller{CompanyId: 5, CompanyName: "Meli", Address: "BR", Telephone: "5501154545454"}

		dataJson, _ := json.Marshal(input)

		mockService.On("Create", mock.Anything, expected.CompanyId, expected.CompanyName, expected.Address, expected.Telephone).Return(expected, nil)

		server := gin.Default()

		serverSellerGroup := server.Group(URL)
		serverSellerGroup.POST("/", handlerSeller.Create)

		req, rr := createRequestTest(http.MethodGet, URL, "")
		server.ServeHTTP(rr, req)

		req, rr = createRequestTest(http.MethodPost, URL, string(dataJson))
		server.ServeHTTP(rr, req)

		assert.Equal(t, 201, rr.Code)
	})
}

func TestSeller_GetOne(t *testing.T) {
	t.Run("Quando o vendedor não existir, um código 404 será devolvido", func(t *testing.T) {
		var id int = 2
		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		expectedError := fmt.Errorf("the id %d does not exists", id)
		mockService.On("GetOne", mock.Anything, id).Return(seller.Seller{}, expectedError)

		server := gin.Default()
		serverSellerGroup := server.Group(URL)
		serverSellerGroup.GET("/:id", handlerSeller.GetOne)

		req, rr := createRequestTest(http.MethodGet, URL+"2", "")
		server.ServeHTTP(rr, req)

		assert.Equal(t, 404, rr.Code)
	})

	t.Run("Quando a solicitação for bem-sucedida, o back-end retornará as informações solicitadas do vendedor", func(t *testing.T) {

		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		sellerOne := seller.Seller{Id: 1, CompanyId: 5, CompanyName: "TestGetOne", Address: "BR", Telephone: "5501154545454"}

		expectedJson, _ := json.Marshal(sellerOne)

		req, rr := createRequestTest(http.MethodGet, URL+"1", string(expectedJson))
		mockService.On("GetOne", mock.Anything, 1).Return(sellerOne, nil)

		server := gin.Default()
		sellerServerGroup := server.Group(URL)
		sellerServerGroup.GET("/:id", handlerSeller.GetOne)

		server.ServeHTTP(rr, req)

		var response responseId

		json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, sellerOne, response.Data)
	})

	t.Run("Quando o id não for um número, deve retornar um erro e código 500", func(t *testing.T) {

		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		sellerOne := seller.Seller{Id: 1, CompanyId: 5, CompanyName: "TestGetOne", Address: "BR", Telephone: "5501154545454"}

		expectedJson, _ := json.Marshal(sellerOne)

		req, rr := createRequestTest(http.MethodGet, URL+"err", string(expectedJson))

		server := gin.Default()
		sellerServerGroup := server.Group(URL)
		sellerServerGroup.GET("/:id", handlerSeller.GetOne)

		server.ServeHTTP(rr, req)

		var response responseId

		json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Equal(t, 500, response.Code)
	})
}

func TestSeller_GetAll(t *testing.T) {
	t.Run("Deverá retornar erro quando a solicitação for mal sucedida", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		sellerList := []seller.Seller{{Id: 1, CompanyId: 5, CompanyName: "TestUpdate", Address: "BR", Telephone: "5501154545454"},
			{Id: 2, CompanyId: 6, CompanyName: "TestGetAll", Address: "BR", Telephone: "5501154545454"}}

		dataJson, _ := json.Marshal(sellerList)

		expectedError := errors.New("erro ao inicializar a lista")

		req, rr := createRequestTest(http.MethodGet, URL, string(dataJson))
		mockService.On("GetAll", mock.Anything).Return([]seller.Seller{}, expectedError)

		server := gin.Default()
		sellerServerGroup := server.Group(URL)
		sellerServerGroup.GET("/", handlerSeller.GetAll)

		server.ServeHTTP(rr, req)

		assert.Equal(t, 404, rr.Code)
	})
	t.Run("Quando a solicitação for bem-sucedida, o back-end retornará uma lista de todos os vendedores existentes.", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		sellerList := []seller.Seller{{Id: 1, CompanyId: 5, CompanyName: "TestUpdate", Address: "BR", Telephone: "5501154545454"},
			{Id: 2, CompanyId: 6, CompanyName: "TestGetAll", Address: "BR", Telephone: "5501154545454"}}

		dataJson, _ := json.Marshal(sellerList)

		req, rr := createRequestTest(http.MethodGet, URL, string(dataJson))
		mockService.On("GetAll", mock.Anything).Return(sellerList, nil)

		server := gin.Default()
		sellerServerGroup := server.Group(URL)
		sellerServerGroup.GET("/", handlerSeller.GetAll)

		var response responseArray

		server.ServeHTTP(rr, req)
		json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, sellerList, response.Data)
		assert.Equal(t, "", response.Error)
	})
}

func TestSeller_Delete(t *testing.T) {
	t.Run("Quando o vendedor não existir, um código 404 será devolvido", func(t *testing.T) {
		var id int = 3
		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		server := gin.Default()
		sellerRouterGroup := server.Group(URL)
		expectedError := fmt.Errorf("the id %d does not exists", id)

		req, rr := createRequestTest(http.MethodDelete, URL+"3", "")
		mockService.On("Delete", mock.Anything, id).Return(expectedError)

		sellerRouterGroup.DELETE("/:id", handlerSeller.Delete)
		server.ServeHTTP(rr, req)

		assert.Equal(t, 404, rr.Code)
	})

	t.Run("Quando a exclusão for bem-sucedida, um código 204 será retornado.", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		server := gin.Default()
		sellerRouterGroup := server.Group(URL)

		req, rr := createRequestTest(http.MethodDelete, URL+"1", "")

		mockService.On("Delete", mock.Anything, 1).Return(nil)

		sellerRouterGroup.DELETE("/:id", handlerSeller.Delete)
		server.ServeHTTP(rr, req)

		assert.Equal(t, 204, rr.Code)
	})

	t.Run("Quando o id não for int, um código 500 deverá devolvido", func(t *testing.T) {
		mockService := mocks.NewService(t)
		handlerSeller := NewSeller(mockService)

		server := gin.Default()
		sellerRouterGroup := server.Group(URL)

		req, rr := createRequestTest(http.MethodDelete, URL+"err", "")

		sellerRouterGroup.DELETE("/:id", handlerSeller.Delete)
		server.ServeHTTP(rr, req)

		assert.Equal(t, 500, rr.Code)
	})
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", os.Getenv("TOKEN"))
	return req, httptest.NewRecorder()
}

func TestSeller_ValidateFields(t *testing.T) {

	t.Run("Deve retornar a mensagem de erro do campo inválido", func(t *testing.T) {
		sellerCompanyId := requestSeller{CompanyName: "TestUpdate", Address: "BR", Telephone: "5501154545454"}
		sellerCompanyName := requestSeller{CompanyId: 5, Address: "BR", Telephone: "5501154545454"}
		sellerAddress := requestSeller{CompanyId: 5, CompanyName: "TestUpdate", Telephone: "5501154545454"}
		sellerTelephone := requestSeller{CompanyId: 5, CompanyName: "TestUpdate", Address: "BR"}

		errSellerCompanyId := validateFields(sellerCompanyId)
		errSellerCompanyName := validateFields(sellerCompanyName)
		errSellersellerAddress := validateFields(sellerAddress)
		errSellerTelephone := validateFields(sellerTelephone)

		assert.NotEqual(t, "field cid is required", errSellerCompanyId)
		assert.NotEqual(t, "field company_name is required", errSellerCompanyName)
		assert.NotEqual(t, "field address is required", errSellersellerAddress)
		assert.NotEqual(t, "field telephone is required", errSellerTelephone)
	})

	t.Run("Deve retornar nil quando não houver erro", func(t *testing.T) {
		sellerOk := requestSeller{CompanyId: 5, CompanyName: "TestUpdate", Address: "BR", Telephone: "5501154545454"}
		errSellerOk := validateFields(sellerOk)

		assert.Nil(t, errSellerOk)
	})

}
