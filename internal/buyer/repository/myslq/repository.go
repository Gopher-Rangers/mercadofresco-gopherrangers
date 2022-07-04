package myslq

import (
	"context"
	"database/sql"
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

	rows, err := r.db.QueryContext(ctx, sqlGetAll)
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

	rows, err := r.db.QueryContext(ctx, sqlGetById, id)
	if err != nil {
		return domain.Buyer{}, err
	}

	defer rows.Close() // Impedir vazamento de memória

	for rows.Next() {
		err := rows.Scan(&buyer.ID, &buyer.CardNumberId, &buyer.FirstName, &buyer.LastName)
		if err != nil {
			return domain.Buyer{}, err
		}

	}

	return buyer, nil
}

func (r repository) Create(ctx context.Context, buyer domain.Buyer) (domain.Buyer, error) {
	res, err := r.db.ExecContext(ctx, sqlStore, &buyer.CardNumberId, &buyer.FirstName, &buyer.LastName)
	if err != nil {
		return domain.Buyer{}, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return buyer, err
	}

	buyer.ID = int(lastID)

	return buyer, nil
}

func (r repository) Update(ctx context.Context, buyer domain.Buyer) (domain.Buyer, error) {
	_, err := r.db.ExecContext(
		ctx,
		sqlUpdate,
		&buyer.CardNumberId,
		&buyer.FirstName,
		&buyer.LastName,
		&buyer.ID,
	)
	if err != nil {
		return buyer, err
	}

	return buyer, nil
}

func (r repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		return err
	}

	return nil
}
