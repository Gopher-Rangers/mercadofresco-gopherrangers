package productrecord

import (
	"database/sql"
	"fmt"
)

const (
	GET    = `SELECT * FROM product_records`
	GETALL = `SELECT p.id, p.description, COUNT(pr.product_id)
				AS records_count
				FROM product_records pr 
				JOIN products p ON p.id = pr.product_id
				GROUP BY p.id`
	GETBYID = `SELECT pr.product_id, p.description, COUNT(pr.product_id)
				AS records_count
				FROM product_records pr 
				JOIN products p ON p.id = pr.product_id
				WHERE pr.product_id = ?
				GROUP BY pr.product_id`
	STORE = `INSERT INTO product_records (last_update_date, purchase_price,
		sale_price, product_id) VALUES (?, ?, ?, ?)`
	LAST_ID = "SELECT MAX(id) as last_id FROM product_records"
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
	LastId() (int, error)
	Store(prod ProductRecord, id int) (ProductRecord, error)
	GetById(id int) (ProductRecordGet, error)
	GetAll() ([]ProductRecordGet, error)
	Get() ([]ProductRecord, error)
}

type repository struct {
	db *sql.DB
}

func (r *repository) LastId() (int, error) {
	var lastId int
	row := r.db.QueryRow(LAST_ID)
	err := row.Scan(&lastId)
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Store(prod ProductRecord, id int) (ProductRecord, error) {
	stmt, err := r.db.Prepare(STORE)
	if err != nil {
		return ProductRecord{}, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(&prod.LastUpdateDate, &prod.PurchasePrice,
		&prod.SalePrice, &prod.ProductId)
	if err != nil {
		return ProductRecord{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return ProductRecord{}, fmt.Errorf("falha ao salvar")
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return ProductRecord{}, err
	}
	prod.ID = int(lastId)
	return prod, nil
}

func (r *repository) GetById(id int) (ProductRecordGet, error) {
	var prod ProductRecordGet
	stmt, err := r.db.Prepare(GETBYID)
	if err != nil {
		return ProductRecordGet{}, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&prod.ProductId, &prod.Description,
		&prod.RecordsCount)
	if err != nil {
		return ProductRecordGet{}, fmt.Errorf("produt record %d não encontrado", id)
	}
	return prod, nil
}

func (r *repository) GetAll() ([]ProductRecordGet, error) {
	var ps []ProductRecordGet
	rows, err := r.db.Query(GETALL)
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

func (r *repository) Get() ([]ProductRecord, error) {
	var ps []ProductRecord
	rows, err := r.db.Query(GETALL)
	if err != nil {
		return ps, err
	}
	defer rows.Close()
	for rows.Next() {
		var prod ProductRecord
		err := rows.Scan(&prod.ID, &prod.LastUpdateDate, &prod.PurchasePrice,
			&prod.SalePrice, &prod.ProductId)
		if err != nil {
			return ps, err
		}
		ps = append(ps, prod)
	}
	return ps, nil
}

