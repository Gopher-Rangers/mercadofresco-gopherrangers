package products

import (
	"database/sql"
)

const (
	GETALL = "SELECT * FROM products"
)

type sqlDbRepository struct {
	db *sql.DB
}

func NewDBRepository(db *sql.DB) Repository {
	return &sqlDbRepository{db: db}
}

func (r *sqlDbRepository) LastID() (int, error) {
	return 0, nil
}

func (r *sqlDbRepository) Store(prod Product, id int) (Product, error) {
	return prod, nil
}

func (r *sqlDbRepository) GetAll() ([]Product, error) {
	var ps []Product
	rows, err := r.db.Query(GETALL)
	if err != nil {
		return ps, nil
	}
	defer rows.Close()
	for rows.Next() {
		var prod Product
		err := rows.Scan(&prod.ID, &prod.ProductCode, &prod.Description,
			&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
			&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
			&prod.FreezingRate, &prod.ProductTypeId, &prod.SellerId)
		if err != nil {
			return ps, err
		}
		ps = append(ps, prod)
	}
	return ps, nil
}

func (r *sqlDbRepository) GetById(id int) (Product, error) {
	return Product{}, nil
}

func (r *sqlDbRepository) Update(prod Product, id int) (Product, error) {
	return Product{}, nil
}

func (r *sqlDbRepository) Delete(id int) error {
	return nil
}
