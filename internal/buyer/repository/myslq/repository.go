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
	if err != nil {
		return buyer, err
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

	total, err := r.getTotalOrderByBuyer(ctx, id)
	if err != nil {
		return domain.BuyerTotalOrders{}, err
	}

	var buyerData domain.Buyer

	rows, err := r.db.QueryContext(ctx, SqlGetById, id)
	if err != nil {
		return domain.BuyerTotalOrders{}, err
	}

	defer rows.Close() // Impedir vazamento de memória

	for rows.Next() {
		err := rows.Scan(&buyerData.ID, &buyerData.CardNumberId, &buyerData.FirstName, &buyerData.LastName)
		if err != nil {
			return domain.BuyerTotalOrders{}, fmt.Errorf("buyer with id (%d) not founded", id)
		}

	}

	fmt.Printf("test id final %d \n", id)

	fmt.Println(buyerData)

	return domain.BuyerTotalOrders{ID: buyerData.ID, CardNumberId: buyerData.CardNumberId, FirstName: buyerData.FirstName, LastName: buyerData.LastName, PurchaseOrdersCount: total}, nil
}

func (r repository) getTotalOrderByBuyer(ctx context.Context, id int) (int, error) {
	var buyerOrders int

	query := r.db.QueryRow(SqlCountOrdersByBuyerId, id)

	err := query.Scan(&buyerOrders)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Total order buyers %d ", buyerOrders)

	return buyerOrders, nil
}
