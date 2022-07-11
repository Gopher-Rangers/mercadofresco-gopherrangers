package inboundorders

import (
	"database/sql"
	"fmt"
)

type InboundOrder struct {
	ID             int    `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeId     int    `json:"employee_id"`
	ProductBatchId int    `json:"product_batch_id"`
	WarehouseId    int    `json:"warehouse_id"`
}

type Repository interface {
	Create(orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (InboundOrder, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) Create(orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (InboundOrder, error) {
	res, err := r.db.Exec(SqlCreate, orderDate, orderNumber, employeeId, productBatchId, warehouseId)
	if err != nil {
		return InboundOrder{}, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected <= 0 {
		return InboundOrder{}, fmt.Errorf("rows not affected")
	}

	lastID, _ := res.LastInsertId()

	emp := InboundOrder{int(lastID), orderDate, orderNumber, employeeId, productBatchId, warehouseId}
	return emp, nil
}
