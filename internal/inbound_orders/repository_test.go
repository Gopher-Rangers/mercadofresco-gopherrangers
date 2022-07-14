package inboundorders_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	inboundorders "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/inbound_orders"
	"github.com/stretchr/testify/assert"
)

// func mockRowsArray() *sqlmock.Rows {
// 	ios := createInboundOrdersArray()
// 	rows := sqlmock.NewRows([]string{
// 		"id", "order_date", "order_number",
// 		"employee_id", "product_batch_id", "warehouse_id"}).AddRow(
// 		ios[0].ID, ios[0].OrderDate, ios[0].OrderNumber, ios[0].EmployeeId,
// 		ios[0].ProductBatchId, ios[0].WarehouseId).AddRow(
// 		ios[1].ID, ios[1].OrderDate, ios[1].OrderNumber, ios[1].EmployeeId,
// 		ios[1].ProductBatchId, ios[1].WarehouseId)
// 	return rows
// }

// func mockRow() *sqlmock.Rows {
// 	ios := createInboundOrdersArray()
// 	rows := sqlmock.NewRows([]string{
// 		"id", "card_number_id", "first_name",
// 		"last_name", "warehouse_id"}).AddRow(
// 		ios[0].ID, ios[0].OrderDate, ios[0].OrderNumber, ios[0].EmployeeId,
// 		ios[0].ProductBatchId, ios[0].WarehouseId)
// 	return rows
// }

func TestRepositoryCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		io := createInboundOrdersArray()[0]
		mock.ExpectExec(regexp.QuoteMeta(inboundorders.SqlCreate)).WithArgs(&io.OrderDate, &io.OrderNumber,
			&io.EmployeeId, &io.ProductBatchId, &io.WarehouseId).WillReturnResult(sqlmock.NewResult(1, 1))
		inboundordersRepo := inboundorders.NewRepository(db)
		result, err := inboundordersRepo.Create(io.OrderDate, io.OrderNumber, io.EmployeeId, io.ProductBatchId, io.WarehouseId)
		assert.NoError(t, err)
		assert.Equal(t, result, io)
	})
	t.Run("create_employee_id_error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		io := createInboundOrdersArray()[0]
		mock.ExpectExec(regexp.QuoteMeta(inboundorders.SqlCreate)).WithArgs(&io.OrderDate, &io.OrderNumber,
			&io.EmployeeId, &io.ProductBatchId, &io.WarehouseId).WillReturnResult(sqlmock.NewResult(0, 0))
		inboundordersRepo := inboundorders.NewRepository(db)
		_, err = inboundordersRepo.Create(io.OrderDate, io.OrderNumber, 100, io.ProductBatchId, io.WarehouseId)
		assert.Equal(t, err, fmt.Errorf("funcionario nao existe"))
	})
}

func TestGetCountByEmployee(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		io := createInboundOrdersArray()[0]
		mock.ExpectExec(regexp.QuoteMeta(inboundorders.SqlCreate)).WithArgs(&io.OrderDate, &io.OrderNumber,
			&io.EmployeeId, &io.ProductBatchId, &io.WarehouseId).WillReturnResult(sqlmock.NewResult(1, 1))
		inboundordersRepo := inboundorders.NewRepository(db)
		result, err := inboundordersRepo.Create(io.OrderDate, io.OrderNumber, io.EmployeeId, io.ProductBatchId, io.WarehouseId)
		assert.NoError(t, err)
		assert.Equal(t, result, io)
	})
}
