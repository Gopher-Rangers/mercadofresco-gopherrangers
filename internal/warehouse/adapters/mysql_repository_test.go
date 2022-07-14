package adapters_test

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
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

	defer db.Close()

	repository := adapters.NewMySqlRepository(db)

	t.Run("Deve retornar todas as Warehouses, se a query estiver correta.", func(t *testing.T) {

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

	t.Run("Deve retornar um Warehouse vazio, se a query estiver incorreta.", func(t *testing.T) {

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM warehouse`)).WillReturnError(sql.ErrNoRows)

		result := repository.GetAll()

		expected := []domain.Warehouse{}

		assert.Equal(t, expected, result)

	})
}

func Test_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New() // cria mock banco de dados..

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repository := adapters.NewMySqlRepository(db)

	t.Run("Deve retornar um Warehouse, se a query estiver correta.", func(t *testing.T) {

		row := sqlmock.NewRows([]string{
			"id", "warehouse_code", "address", "telephone", "locality_id",
		}).AddRow(
			validWarehouse.ID,
			validWarehouse.WarehouseCode,
			validWarehouse.Address,
			validWarehouse.Telephone,
			validWarehouse.LocalityID,
		)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM warehouse WHERE id=?`)).WillReturnRows(row)

		result, err := repository.GetByID(validWarehouse.ID)

		expected := validWarehouse

		assert.NoError(t, err)
		assert.Equal(t, expected, result)

	})

}

func Test_CreateWarehouse(t *testing.T) {
	db, mock, err := sqlmock.New() // cria mock do banco de dados

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repository := adapters.NewMySqlRepository(db)

	t.Run("Deve preparar uma query e executar corretamente.", func(t *testing.T) {

		mock.ExpectPrepare("INSERT INTO ").ExpectExec().WithArgs(validWarehouse.WarehouseCode, validWarehouse.Address, validWarehouse.Telephone, validWarehouse.LocalityID).WillReturnResult(sqlmock.NewResult(1, 1))

		_, err = repository.CreateWarehouse(validWarehouse.WarehouseCode, validWarehouse.Address, validWarehouse.Telephone, validWarehouse.LocalityID)

		err := mock.ExpectationsWereMet() // Verifica essas expectativas anteriormente foram cumpridas com sucesso.

		assert.Nil(t, err)
	})

	t.Run("Deve retornar um erro ao preparar uma query e executa-la.", func(t *testing.T) {

		mock.ExpectPrepare("INSERT INTO ").WillReturnError(fmt.Errorf("erro ao preparar a query"))

		_, err = repository.CreateWarehouse(validWarehouse.WarehouseCode, validWarehouse.Address, validWarehouse.Telephone, validWarehouse.LocalityID)

		assert.NotNil(t, err)

	})

	t.Run("Deve retornar um erro ao obter o id criado.", func(t *testing.T) {

		mock.ExpectPrepare("INSERT INTO ").ExpectExec().WithArgs(validWarehouse.WarehouseCode, validWarehouse.Address, validWarehouse.Telephone, validWarehouse.LocalityID).WillReturnResult(driver.ResultNoRows)

		_, err = repository.CreateWarehouse(validWarehouse.WarehouseCode, validWarehouse.Address, validWarehouse.Telephone, validWarehouse.LocalityID)

		assert.Error(t, err)

	})

	t.Run("Deve retornar um erro ao executar a query", func(t *testing.T) {

		mock.ExpectPrepare("INSERT INTO ").ExpectExec().WithArgs(validWarehouse.WarehouseCode, validWarehouse.Address, validWarehouse.Telephone, validWarehouse.LocalityID)

		_, err = repository.CreateWarehouse(validWarehouse.WarehouseCode, validWarehouse.Address, validWarehouse.Telephone, validWarehouse.LocalityID)

		assert.Error(t, err)

	})
}

