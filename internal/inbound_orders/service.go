package inboundorders

type Services interface {
	Create(orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (InboundOrder, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Services {
	s := service{r}
	return &s
}

func (s *service) Create(orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (InboundOrder, error) {
	inboundOrder, err := s.repository.Create(orderDate, orderNumber, employeeId, productBatchId, warehouseId)
	if err != nil {
		return InboundOrder{}, err
	}
	return inboundOrder, nil
}
