package adapters

import (
	"database/sql"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carries/usecases"
)

type mysqlCarryRepository struct {
	db *sql.DB
}

func NewMySqlCarryRepository(db *sql.DB) usecases.RepositoryCarry {
	return &mysqlCarryRepository{db: db}
}
