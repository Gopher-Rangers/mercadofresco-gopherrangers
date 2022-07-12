package adapters_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/adapters"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
	"github.com/stretchr/testify/assert"
)

var validLocalityCarry = domain.Locality{
	ID:    1,
	Name:  "Florianopolis",
	Count: 5,
}

func Test_GetAllCarriesLocality(t *testing.T) {

	db, mock, err := sqlmock.New() // cria mock do banco de dados

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repository := adapters.NewMySqlLocalityRepository(db)

	t.Run("Deve retornar todas as Carry Locality, se a query estiver correta", func(t *testing.T) {

		row := sqlmock.NewRows([]string{
			"id", "locality_name", "carries_count",
		}).AddRow(
			validLocalityCarry.ID,
			validLocalityCarry.Name,
			validLocalityCarry.Count,
		)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT ca.locality_id, lo.locality_name, COUNT(*) AS carries_count 
		FROM carriers AS ca
		INNER JOIN localities AS lo
		ON ca.locality_id = lo.id
		GROUP BY ca.locality_id
		`)).WillReturnRows(row)

		result, err := repository.GetAllCarriesLocality()

		expected := []domain.Locality{
			{
				ID:    1,
				Name:  "Florianopolis",
				Count: 5,
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, expected, result)

	})

	t.Run("Deve retornar um erro se a query estiver incorreta.", func(t *testing.T) {

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT ca.locality_id, lo.locality_name, COUNT(*) AS carries_count 
		FROM carriers AS ca
		INNER JOIN localities AS la
		ON ca.locality_id = lo.id
		GROUP BY ca.locality_id
		`)).WillReturnError(sql.ErrNoRows)

		result, err := repository.GetAllCarriesLocality()

		expected := []domain.Locality{}

		assert.Error(t, err)
		assert.Equal(t, expected, result)

	})
}
