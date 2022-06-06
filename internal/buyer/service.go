package buyer

type Service interface {
	GetAll() ([]Buyer, error)
	Save(id int, cardNumberId string, firstName string, lastName string) (Buyer, error)
}

type service struct {
	repository Repository
}

func NewService() Service {
	return &service{repository: NewRepository()}
}

func (s *service) Save(id int, cardNumberId string, firstName string, lastName string) (Buyer, error) {
	buyer, err := s.repository.Save(id, cardNumberId, firstName, lastName)
	if err != nil {
		return Buyer{}, err
	}
	return buyer, nil
}

func (s *service) GetAll() ([]Buyer, error) {
	buyers, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return buyers, nil
}
