package productrecord

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	GETALL = `SELECT p.id, p.description, COUNT(pr.product_id)
				FROM product_records pr 
				JOIN products p ON p.id = pr.product_id
				GROUP BY p.id`
	GETBYID = `SELECT pr.product_id, p.description, COUNT(pr.product_id)
				FROM product_records pr 
				JOIN products p ON p.id = pr.product_id
				WHERE pr.product_id = ?
				GROUP BY pr.product_id`
	STORE = `INSERT INTO product_records (last_update_date, purchase_price,
				sale_price, product_id) VALUES (?, ?, ?, ?)`
)

type ProductRecord struct {
	ID             int     `json:"id"`
	LastUpdateDate string  `json:"last_update_date" validate:"required"`
	PurchasePrice  float64 `json:"purchase_price" validate:"required"`
	SalePrice      float64 `json:"sale_price" validate:"required,gt=0"`
	ProductId      int     `json:"product_id" validate:"required,gt=0"`
}

type ProductRecordGet struct {
	ProductId    int    `json:"product_id"`
	Description  string `json:"description"`
	RecordsCount int64  `json:"records_count"`
}

type Repository interface {
	Store(ctx context.Context, prod ProductRecord) (ProductRecord, error)
	GetById(ctx context.Context, id int) (ProductRecordGet, error)
	GetAll(ctx context.Context) ([]ProductRecordGet, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Store(ctx context.Context, prod ProductRecord) (
	ProductRecord, error) {
	stmt, err := r.db.PrepareContext(ctx, STORE)
	if err != nil {
		return ProductRecord{}, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, &prod.LastUpdateDate,
		&prod.PurchasePrice, &prod.SalePrice, &prod.ProductId)
	if err != nil {
		return ProductRecord{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return ProductRecord{}, fmt.Errorf("fail to save")
	}
	lastId, _ := result.LastInsertId()
	prod.ID = int(lastId)
	return prod, nil
}

func (r *repository) GetById(ctx context.Context, id int) (
	ProductRecordGet, error) {
	var prod ProductRecordGet
	stmt, err := r.db.PrepareContext(ctx, GETBYID)
	if err != nil {
		return ProductRecordGet{}, err
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(ctx, id).Scan(&prod.ProductId, &prod.Description,
		&prod.RecordsCount)
	if err != nil {
		return ProductRecordGet{}, fmt.Errorf("product record %d not found", id)
	}
	return prod, nil
}

func (r *repository) GetAll(ctx context.Context) ([]ProductRecordGet, error) {
	var ps []ProductRecordGet
	rows, err := r.db.QueryContext(ctx, GETALL)
	if err != nil {
		return ps, err
	}
	defer rows.Close()
	for rows.Next() {
		var prod ProductRecordGet
		err := rows.Scan(&prod.ProductId, &prod.Description,
			&prod.RecordsCount)
		if err != nil {
			return ps, err
		}
		ps = append(ps, prod)
	}
	return ps, nil
}
