package product_batches_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/handlers/product_batches"
	productbatch "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_batch"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_batch/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

const (
	URL_PRODUCTS_BATCH = "/api/v1/productBatches"
	URL_SECTION_REPORT = "/api/v1/sections/reportProducts"
)

const (
	ERROR_BIND          = "Key: 'ProductBatch.ProductTypeID' Error:Field validation for 'ProductTypeID' failed on the 'required' tag\nKey: 'ProductBatch.SectionID' Error:Field validation for 'SectionID' failed on the 'required' tag"
	ERROR_CONFLICT_SEC  = "Error 1452: Cannot add or update a child row: a foreign key constraint fails (`mercado-fresco`.`product_batches`, CONSTRAINT `FK_PRODUCT_BATCHES_SECTION` FOREIGN KEY (`section_id`) REFERENCES `section` (`id`))"
	ERROR_CONFLICT_PROD = "Error 1452: Cannot add or update a child row: a foreign key constraint fails (`mercado-fresco`.`product_batches`, CONSTRAINT `FK_PRODUCT_BATCHES_PRODUCT` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`))"
)

func CreateReportArray() []productbatch.Report {
	var exp = []productbatch.Report{
		{SecID: 1, SecNum: 22, ProdCount: 5},
		{SecID: 3, SecNum: 483, ProdCount: 28},
		{SecID: 5, SecNum: 7843, ProdCount: 90},
	}

	return exp
}

func InitTest(t *testing.T) (*gin.Engine, *mocks.Repository, product_batches.ProductBatch) {
	mockRepository := mocks.NewRepository(t)
	prod_b := productbatch.NewService(mockRepository)
	pb := product_batches.NewProductBatch(prod_b)

	rec := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rec)

	return engine, mockRepository, pb
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

type ExpectedErrorJSON struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func TestBatchCreate(t *testing.T) {
	engine, mockRepository, pb := InitTest(t)
	exp := productbatch.ProductBatch{
		ID:              1,
		BatchNumber:     111,
		CurQuantity:     200,
		CurTemperature:  20,
		DueDate:         "2022-04-04",
		InitialQuantity: 10,
		ManufactDate:    "2020-04-04",
		ManufactHour:    10,
		MinTemperature:  5,
		ProductTypeID:   1,
		SectionID:       1,
	}

	engine.POST(URL_PRODUCTS_BATCH, pb.Create())

	t.Run("create_ok", func(t *testing.T) {
		mockRepository.On("GetByBatchNum", mock.Anything, mock.Anything).Return(productbatch.ProductBatch{}, errors.New(""))
		mockRepository.On("Create", mock.Anything, exp).Return(exp, nil)

		expected, _ := json.Marshal(exp)
		req, w := InitServer(http.MethodPost, URL_PRODUCTS_BATCH, expected)

		engine.ServeHTTP(w, req)

		exp := ExpectedJSON{201, exp}
		expJSON, _ := json.Marshal(exp)

		assert.Equal(t, exp.Code, w.Code)
		assert.Equal(t, string(expJSON), w.Body.String())
	})

	t.Run("create_fail_bind", func(t *testing.T) {
		exp.SectionID = 0
		exp.ProductTypeID = 0

		expected, _ := json.Marshal(exp)
		req, w := InitServer(http.MethodPost, URL_PRODUCTS_BATCH, expected)
		engine.ServeHTTP(w, req)

		exp := ExpectedErrorJSON{422, ERROR_BIND}
		expectedJSON, _ := json.Marshal(exp)

		assert.Equal(t, exp.Code, w.Code)
		assert.Equal(t, string(expectedJSON), w.Body.String())
	})

	t.Run("create_fail_conflict_sec", func(t *testing.T) {
		exp.SectionID = 99
		exp.ProductTypeID = 1

		expected, _ := json.Marshal(exp)
		req, w := InitServer(http.MethodPost, URL_PRODUCTS_BATCH, expected)

		mockRepository.On("GetByBatchNum", mock.Anything, mock.Anything).Return(productbatch.ProductBatch{}, errors.New(""))
		mockRepository.On("Create", mock.Anything, exp).Return(productbatch.ProductBatch{}, errors.New(ERROR_CONFLICT_SEC))
		engine.ServeHTTP(w, req)

		exp := ExpectedErrorJSON{409, ERROR_CONFLICT_SEC}
		expectedJSON, _ := json.Marshal(exp)

		assert.Equal(t, exp.Code, w.Code)
		assert.Equal(t, string(expectedJSON), w.Body.String())
	})

	t.Run("create_fail_conflict_prod", func(t *testing.T) {
		exp.SectionID = 1
		exp.ProductTypeID = 99

		mockRepository.On("GetByBatchNum", mock.Anything, mock.Anything).Return(productbatch.ProductBatch{}, errors.New(""))
		mockRepository.On("Create", mock.Anything, exp).Return(productbatch.ProductBatch{}, errors.New(ERROR_CONFLICT_PROD))

		expected, _ := json.Marshal(exp)
		req, w := InitServer(http.MethodPost, URL_PRODUCTS_BATCH, expected)
		engine.ServeHTTP(w, req)

		exp := ExpectedErrorJSON{409, ERROR_CONFLICT_PROD}
		expectedJSON, _ := json.Marshal(exp)

		assert.Equal(t, exp.Code, w.Code)
		assert.Equal(t, string(expectedJSON), w.Body.String())
	})
}

