package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

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

func TestGetAll(t *testing.T) {
	gin.SetMode("release")
	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)
	sec := handlers.NewSection(service)
	exp := createSectionArray()

	router := gin.Default()
	router.GET("/api/v1/sections/", sec.GetAll())

	t.Run("find_all", func(t *testing.T) {
		mockRepository.On("GetAll").Return(exp)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/sections/", nil)
		router.ServeHTTP(w, req)

		exp, _ := json.Marshal(exp)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, string(exp), w.Body.String()[19:len(w.Body.String())-1])
	})
}

func TestGetByID(t *testing.T) {
	gin.SetMode("release")
	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)
	sec := handlers.NewSection(service)
	exp := createSectionArray()

	router := gin.Default()
	router.GET("/api/v1/sections/:id", sec.GetByID())

	t.Run("find_by_id_non_existent", func(t *testing.T) {
		mockRepository.On("GetByID", 90).Return(section.Section{}, errors.New("seção 90 não encontrada"))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/sections/90", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, "{\"code\":404,\"error\":\"seção 90 não encontrada\"}", w.Body.String())
	})

	t.Run("find_by_id_existent", func(t *testing.T) {
		mockRepository.On("GetByID", 1).Return(exp[0], nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/sections/1", nil)
		router.ServeHTTP(w, req)

		exp, _ := json.Marshal(exp[0])

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, string(exp), w.Body.String()[19:len(w.Body.String())-1])
	})
}

func TestCreate(t *testing.T) {
	gin.SetMode("release")
	router := gin.Default()

	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)
	sec := handlers.NewSection(service)
	router.POST("/api/v1/sections/", sec.CreateSection())

	t.Run("create_ok", func(t *testing.T) {
		secs := createSectionArray()
		exp := secs[0]
		secs = append([]section.Section{}, secs[1:]...)
		mockRepository.On("GetAll").Return(secs)
		mockRepository.On("Create", 1, exp.SectionNumber, exp.CurTemperature, exp.MinTemperature, exp.CurCapacity,
			exp.MinCapacity, exp.MaxCapacity, exp.WareHouseID, exp.ProductTypeID).Return(exp, nil)

		w := httptest.NewRecorder()
		expJSON, _ := json.Marshal(exp)
		req, _ := http.NewRequest("POST", "/api/v1/sections/", bytes.NewBuffer(expJSON))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		defer req.Body.Close()

		assert.Equal(t, 201, w.Code)
		assert.Equal(t, string(expJSON), w.Body.String()[19:len(w.Body.String())-1])
	})

	t.Run("create_fail", func(t *testing.T) {
		secs := createSectionArray()
		exp := secs[0]
		exp.ProductTypeID = 0
		secs = append([]section.Section{}, secs[1:]...)

		w := httptest.NewRecorder()
		expJSON, _ := json.Marshal(exp)
		req, _ := http.NewRequest("POST", "/api/v1/sections/", bytes.NewBuffer(expJSON))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		defer req.Body.Close()

		assert.Equal(t, 422, w.Code)
		assert.Equal(t, "{\"code\":422,\"error\":\"Key: 'sectionRequest.ProductTypeID'"+
			" Error:Field validation for 'ProductTypeID' failed on the 'required' tag\"}", w.Body.String())
	})
}

func TestCreateConflict(t *testing.T) {
	gin.SetMode("release")
	router := gin.Default()

	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)
	sec := handlers.NewSection(service)
	router.POST("/api/v1/sections/", sec.CreateSection())

	t.Run("create_conflict", func(t *testing.T) {
		secs := createSectionArray()
		exp := secs[0]

		mockRepository.On("GetAll").Return(secs)

		w := httptest.NewRecorder()
		expJSON, _ := json.Marshal(exp)
		req, _ := http.NewRequest("POST", "/api/v1/sections/", bytes.NewBuffer(expJSON))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		defer req.Body.Close()

		assert.Equal(t, 409, w.Code)
		assert.Equal(t, "{\"code\":409,\"error\":\"seção com sectionNumber: 40 já existe no banco de dados\"}", w.Body.String())
	})
}

func TestUpdateSecID(t *testing.T) {
	gin.SetMode("release")
	router := gin.Default()

	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)
	sec := handlers.NewSection(service)
	router.PATCH("/api/v1/sections/:id", sec.UpdateSecID())

	t.Run("update_ok", func(t *testing.T) {
		secs := createSectionArray()
		exp := secs[0]
		exp.SectionNumber = 50

		mockRepository.On("GetAll").Return(secs)
		mockRepository.On("UpdateSecID", 1, exp.SectionNumber).Return(exp, section.CodeError{})

		w := httptest.NewRecorder()

		expJSON, _ := json.Marshal(exp)
		req, _ := http.NewRequest("PATCH", "/api/v1/sections/1", bytes.NewBuffer(expJSON))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		defer req.Body.Close()

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, string(expJSON), w.Body.String()[19:len(w.Body.String())-1])
	})

	t.Run("update_non_existent", func(t *testing.T) {
		secs := createSectionArray()
		exp := secs[0]
		exp.SectionNumber = 50

		mockRepository.On("GetAll").Return(secs)
		mockRepository.On("UpdateSecID", 99, exp.SectionNumber).Return(section.Section{},
			section.CodeError{Code: 404, Message: errors.New("seção 99 não encontrada")})

		w := httptest.NewRecorder()

		expJSON, _ := json.Marshal(exp)
		req, _ := http.NewRequest("PATCH", "/api/v1/sections/99", bytes.NewBuffer(expJSON))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		defer req.Body.Close()

		assert.Equal(t, 404, w.Code)
		assert.Equal(t, "{\"code\":404,\"error\":\"seção 99 não encontrada\"}", w.Body.String())
	})

	t.Run("update_conflict", func(t *testing.T) {
		secs := createSectionArray()
		exp := secs[0]

		mockRepository.On("GetAll").Return(secs)

		w := httptest.NewRecorder()
		expJSON, _ := json.Marshal(exp)
		req, _ := http.NewRequest("PATCH", "/api/v1/sections/1", bytes.NewBuffer(expJSON))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		defer req.Body.Close()

		assert.Equal(t, 409, w.Code)
		assert.Equal(t, "{\"code\":409,\"error\":\"seção com section_number: 40 já existe no banco de dados\"}", w.Body.String())
	})
}

func TestDelete(t *testing.T) {
	gin.SetMode("release")
	router := gin.Default()

	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)
	sec := handlers.NewSection(service)
	router.DELETE("/api/v1/sections/:id", sec.DeleteSection())

	t.Run("delete_non_existent", func(t *testing.T) {
		//secs := createSectionArray()

		mockRepository.On("DeleteSection", 99).Return(errors.New("seção 99 não encontrada"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/api/v1/sections/99", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
		assert.Equal(t, "{\"code\":404,\"error\":\"seção 99 não encontrada\"}", w.Body.String())
	})

	t.Run("delete_ok", func(t *testing.T) {
		//secs := createSectionArray()

		mockRepository.On("DeleteSection", 1).Return(nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/api/v1/sections/1", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 204, w.Code)
	})
}
