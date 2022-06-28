package test

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Update(t *testing.T) {
	t.Run("Deve realizar o update dos dados do seller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mockSellers := []seller.Seller{{Id: 1, CompanyId: 1, CompanyName: "Meli", Address: "Osasco", Telephone: "99999"},
			{Id: 2, CompanyId: 2, CompanyName: "Lojinha", Address: "Barueri", Telephone: "000000"}}

		stmt := mock.ExpectPrepare("UPDATE seller")
		stmt.ExpectExec().WithArgs(3, "Melii", "Osascão", "9999", 1).WillReturnResult(sqlmock.NewResult(1, 1))

		sellerRepo := seller.NewMariaDBRepository(db)
		result, err := sellerRepo.Update(3, "Melii", "Osascão", "9999", mockSellers[0])

		assert.Equal(t, "Melii", result.CompanyName)
	})

	t.Run("Deve retornar erro ao executar a query com parametro errado", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mockSellers := []seller.Seller{{Id: 1, CompanyId: 1, CompanyName: "Meli", Address: "Osasco", Telephone: "99999"},
			{Id: 2, CompanyId: 2, CompanyName: "Lojinha", Address: "Barueri", Telephone: "000000"}}

		stmt := mock.ExpectPrepare("UPDATE seller")
		stmt.ExpectExec().WithArgs(2, "Melii", "Osascão", "9999", 1).WillReturnError(fmt.Errorf("error"))

		sellerRepo := seller.NewMariaDBRepository(db)
		_, err = sellerRepo.Update(2, "Melii", "Osascão", "9999", mockSellers[1])

		assert.Error(t, err)
	})
}

func Test_Delete(t *testing.T) {
	t.Run("Deve excluir o seller se existir", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		stmt := mock.ExpectPrepare("DELETE FROM seller")
		stmt.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

		sellerRepo := seller.NewMariaDBRepository(db)
		err = sellerRepo.Delete(1)

		assert.Nil(t, err)
	})

	t.Run("Deve retornar erro quando a query estiver errada", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		stmt := mock.ExpectPrepare("DELET FROM seller")
		stmt.ExpectExec().WithArgs(1).WillReturnError(fmt.Errorf("error"))

		sellerRepo := seller.NewMariaDBRepository(db)
		err = sellerRepo.Delete(1)

		assert.Error(t, err)
	})

	t.Run("Deve retornar erro quando ao executar query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		stmt := mock.ExpectPrepare("DELET FROM seller")
		stmt.ExpectExec().WithArgs("1").WillReturnResult(sqlmock.NewResult(1, 1))

		sellerRepo := seller.NewMariaDBRepository(db)
		err = sellerRepo.Delete(1)

		assert.Error(t, err)
	})
}

func Test_Store(t *testing.T) {
	t.Run("Deve criar um seller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		input := seller.Seller{CompanyId: 1, CompanyName: "Meli", Address: "A", Telephone: "9999999"}

		rows := sqlmock.NewRows([]string{"id", "cid", "company_name", "address", "telephone"})

		mock.ExpectQuery("SELECT \\* FROM seller").WillReturnRows(rows)

		stmt := mock.ExpectPrepare("INSERT INTO seller")
		stmt.ExpectExec().WithArgs(input.CompanyId, input.CompanyName, input.Address, input.Telephone).
			WillReturnResult(sqlmock.NewResult(1, 1))

		sellerRepo := seller.NewMariaDBRepository(db)
		result, err := sellerRepo.Create(1, "Meli", "A", "9999999")
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		assert.Equal(t, input.CompanyId, result.CompanyId)
		assert.Equal(t, input.CompanyName, result.CompanyName)
		assert.Equal(t, input.Address, result.Address)
	})

	t.Run("Não deve criar novo seller se o cid já existir", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		input := seller.Seller{CompanyId: 1, CompanyName: "Meli", Address: "A", Telephone: "9999999"}

		rows := sqlmock.NewRows([]string{"id", "cid", "company_name", "address", "telephone"}).AddRow(1, 1, "Meli", "Osasco", "9999999")

		mock.ExpectQuery("SELECT \\* FROM seller").WillReturnRows(rows)

		stmt := mock.ExpectPrepare("INSERT INTO seller")
		stmt.ExpectExec().WithArgs(input.CompanyId, input.CompanyName, input.Address, input.Telephone).WillReturnError(fmt.Errorf("error"))

		mock.ExpectCommit()

		sellerRepo := seller.NewMariaDBRepository(db)
		result, err := sellerRepo.Create(1, "Meli", "A", "9999999")

		assert.Error(t, err)
		assert.Equal(t, seller.Seller{}, result)
		assert.Equal(t, "the cid already exists", err.Error())
	})

	t.Run("Deve retornar erro com sql query errada", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		input := seller.Seller{CompanyId: 19, CompanyName: "Meli", Address: "A", Telephone: "9999999"}

		rows := sqlmock.NewRows([]string{"id", "cid", "company_name", "address", "telephone"})

		mock.ExpectQuery("SELECT \\* FROM seller").WillReturnRows(rows)

		stmt := mock.ExpectPrepare("INSER INTO seller")
		stmt.ExpectExec().WithArgs(input.CompanyId, input.CompanyName, input.Address, input.Telephone).WillReturnError(fmt.Errorf("error"))

		mock.ExpectCommit()

		sellerRepo := seller.NewMariaDBRepository(db)
		_, err = sellerRepo.Create(10, "Meli", "A", "9999999")

		assert.Error(t, err)
	})

	t.Run("Deve retornar erro com input com type errado", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		input := seller.Seller{CompanyId: 1, CompanyName: "Meli", Address: "A", Telephone: "9999999"}

		rows := sqlmock.NewRows([]string{"id", "cid", "company_name", "address", "telephone"})

		mock.ExpectQuery("SELECT \\* FROM seller").WillReturnRows(rows)

		stmt := mock.ExpectPrepare("INSERT INTO seller")
		stmt.ExpectExec().WithArgs("", input.CompanyName, input.Address, input.Telephone).WillReturnError(fmt.Errorf("error"))

		mock.ExpectCommit()

		sellerRepo := seller.NewMariaDBRepository(db)
		_, err = sellerRepo.Create(12, "Meli", "A", "9999999")

		assert.Error(t, err)
	})

	t.Run("Deve retornar erro ao chamar o GetAll", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		input := seller.Seller{CompanyId: 1, CompanyName: "Meli", Address: "A", Telephone: "9999999"}

		mock.ExpectQuery("SELECT \\* FROM seller").WillReturnError(fmt.Errorf("error"))

		stmt := mock.ExpectPrepare("INSERT INTO seller")
		stmt.ExpectExec().WithArgs(input.CompanyId, input.CompanyName, input.Address, input.Telephone).WillReturnError(fmt.Errorf("error"))

		mock.ExpectCommit()

		sellerRepo := seller.NewMariaDBRepository(db)
		_, err = sellerRepo.Create(12, "Meli", "A", "9999999")

		assert.Error(t, err)
	})

}

