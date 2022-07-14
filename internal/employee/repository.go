package employee

import (
	"database/sql"
	"fmt"
)

type Employee struct {
	ID          int    `json:"id"`
	CardNumber  int    `json:"card_number_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	WareHouseID int    `json:"warehouse_id"`
}

type EmployeeInboundOrders struct {
	ID                 int    `json:"id"`
	CardNumber         int    `json:"card_number_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	WareHouseID        int    `json:"warehouse_id"`
	InboundOrdersCount int    `json:"inbound_orders_count"`
}

type Repository interface {
	Create(cardNum int, firstName string, lastName string, warehouseId int) (Employee, error)
	GetAll() ([]Employee, error)
	Delete(id int) error
	GetById(id int) (Employee, error)
	Update(id int, firstName string, lastName string, warehouseId int) (Employee, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) Create(cardNum int, firstName string, lastName string, warehouseId int) (Employee, error) {
	res, err := r.db.Exec(SqlCreate, cardNum, firstName, lastName, warehouseId)
	if err != nil {
		return Employee{}, fmt.Errorf("galpao %d nao existe", warehouseId)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected <= 0 {
		return Employee{}, fmt.Errorf("employee not created")
	}

	lastID, _ := res.LastInsertId()

	emp := Employee{int(lastID), cardNum, firstName, lastName, warehouseId}
	return emp, nil
}

func (r repository) GetAll() ([]Employee, error) {
	var employees []Employee

	rows, err := r.db.Query(SqlGetAll)

	if err != nil {
		return employees, err
	}

	defer rows.Close()

	for rows.Next() {
		var emp Employee

		err := rows.Scan(&emp.ID, &emp.CardNumber, &emp.FirstName, &emp.LastName, &emp.WareHouseID)

		if err != nil {
			return employees, err
		}

		employees = append(employees, emp)
	}

	return employees, err
}

func (r repository) Update(id int, firstName string, lastName string, warehouseId int) (Employee, error) {
	olderProduct, _ := r.GetById(id)
	newProduct := Employee{id, olderProduct.CardNumber, firstName, lastName, warehouseId}
	res, err := r.db.Exec(SqlUpdate, firstName, lastName, warehouseId, id)
	if err != nil {
		return Employee{}, fmt.Errorf("funcionario nao existe")
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected <= 0 && newProduct != olderProduct {
		return Employee{}, fmt.Errorf("rows not affected")
	}

	emp, _ := r.GetById(id)
	return emp, nil
}

func (r repository) Delete(id int) error {
	res, err := r.db.Exec(SqlDelete, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected <= 0 {
		return fmt.Errorf("row not affected")
	}

	return nil
}

func (r repository) GetById(id int) (Employee, error) {
	var emp Employee

	rows, err := r.db.Query(SqlGetById, id)
	if err != nil {
		return Employee{}, fmt.Errorf("funcionario nao existe")
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&emp.ID, &emp.CardNumber, &emp.FirstName, &emp.LastName, &emp.WareHouseID)
		if err != nil {
			return Employee{}, err
		}
	}

	return emp, nil
}
