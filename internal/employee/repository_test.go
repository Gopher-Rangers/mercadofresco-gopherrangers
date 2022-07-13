package employee_test

import (
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
		emp[0].WareHouseID).AddRow(
		emp[1].ID, emp[1].CardNumber, emp[1].FirstName, emp[1].LastName,
		emp[1].WareHouseID)
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

// func TestRepositoryUpdate(t *testing.T) {
// 	t.Run("update_ok", func(t *testing.T) {
// 		db, mock, err := sqlmock.New()

// 		assert.NoError(t, err)
// 		defer db.Close()
// 		emp := createEmployeeArray()[1]
// 		mock.ExpectExec(regexp.QuoteMeta(employees.SqlUpdate)).WithArgs(&emp.FirstName,
// 			&emp.LastName, &emp.WareHouseID, emp.ID).WillReturnResult(sqlmock.NewResult(1, 1))
// 		employeesRepo := employees.NewRepository(db)
// 		result, err := employeesRepo.Update(1, emp.FirstName, emp.LastName, emp.WareHouseID)
// 		assert.NoError(t, err)
// 		assert.Equal(t, result, emp)
// 	})
// }
