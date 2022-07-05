package domain

import "context"

type Buyer struct {
	ID           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type BuyerTotalOrders struct {
	ID                  int    `json:"id"`
	CardNumberId        string `json:"card_number_id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	PurchaseOrdersCount int    `json:"purchase_orders_count"`
}

type Repository interface {
	GetAll(ctx context.Context) ([]Buyer, error)
	Create(ctx context.Context, buyer Buyer) (Buyer, error)
	Update(ctx context.Context, buyer Buyer) (Buyer, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (Buyer, error)
	GetBuyerOrdersById(ctx context.Context, id int) (BuyerTotalOrders, error)
}

type Service interface {
	GetAll(ctx context.Context) ([]Buyer, error)
	Create(ctx context.Context, buyer Buyer) (Buyer, error)
	Update(ctx context.Context, buyer Buyer) (Buyer, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (Buyer, error)
	GetBuyerOrdersById(ctx context.Context, id int) (BuyerTotalOrders, error)
}