func TestBatchReport(t *testing.T) {
	engine, mockRepository, pb := InitTest(t)
	exp := CreateReportArray()

	engine.GET(URL_SECTION_REPORT, pb.Report())

	t.Run("report_all_ok", func(t *testing.T) {
		mockRepository.On("Report", mock.Anything).Return(exp, nil).Once()
		req, w := InitServer(http.MethodGet, URL_SECTION_REPORT, nil)

		engine.ServeHTTP(w, req)

		exp := ExpectedJSON{200, exp}
		expJSON, _ := json.Marshal(exp)

		assert.Equal(t, exp.Code, w.Code)
		assert.Equal(t, string(expJSON), w.Body.String())
	})

	t.Run("report_all_fail_db", func(t *testing.T) {
		mockRepository.On("Report", mock.Anything).Return([]productbatch.Report{}, errors.New("sql: connection failed"))
		req, w := InitServer(http.MethodGet, URL_SECTION_REPORT, nil)

		engine.ServeHTTP(w, req)

		exp := ExpectedErrorJSON{400, "sql: connection failed"}
		ExpectedJSON, _ := json.Marshal(exp)

		assert.Equal(t, exp.Code, w.Code)
		assert.Equal(t, string(ExpectedJSON), w.Body.String())
	})
}

func TestBatchReportID(t *testing.T) {
	engine, mockRepository, pb := InitTest(t)
	exp := CreateReportArray()[0]

	engine.GET(URL_SECTION_REPORT, pb.Report())

	t.Run("report_id_ok", func(t *testing.T) {
		mockRepository.On("ReportByID", mock.Anything, 1).Return(exp, nil)
		req, w := InitServer(http.MethodGet, URL_SECTION_REPORT+"?id=1", nil)
		engine.ServeHTTP(w, req)

		exp := ExpectedJSON{200, exp}
		expJSON, _ := json.Marshal(exp)

		assert.Equal(t, exp.Code, w.Code)
		assert.Equal(t, string(expJSON), w.Body.String())
	})

	t.Run("report_fail_not_found", func(t *testing.T) {
		mockRepository.On("ReportByID", mock.Anything, 99).Return(productbatch.Report{}, errors.New("sql: no rows in result set"))
		req, w := InitServer(http.MethodGet, URL_SECTION_REPORT+"?id=99", nil)
		engine.ServeHTTP(w, req)

		exp := ExpectedErrorJSON{404, "sql: no rows in result set"}
		ExpectedJSON, _ := json.Marshal(exp)

		assert.Equal(t, exp.Code, w.Code)
		assert.Equal(t, string(ExpectedJSON), w.Body.String())
	})
}
