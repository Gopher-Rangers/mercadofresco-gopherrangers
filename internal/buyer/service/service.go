package buyer

import (
	"fmt"

	buyerDomain "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/domain"
)

type service struct {
	repository buyerDomain.Repository
}

func NewService(r buyerDomain.Repository) buyerDomain.Service {
	return &service{r}
}

func (s *service) Create(buyer buyerDomain.Buyer) (buyerDomain.Buyer, error) {
	err := validateCardNumber(buyer.CardNumberId, s)
	if err != nil {
		return buyerDomain.Buyer{}, err
	}

	buyer.Id = s.repository.GetValidId()

	newBuyer, err := s.repository.Create(buyer)
	if err != nil {
		return buyerDomain.Buyer{}, err
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

func (s *service) Update(buyer buyerDomain.Buyer) (buyerDomain.Buyer, error) {
	err := validateCardNumber(buyer.CardNumberId, s)
	if err != nil {
		return buyerDomain.Buyer{}, err
	}

	updatedBuyer, err := s.repository.Update(buyer)
	if err != nil {
		return buyerDomain.Buyer{}, err
	}

	return updatedBuyer, nil
}

func (s *service) GetAll() ([]buyerDomain.Buyer, error) {
	buyers, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return buyers, nil
}

func (s *service) GetById(id int) (buyerDomain.Buyer, error) {
	buyer, err := s.repository.GetById(id)
	if err != nil {
		return buyerDomain.Buyer{}, err
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
