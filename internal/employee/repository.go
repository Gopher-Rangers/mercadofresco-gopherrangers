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
	// LastID() int
	GetAll() ([]Employee, error)
	// Delete(id int) error
	// GetById(id int) (Employee, error)
	// Update(emp Employee, id int) (Employee, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// func (r repository) Create(id int, cardNum int, firstName string, lastName string, warehouseId int) (Employee, error) {
// 	var Employees []Employee
// 	r.db.Read(&Employees)

// 	p := Employee{id, cardNum, firstName, lastName, warehouseId}

// 	for i := range Employees {
// 		if Employees[i].ID+1 == id {
// 			post := make([]Employee, len(Employees[i+1:]))
// 			copy(post, Employees[i+1:])

// 			Employees = append(Employees[:i+1], p)
// 			Employees = append(Employees, post...)
// 			break
// 		}
// 	}

// 	if id == 1 {
// 		emp := []Employee{p}
// 		Employees = append(emp, Employees...)
// 	}
// 	r.db.Write(Employees)
// 	return p, nil
// }

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

// func (r repository) LastID() int {
// 	var Employees []Employee
// 	r.db.Read(&Employees)

// 	if len(Employees) == 0 {
// 		return 1
// 	}
// 	return Employees[len(Employees)-1].ID + 1
// }

func (r repository) GetAll() ([]Employee, error) {
	var employees []Employee

	rows, err := r.db.Query(SqlGetAll)
	if err != nil {
		return employees, err
	}

	defer rows.Close()

	for rows.Next() {
		var emp Employee

		err := rows.Scan(&emp.ID, &emp.CardNumber, &emp.FirstName, &emp.WareHouseID)
		if err != nil {
			return employees, err
		}

		employees = append(employees, emp)
	}

	return employees, nil
}

// func (r *repository) Delete(id int) error {
// 	var Employees []Employee
// 	r.db.Read(&Employees)

// 	for i := range Employees {
// 		if Employees[i].ID == id {
// 			Employees = append(Employees[:i], Employees[i+1:]...)
// 			r.db.Write(Employees)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("usuário de ID: %d não existe", id)
// }

// func (r repository) GetById(id int) (Employee, error) {
// 	var Employees []Employee
// 	r.db.Read(&Employees)

// 	for i := range Employees {
// 		if Employees[i].ID == id {
// 			return Employees[i], nil
// 		}
// 	}
// 	return Employee{}, fmt.Errorf("o funcionário não existe")
// }

// func (r *repository) Update(emp Employee, id int) (Employee, error) {
// 	var employees []Employee
// 	r.db.Read(&employees)

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
