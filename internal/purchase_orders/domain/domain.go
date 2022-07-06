package domain

import "context"

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
	ValidadeOrderNumber(orderNumber string) (bool, error)
}

type Service interface {
	Create(ctx context.Context, purchaseOrder PurchaseOrders) (PurchaseOrders, error)
	GetById(ctx context.Context, id int) (PurchaseOrders, error)
}
