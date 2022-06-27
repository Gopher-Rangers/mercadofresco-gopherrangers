package seller

import (
	"database/sql"
	"errors"
)

type Repository interface {
	GetOne(id int) (Seller, error)
	GetAll() ([]Seller, error)
	Create(cid int, companyName, address, telephone string) (Seller, error)
	Update(cid int, companyName, address, telephone string, seller Seller) (Seller, error)
	Delete(id int) error
}

type mariaDBRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) Repository {
	return &mariaDBRepository{db: db}
}

func (m mariaDBRepository) GetOne(id int) (Seller, error) {
	var seller Seller

	rows, err := m.db.Query("SELECT *  FROM seller WHERE id=?", id)

	if err != nil {
		return seller, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&seller.Id, &seller.CompanyId, &seller.CompanyName, &seller.Address, &seller.Telephone)

		if err != nil {
			return seller, err
		}
	}

	err = rows.Err()
	if err != nil {
		return Seller{}, err
	}

	return seller, nil
}

func (m *mariaDBRepository) GetAll() ([]Seller, error) {
	var sellerList []Seller

	if len(sellerList) < 0 {
		return sellerList, errors.New("erro ao inicializar a lista")
	}

	rows, err := m.db.Query("SELECT * FROM seller")

	if err != nil {
		return sellerList, err
	}

	defer rows.Close()

	for rows.Next() {
		var seller Seller

		err := rows.Scan(&seller.Id, &seller.CompanyId, &seller.CompanyName, &seller.Address, &seller.Telephone)

		if err != nil {
			return sellerList, err
		}

		sellerList = append(sellerList, seller)
	}

	return sellerList, err
}

func (m *mariaDBRepository) Create(cid int, companyName, address, telephone string) (Seller, error) {
	var seller Seller

	sellerList, err := m.GetAll()
	if err != nil {
		return seller, err
	}

	for i := range sellerList {
		if sellerList[i].CompanyId == cid {
			return seller, errors.New("the cid already exists")
		}
	}

	seller = Seller{CompanyId: cid, CompanyName: companyName, Address: address, Telephone: telephone}

	stmt, err := m.db.Prepare("INSERT INTO seller (cid, company_name, address, telephone) VALUES (?,?,?,?)")

	if err != nil {
		return seller, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(&seller.CompanyId, &seller.CompanyName, &seller.Address, &seller.Telephone)

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

func (m *mariaDBRepository) Update(cid int, companyName, address, telephone string, seller Seller) (Seller, error) {

	seller, err := m.GetOne(seller.Id)

	if err != nil {
		return seller, err
	}

	seller.CompanyId = cid
	seller.CompanyName = companyName
	seller.Address = address
	seller.Telephone = telephone

	stmt, err := m.db.Prepare("UPDATE seller SET cid=?, company_name=?, address=?, telephone=? WHERE id=?")

	if err != nil {
		return seller, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&seller.CompanyId, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.Id)

	if err != nil {
		return seller, err
	}

	return seller, nil
}

func (m *mariaDBRepository) Delete(id int) error {
	stmt, err := m.db.Prepare("DELETE FROM seller WHERE id=?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
