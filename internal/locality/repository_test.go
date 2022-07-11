package locality_test

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"regexp"
	"testing"
)

func TestRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	mockLocality := []locality.Locality{{Id: 1, ZipCode: "6700", LocalityName: "Gru", ProvinceName: "SP", CountryName: "BRA"},
		{Id: 1, ZipCode: "6701", LocalityName: "Manaus", ProvinceName: "Amazonia", CountryName: "BRA"}}

	t.Run("Deve retornar lista de localities com sucesso", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "zip_code", "locality_name", "province_name", "country_name",
		}).AddRow(mockLocality[0].Id, mockLocality[0].ZipCode, mockLocality[0].LocalityName, mockLocality[0].ProvinceName, mockLocality[0].CountryName).
			AddRow(mockLocality[1].Id, mockLocality[1].ZipCode, mockLocality[1].LocalityName, mockLocality[1].ProvinceName, mockLocality[1].CountryName)

		query := "SELECT \\* FROM localities"

		mock.ExpectQuery(query).WillReturnRows(rows)

		localityRepo := locality.NewMariaDBRepository(db)

		localityList, _ := localityRepo.GetAll(context.Background())

		assert.Equal(t, localityList, mockLocality)
	})
	t.Run("Deve retornar erro no Scan", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{
			"id", "zip_code", "locality_name", "province_name", "country_name",
		}).AddRow("", "", "", "", "")

		query := "SELECT \\* FROM localities"

		mock.ExpectQuery(query).WillReturnRows(rows)

		localityRepo := locality.NewMariaDBRepository(db)

		_, err = localityRepo.GetAll(context.Background())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Deve retorar erro ao realizar o select", func(t *testing.T) {

		query := "SELECT \\* FROM localities"
		mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)

		localityRepo := locality.NewMariaDBRepository(db)

		_, err = localityRepo.GetAll(context.Background())
		assert.Error(t, err)
	})
}

func TestRepository_GetById(t *testing.T) {

	t.Run("Deve retornar a locality com sucesso", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		localityOne := locality.Locality{Id: 1, ZipCode: "6700", LocalityName: "Gru", ProvinceName: "SP", CountryName: "BRA"}

		rows := sqlmock.NewRows([]string{
			"id", "zip_code", "locality_name", "province_name", "country_name",
		}).AddRow(localityOne.Id, localityOne.ZipCode, localityOne.LocalityName, localityOne.ProvinceName, localityOne.CountryName)

		mock.ExpectQuery("SELECT").WithArgs(1).WillReturnRows(rows)

		localityRepo := locality.NewMariaDBRepository(db)
		result, err := localityRepo.GetById(context.Background(), 1)

		assert.NoError(t, err)

		assert.Equal(t, result, localityOne)
	})

	t.Run("Deve retornar error quando o id não existir", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectQuery("SELECT \\* FROM localities WHERE id=?").WithArgs(2).
			WillReturnError(fmt.Errorf("id does not exists"))

		localityRepo := locality.NewMariaDBRepository(db)
		_, err = localityRepo.GetById(context.Background(), 2)

		assert.Error(t, err)
	})

	t.Run("Deve retornar error ao executar o Scan", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id", "zip_code", "locality_name", "province_name", "country_name",
		}).AddRow("", "", "", "", "")

		mock.ExpectQuery("SELECT").WithArgs(1).WillReturnRows(rows)

		localityRepo := locality.NewMariaDBRepository(db)
		_, err = localityRepo.GetById(context.Background(), 1)

		assert.Error(t, err)
	})

	t.Run("Deve retorar erro ao realizar o select", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		query := "SELECT \\*  FROM localities WHERE id=?"
		mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)

		localityRepo := locality.NewMariaDBRepository(db)
		_, err = localityRepo.GetAll(context.Background())
		assert.Error(t, err)
	})
}

func TestMariaDBRepository_Create(t *testing.T) {
	t.Run("Deve criar uma locality com sucesso", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		expected := locality.Locality{Id: 1, ZipCode: "6700", LocalityName: "Gru", ProvinceName: "SP", CountryName: "BRA"}
		input := locality.Locality{ZipCode: "6700", LocalityName: "Gru", ProvinceName: "SP", CountryName: "BRA"}

		mock.ExpectExec("INSERT INTO localities").WithArgs(input.ZipCode, input.LocalityName, input.ProvinceName, input.CountryName).
			WillReturnResult(sqlmock.NewResult(1, 1))

		localityRepo := locality.NewMariaDBRepository(db)
		result, err := localityRepo.Create(context.Background(), input.ZipCode, input.LocalityName, input.ProvinceName, input.CountryName)

		assert.NoError(t, err)
		assert.Equal(t, result, expected)
	})
	t.Run("Deve retornar error no executar query", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		input := locality.Locality{ZipCode: "6700", LocalityName: "Gru", ProvinceName: "SP", CountryName: "BRA"}

		mock.ExpectExec("INSERT INTO locality").WithArgs(input.ZipCode, input.LocalityName, input.ProvinceName, input.CountryName).
			WillReturnError(fmt.Errorf("error"))

		localityRepo := locality.NewMariaDBRepository(db)
		_, err = localityRepo.Create(context.Background(), input.ZipCode, input.LocalityName, input.ProvinceName, input.CountryName)

		assert.Error(t, err)
	})
}
func TestRepository_ReportSellers(t *testing.T) {
	t.Run("Deve retornar o report sellers com sucesso", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mockReportSellers := locality.ReportSeller{LocalityID: 1, LocalityName: "São Paulo", SellersCount: 2}

		rows := sqlmock.NewRows([]string{
			"locality_id", "locality_name", "sellers_count",
		}).AddRow(mockReportSellers.LocalityID, mockReportSellers.LocalityName, mockReportSellers.SellersCount)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT l.id, l.locality_name, COUNT(seller.id) FROM localities l LEFT JOIN seller ON l.id=seller.locality_id WHERE l.id = ?")).
			WithArgs(1).WillReturnRows(rows)

		localityRepo := locality.NewMariaDBRepository(db)
		result, err := localityRepo.ReportSellers(context.Background(), 1)

		assert.NoError(t, err)
		assert.Equal(t, result, mockReportSellers)
	})

	t.Run("Deve retornar erro ao executar a query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta("SELECT l.id, l.locality_name, COUNT(seller.id) FROM localities l LEFT JOIN seller ON l.id=seller.locality_id WHERE l.id = ?")).
			WithArgs(1).WillReturnError(fmt.Errorf("error"))

		localityRepo := locality.NewMariaDBRepository(db)
		_, err = localityRepo.ReportSellers(context.Background(), 1)

		assert.Error(t, err)

	})
}
