package adapters_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/adapters"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/domain"
	"github.com/stretchr/testify/assert"
)

var validWarehouse = domain.Warehouse{
	ID:            1,
	WarehouseCode: "caju",
	Address:       "Rua das Rendeiras",
	Telephone:     "333333",
	LocalityID:    1,
}

func Test_GetAll(t *testing.T) {

	db, mock, err := sqlmock.New() // cria mock banco de dados..

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repository := adapters.NewMySqlRepository(db)

	t.Run("Deve retornar todas as Warehouses, se a query estiver correta", func(t *testing.T) {

		row := sqlmock.NewRows([]string{
			"id", "warehouse_code", "address", "telephone", "locality_id",
		}).AddRow(
			validWarehouse.ID,
			validWarehouse.WarehouseCode,
			validWarehouse.Address,
			validWarehouse.Telephone,
			validWarehouse.LocalityID,
		)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM warehouse
		`)).WillReturnRows(row)

		result := repository.GetAll()

		expected := []domain.Warehouse{
			{
				ID:            validWarehouse.ID,
				WarehouseCode: validWarehouse.WarehouseCode,
				Address:       validWarehouse.Address,
				Telephone:     validWarehouse.Telephone,
				LocalityID:    validWarehouse.LocalityID,
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, expected, result)

	})
}
