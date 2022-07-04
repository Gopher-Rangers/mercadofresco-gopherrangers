package section_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section"
	"github.com/stretchr/testify/assert"
)

const (
	WithValue = true
	FailScan  = false
)

func InitTest(t *testing.T) (sqlmock.Sqlmock, section.Repository, *sqlmock.Rows) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	mockRepository := section.NewRepository(db)

	rows := mockRow(WithValue)

	return mock, mockRepository, rows
}

func mockRowsArray() *sqlmock.Rows {
	sec := createSectionArray()

	rows := sqlmock.NewRows([]string{
		"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity",
		"minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"})

	for i := range sec {
		rows.AddRow(sec[i].ID, sec[i].SectionNumber, sec[i].CurTemperature, sec[i].MinTemperature,
			sec[i].CurCapacity, sec[i].MinCapacity, sec[i].MaxCapacity, sec[i].WareHouseID, sec[i].ProductTypeID)
	}

	return rows
}

func mockRow(flag bool) *sqlmock.Rows {
	sec := createSectionArray()

	rows := sqlmock.NewRows([]string{
		"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity",
		"minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"})

	if !flag {
		rows.AddRow("", "", "", "", "", "", "", "", "")
		return rows
	}

	rows.AddRow(sec[0].ID, sec[0].SectionNumber, sec[0].CurTemperature, sec[0].MinTemperature,
		sec[0].CurCapacity, sec[0].MinCapacity, sec[0].MaxCapacity, sec[0].WareHouseID, sec[0].ProductTypeID)

	return rows
}

func TestRepositoryGetAll(t *testing.T) {
	mock, mockRepository, _ := InitTest(t)

	rows := mockRowsArray()
	exp := createSectionArray()

	t.Run("find_all", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(section.SqlGetAll)).WillReturnRows(rows)
		sections, err := mockRepository.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, exp, sections)
	})

	t.Run("find_all_fail_query", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(section.SqlGetAll)).WillReturnError(sql.ErrNoRows)
		sections, err := mockRepository.GetAll()

		assert.Equal(t, []section.Section(nil), sections)
		assert.Error(t, err)
	})

	t.Run("find_all_fail_scan", func(t *testing.T) {
		row := mockRow(FailScan)
		mock.ExpectQuery(regexp.QuoteMeta(section.SqlGetAll)).WillReturnRows(row)
		sec, err := mockRepository.GetAll()

		assert.Equal(t, []section.Section(nil), sec)
		assert.Error(t, err)
	})
}

func TestRepositoryGetID(t *testing.T) {
	mock, mockRepository, row := InitTest(t)
	sec := createSectionArray()

	t.Run("find_by_id_existent", func(t *testing.T) {
		exp := sec[0]
		mock.ExpectQuery(regexp.QuoteMeta(section.SqlGetById)).WillReturnRows(row)
		sections, err := mockRepository.GetByID(1)

		assert.NoError(t, err)
		assert.Equal(t, exp, sections)
	})

	t.Run("find_by_id_fail_query", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(section.SqlGetById)).WillReturnError(sql.ErrNoRows)
		sections, err := mockRepository.GetByID(1)

		assert.Equal(t, section.Section{}, sections)
		assert.Error(t, err)
	})

	t.Run("find_all_fail_scan", func(t *testing.T) {
		row := mockRow(FailScan)
		mock.ExpectQuery(regexp.QuoteMeta(section.SqlGetById)).WillReturnRows(row)
		sec, err := mockRepository.GetByID(1)

		assert.Equal(t, section.Section{}, sec)
		assert.Error(t, err)
	})
}

