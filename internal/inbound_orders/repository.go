package inboundorders

import (
	"database/sql"
	"fmt"
	"strings"
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
	GetCountByEmployee(id int) (count int)
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
		if strings.Contains(err.Error(), "employee_id") {
			return InboundOrder{}, fmt.Errorf("funcionario nao existe")
		}
		if strings.Contains(err.Error(), "product_batch_id") {
			return InboundOrder{}, fmt.Errorf("product batch nao existe")
		}
		if strings.Contains(err.Error(), "warehouse_id") {
			return InboundOrder{}, fmt.Errorf("warehouse nao existe")
		}
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

// func (r repository) GetCountByEmployee(id int) (count int) {
// 	var ordersArray []InboundOrder

// 	rows, err := r.db.Query(SqlGetAllbyId, id)

// 	if err != nil {
// 		return id
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var order InboundOrder

// 		_ = rows.Scan(&order.ID)

// 		if err != nil {
// 			return id
// 		}

// 		ordersArray = append(ordersArray, order)
// 	}

// 	return len(ordersArray)
// }

func (r repository) GetCountByEmployee(id int) (count int) {
	var counter int
	row := r.db.QueryRow(SqlCountByEmployee, id)

	_ = row.Scan(&counter)

	return counter
}
