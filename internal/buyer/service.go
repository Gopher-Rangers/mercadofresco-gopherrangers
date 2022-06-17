package buyer

import (
	"fmt"

	"github.com/google/uuid"
)

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
	err := validateCardNumber(cardNumberId, s)
	if err != nil {
		return Buyer{}, err
	}

	validId := getNewValidId(s)

	buyer, err := s.repository.Create(validId, cardNumberId, firstName, lastName)
	if err != nil {
		return Buyer{}, err
	}
	return buyer, nil
}

func validateCardNumber(cardNumberId string, s *service) error {

	entities, _ := s.repository.GetAll()

	for i := 0; i < len(entities); i++ {
		if entities[i].CardNumberId == cardNumberId {
			return fmt.Errorf("buyer with card_number_id %s already exists", cardNumberId)
		}
	}

	return nil
}

func getNewValidId(s *service) int {
	generatedId := int(uuid.New().ID())

	entities, _ := s.repository.GetAll()

	for i := 0; i < len(entities); i++ {
		if entities[i].Id == generatedId {
			generatedId = int(uuid.New().ID())
			i = 0
		}
	}

	return generatedId
}

func (s *service) Update(id int, cardNumberId string, firstName string, lastName string) (Buyer, error) {
	entities, _ := s.repository.GetAll()

	for i := 0; i < len(entities); i++ {
		if entities[i].CardNumberId == cardNumberId && entities[i].Id != id {
			return Buyer{}, fmt.Errorf("buyer with card_number_id %s already exists", cardNumberId)
		}
	}

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
