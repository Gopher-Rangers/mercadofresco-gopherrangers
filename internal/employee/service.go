package employee

import (
	"fmt"
)

type EmployeeOrderCount struct {
	Employee
	count int
}

type Services interface {
	Create(cardNum int, firstName string, lastName string, warehouseId int) (Employee, error)
	GetAll() ([]Employee, error)
	Delete(id int) error
	GetById(id int) (Employee, error)
	Update(emp Employee, id int) (Employee, error)
	GetCount(id, counter int) (EmployeeOrderCount, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Services {
	s := service{r}
	return &s
}

func (s *service) validateCardNumber(cardNum int) bool {
	employees, _ := s.GetAll()
	for i := range employees {
		if employees[i].CardNumber == cardNum {
			return false
		}
	}
	return true
}

func (s *service) Create(cardNum int, firstName string, lastName string, warehouseId int) (Employee, error) {
	if !s.validateCardNumber(cardNum) {
		return Employee{}, fmt.Errorf("funcionário com cartão nº: %d já existe no banco de dados", cardNum)
	}
	emps, err := s.repository.Create(cardNum, firstName, lastName, warehouseId)
	if err != nil {
		return Employee{}, err
	}
	return emps, nil
}

func (s *service) GetAll() ([]Employee, error) {
	emps, err := s.repository.GetAll()

	if err != nil {
		return emps, err
	}

	fmt.Println("chegou no service")
	return emps, nil
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	fmt.Println("chegou no service delete")
	return nil
}

func (s service) GetById(id int) (Employee, error) {
	employee, err := s.repository.GetById(id)
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func (s *service) Update(emp Employee, id int) (Employee, error) {
	empToMatch, _ := s.repository.GetById(id)

	if emp.FirstName == "" {
		emp.FirstName = empToMatch.FirstName
	}

	if emp.LastName == "" {
		emp.LastName = empToMatch.LastName
	}

	if emp.WareHouseID == 0 {
		emp.WareHouseID = empToMatch.WareHouseID
	}

	employee, err := s.repository.Update(id, emp.FirstName, emp.LastName, emp.WareHouseID)
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func (s *service) GetCount(id, counter int) (EmployeeOrderCount, error) {
	employee, _ := s.repository.GetById(id)

	return EmployeeOrderCount{employee, counter}, nil
}
