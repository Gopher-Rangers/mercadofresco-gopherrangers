package inboundorders

type Services interface {
	Create(orderDate, orderNumber string, employeeId, productBatchId, warehouseId int) (InboundOrder, error)
	GetCounterByEmployee(id int) (counter int)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Services {
	s := service{r}
	return &s
}

func (s *service) Create(orderDate, orderNumber string, employeeId, productBatchId, warehouseId int) (InboundOrder, error) {
	inboundOrder, err := s.repository.Create(orderDate, orderNumber, employeeId, productBatchId, warehouseId)
	if err != nil {
		return InboundOrder{}, err
	}
	return inboundOrder, nil
}

func (s *service) GetCounterByEmployee(id int) (counter int) {
	counter = s.repository.GetCountByEmployee(id)

	return counter
}
