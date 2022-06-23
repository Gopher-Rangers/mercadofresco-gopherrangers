package products_test

import (
		"fmt"
		"testing"

		"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
		"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store/mocks"

		"github.com/stretchr/testify/assert"
)

func TestRepositoryDelete(t *testing.T) {
		t.Run("delete_non_existent", func(t *testing.T) {
				mockStore := mocks.NewStore(t)
				repository := products.NewRepository(mockStore)

				var ps []products.Product
				mockStore.On("Read", &ps).Return(nil)
				err := repository.Delete(1)
				assert.Equal(t, err, fmt.Errorf("produto 1 não encontrado"))
		})
}

func TestRepositoryFindById(t *testing.T) {
		t.Run("find_by_id_non_existent", func(t *testing.T) {
				mockStore := mocks.NewStore(t)
				repository := products.NewRepository(mockStore)

				var ps []products.Product
				mockStore.On("Read", &ps).Return(nil)
				_, err := repository.GetById(1)
				assert.Equal(t, err, fmt.Errorf("produto 1 não encontrado"))
		})
}

func TestRepositoryGetAll(t *testing.T) {
		t.Run("find_all_empty", func(t *testing.T) {
				mockStore := mocks.NewStore(t)
				repository := products.NewRepository(mockStore)

				var ps []products.Product
				mockStore.On("Read", &ps).Return(nil)
				prod, err := repository.GetAll()
				assert.Equal(t, prod, []products.Product([]products.Product(nil)))
				assert.Nil(t, err)
		})
}

func TestRepositoryUpdate(t *testing.T) {
		t.Run("update_non_existent", func(t *testing.T) {
				mockStore := mocks.NewStore(t)
				repository := products.NewRepository(mockStore)

				var ps []products.Product
				mockStore.On("Read", &ps).Return(nil)
				prod := createProductsArray()
				_, err := repository.Update(prod[0], 1)
				assert.Equal(t, err, fmt.Errorf("produto 1 não encontrado"))
		})
}

func TestLatsId(t *testing.T) {
		t.Run("last_id_empty", func(t *testing.T) {
				mockStore := mocks.NewStore(t)
				repository := products.NewRepository(mockStore)

				var ps []products.Product
				mockStore.On("Read", &ps).Return(nil)
				id, err := repository.LastID()
				assert.Equal(t, id, 0)
				assert.Nil(t, err)
		})
}