func Test_UpdateWarehouse(t *testing.T) {
	db, mock, err := sqlmock.New() // cria mock do banco de dados

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repository := adapters.NewMySqlRepository(db)

	t.Run("Deve preparar uma query e executar corretamente.", func(t *testing.T) {

		row := sqlmock.NewRows([]string{
			"id", "warehouse_code", "address", "telephone", "locality_id",
		}).AddRow(
			validWarehouse.ID,
			validWarehouse.WarehouseCode,
			validWarehouse.Address,
			validWarehouse.Telephone,
			validWarehouse.LocalityID,
		)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM warehouse WHERE id=?")).WillReturnRows(row)

		mock.ExpectPrepare("UPDATE warehouse SET").ExpectExec().WithArgs(validWarehouse.WarehouseCode, validWarehouse.ID).WillReturnResult(sqlmock.NewResult(1, 1))

		repository.UpdatedWarehouseID(validWarehouse.ID, validWarehouse.WarehouseCode)

		err = mock.ExpectationsWereMet() // Verifica se as expectativas anteriormente foram cumpridas com sucesso.

		assert.Nil(t, err)
	})

	t.Run("Deve retornar um erro quando o prepare retornar um erro.", func(t *testing.T) {

		row := sqlmock.NewRows([]string{
			"id", "warehouse_code", "address", "telephone", "locality_id",
		}).AddRow(
			validWarehouse.ID,
			validWarehouse.WarehouseCode,
			validWarehouse.Address,
			validWarehouse.Telephone,
			validWarehouse.LocalityID,
		)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM warehouse WHERE id=?")).WillReturnRows(row)

		mock.ExpectPrepare("UPDATE warehouse SET").WillReturnError(fmt.Errorf("erro ao preparar a query"))

		result, err := repository.UpdatedWarehouseID(validWarehouse.ID, validWarehouse.WarehouseCode)

		assert.NotNil(t, err)
		assert.Equal(t, domain.Warehouse{}, result)
	})

	t.Run("Deve retornar um erro ao executar a query", func(t *testing.T) {

		row := sqlmock.NewRows([]string{
			"id", "warehouse_code", "address", "telephone", "locality_id",
		}).AddRow(
			validWarehouse.ID,
			validWarehouse.WarehouseCode,
			validWarehouse.Address,
			validWarehouse.Telephone,
			validWarehouse.LocalityID,
		)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM warehouse WHERE id=?")).WillReturnRows(row)

		mock.ExpectPrepare("UPDATE warehouse SET").ExpectExec().WithArgs(validWarehouse.WarehouseCode, validWarehouse.ID).WillReturnError(fmt.Errorf("erro ao executar a query"))

		_, err = repository.UpdatedWarehouseID(validWarehouse.ID, validWarehouse.WarehouseCode)

		assert.Error(t, err)

	})
}

func Test_DeleteWarehouse(t *testing.T) {
	db, mock, err := sqlmock.New() // cria mock do banco de dados

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repository := adapters.NewMySqlRepository(db)

	t.Run("Deve preparar uma query e executar corretamente.", func(t *testing.T) {

		mock.ExpectPrepare("DELETE FROM warehouse").ExpectExec().WithArgs(validWarehouse.ID).WillReturnResult(sqlmock.NewResult(1, 1))

		err = repository.DeleteWarehouse(validWarehouse.ID)

		err := mock.ExpectationsWereMet()

		assert.Nil(t, err)
	})

	t.Run("Deve retornar um erro ao preparar uma query e executa-la.", func(t *testing.T) {

		mock.ExpectPrepare("DELETE FROM warehouse WHERE id=?").WillReturnError(fmt.Errorf("erro ao preparar a query"))

		err = repository.DeleteWarehouse(validWarehouse.ID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "erro ao preparar a query: erro ao preparar a query")

	})

	t.Run("Deve retornar um erro se não existir o id no banco de dados.", func(t *testing.T) {

		mock.ExpectPrepare("DELETE * FROM warehouse WHERE id=?").ExpectExec().WithArgs(validWarehouse.ID).WillReturnResult(sqlmock.NewResult(1, 0))

		result := repository.DeleteWarehouse(validWarehouse.ID)

		err := mock.ExpectationsWereMet()

		assert.Error(t, result)
		assert.Nil(t, err)
	})

	t.Run("Deve retornar um erro ao executar a query", func(t *testing.T) {

		mock.ExpectPrepare("DELETE FROM warehouse").ExpectExec().WithArgs(validWarehouse.ID).WillReturnError(fmt.Errorf("xablau"))

		err := repository.DeleteWarehouse(validWarehouse.ID)

		assert.Error(t, err)
		assert.EqualError(t, err, "erro ao executar query: xablau")
	})
}
