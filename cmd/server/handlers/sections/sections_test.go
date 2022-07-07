package sections_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	sections "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers/sections"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section/mocks"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const URL_SECTIONS = "/api/v1/sections/"

func createSectionArray() []section.Section {
	var sec []section.Section
	sec1 := section.Section{
		ID:             1,
		SectionNumber:  40,
		CurTemperature: 90,
		MinTemperature: 10,
		CurCapacity:    20,
		MinCapacity:    10,
		MaxCapacity:    999,
		WareHouseID:    9876,
		ProductTypeID:  7659,
	}

	sec2 := section.Section{
		ID:             2,
		SectionNumber:  20,
		CurTemperature: 32,
		MinTemperature: 15,
		CurCapacity:    100,
		MinCapacity:    20,
		MaxCapacity:    500,
		WareHouseID:    9876,
		ProductTypeID:  3747,
	}

	sec = append(sec, sec1, sec2)
	return sec
}

func InitTest(t *testing.T) (*gin.Engine, *mocks.Repository, sections.Section) {
	gin.SetMode("release")
	router := gin.Default()

	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)
	sec := sections.NewSection(service)

	return router, mockRepository, sec
}

func InitServer(method string, url string, body []byte) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	defer req.Body.Close()

	return req, httptest.NewRecorder()
}

type ExpectedJSON struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func TestSectionGetAll(t *testing.T) {
	router, mockRepository, sec := InitTest(t)
	exp := createSectionArray()
	router.GET(URL_SECTIONS, sec.GetAll())
	mockRepository.On("GetAll").Return(exp, nil)

	t.Run("find_all", func(t *testing.T) {
		req, w := InitServer(http.MethodGet, URL_SECTIONS, nil)
		router.ServeHTTP(w, req)

		exp := ExpectedJSON{200, exp}
		expJSON, _ := json.Marshal(exp)
		assert.Equal(t, exp.Code, w.Code)
		assert.Equal(t, string(expJSON), w.Body.String())
	})
}

func TestSectionGetByID(t *testing.T) {
	router, mockRepository, sec := InitTest(t)
	exp := createSectionArray()
	router.GET(URL_SECTIONS+":id", sec.GetByID())
	mockRepository.On("GetAll").Return(exp, nil)

	t.Run("find_by_id_non_existent", func(t *testing.T) {
		req, w := InitServer(http.MethodGet, URL_SECTIONS+"90", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, "{\"code\":404,\"error\":\"seção com id: 90 não existe no banco de dados\"}", w.Body.String())
	})

	t.Run("find_by_id_existent", func(t *testing.T) {
		mockRepository.On("GetByID", 1).Return(exp[0], nil)
		req, w := InitServer(http.MethodGet, URL_SECTIONS+"1", nil)
		router.ServeHTTP(w, req)

		expJSON := ExpectedJSON{200, exp[0]}
		expectedJSON, _ := json.Marshal(expJSON)

		assert.Equal(t, expJSON.Code, w.Code)
		assert.Equal(t, string(expectedJSON), w.Body.String())
	})
}

func TestSectionCreate(t *testing.T) {
	router, mockRepository, sec := InitTest(t)
	router.POST(URL_SECTIONS, sec.CreateSection())
	secs := createSectionArray()
	exp := secs[0]
	secs = append([]section.Section{}, secs[1:]...)

	t.Run("create_ok", func(t *testing.T) {
		mockRepository.On("GetAll").Return(secs, nil)
		mockRepository.On("Create", exp.SectionNumber, exp.CurTemperature, exp.MinTemperature, exp.CurCapacity,
			exp.MinCapacity, exp.MaxCapacity, exp.WareHouseID, exp.ProductTypeID).Return(exp, nil)

		expected, _ := json.Marshal(exp)
		req, w := InitServer(http.MethodPost, URL_SECTIONS, expected)
		router.ServeHTTP(w, req)

		expJSON := ExpectedJSON{201, exp}
		expectedJSON, _ := json.Marshal(expJSON)
		assert.Equal(t, expJSON.Code, w.Code)
		assert.Equal(t, string(expectedJSON), w.Body.String())
	})

	t.Run("create_conflict", func(t *testing.T) {
		exp = secs[0]
		mockRepository.On("GetAll").Return(secs)

		expJSON, _ := json.Marshal(exp)
		req, w := InitServer(http.MethodPost, URL_SECTIONS, expJSON)
		router.ServeHTTP(w, req)

		assert.Equal(t, 409, w.Code)
		assert.Equal(t, "{\"code\":409,\"error\":\"seção com sectionNumber: 20 já existe no banco de dados\"}", w.Body.String())
	})

	t.Run("create_fail", func(t *testing.T) {
		exp.ProductTypeID = 0

		expJSON, _ := json.Marshal(exp)
		req, w := InitServer(http.MethodPost, URL_SECTIONS, expJSON)
		router.ServeHTTP(w, req)

		assert.Equal(t, 422, w.Code)
		assert.Equal(t, "{\"code\":422,\"error\":\"Key: 'sectionRequest.ProductTypeID'"+
			" Error:Field validation for 'ProductTypeID' failed on the 'required' tag\"}", w.Body.String())
	})
}

