package adapters

import (
	"database/sql"
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
)

const (
	queryCreateCarry = "INSERT INTO carriers (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	queryGetByCid    = "SELECT * FROM carriers WHERE cid=? "
)

type mysqlCarryRepository struct {
	db *sql.DB
}

func NewMySqlCarryRepository(db *sql.DB) usecases.RepositoryCarry {
	return &mysqlCarryRepository{db: db}
}

func (r *mysqlCarryRepository) CreateCarry(carry domain.Carry) (domain.Carry, error) {
	stmt, err := r.db.Prepare(queryCreateCarry)

	if err != nil {
		return domain.Carry{}, fmt.Errorf("erro ao preparar a query")
	}

	defer stmt.Close()

	result, err := stmt.Exec(carry.Cid, carry.Name, carry.Address, carry.Telephone, carry.LocalityID)

	if err != nil {
		return domain.Carry{}, fmt.Errorf("erro ao executar a query")
	}

	id, err := result.LastInsertId()

	if err != nil {
		return domain.Carry{}, fmt.Errorf("falha ao obter o id no banco de dados")
	}

	return domain.Carry{
		ID:         int(id),
		Cid:        carry.Cid,
		Name:       carry.Name,
		Address:    carry.Address,
		Telephone:  carry.Telephone,
		LocalityID: carry.LocalityID,
	}, nil

}

func (r mysqlCarryRepository) GetCaryPerLocality(id int) (domain.Carry, error) {
	return domain.Carry{}, nil
}

func (r mysqlCarryRepository) GetCarryByCid(cid string) (domain.Carry, error) {
	var carry domain.Carry

	stmt := r.db.QueryRow(queryGetByCid, cid)

	err := stmt.Scan(&carry.ID, &carry.Cid, &carry.Name, &carry.Address, &carry.Telephone, &carry.LocalityID)

	if err != nil {
		return domain.Carry{}, fmt.Errorf("o carry com esse `cid`: %s não foi encontrado", carry.Cid)
	}

	return domain.Carry{}, nil
}
