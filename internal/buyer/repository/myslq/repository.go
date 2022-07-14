package myslq

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	var buyers []domain.Buyer

	rows, err := r.db.QueryContext(ctx, SqlGetAll)
	if err != nil {
		return buyers, err
	}

	defer rows.Close() // Impedir vazamento de memória

	for rows.Next() {
		var buyer domain.Buyer

		err := rows.Scan(&buyer.ID, &buyer.CardNumberId, &buyer.FirstName, &buyer.LastName)
		if err != nil {
			return buyers, err
		}

		buyers = append(buyers, buyer)
	}

	return buyers, nil
}

func (r *repository) GetById(ctx context.Context, id int) (domain.Buyer, error) {
	var buyer domain.Buyer

	rows, err := r.db.QueryContext(ctx, SqlGetById, id)
	if err != nil {
		return domain.Buyer{}, err
	}

	defer rows.Close() // Impedir vazamento de memória

	for rows.Next() {
		err := rows.Scan(&buyer.ID, &buyer.CardNumberId, &buyer.FirstName, &buyer.LastName)
		if err != nil {
			return domain.Buyer{}, fmt.Errorf("buyer with id (%d) not founded", id)
		}

	}

	return buyer, nil
}

func (r repository) Create(ctx context.Context, buyer domain.Buyer) (domain.Buyer, error) {
	res, err := r.db.ExecContext(ctx, SqlStore, &buyer.CardNumberId, &buyer.FirstName, &buyer.LastName)
	if err != nil {
		return domain.Buyer{}, err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return domain.Buyer{}, fmt.Errorf("error while saving")
	}

	lastID, err := res.LastInsertId()
	if err != nil || lastID < 1 {
		return domain.Buyer{}, err
	}

	buyer.ID = int(lastID)

	return buyer, nil
}

func (r repository) Update(ctx context.Context, buyer domain.Buyer) (domain.Buyer, error) {
	res, err := r.db.ExecContext(
		ctx,
		SqlUpdate,
		&buyer.CardNumberId,
		&buyer.FirstName,
		&buyer.LastName,
		&buyer.ID,
	)
	if err != nil {
		return domain.Buyer{}, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return domain.Buyer{}, fmt.Errorf("buyer wiht id (%d) not founded", buyer.ID)
	}

	return buyer, nil
}

func (r repository) Delete(ctx context.Context, id int) error {
	res, err := r.db.ExecContext(ctx, SqlDelete, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("buyer with id (%d) not founded", id)
	}

	return nil
}

func (r *repository) GetBuyerOrdersById(ctx context.Context, id int) (domain.BuyerTotalOrders, error) {

	var buyerData domain.BuyerTotalOrders

	buyerData.ID = 0

	rows, err := r.db.QueryContext(ctx, SqlBuyerWithOrdersById, id)
	if err != nil {
		return domain.BuyerTotalOrders{}, err
	}

	defer rows.Close() // Impedir vazamento de memória

	for rows.Next() {
		err := rows.Scan(&buyerData.ID, &buyerData.CardNumberId, &buyerData.FirstName, &buyerData.LastName, &buyerData.PurchaseOrdersCount)
		if err != nil {
			return domain.BuyerTotalOrders{}, fmt.Errorf("buyer with id (%d) not founded", id)
		}
	}

	if buyerData.ID == 0 {
		return domain.BuyerTotalOrders{}, fmt.Errorf("buyer with id (%d) not founded", id)
	}

	return buyerData, nil
}

func (r *repository) GetBuyerTotalOrders(ctx context.Context) ([]domain.BuyerTotalOrders, error) {

	var buyersData []domain.BuyerTotalOrders

	rows, err := r.db.QueryContext(ctx, SqlBuyersWithOrders)
	if err != nil {
		return nil, err
	}

	defer rows.Close() // Impedir vazamento de memória

	for rows.Next() {
		var rowData domain.BuyerTotalOrders
		err := rows.Scan(&rowData.ID, &rowData.CardNumberId, &rowData.FirstName, &rowData.LastName, &rowData.PurchaseOrdersCount)
		if err != nil {
			return nil, err
		}
		buyersData = append(buyersData, rowData)
	}

	return buyersData, nil
}

func (r *repository) ValidadeCardNumberId(ctx context.Context, id int, cardNumber string) (bool, error) {

	var idExists int
	stmt, err := r.db.PrepareContext(ctx, SqlUniqueCardNumberId)

	if err != nil {
		return false, err
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(ctx, id, cardNumber).Scan(&idExists)

	return idExists == 0, nil
}
