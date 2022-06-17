package warehouse

type Service interface {
	GetAll() []Warehouse
	GetByID(id int) (Warehouse, error)
	CreateWarehouse(
		code,
		address,
		tel string,
		minCap,
		minTemp int) (Warehouse, error)
	UpdatedWarehouseID(id int, code string) (Warehouse, error)
	DeleteWarehouse(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s service) GetAll() []Warehouse {
	return s.repository.GetAll()
}

func (s service) GetByID(id int) (Warehouse, error) {
	warehouse, err := s.repository.GetByID(id)

	if err != nil {
		return Warehouse{}, err
	}

	return warehouse, nil

}

func (s service) CreateWarehouse(code, address, tel string, minCap, minTemp int) (Warehouse, error) {

	id := s.repository.IncrementID()

	warehouse, err := s.repository.CreateWarehouse(id, code, address, tel, minCap, minTemp)

	if err != nil {
		return Warehouse{}, err
	}

	return warehouse, nil
}

func (s service) UpdatedWarehouseID(id int, code string) (Warehouse, error) {
	warehouse, err := s.repository.UpdatedWarehouseID(id, code)

	if err != nil {
		return Warehouse{}, err
	}

	return warehouse, nil
}

func (s service) DeleteWarehouse(id int) error {
	err := s.repository.DeleteWarehouse(id)

	if err != nil {
		return err
	}

	return nil
}

func (s service) IncrementID() int {
	return s.repository.IncrementID()
}
