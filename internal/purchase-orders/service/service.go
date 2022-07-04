package service

import (
	"context"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase-orders/domain"
)

type service struct {
	repository domain.Repository
}

func NewService(r domain.Repository) domain.Service {
	return &service{r}
}

func (s *service) GetById(ctx context.Context, id int) (domain.PurchaseOrders, error) {
	buyer, err := s.repository.GetById(ctx, id)
	if err != nil {
		return domain.PurchaseOrders{}, err
	}
	return buyer, nil
}

func (s *service) Create(ctx context.Context, purchaseOrder domain.PurchaseOrders) (domain.PurchaseOrders, error) {
	newBuyer, err := s.repository.Create(ctx, purchaseOrder)
	if err != nil {
		return domain.PurchaseOrders{}, err
	}
	return newBuyer, nil
}
