package domain

import (
	"context"
)

const (
	ERROR_UNIQUE_ORDER_NUMBER = "the order number must be unique"
	ERROR_WHILE_SAVING        = "the order number must be unique"
)

type PurchaseOrders struct {
	ID              int    `json:"id"`
	OrderNumber     string `json:"order_number"`
	OrderDate       string `json:"order_date"`
	TrackingCode    string `json:"tracking_code"`
	BuyerId         int    `json:"buyer_id"`
	ProductRecordId int    `json:"product_record_id"`
	OrderStatusId   int    `json:"order_status_id"`
}

type Repository interface {
	Create(ctx context.Context, purchaseOrder PurchaseOrders) (PurchaseOrders, error)
	GetById(ctx context.Context, id int) (PurchaseOrders, error)
	ValidadeOrderNumber(ctx context.Context, orderNumber string) (bool, error)
}

type Service interface {
	Create(ctx context.Context, purchaseOrder PurchaseOrders) (PurchaseOrders, error)
	GetById(ctx context.Context, id int) (PurchaseOrders, error)
}
