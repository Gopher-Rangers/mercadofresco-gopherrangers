package employee

type Services interface {
	Create(cardNum int, firstName string, lastName string, warehouseId int) (Employee, error)
	LastID() int
	GetAll() []Employee
	Delete(id int) error
	GetById(id int) (Employee, error)
	Update(id int, cardNum int) (Employee, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Services {
	s := service{r}
	return &s
}

func (s service) LastID() int {
	return s.repository.LastID()
}

func (s *service) Create(cardNum int, firstName string, lastName string, warehouseId int) (Employee, error) {
	id := s.repository.AvailableID()
	ps, err := s.repository.Create(id, cardNum, firstName, lastName, warehouseId)
	if err != nil {
		return Employee{}, err
	}
	return ps, nil
}

func (s service) GetAll() []Employee {
	return s.repository.GetAll()
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s service) GetById(id int) (Employee, error) {
	employee, err := s.repository.GetById(id)
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func (s *service) Update(id, cardNum int) (Employee, error) {
	employee, err := s.repository.Update(id, cardNum)
	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}
