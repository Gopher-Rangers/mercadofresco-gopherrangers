package products_test

import (
	"fmt"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product/mocks"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	t.Run("delete_non_existent", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		mockRepository.On("Delete", 9).Return(fmt.Errorf("produto não encontrado"))

		service := products.NewService(mockRepository)
		err := service.Delete(9)

		assert.Equal(t, fmt.Errorf("produto não encontrado"), err)
	})
}
