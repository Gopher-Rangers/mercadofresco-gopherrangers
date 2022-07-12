package productbatch_test

import (
	"context"
	"errors"
	"testing"

	productbatch "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_batch"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_batch/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

		mockRepository.On("ReportByID", mock.Anything, mock.Anything).Return(productbatch.Report{}, errors.New(""))
		mockRepository.On("Create", mock.Anything, exp).Return(exp, nil)
		pb, err := service.Create(context.TODO(), exp)

		assert.NoError(t, err)
		assert.Equal(t, exp, pb)
	})

	t.Run("create_fail", func(t *testing.T) {
		service, mockRepository := InitTestService(t)
		exp := productbatch.ProductBatch{1, 111, 200, 20, "2022-04-04", 10, "2020-04-04", 10, 5, 1, 1}

		mockRepository.On("ReportByID", mock.Anything, mock.Anything).Return(productbatch.Report{}, errors.New(""))
		mockRepository.On("Create", mock.Anything, exp).Return(productbatch.ProductBatch{}, errors.New("sql: rows not affected"))
		pb, err := service.Create(context.TODO(), exp)

		assert.Error(t, err)
		assert.Equal(t, errors.New("sql: rows not affected"), err)
		assert.Equal(t, productbatch.ProductBatch{}, pb)
	})

	t.Run("create_already_exists", func(t *testing.T) {
		service, mockRepository := InitTestService(t)
		exp := productbatch.ProductBatch{1, 111, 200, 20, "2022-04-04", 10, "2020-04-04", 10, 5, 1, 1}

		mockRepository.On("ReportByID", mock.Anything, mock.Anything).Return(productbatch.Report{}, nil)
		pb, err := service.Create(context.TODO(), exp)

		assert.Error(t, err)
		assert.Equal(t, errors.New("error: batch number '111' already exists in BD"), err)
		assert.Equal(t, productbatch.ProductBatch{}, pb)
	})
}

func TestServiceReport(t *testing.T) {
	t.Run("report_ok", func(t *testing.T) {
		service, mockRepository := InitTestService(t)
		exp := CreateReportArray()

		mockRepository.On("Report", mock.Anything).Return(exp, nil)
		pb, err := service.Report(context.TODO())

		assert.NoError(t, err)
		assert.Equal(t, exp, pb)
	})

	t.Run("report_fail", func(t *testing.T) {
		service, mockRepository := InitTestService(t)

		mockRepository.On("Report", mock.Anything).Return([]productbatch.Report{}, errors.New("sql: rows not affected"))
		pb, err := service.Report(context.TODO())

		assert.Error(t, err)
		assert.Equal(t, errors.New("sql: rows not affected"), err)
		assert.Equal(t, []productbatch.Report{}, pb)
	})
}

func TestServiceReportByID(t *testing.T) {
	t.Run("report_id_ok", func(t *testing.T) {
		service, mockRepository := InitTestService(t)
		exp := CreateReportArray()[0]

		mockRepository.On("ReportByID", mock.Anything, 1).Return(exp, nil)
		pb, err := service.ReportByID(context.TODO(), 1)

		assert.NoError(t, err)
		assert.Equal(t, exp, pb)
	})

	t.Run("report_fail", func(t *testing.T) {
		service, mockRepository := InitTestService(t)

		mockRepository.On("ReportByID", mock.Anything, 1).Return(productbatch.Report{}, errors.New("sql: rows not affected"))
		pb, err := service.ReportByID(context.TODO(), 1)

		assert.Error(t, err)
		assert.Equal(t, errors.New("sql: rows not affected"), err)
		assert.Equal(t, productbatch.Report{}, pb)
	})
}
