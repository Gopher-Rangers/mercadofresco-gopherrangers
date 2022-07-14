package service

import (
	"context"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/domain"
)

type service struct {
	repository domain.Repository
}

func NewService(r domain.Repository) domain.Service {
	return &service{r}
}

func (s *service) GetById(ctx context.Context, id int) (domain.PurchaseOrders, error) {
	purchaseOrders, err := s.repository.GetById(ctx, id)
	if err != nil {
		return domain.PurchaseOrders{}, err
	}
	return purchaseOrders, nil
}

func (s *service) Create(ctx context.Context, purchaseOrder domain.PurchaseOrders) (domain.PurchaseOrders, error) {
	isValid := s.repository.ValidadeOrderNumber(ctx, purchaseOrder.OrderNumber)
	if !isValid {
		return domain.PurchaseOrders{}, fmt.Errorf(domain.ERROR_UNIQUE_ORDER_NUMBER)
	}

	newPurchaseOrder, err := s.repository.Create(ctx, purchaseOrder)
	if err != nil {
		return domain.PurchaseOrders{}, err
	}
	return newPurchaseOrder, nil
}
