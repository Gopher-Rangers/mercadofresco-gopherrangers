package usecases_test

import (
	"fmt"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases/mock/mock_repository_carry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeValidDBCarry() domain.Carry {
	return domain.Carry{
		ID:         1,
		Cid:        "CID#5",
		Name:       "mercado-livre",
		Address:    "Criciuma, 666",
		Telephone:  "99999999",
		LocalityID: 2,
	}
}

func Test_CreateCarry(t *testing.T) {
	t.Run("Deve conter os campos necessários para ser criado uma Carry.", func(t *testing.T) {
		mockRepository := mock_repository_carry.NewRepositoryCarry(t)
		service := usecases.NewServiceCarry(mockRepository)

		data := domain.Carry{
			Cid:        "CID#5",
			Name:       "mercado-livre",
			Address:    "Criciuma, 666",
			Telephone:  "99999999",
			LocalityID: 2,
		}

		expected := makeValidDBCarry()

		mockRepository.On("GetCarryByCid", mock.AnythingOfType("string")).Return(domain.Carry{},
			fmt.Errorf("o carry com esse `cid`: %s não foi encontrado", expected.Cid))

		mockRepository.On("CreateCarry", data).Return(expected, nil)

		result, err := service.CreateCarry(data)

		assert.Nil(t, err)
		assert.Equal(t, result, expected)

	})

	t.Run("Deve retornar uma Carry vazia se já existir um `cid` no banco de dados.", func(t *testing.T) {
		mockRepository := mock_repository_carry.NewRepositoryCarry(t)
		service := usecases.NewServiceCarry(mockRepository)

		data := domain.Carry{
			ID:         1,
			Cid:        "CID#5",
			Name:       "mercado-livre",
			Address:    "Criciuma, 666",
			Telephone:  "99999999",
			LocalityID: 2,
		}

		carry := makeValidDBCarry()

		expected := domain.Carry{}

		mockRepository.On("GetCarryByCid", mock.AnythingOfType("string")).Return(carry, nil)

		result, err := service.CreateCarry(data)

		assert.Equal(t, result, expected)
		assert.Equal(t, err, fmt.Errorf("o `cid` já está em uso"))
		assert.Error(t, err)

	})

	t.Run("Deve retornar um erro caso CreateCarry, retorne um error", func(t *testing.T) {
		mockRepository := mock_repository_carry.NewRepositoryCarry(t)
		service := usecases.NewServiceCarry(mockRepository)

		data := domain.Carry{
			ID:         1,
			Cid:        "CID#5",
			Name:       "mercado-livre",
			Address:    "Criciuma, 666",
			Telephone:  "99999999",
			LocalityID: 2,
		}

		expected := domain.Carry{}

		mockRepository.On("GetCarryByCid", mock.AnythingOfType("string")).Return(expected, fmt.Errorf("a Carry com esse `cid`: %s não foi encontrada", data.Cid))

		mockRepository.On("CreateCarry", data).Return(expected, fmt.Errorf("erro ao preparar a query"))

		result, err := service.CreateCarry(data)

		assert.Equal(t, result, expected)
		assert.Equal(t, err, fmt.Errorf("erro ao preparar a query"))
		assert.Error(t, err)

	})
}
