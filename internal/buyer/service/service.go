package service

import (
	"context"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/domain"
)

type service struct {
	repository domain.Repository
}

func NewService(r domain.Repository) domain.Service {
	return &service{r}
}

func (s *service) Create(ctx context.Context, buyer domain.Buyer) (domain.Buyer, error) {
	isValid, err := s.repository.ValidateCardNumberId(ctx, buyer.ID, buyer.CardNumberId)
	if err != nil {
		return domain.Buyer{}, err
	}
	if !isValid {
		return domain.Buyer{}, fmt.Errorf(domain.ERROR_UNIQUE_CARD_NUMBER_ID)
	}

	newBuyer, err := s.repository.Create(ctx, buyer)
	if err != nil {
		return domain.Buyer{}, err
	}
	return newBuyer, nil
}

func (s *service) Update(ctx context.Context, buyer domain.Buyer) (domain.Buyer, error) {
	isValid, err := s.repository.ValidateCardNumberId(ctx, buyer.ID, buyer.CardNumberId)
	if err != nil {
		return domain.Buyer{}, err
	}
	if !isValid {
		return domain.Buyer{}, fmt.Errorf(domain.ERROR_UNIQUE_CARD_NUMBER_ID)
	}

	updatedBuyer, err := s.repository.Update(ctx, buyer)
	if err != nil {
		return domain.Buyer{}, err
	}

	return updatedBuyer, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	buyers, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return buyers, nil
}

func (s *service) GetById(ctx context.Context, id int) (domain.Buyer, error) {
	buyer, err := s.repository.GetById(ctx, id)
	if err != nil {
		return domain.Buyer{}, err
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

func (s *service) GetBuyerOrdersById(ctx context.Context, id int) (domain.BuyerTotalOrders, error) {
	buyerWithOrders, err := s.repository.GetBuyerOrdersById(ctx, id)
	if err != nil {
		return domain.BuyerTotalOrders{}, err
	}
	return buyerWithOrders, nil
}

func (s *service) GetBuyerTotalOrders(ctx context.Context) ([]domain.BuyerTotalOrders, error) {
	buyerWithOrders, err := s.repository.GetBuyerTotalOrders(ctx)
	if err != nil {
		return nil, err
	}
	return buyerWithOrders, nil
}
