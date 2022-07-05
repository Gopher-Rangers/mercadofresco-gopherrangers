package productbatch_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	productbatch "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_batch"
	"github.com/stretchr/testify/assert"
)

func InitTest(t *testing.T) (sqlmock.Sqlmock, productbatch.Repository) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	mockRepository := productbatch.NewRepository(db)

	return mock, mockRepository
}

func TestRepositoryCreate(t *testing.T) {
	mock, mockRepository := InitTest(t)

	exp := productbatch.ProductBatch{1, 111, 200, 20, "2022-04-04", 10, "2020-04-04", 10, 5, 1, 1}

	t.Run("create_ok", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(productbatch.SqlCreateBatch)).WithArgs(&exp.BatchNumber, &exp.CurQuantity,
			&exp.CurTemperature, &exp.DueDate, &exp.InitialQuantity, &exp.ManufactDate, &exp.ManufactHour,
			&exp.MinTemperature, &exp.ProductTypeID, &exp.SectionID).WillReturnResult(sqlmock.NewResult(15, 1))

		pb, err := mockRepository.Create(exp)

		exp.ID = 15
		assert.Equal(t, exp, pb)
		assert.NoError(t, err)
	})

	t.Run("create_fail_exec", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(productbatch.SqlCreateBatch)).WithArgs(&exp.BatchNumber, &exp.CurQuantity,
			&exp.CurTemperature, &exp.DueDate, &exp.InitialQuantity, &exp.ManufactDate, &exp.ManufactHour,
			&exp.MinTemperature, &exp.ProductTypeID, &exp.SectionID).WillReturnError(sql.ErrNoRows)

		pb, err := mockRepository.Create(exp)

		assert.Equal(t, productbatch.ProductBatch{}, pb)
		assert.Error(t, err)
		assert.Equal(t, errors.New("sql: no rows in result set"), err)
	})

	t.Run("create_fail_zero_rows_affected", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(productbatch.SqlCreateBatch)).WithArgs(&exp.BatchNumber, &exp.CurQuantity,
			&exp.CurTemperature, &exp.DueDate, &exp.InitialQuantity, &exp.ManufactDate, &exp.ManufactHour,
			&exp.MinTemperature, &exp.ProductTypeID, &exp.SectionID).WillReturnResult(sqlmock.NewResult(1, 0))

		pb, err := mockRepository.Create(exp)

		assert.Equal(t, productbatch.ProductBatch{}, pb)
		assert.Error(t, err)
		assert.Equal(t, errors.New("sql: rows not affected"), err)
	})
}
