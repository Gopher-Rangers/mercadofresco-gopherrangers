package buyer

type Service interface {
	GetAll() ([]Buyer, error)
	Create(cardNumberId string, firstName string, lastName string) (Buyer, error)
	Update(id int, cardNumberId string, firstName string, lastName string) (Buyer, error)
	Delete(id int) error
	GetById(id int) (Buyer, error)
}

type service struct {
	repository Repository
}

func NewService() Service {
	return &service{repository: NewRepository()}
}

func (s *service) Create(cardNumberId string, firstName string, lastName string) (Buyer, error) {
	buyer, err := s.repository.Create(cardNumberId, firstName, lastName)
	if err != nil {
		return Buyer{}, err
	}
	return buyer, nil
}

func (s *service) Update(id int, cardNumberId string, firstName string, lastName string) (Buyer, error) {
	updatedBuyer, err := s.repository.Update(id, cardNumberId, firstName, lastName)
	if err != nil {
		return Buyer{}, err
	}

	return updatedBuyer, nil
}

func (s *service) GetAll() ([]Buyer, error) {
	buyers, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return buyers, nil
}

func (s *service) GetById(id int) (Buyer, error) {
	buyer, err := s.repository.GetById(id)
	if err != nil {
		return Buyer{}, err
	}
	return buyer, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
