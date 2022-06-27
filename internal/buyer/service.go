package buyer

import (
	"context"
	"fmt"
)

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(ctx context.Context, buyer Buyer) (Buyer, error) {
	err := validateCardNumber(ctx, buyer.CardNumberId, s)
	if err != nil {
		return Buyer{}, err
	}

	newBuyer, err := s.repository.Create(ctx, buyer)
	if err != nil {
		return Buyer{}, err
	}
	return newBuyer, nil
}

func validateCardNumber(ctx context.Context, cardNumberId string, s *service) error {

	entities, _ := s.repository.GetAll(ctx)

	for i := 0; i < len(entities); i++ {
		if entities[i].CardNumberId == cardNumberId {
			return fmt.Errorf("buyer with card_number_id %s already exists", cardNumberId)
		}
	}

	return nil
}

func (s *service) Update(ctx context.Context, buyer Buyer) (Buyer, error) {
	err := validateCardNumber(ctx, buyer.CardNumberId, s)
	if err != nil {
		return Buyer{}, err
	}

	updatedBuyer, err := s.repository.Update(ctx, buyer)
	if err != nil {
		return Buyer{}, err
	}

	return updatedBuyer, nil
}

func (s *service) GetAll(ctx context.Context) ([]Buyer, error) {
	buyers, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return buyers, nil
}

func (s *service) GetById(ctx context.Context, id int) (Buyer, error) {
	buyer, err := s.repository.GetById(ctx, id)
	if err != nil {
		return Buyer{}, err
	}
	return buyer, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
