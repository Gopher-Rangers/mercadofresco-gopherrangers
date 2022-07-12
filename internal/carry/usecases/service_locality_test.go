package usecases_test

import (
	"fmt"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases/mock/mock_repository_locality"
	"github.com/stretchr/testify/assert"
)

func makeValidDBLocality() domain.Locality {
	return domain.Locality{
		ID:    1,
		Name:  "Florianopolis",
		Count: 5,
	}
}

func Test_GetCarryLocalityByID(t *testing.T) {
	t.Run("Deve retornar um Locality Cary vazio e um erro, se um elemento com o id especifíco não existir.", func(t *testing.T) {
		mockRepository := mock_repository_locality.NewRepositoryLocality(t)
		service := usecases.NewServiceLocality(mockRepository)

		mockRepository.On("GetCarryLocalityByID", 1).Return(domain.Locality{}, fmt.Errorf("o id: %d não foi encontrado", 1))

		result, err := service.GetCarryLocalityByID(1)

		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, result, domain.Locality{})
		assert.Empty(t, result)
	})

	t.Run("Deve retornar um LocalityCarry, com o id solicitado.", func(t *testing.T) {
		mockRepository := mock_repository_locality.NewRepositoryLocality(t)
		service := usecases.NewServiceLocality(mockRepository)

		expected := makeValidDBLocality()

		mockRepository.On("GetCarryLocalityByID", 1).Return(expected, nil)

		result, err := service.GetCarryLocalityByID(1)

		assert.Nil(t, err)
		assert.Equal(t, result, expected)
		assert.NotEmpty(t, result)
	})
}
