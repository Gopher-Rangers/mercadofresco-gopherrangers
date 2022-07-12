package adapters_test

import (
	"database/sql/driver"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/adapters"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
	"github.com/stretchr/testify/assert"
)

var validCarry = domain.Carry{
	ID:         1,
	Cid:        "CID#5",
	Name:       "mercado-livre",
	Address:    "Criciuma, 666",
	Telephone:  "99999999",
	LocalityID: 2,
}

func Test_CreateCarry(t *testing.T) {
	db, mock, err := sqlmock.New() // cria mock do banco de dados

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repository := adapters.NewMySqlCarryRepository(db)

	t.Run("Deve preparar uma query e executar corretamente.", func(t *testing.T) {

		mock.ExpectPrepare("INSERT INTO carriers").ExpectExec().WithArgs("CID#5", "mercado-livre", "Criciuma, 666", "99999999", 2).WillReturnResult(sqlmock.NewResult(1, 1))

		repository.CreateCarry(validCarry)

		err := mock.ExpectationsWereMet()

		assert.Nil(t, err)
	})

	t.Run("Deve retornar um erro ao preparar uma query e executa-la.", func(t *testing.T) {

		mock.ExpectPrepare("INSERT INTO carriers").WillReturnError(fmt.Errorf("erro ao preparar a query"))

		_, err = repository.CreateCarry(validCarry)

		assert.NotNil(t, err)

	})

	t.Run("Deve retornar um erro ao obter o id criado.", func(t *testing.T) {

		mock.ExpectPrepare("INSERT INTO carriers").ExpectExec().WithArgs("CID#5", "mercado-livre", "Criciuma, 666", "99999999", 2).WillReturnResult(driver.ResultNoRows)

		_, err = repository.CreateCarry(validCarry)

		assert.Error(t, err)

	})

	t.Run("Deve retornar um erro ao executar a query", func(t *testing.T) {

		mock.ExpectPrepare("INSERT INTO carriers").ExpectExec().WithArgs("CID#5", "mercado-livre", "Criciuma, 666", "99999999")

		_, err = repository.CreateCarry(validCarry)

		assert.Error(t, err)

	})
}

func Test_GetCarryByCid(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repository := adapters.NewMySqlCarryRepository(db)

	t.Run("Deve retornar uma Carry com o cid pesquisado", func(t *testing.T) {

		row := sqlmock.NewRows([]string{
			"id", "cid", "company_name", "address", "telephone", "locality_id",
		}).AddRow(
			validCarry.ID,
			validCarry.Cid,
			validCarry.Name,
			validCarry.Address,
			validCarry.Telephone,
			validCarry.LocalityID,
		)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM carriers WHERE cid=?")).WithArgs(validCarry.Cid).WillReturnRows(row)

		result, err := repository.GetCarryByCid(validCarry.Cid)

		assert.NoError(t, err)
		assert.Equal(t, validCarry, result)

	})

	t.Run("Deve retornar um erro caso o `cid` não for encontrado.", func(t *testing.T) {

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM carriers WHERE cid=?")).WithArgs(validCarry.Cid).WillReturnError(fmt.Errorf("a carry com esse `cid`: %s não foi encontrada", validCarry.Cid))

		_, err := repository.GetCarryByCid(validCarry.Cid)

		assert.Error(t, err)
	})
}
