package employee

import (
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
)

type Employee struct {
	ID          int    `json:"id"`
	CardNumber  int    `json:"card_number_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	WareHouseID int    `json:"warehouse_id"`
}

type Repository interface {
	Create(id int, cardNum int, firstName string, lastName string, warehouseId int) (Employee, error)
	LastID() int
	AvailableID() int
	GetAll() []Employee
	Delete(id int) error
	GetById(id int) (Employee, error)
	Update(emp Employee, id int) (Employee, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r repository) Create(id int, cardNum int, firstName string, lastName string, warehouseId int) (Employee, error) {
	var Employees []Employee
	r.db.Read(&Employees)

	p := Employee{id, cardNum, firstName, lastName, warehouseId}

	for i := range Employees {
		if Employees[i].CardNumber == cardNum {
			return Employee{}, fmt.Errorf("seção com cartão nº: %d já existe no banco de dados", cardNum)
		}
	}

	for i := range Employees {
		if Employees[i].ID+1 == id {
			post := make([]Employee, len(Employees[i+1:]))
			copy(post, Employees[i+1:])

			Employees = append(Employees[:i+1], p)
			Employees = append(Employees, post...)
			break
		}
	}

	if id == 1 {
		emp := []Employee{p}
		Employees = append(emp, Employees...)
	}
	r.db.Write(Employees)
	return p, nil
}

func (r repository) AvailableID() int {
	var Employees []Employee
	r.db.Read(&Employees)

	if len(Employees) == 0 {
		return 1
	}

	for prevI := range Employees[:len(Employees)-1] {
		i := prevI + 1
		if Employees[i].ID != (Employees[prevI].ID + 1) {
			id := Employees[prevI].ID + 1
			return id
		}
	}
	return r.LastID()
}

func (r repository) LastID() int {
	var Employees []Employee
	r.db.Read(&Employees)

	if len(Employees) == 0 || Employees[0].ID != 1 {
		return 1
	}
	return Employees[len(Employees)-1].ID + 1
}

func (r repository) GetAll() []Employee {
	var Employees []Employee
	r.db.Read(&Employees)

	return Employees
}

func (r *repository) Delete(id int) error {
	var Employees []Employee
	r.db.Read(&Employees)

	for i := range Employees {
		if Employees[i].ID == id {
			Employees = append(Employees[:i], Employees[i+1:]...)
			r.db.Write(Employees)
			return nil
		}
	}
	return fmt.Errorf("usuário de ID: %d não existe", id)
}

func (r repository) GetById(id int) (Employee, error) {
	var Employees []Employee
	r.db.Read(&Employees)

	for i := range Employees {
		if Employees[i].ID == id {
			return Employees[i], nil
		}
	}
	return Employee{}, fmt.Errorf("o funcionário não foi encontrado")
}

func (r repository) Update(emp Employee, id int) (Employee, error) {
	var employees []Employee
	r.db.Read(&employees)

	if emp.ID == 0 {
		emp.ID = employees[id-1].ID
	} else if employees[id-1].ID != emp.ID {
		return Employee{}, fmt.Errorf("você não pode trocar o ID do funcionário")
	}

	for i := range employees {
		if employees[i].ID == id {
			employees[i].FirstName = emp.FirstName
			employees[i].LastName = emp.LastName
			employees[i].WareHouseID = emp.WareHouseID
			if err := r.db.Write(&employees); err != nil {
				return Employee{}, err
			}
			return emp, nil
		}
	}
	return Employee{}, fmt.Errorf("funcionário %d não encontrado", id)
}
