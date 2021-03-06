package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{db: db}
}

func (r *repository) GetById(ctx context.Context, id int) (domain.PurchaseOrders, error) {
	var purchaseOrder domain.PurchaseOrders

	stmt, err := r.db.PrepareContext(ctx, SqlGetById)

	if err != nil {
		return domain.PurchaseOrders{}, err
	}

	defer stmt.Close()
	err = stmt.QueryRowContext(ctx, id).Scan(&purchaseOrder.ID, &purchaseOrder.OrderNumber, &purchaseOrder.OrderDate,
		&purchaseOrder.TrackingCode, &purchaseOrder.BuyerId, &purchaseOrder.ProductRecordId, &purchaseOrder.OrderStatusId)
	if err != nil {
		return domain.PurchaseOrders{}, fmt.Errorf("purchase order with id (%d) not founded", id)
	}
	return purchaseOrder, nil
}

func (r repository) Create(ctx context.Context, purchaseOrder domain.PurchaseOrders) (domain.PurchaseOrders, error) {
	res, err := r.db.ExecContext(ctx, SqlCreate, &purchaseOrder.OrderNumber, &purchaseOrder.OrderDate,
		&purchaseOrder.TrackingCode, &purchaseOrder.BuyerId, &purchaseOrder.ProductRecordId, &purchaseOrder.OrderStatusId)
	if err != nil {
		return domain.PurchaseOrders{}, err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return domain.PurchaseOrders{}, fmt.Errorf(domain.ERROR_WHILE_SAVING)
	}

	lastID, err := res.LastInsertId()
	if err != nil || lastID < 1 {
		return domain.PurchaseOrders{}, fmt.Errorf(domain.ERROR_WHILE_SAVING)
	}

	purchaseOrder.ID = int(lastID)

	return purchaseOrder, nil
}

func (r *repository) ValidadeOrderNumber(ctx context.Context, orderNumber string) (bool, error) {

	var orderExistent string
	stmt, err := r.db.PrepareContext(ctx, SqlOrderNumber)

	if err != nil {
		return false, err
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(ctx, orderNumber).Scan(&orderExistent)

	return orderExistent == "", nil
}