func TestRepositoryCreate(t *testing.T) {
	mock, mockRepository, _ := InitTest(t)

	sec := createSectionArray()
	exp := sec[0]
	sec = append([]section.Section{}, sec[1:]...)

	t.Run("create_ok", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(section.SqlStore)).WithArgs(&exp.SectionNumber, &exp.CurTemperature,
			&exp.MinTemperature, &exp.CurCapacity, &exp.MinCapacity, &exp.MaxCapacity, &exp.WareHouseID,
			&exp.ProductTypeID).WillReturnResult(sqlmock.NewResult(1, 1))

		sec, err := mockRepository.Create(exp.SectionNumber, exp.CurTemperature, exp.MinTemperature,
			exp.CurCapacity, exp.MinCapacity, exp.MaxCapacity, exp.WareHouseID, exp.ProductTypeID)

		assert.Equal(t, exp, sec)
		assert.NoError(t, err)
	})

	t.Run("create_fail_exec", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(section.SqlStore)).WithArgs(&exp.SectionNumber, &exp.CurTemperature,
			&exp.MinTemperature, &exp.CurCapacity, &exp.MinCapacity, &exp.MaxCapacity, &exp.WareHouseID,
			&exp.ProductTypeID).WillReturnError(sql.ErrNoRows)

		sec, err := mockRepository.Create(exp.SectionNumber, exp.CurTemperature, exp.MinTemperature,
			exp.CurCapacity, exp.MinCapacity, exp.MaxCapacity, exp.WareHouseID, exp.ProductTypeID)

		assert.Equal(t, section.Section{}, sec)
		assert.Error(t, err)
	})

	t.Run("create_fail_zero_rows_affected", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(section.SqlStore)).WithArgs(&exp.SectionNumber, &exp.CurTemperature,
			&exp.MinTemperature, &exp.CurCapacity, &exp.MinCapacity, &exp.MaxCapacity, &exp.WareHouseID,
			&exp.ProductTypeID).WillReturnResult(sqlmock.NewResult(1, 0))

		sec, err := mockRepository.Create(exp.SectionNumber, exp.CurTemperature, exp.MinTemperature,
			exp.CurCapacity, exp.MinCapacity, exp.MaxCapacity, exp.WareHouseID, exp.ProductTypeID)

		assert.Equal(t, section.Section{}, sec)
		assert.Error(t, err)
		assert.Equal(t, errors.New("rows not affected"), err)
	})
}

func TestRepositoryUpdate(t *testing.T) {
	mock, mockRepository, _ := InitTest(t)

	sec := createSectionArray()
	exp := sec[0]

	t.Run("update_existent", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(section.SqlUpdateSecID)).WithArgs(50, 1).WillReturnResult(sqlmock.NewResult(1, 1))

		row := sqlmock.NewRows([]string{
			"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity",
			"minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"})

		row.AddRow(exp.ID, 50, exp.CurTemperature, exp.MinTemperature, exp.CurCapacity,
			exp.MinCapacity, exp.MaxCapacity, exp.WareHouseID, exp.ProductTypeID)

		mock.ExpectQuery(regexp.QuoteMeta(section.SqlGetById)).WillReturnRows(row)

		sec, err := mockRepository.UpdateSecID(1, 50)

		exp.SectionNumber = 50
		assert.Equal(t, exp, sec)
		assert.Equal(t, section.CodeError{200, nil}, err)
	})

	t.Run("update_fail_update_query", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(section.SqlUpdateSecID)).WithArgs(50, 1).WillReturnError(sql.ErrNoRows)

		sec, err := mockRepository.UpdateSecID(1, 50)

		assert.Equal(t, section.Section{}, sec)
		assert.Equal(t, section.CodeError{500, errors.New("sql: no rows in result set")}, err)
	})

	t.Run("update_fail_zero_rows_affected", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(section.SqlUpdateSecID)).WithArgs(50, 1).WillReturnResult(sqlmock.NewResult(1, 0))

		sec, err := mockRepository.UpdateSecID(1, 50)

		assert.Equal(t, section.Section{}, sec)
		assert.Equal(t, section.CodeError{500, errors.New("rows not affected")}, err)
	})
}

func TestRepositoryDelete(t *testing.T) {
	mock, mockRepository, _ := InitTest(t)

	t.Run("delete_ok", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(section.SqlDelete)).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
		err := mockRepository.DeleteSection(1)

		assert.NoError(t, err)
	})

	t.Run("delete_fail_query", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(section.SqlDelete)).WithArgs(1).WillReturnError(sql.ErrNoRows)
		err := mockRepository.DeleteSection(1)

		assert.Error(t, err)
		assert.Equal(t, errors.New("sql: no rows in result set"), err)
	})

	t.Run("delete_fail_zero_rows_affected", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(section.SqlDelete)).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 0))
		err := mockRepository.DeleteSection(1)

		assert.Error(t, err)
		assert.Equal(t, errors.New("rows not affected"), err)
	})
}
