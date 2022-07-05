package product_batches_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers/product_batches"
	productbatch "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_batch"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_batch/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

const URL_PRODUCTS_BATCH = "/api/v1/"

func CreateReportArray() []productbatch.Report {
	var exp = []productbatch.Report{
		{SecID: 1, SecNum: 22, ProdCount: 5},
		{SecID: 3, SecNum: 483, ProdCount: 28},
		{SecID: 5, SecNum: 7843, ProdCount: 90},
	}

	return exp
}

func InitTest(t *testing.T) (*gin.Engine, *mocks.Repository, product_batches.ProductBatch) {
	gin.SetMode("release")
	router := gin.Default()

	mockRepository := mocks.NewRepository(t)
	prod_b := productbatch.NewService(mockRepository)
	pb := product_batches.NewProductBatch(prod_b)

	return router, mockRepository, pb
}

func InitServer(method string, url string, body []byte) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	defer req.Body.Close()

	return req, httptest.NewRecorder()
}

type ExpectedAllJSON struct {
	Code int                   `json:"code"`
	Data []productbatch.Report `json:"data"`
}

type ExpectedIdJSON struct {
	Code int                 `json:"code"`
	Data productbatch.Report `json:"data"`
}

type ExpectedCreateJSON struct {
	Code int                          `json:"code"`
	Data product_batches.ProductBatch `json:"data"`
}

func TestBatchReport(t *testing.T) {
	router, mockRepository, pb := InitTest(t)
	exp := CreateReportArray()

	router.GET(URL_PRODUCTS_BATCH+"sections/reportProducts", pb.Report())

	t.Run("report_all_ok", func(t *testing.T) {
		mockRepository.On("Report").Return(exp, nil)
		req, w := InitServer(http.MethodGet, URL_PRODUCTS_BATCH+"sections/reportProducts", nil)
		router.ServeHTTP(w, req)

		exp := ExpectedAllJSON{200, exp}
		expJSON, _ := json.Marshal(exp)

		assert.Equal(t, exp.Code, w.Code)
		assert.Equal(t, string(expJSON), w.Body.String())
	})
}

func TestBatchReportID(t *testing.T) {
	router, mockRepository, pb := InitTest(t)
	exp := CreateReportArray()[0]

	router.GET(URL_PRODUCTS_BATCH+"sections/reportProducts", pb.Report())

	t.Run("report_id_ok", func(t *testing.T) {
		mockRepository.On("ReportByID", 1).Return(exp, nil)
		req, w := InitServer(http.MethodGet, URL_PRODUCTS_BATCH+"sections/reportProducts?id=1", nil)
		router.ServeHTTP(w, req)

		exp := ExpectedIdJSON{200, exp}
		expJSON, _ := json.Marshal(exp)

		assert.Equal(t, exp.Code, w.Code)
		assert.Equal(t, string(expJSON), w.Body.String())
	})

	t.Run("report_fail_not_found", func(t *testing.T) {
		mockRepository.On("ReportByID", 99).Return(productbatch.Report{}, errors.New("sql: no rows in result set"))
		req, w := InitServer(http.MethodGet, URL_PRODUCTS_BATCH+"sections/reportProducts?id=99", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
		assert.Equal(t, "{\"code\":404,\"error\":\"sql: no rows in result set\"}", w.Body.String())
	})
}