func TestSectionUpdateSecID(t *testing.T) {
	router, mockRepository, sec := InitTest(t)
	router.PATCH(URL_SECTIONS+":id", sec.UpdateSecID())

	secs := createSectionArray()
	exp := secs[0]

	mockRepository.On("GetAll").Return(secs, nil)

	t.Run("update_ok", func(t *testing.T) {
		exp.SectionNumber = 50

		mockRepository.On("UpdateSecID", 1, exp.SectionNumber).Return(exp, section.CodeError{Code: 200, Message: nil})

		expected, _ := json.Marshal(exp)
		req, w := InitServer(http.MethodPatch, URL_SECTIONS+"1", expected)
		router.ServeHTTP(w, req)

		expJSON := ExpectedJSON{200, exp}
		expectedJSON, _ := json.Marshal(expJSON)

		assert.Equal(t, expJSON.Code, w.Code)
		assert.Equal(t, string(expectedJSON), w.Body.String())
	})

	t.Run("update_non_existent", func(t *testing.T) {
		exp.SectionNumber = 50

		mockRepository.On("UpdateSecID", 99, exp.SectionNumber).Return(section.Section{},
			section.CodeError{Code: 404, Message: errors.New("seção 99 não encontrada")})

		expJSON, _ := json.Marshal(exp)
		req, w := InitServer(http.MethodPatch, URL_SECTIONS+"99", expJSON)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
		assert.Equal(t, "{\"code\":404,\"error\":\"seção 99 não encontrada\"}", w.Body.String())
	})

	t.Run("update_conflict", func(t *testing.T) {
		exp.SectionNumber = 40

		expJSON, _ := json.Marshal(exp)
		req, w := InitServer(http.MethodPatch, URL_SECTIONS+"1", expJSON)
		router.ServeHTTP(w, req)

		assert.Equal(t, 409, w.Code)
		assert.Equal(t, "{\"code\":409,\"error\":\"seção com section_number: 40 já existe no banco de dados\"}", w.Body.String())
	})

	t.Run("update_fail", func(t *testing.T) {
		exp.SectionNumber = 0

		expJSON, _ := json.Marshal(exp)
		req, w := InitServer(http.MethodPatch, URL_SECTIONS+"1", expJSON)
		router.ServeHTTP(w, req)

		assert.Equal(t, 422, w.Code)
		assert.Equal(t, "{\"code\":422,\"error\":\"Key: 'sectionRequest.SectionNumber'"+
			" Error:Field validation for 'SectionNumber' failed on the 'required' tag\"}", w.Body.String())
	})
}

func TestSectionDelete(t *testing.T) {
	router, mockRepository, sec := InitTest(t)
	router.DELETE(URL_SECTIONS+":id", sec.DeleteSection())

	secs := createSectionArray()
	mockRepository.On("GetAll").Return(secs, nil)

	t.Run("delete_non_existent", func(t *testing.T) {
		req, w := InitServer(http.MethodDelete, URL_SECTIONS+"99", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
		assert.Equal(t, "{\"code\":404,\"error\":\"seção com id: 99 não existe no banco de dados\"}", w.Body.String())
	})

	t.Run("delete_ok", func(t *testing.T) {
		mockRepository.On("DeleteSection", 1).Return(nil)

		req, w := InitServer(http.MethodDelete, URL_SECTIONS+"1", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 204, w.Code)
	})
}

func TestTokenAuth(t *testing.T) {
	router, mockRepository, sec := InitTest(t)
	exp := createSectionArray()
	router.Use(sec.TokenAuthMiddleware)
	router.GET(URL_SECTIONS, sec.GetAll())
	godotenv.Load("../../../../.env")

	mockRepository.On("GetAll").Return(exp, nil)

	t.Run("token_ok", func(t *testing.T) {
		req, w := InitServer(http.MethodGet, URL_SECTIONS, nil)
		req.Header.Add("Token", os.Getenv("TOKEN"))

		router.ServeHTTP(w, req)

		exp := ExpectedJSON{200, exp}
		expJSON, _ := json.Marshal(exp)
		assert.Equal(t, exp.Code, w.Code)
		assert.Equal(t, string(expJSON), w.Body.String())
	})

	t.Run("token_invalid", func(t *testing.T) {
		req, w := InitServer(http.MethodGet, URL_SECTIONS, nil)
		req.Header.Add("Token", "XXXXXX")

		router.ServeHTTP(w, req)

		assert.Equal(t, 401, w.Code)
		assert.Equal(t, "{\"code\":401,\"error\":\"token inválido\"}", w.Body.String())
	})

	t.Run("token_empty", func(t *testing.T) {
		req, w := InitServer(http.MethodGet, URL_SECTIONS, nil)
		req.Header.Add("Token", "")

		router.ServeHTTP(w, req)

		assert.Equal(t, 401, w.Code)
		assert.Equal(t, "{\"code\":401,\"error\":\"token vazio\"}", w.Body.String())
	})
}

func TestVerificatorID(t *testing.T) {
	router, mockRepository, sec := InitTest(t)
	router.Use(sec.IdVerificatorMiddleware)
	router.DELETE(URL_SECTIONS+":id", sec.DeleteSection())
	godotenv.Load("../../../../.env")

	secs := createSectionArray()
	mockRepository.On("GetAll").Return(secs, nil)

	t.Run("id_ok", func(t *testing.T) {
		mockRepository.On("DeleteSection", 1).Return(nil)

		req, w := InitServer(http.MethodDelete, URL_SECTIONS+"1", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 204, w.Code)
	})

	t.Run("id_invalid", func(t *testing.T) {
		mockRepository.On("DeleteSection", 1).Return(nil)

		req, w := InitServer(http.MethodDelete, URL_SECTIONS+"XXXXX", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, "{\"code\":400,\"error\":\"id não é alphanumérico\"}", w.Body.String())
	})

	t.Run("id_negative", func(t *testing.T) {
		mockRepository.On("DeleteSection", 1).Return(nil)

		req, w := InitServer(http.MethodDelete, URL_SECTIONS+"-99", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
		assert.Equal(t, "{\"code\":404,\"error\":\"id negativo inválido\"}", w.Body.String())
	})
}