func Test_FindOne(t *testing.T) {
	t.Run("Deve retornar um seller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		expectedResult := seller.Seller{Id: 1, CompanyId: 1, CompanyName: "Meli", Address: "Osasco", Telephone: "999999"}

		rows := sqlmock.NewRows([]string{
			"id", "cid", "company_name", "address", "telephone",
		}).AddRow(expectedResult.Id, expectedResult.CompanyId, expectedResult.CompanyName, expectedResult.Address, expectedResult.Telephone)

		query := "SELECT \\*  FROM seller WHERE id=?"
		mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

		sellerRepo := seller.NewMariaDBRepository(db)
		result, err := sellerRepo.GetOne(1)
		assert.NoError(t, err)

		assert.Equal(t, expectedResult.Id, result.Id)
	})

	t.Run("Quando o id não existir, deve retornar um erro", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		query := "SELECT \\*  FROM seller WHERE id=?"
		mock.ExpectQuery(query).WithArgs(2).WillReturnError(fmt.Errorf("the id %d does not exists", 2))

		sellerRepo := seller.NewMariaDBRepository(db)
		result, err := sellerRepo.GetOne(2)

		assert.Error(t, err, "the id 2 does not exists")
		assert.Equal(t, seller.Seller{}, result)
		assert.Equal(t, seller.Seller{}, result)
	})

	t.Run("Deve retornar erro no Scan", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id", "cid", "company_id", "address", "telephone"}).AddRow("", "", "", "", "")

		query := "SELECT \\*  FROM seller WHERE id=?"

		mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

		sellerRepo := seller.NewMariaDBRepository(db)

		_, err = sellerRepo.GetOne(1)

		assert.Error(t, err)
	})

	t.Run("Deve retorar erro ao realizar o select", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		query := "SELECT \\*  FROM seller WHERE id=?"
		mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)

		sellerRepo := seller.NewMariaDBRepository(db)
		_, err = sellerRepo.GetAll()
		assert.Error(t, err)
	})
}

func Test_GetAll(t *testing.T) {
	t.Run("Deve retornar ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mockSellers := []seller.Seller{{Id: 1, CompanyId: 1, CompanyName: "Meli", Address: "Osasco", Telephone: "99999"},
			{Id: 2, CompanyId: 2, CompanyName: "Lojinha", Address: "Barueri", Telephone: "000000"}}

		rows := sqlmock.NewRows([]string{
			"id", "cid", "company_id", "address", "telephone",
		}).AddRow(mockSellers[0].Id, mockSellers[0].CompanyId, mockSellers[0].CompanyName, mockSellers[0].Address, mockSellers[0].Telephone).
			AddRow(mockSellers[1].Id, mockSellers[1].CompanyId, mockSellers[1].CompanyName, mockSellers[1].Address, mockSellers[1].Telephone)

		query := "SELECT \\* FROM seller"

		mock.ExpectQuery(query).WillReturnRows(rows)

		sellerRepo := seller.NewMariaDBRepository(db)

		_, err = sellerRepo.GetAll()

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		//assert.NoError(t, err)
		//
		//assert.Equal(t, result[0].Id, 1)
		//assert.Equal(t, result[1].Id, 2)
	})

	t.Run("Deve retornar erro no Scan", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id", "cid", "company_id", "address", "telephone"}).AddRow("", "", "", "", "")

		query := "SELECT \\* FROM seller"

		mock.ExpectQuery(query).WillReturnRows(rows)

		sellerRepo := seller.NewMariaDBRepository(db)

		_, err = sellerRepo.GetAll()

		assert.Error(t, err)
	})

	t.Run("Deve retorar erro ao realizar o select", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		query := "SELECT \\* FROM seller"
		mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)

		sellerRepo := seller.NewMariaDBRepository(db)
		_, err = sellerRepo.GetAll()
		assert.Error(t, err)
	})
}
