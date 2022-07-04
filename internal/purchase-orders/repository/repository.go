package repository

import (
	"context"
	"database/sql"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase-orders/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{db: db}
}

func (r *repository) GetById(ctx context.Context, id int) (domain.PurchaseOrders, error) {
	var purchaseOrder domain.PurchaseOrders

	rows, err := r.db.QueryContext(ctx, sqlGetById, id)
	if err != nil {
		return domain.PurchaseOrders{}, err
	}

	defer rows.Close() // Impedir vazamento de memória

	for rows.Next() {
		err := rows.Scan(&purchaseOrder.ID, &purchaseOrder.OrderNumber, &purchaseOrder.OrderDate,
			&purchaseOrder.TrackingCode, &purchaseOrder.BuyerId, &purchaseOrder.ProductRecordId, &purchaseOrder.OrderStatusId)
		if err != nil {
			return domain.PurchaseOrders{}, err
		}

	}

	return purchaseOrder, nil
}

func (r repository) Create(ctx context.Context, purchaseOrder domain.PurchaseOrders) (domain.PurchaseOrders, error) {
	res, err := r.db.ExecContext(ctx, sqlCreate, &purchaseOrder.OrderNumber, &purchaseOrder.OrderDate,
		&purchaseOrder.TrackingCode, &purchaseOrder.BuyerId, &purchaseOrder.ProductRecordId, &purchaseOrder.OrderStatusId)
	if err != nil {
		return domain.PurchaseOrders{}, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return purchaseOrder, err
	}

	purchaseOrder.ID = int(lastID)

	return purchaseOrder, nil
}
