package buyer

import (
	"fmt"
)

type Service interface {
	GetAll() ([]Buyer, error)
	Create(buyer Buyer) (Buyer, error)
	Update(buyer Buyer) (Buyer, error)
	Delete(id int) error
	GetById(id int) (Buyer, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(buyer Buyer) (Buyer, error) {
	err := validateCardNumber(buyer.CardNumberId, s)
	if err != nil {
		return Buyer{}, err
	}

	buyer.Id = s.repository.GetValidId()

	newBuyer, err := s.repository.Create(buyer)
	if err != nil {
		return Buyer{}, err
	}
	return newBuyer, nil
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

func (s *service) Update(buyer Buyer) (Buyer, error) {
	err := validateCardNumber(buyer.CardNumberId, s)
	if err != nil {
		return Buyer{}, err
	}

	updatedBuyer, err := s.repository.Update(buyer)
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
