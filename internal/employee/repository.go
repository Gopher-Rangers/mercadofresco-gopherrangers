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

type Repository interface {
	Create(cardNum int, firstName string, lastName string, warehouseId int) (Employee, error)
	GetAll() ([]Employee, error)
	Delete(id int) error
	GetById(id int) (Employee, error)
	// Update(emp Employee, id int) (Employee, error)
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
		return Employee{}, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected <= 0 {
		return Employee{}, fmt.Errorf("rows not affected")
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

// func (r *repository) Update(emp Employee, id int) (Employee, error) {
// 	var employee Employee

// 	r.db.Exec(SqlUpdateFirstName)
// 	for i := range employees {
// 		if emp.ID == 0 {
// 			emp.ID = id
// 		}
// 		if employees[i].ID == id {
// 			if emp.FirstName == "" {
// 				emp.FirstName = employees[i].FirstName
// 			} else {
// 				employees[i].FirstName = emp.FirstName
// 			}
// 			if emp.LastName == "" {
// 				emp.LastName = employees[i].LastName
// 			} else {
// 				employees[i].LastName = emp.LastName
// 			}
// 			if emp.WareHouseID == 0 {
// 				emp.WareHouseID = employees[i].WareHouseID
// 			} else {
// 				employees[i].WareHouseID = emp.WareHouseID
// 			}
// 			if err := r.db.Write(&employees); err != nil {
// 				return Employee{}, err
// 			}
// 			return employees[i], nil
// 		}
// 	}
// 	return Employee{}, fmt.Errorf("funcionário não foi encontrado")
// }

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
		return Employee{}, err
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
