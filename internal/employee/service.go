package employee

type Services interface {
	Create(cardNum int, firstName string, lastName string, warehouseId int) (Employee, error)
	LastID() int
	GetAll() []Employee
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
