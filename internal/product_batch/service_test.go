package productbatch_test

import (
	"errors"
	"testing"

	productbatch "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_batch"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_batch/mocks"
	"github.com/stretchr/testify/assert"
)

func InitTestService(t *testing.T) (productbatch.Services, *mocks.Repository) {
	mockRepository := mocks.NewRepository(t)
	service := productbatch.NewService(mockRepository)

	return service, mockRepository
}

func TestServiceCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		service, mockRepository := InitTestService(t)
		exp := productbatch.ProductBatch{1, 111, 200, 20, "2022-04-04", 10, "2020-04-04", 10, 5, 1, 1}

		mockRepository.On("Create", exp).Return(exp, nil)
		pb, err := service.Create(exp)

		assert.NoError(t, err)
		assert.Equal(t, exp, pb)
	})

	t.Run("create_fail", func(t *testing.T) {
		service, mockRepository := InitTestService(t)
		exp := productbatch.ProductBatch{1, 111, 200, 20, "2022-04-04", 10, "2020-04-04", 10, 5, 1, 1}

		mockRepository.On("Create", exp).Return(productbatch.ProductBatch{}, errors.New("sql: rows not affected"))
		pb, err := service.Create(exp)

		assert.Error(t, err)
		assert.Equal(t, errors.New("sql: rows not affected"), err)
		assert.Equal(t, productbatch.ProductBatch{}, pb)
	})
}

func TestServiceReport(t *testing.T) {
	t.Run("report_ok", func(t *testing.T) {
		service, mockRepository := InitTestService(t)
		exp := CreateReportArray()

		mockRepository.On("Report").Return(exp, nil)
		pb, err := service.Report()

		assert.NoError(t, err)
		assert.Equal(t, exp, pb)
	})

	t.Run("report_fail", func(t *testing.T) {
		service, mockRepository := InitTestService(t)

		mockRepository.On("Report").Return([]productbatch.Report{}, errors.New("sql: rows not affected"))
		pb, err := service.Report()

		assert.Error(t, err)
		assert.Equal(t, errors.New("sql: rows not affected"), err)
		assert.Equal(t, []productbatch.Report{}, pb)
	})
}

func TestServiceReportByID(t *testing.T) {
	t.Run("report_id_ok", func(t *testing.T) {
		service, mockRepository := InitTestService(t)
		exp := CreateReportArray()[0]

		mockRepository.On("ReportByID", 1).Return(exp, nil)
		pb, err := service.ReportByID(1)

		assert.NoError(t, err)
		assert.Equal(t, exp, pb)
	})

	t.Run("report_fail", func(t *testing.T) {
		service, mockRepository := InitTestService(t)

		mockRepository.On("ReportByID", 1).Return(productbatch.Report{}, errors.New("sql: rows not affected"))
		pb, err := service.ReportByID(1)

		assert.Error(t, err)
		assert.Equal(t, errors.New("sql: rows not affected"), err)
		assert.Equal(t, productbatch.Report{}, pb)
	})
}
