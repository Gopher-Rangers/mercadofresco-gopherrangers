package seller

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	GetOne(ctx context.Context, id int) (Seller, error)
	GetAll(ctx context.Context) ([]Seller, error)
	Create(ctx context.Context, cid int, companyName, address, telephone string, locality int) (Seller, error)
	Update(ctx context.Context, cid int, companyName, address, telephone string, seller Seller) (Seller, error)
	Delete(ctx context.Context, id int) error
}

type mariaDBRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) Repository {
	return &mariaDBRepository{db: db}
}

func (m mariaDBRepository) GetOne(ctx context.Context, id int) (Seller, error) {
	var seller Seller

	rows, err := m.db.QueryContext(ctx, "SELECT localities.id FROM seller LEFT JOIN localities ON seller.locality_id= localities.id WHERE seller.id = ?", id)

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

	rows, err := m.db.QueryContext(ctx, "SELECT * FROM seller")

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

	stmt, err := m.db.PrepareContext(ctx, "INSERT INTO seller (cid, company_name, address, telephone, locality_id) VALUES (?,?,?,?,?)")

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

func (m *mariaDBRepository) Update(ctx context.Context, cid int, companyName, address, telephone string, seller Seller) (Seller, error) {

	seller.CompanyId = cid
	seller.CompanyName = companyName
	seller.Address = address
	seller.Telephone = telephone

	stmt, err := m.db.PrepareContext(ctx, "UPDATE seller SET cid=?, company_name=?, address=?, telephone=? WHERE id=?")

	if err != nil {
		return seller, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, &seller.CompanyId, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.Id)

	if err != nil {
		return seller, err
	}

	return seller, nil
}

func (m *mariaDBRepository) Delete(ctx context.Context, id int) error {
	stmt, err := m.db.PrepareContext(ctx, "DELETE FROM seller WHERE id=?")

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
