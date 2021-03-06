package seller

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	GetOne(ctx context.Context, id int) (Seller, error)
	GetAll(ctx context.Context) ([]Seller, error)
	Create(ctx context.Context, cid int, companyName, address, telephone string, localityID int) (Seller, error)
	Update(ctx context.Context, cid int, companyName, address, telephone string, localityID int, seller Seller) (Seller, error)
	Delete(ctx context.Context, id int) error
}

const (
	GETALL  = "SELECT * FROM sellers"
	GETBYID = "SELECT * FROM sellers WHERE id=?"
	INSERT  = "INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (?,?,?,?,?)"
	UPDATE  = "UPDATE sellers SET cid=?, company_name=?, address=?, telephone=?, locality_id=? WHERE id=?"
	DELETE  = "DELETE FROM sellers WHERE id=?"
)

type mariaDBRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) Repository {
	return &mariaDBRepository{db: db}
}

func (m mariaDBRepository) GetOne(ctx context.Context, id int) (Seller, error) {
	var seller Seller

	rows, err := m.db.QueryContext(ctx, GETBYID, id)

	if err != nil {
		return seller, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&seller.Id, &seller.CompanyId, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityID)

		if err != nil {
			return seller, err
		}

		return seller, nil
	}

	err = rows.Err()

	if err != nil {
		return Seller{}, err
	}

	return seller, fmt.Errorf("id does not exists")
}

func (m *mariaDBRepository) GetAll(ctx context.Context) ([]Seller, error) {
	var sellerList []Seller

	rows, err := m.db.QueryContext(ctx, GETALL)

	if err != nil {
		return sellerList, err
	}

	defer rows.Close()

	for rows.Next() {
		var seller Seller

		err := rows.Scan(&seller.Id, &seller.CompanyId, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityID)

		if err != nil {
			return sellerList, err
		}

		sellerList = append(sellerList, seller)
	}

	return sellerList, err
}

func (m *mariaDBRepository) Create(ctx context.Context, cid int, companyName, address, telephone string, localityID int) (Seller, error) {
	var seller Seller

	seller = Seller{CompanyId: cid, CompanyName: companyName, Address: address, Telephone: telephone, LocalityID: localityID}

	stmt, err := m.db.PrepareContext(ctx, INSERT)

	if err != nil {
		return seller, err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, &seller.CompanyId, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityID)

	if err != nil {
		return seller, err
	}

	lastID, err := res.LastInsertId()

	if err != nil {
		return seller, err
	}

	seller.Id = int(lastID)

	return seller, nil
}

func (m *mariaDBRepository) Update(ctx context.Context, cid int, companyName, address, telephone string, localityID int, seller Seller) (Seller, error) {

	seller.CompanyId = cid
	seller.CompanyName = companyName
	seller.Address = address
	seller.Telephone = telephone
	seller.LocalityID = localityID

	stmt, err := m.db.PrepareContext(ctx, UPDATE)

	if err != nil {
		return seller, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, &seller.CompanyId, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityID, &seller.Id)

	if err != nil {
		return seller, err
	}

	return seller, nil
}

func (m *mariaDBRepository) Delete(ctx context.Context, id int) error {

	stmt, err := m.db.PrepareContext(ctx, DELETE)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
