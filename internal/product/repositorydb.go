package products

import (
	"database/sql"
	"fmt"
)

const (
	GETALL  = "SELECT * FROM products"
	GETBYID = "SELECT * FROM products WHERE id=?"
	STORE   = `INSERT INTO products (product_code, description,
		width, height, length, net_weight, expiration_rate,
		recommended_freezing_temperature, freezing_rate,
		product_type_id, seller_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	UPDATE = `UPDATE products SET
		product_code=?, description=?, width=?, height=?,
		length=?, net_weight=?, expiration_rate=?,
		recommended_freezing_temperature=?, freezing_rate=?,
		product_type_id=?, seller_id=?
		WHERE id=?`
	DELETE  = "DELETE FROM products WHERE id=?"
	LAST_ID = "SELECT MAX(id) as last_id FROM products"
)

type sqlDbRepository struct {
	db *sql.DB
}

func NewDBRepository(db *sql.DB) Repository {
	return &sqlDbRepository{db: db}
}

func (r *sqlDbRepository) LastID() (int, error) {
	var lastId int
	row := r.db.QueryRow(LAST_ID)
	err := row.Scan(&lastId)
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (r *sqlDbRepository) Store(prod Product, id int) (Product, error) {
	stmt, err := r.db.Prepare(STORE)
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(&prod.ProductCode, &prod.Description,
		&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
		&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
		&prod.FreezingRate, &prod.ProductTypeId, &prod.SellerId)
	if err != nil {
		return Product{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Product{}, fmt.Errorf("falha ao salvar")
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return Product{}, err
	}
	prod.ID = int(lastId)
	return prod, nil
}

func (r *sqlDbRepository) GetAll() ([]Product, error) {
	var ps []Product
	rows, err := r.db.Query(GETALL)
	if err != nil {
		return ps, err
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
	var prod Product
	stmt, err := r.db.Prepare(GETBYID)
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&prod.ID, &prod.ProductCode, &prod.Description,
		&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
		&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
		&prod.FreezingRate, &prod.ProductTypeId, &prod.SellerId)
	if err != nil {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	return prod, nil
}

func (r *sqlDbRepository) Update(prod Product, id int) (Product, error) {
	stmt, err := r.db.Prepare(UPDATE)
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(&prod.ProductCode, &prod.Description,
		&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
		&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
		&prod.FreezingRate, &prod.ProductTypeId, &prod.SellerId, id)
	if err != nil {
		return Product{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	return prod, nil
}

func (r *sqlDbRepository) Delete(id int) error {
	stmt, err := r.db.Prepare(DELETE)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("produto %d não encontrado", id)
	}
	return nil
}
