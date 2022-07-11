package locality_test

import (
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality/mocks"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
)

func TestService_ReportSellers(t *testing.T) {
	t.Run("Deve retornar o report de sellers com sucesso", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)

		localityOne := locality.Locality{1, "6700", "Gru", "SP", "BRA"}
		reportSeller := locality.ReportSeller{1, "Gru", 3}

		mockRepo.On("GetById", context.Background(), 1).Return(localityOne, nil)
		mockRepo.On("ReportSellers", context.Background(), 1).Return(reportSeller, nil)

		service := locality.NewService(mockRepo)
		result, err := service.ReportSellers(context.Background(), 1)

		assert.NoError(t, err)
		assert.Equal(t, result, reportSeller)
	})

	t.Run("Deve retornar erro no report sellers report", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)

		localityOne := locality.Locality{1, "6700", "Gru", "SP", "BRA"}

		mockRepo.On("GetById", context.Background(), 1).Return(localityOne, nil)
		mockRepo.On("ReportSellers", context.Background(), 1).
			Return(locality.ReportSeller{}, fmt.Errorf("error"))

		service := locality.NewService(mockRepo)
		result, err := service.ReportSellers(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, result, locality.ReportSeller{})
	})

	t.Run("Deve retornar erro ao localizar a locality", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)

		localityOne := locality.Locality{1, "6700", "Gru", "SP", "BRA"}
		//reportSeller := locality.ReportSeller{1, "Gru", 3}

		mockRepo.On("GetById", context.Background(), 1).Return(localityOne, fmt.Errorf("error"))

		service := locality.NewService(mockRepo)
		result, err := service.ReportSellers(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, result, locality.ReportSeller{})
	})
}

func TestService_Create(t *testing.T) {
	t.Run("Deve criar uma locality com sucesso", func(t *testing.T) {

		mockRepo := mocks.NewRepository(t)

		expectedResult := locality.Locality{1, "6700", "Gru", "SP", "BRA"}

		mockRepo.On("GetAll", context.Background()).Return([]locality.Locality{}, nil)
		mockRepo.On("Create", context.Background(), expectedResult.ZipCode, expectedResult.LocalityName, expectedResult.ProvinceName, expectedResult.CountryName).
			Return(expectedResult, nil)

		service := locality.NewService(mockRepo)
		result, err := service.Create(context.Background(), "6700", "Gru", "SP", "BRA")

		assert.NoError(t, err)
		assert.Equal(t, result, expectedResult)
	})

	t.Run("Deve retornar erro quando o zipcode já existir", func(t *testing.T) {

		mockRepo := mocks.NewRepository(t)

		localityList := []locality.Locality{
			{1, "6700", "Gru", "SP", "BRA"},
			{2, "9999", "Rio", "RJ", "BRA"},
		}

		mockRepo.On("GetAll", context.Background()).Return(localityList, nil)

		service := locality.NewService(mockRepo)
		result, err := service.Create(context.Background(), "6700", "Gru", "SP", "BRA")

		assert.Error(t, err)
		assert.Equal(t, result, locality.Locality{})
	})

	t.Run("Deve retornar erro ao consultar lista de localities", func(t *testing.T) {

		mockRepo := mocks.NewRepository(t)

		mockRepo.On("GetAll", context.Background()).Return([]locality.Locality{}, fmt.Errorf("error"))

		service := locality.NewService(mockRepo)
		result, err := service.Create(context.Background(), "6700", "Gru", "SP", "BRA")

		assert.Error(t, err)
		assert.Equal(t, result, locality.Locality{})
	})
}

func TestService_GetById(t *testing.T) {
	t.Run("Deve retornar uma locality com sucesso", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)
		localityOne := locality.Locality{1, "6700", "Gru", "SP", "BRA"}

		mockRepo.On("GetById", context.Background(), 1).Return(localityOne, nil)

		service := locality.NewService(mockRepo)
		result, err := service.GetById(context.Background(), 1)

		assert.NoError(t, err)
		assert.Equal(t, result, localityOne)
	})

	t.Run("Deve retornar erro locality", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)

		mockRepo.On("GetById", context.Background(), 1).
			Return(locality.Locality{}, fmt.Errorf("id does not exists"))

		service := locality.NewService(mockRepo)
		result, err := service.GetById(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, result, locality.Locality{})
	})
}

func TestService_GetAll(t *testing.T) {
	t.Run("Deve retornar lista locality com sucesso", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)

		expectedResult := []locality.Locality{
			{1, "6700", "Gru", "SP", "BRA"},
			{2, "9999", "Rio", "RJ", "BRA"},
		}

		mockRepo.On("GetAll", context.Background()).Return(expectedResult, nil)

		service := locality.NewService(mockRepo)
		result, err := service.GetAll(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, result, expectedResult)
	})

	t.Run("Deve retornar erro", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)

		mockRepo.On("GetAll", context.Background()).Return([]locality.Locality{}, fmt.Errorf("error"))

		service := locality.NewService(mockRepo)
		result, err := service.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, result, []locality.Locality{})
	})
}
