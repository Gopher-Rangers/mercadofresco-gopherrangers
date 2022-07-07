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

const (
	WithValue = true
	FailScan  = false
)

func CreateReportArray() []productbatch.Report {
	var exp = []productbatch.Report{
		{1, 22, 5},
		{3, 483, 28},
		{5, 7843, 90},
	}

	return exp
}

func MockRowsArray(flag bool) *sqlmock.Rows {
	exp := CreateReportArray()

	rows := sqlmock.NewRows([]string{"section_id", "section_number", "products_count"})

	if !flag {
		rows.AddRow("", "", "")
		return rows
	}

	for i := range exp {
		rows.AddRow(exp[i].SecID, exp[i].SecNum, exp[i].ProdCount)
	}

	return rows
}

func MockRow(flag bool) *sqlmock.Rows {
	exp := CreateReportArray()[0]

	rows := sqlmock.NewRows([]string{"section_id", "section_number", "products_count"})

	if !flag {
		rows.AddRow("", "", "")
		return rows
	}

	rows.AddRow(exp.SecID, exp.SecNum, exp.ProdCount)

	return rows
}

func InitTestRepository(t *testing.T) (sqlmock.Sqlmock, productbatch.Repository) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	mockRepository := productbatch.NewRepository(db)

	return mock, mockRepository
}

func TestRepositoryCreate(t *testing.T) {
	mock, mockRepository := InitTestRepository(t)

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

func TestReport(t *testing.T) {
	mock, mockRepository := InitTestRepository(t)
	exp := CreateReportArray()
	rows := MockRowsArray(WithValue)

	t.Run("report_ok", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(productbatch.SqlReportBatchAll)).WillReturnRows(rows)

		pb, err := mockRepository.Report()

		assert.NoError(t, err)
		assert.Equal(t, exp, pb)
	})

	t.Run("report_fail_query", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(productbatch.SqlReportBatchAll)).WillReturnError(sql.ErrNoRows)

		pb, err := mockRepository.Report()

		assert.Equal(t, []productbatch.Report{}, pb)
		assert.Error(t, err)
		assert.Equal(t, errors.New("sql: no rows in result set"), err)
	})

	t.Run("report_fail_scan", func(t *testing.T) {
		rows := MockRowsArray(FailScan)
		mock.ExpectQuery(regexp.QuoteMeta(productbatch.SqlReportBatchAll)).WillReturnRows(rows)

		pb, err := mockRepository.Report()

		assert.Equal(t, []productbatch.Report{}, pb)
		assert.Error(t, err)
	})
}

func TestReportByID(t *testing.T) {
	mock, mockRepository := InitTestRepository(t)
	exp := CreateReportArray()[0]
	row := MockRow(WithValue)

	t.Run("report_ok", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(productbatch.SqlReportBatchByID)).WillReturnRows(row)

		pb, err := mockRepository.ReportByID(1)

		assert.NoError(t, err)
		assert.Equal(t, exp, pb)
	})

	t.Run("report_fail_scan", func(t *testing.T) {
		row := MockRow(FailScan)
		mock.ExpectQuery(regexp.QuoteMeta(productbatch.SqlReportBatchAll)).WillReturnRows(row)

		pb, err := mockRepository.ReportByID(1)

		assert.Equal(t, productbatch.Report{}, pb)
		assert.Error(t, err)
	})
}
