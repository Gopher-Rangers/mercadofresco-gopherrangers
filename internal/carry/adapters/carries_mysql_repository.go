package adapters

import (
	"database/sql"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
)

type mysqlCarryRepository struct {
	db *sql.DB
}

func NewMySqlCarryRepository(db *sql.DB) usecases.RepositoryCarry {
	return &mysqlCarryRepository{db: db}
}

func CreateCarry(carry domain.Carry) (domain.Carry, error) {
	return domain.Carry{}, nil
}

func GetCaryPerLocality(id int) (domain.Carry, error) {
	return domain.Carry{}, nil
}
func GetCarryByCid(cid string) (domain.Carry, error) {
	return domain.Carry{}, nil
}
