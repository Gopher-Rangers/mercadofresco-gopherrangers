package adapters

import (
	"database/sql"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
)

const (
	queryCreateCarry = "INSERT INTO carriers () VALUES (?, ?, ?, ?, ?)"
)

type mysqlCarryRepository struct {
	db *sql.DB
}

func NewMySqlCarryRepository(db *sql.DB) usecases.RepositoryCarry {
	return &mysqlCarryRepository{db: db}
}

func CreateCarry(r *mysqlCarryRepository) (domain.Carry, error) {
	stmt, err := r.db.Prepare()
}

func GetCaryPerLocality(r mysqlCarryRepository) (domain.Carry, error) {
	return domain.Carry{}, nil
}
func GetCarryByCid(r mysqlCarryRepository) (domain.Carry, error) {
	return domain.Carry{}, nil
}
