package employee_test

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	employees "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"
)

func mockRowsArray() *sqlmock.Rows {
	emp := createEmployeeArray()
	rows := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name",
		"last_name", "warehouse_id"}).AddRow(
		emp[0].ID, emp[0].CardNumber, emp[0].FirstName, emp[0].LastName,
		emp[0].WareHouseID).AddRow(
		emp[1].ID, emp[1].CardNumber, emp[1].FirstName, emp[1].LastName,
		emp[1].WareHouseID)
	return rows
}

func mockRow() *sqlmock.Rows {
	emp := createEmployeeArray()
	rows := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name",
		"last_name", "warehouse_id"}).AddRow(
		emp[0].ID, emp[0].CardNumber, emp[0].FirstName, emp[0].LastName,
		emp[0].WareHouseID)
	return rows
}

func TestRepositoryCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		emp := createEmployeeArray()[0]
		mock.ExpectExec(regexp.QuoteMeta(employees.SqlCreate)).WithArgs(&emp.CardNumber, &emp.FirstName,
			&emp.LastName, &emp.WareHouseID).WillReturnResult(sqlmock.NewResult(1, 1))
		employeesRepo := employees.NewRepository(db)
		result, err := employeesRepo.Create(emp.CardNumber, emp.FirstName, emp.LastName, emp.WareHouseID)
		assert.NoError(t, err)
		assert.Equal(t, result, emp)
	})
	t.Run("create_fail_zero_rows_affected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		emp := createEmployeeArray()[0]
		mock.ExpectExec(regexp.QuoteMeta(employees.SqlCreate)).WithArgs(&emp.CardNumber, &emp.FirstName,
			&emp.LastName, &emp.WareHouseID).WillReturnResult(sqlmock.NewResult(1, 0))
		employeesRepo := employees.NewRepository(db)
		result, err := employeesRepo.Create(emp.CardNumber, emp.FirstName, emp.LastName, emp.WareHouseID)
		assert.Equal(t, err, fmt.Errorf("employee not created"))
		assert.Equal(t, result, employees.Employee{})
	})
}

func TestRepositoryUpdate(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := mockRow()
		emp := createEmployeeArray()[0]
		mock.ExpectExec(regexp.QuoteMeta(employees.SqlUpdate)).WithArgs(&emp.FirstName,
			&emp.LastName, &emp.WareHouseID, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery(regexp.QuoteMeta(employees.SqlGetById)).WithArgs(1).WillReturnRows(rows)
		employeesRepo := employees.NewRepository(db)
		result, err := employeesRepo.Update(1, emp.FirstName, emp.LastName, emp.WareHouseID)
		assert.NoError(t, err)
		assert.Equal(t, result, emp)
	})
}

func TestRepositoryGetAll(t *testing.T) {
	t.Run("find_all_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		emp := createEmployeeArray()
		rows := mockRowsArray()
		mock.ExpectQuery(regexp.QuoteMeta(employees.SqlGetAll)).WillReturnRows(rows)
		employeesRepo := employees.NewRepository(db)
		result, err := employeesRepo.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, result[0], emp[0])
		assert.Equal(t, result[1], emp[1])
	})
	t.Run("find_all_fail_scan", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		rows := sqlmock.NewRows([]string{
			"id", "card_number", "first_name",
			"last_name", "warehouse_id"}).AddRow(
			"", "", "", "", "")
		mock.ExpectQuery(regexp.QuoteMeta(employees.SqlGetAll)).WillReturnRows(rows)
		employeesRepo := employees.NewRepository(db)
		emps, err := employeesRepo.GetAll()
		assert.Equal(t, emps, []employees.Employee(nil))
		assert.Error(t, err)
	})
	t.Run("find_all_fail_select", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(regexp.QuoteMeta(employees.SqlGetAll)).WillReturnError(sql.ErrNoRows)
		employeesRepo := employees.NewRepository(db)
		emps, err := employeesRepo.GetAll()
		assert.Equal(t, emps, []employees.Employee(nil))
		assert.Error(t, err)
	})
}

func TestRepositoryDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectExec(regexp.QuoteMeta(employees.SqlDelete)).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
		employeesRepo := employees.NewRepository(db)
		err = employeesRepo.Delete(1)
		assert.NoError(t, err)
	})
	t.Run("delete_non_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectExec(regexp.QuoteMeta(employees.SqlDelete)).WithArgs(40).WillReturnResult(sqlmock.NewResult(2, 0))
		employeesRepo := employees.NewRepository(db)
		err = employeesRepo.Delete(40)
		assert.Equal(t, "row not affected", err.Error())

	})
	t.Run("delete_fail_exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectExec(regexp.QuoteMeta(employees.SqlDelete)).WithArgs(1).WillReturnError(sql.ErrNoRows)
		employeesRepo := employees.NewRepository(db)
		err = employeesRepo.Delete(1)
		assert.Error(t, err)

	})
}

func TestRepositoryGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		emps := createEmployeeArray()
		rows := mockRow()

		mock.ExpectQuery(regexp.QuoteMeta(employees.SqlGetById)).WithArgs(1).WillReturnRows(rows)
		employeesRepo := employees.NewRepository(db)
		result, err := employeesRepo.GetById(emps[0].ID)
		assert.NoError(t, err)
		assert.Equal(t, result, emps[0])
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		expectedError := fmt.Errorf("funcionario nao existe")
		mock.ExpectQuery(regexp.QuoteMeta(employees.SqlGetById)).WithArgs(13).WillReturnError(expectedError)
		employeesRepo := employees.NewRepository(db)
		result, err := employeesRepo.GetById(13)

		assert.Equal(t, err, expectedError)
		assert.Equal(t, result, employees.Employee{})
	})

}
