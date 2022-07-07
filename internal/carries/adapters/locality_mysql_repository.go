package adapters

import (
	"database/sql"
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carries/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carries/usecases"
)

const (
	queryGetLocalityByID = "SELECT COUNT(*) AS carries_count FROM localities WHERE id=? "
)

type mysqlLocalityRepository struct {
	db *sql.DB
}

func NewMySqlLocalityRepository(db *sql.DB) usecases.RepositoryLocality {
	return &mysqlLocalityRepository{db: db}
}

func (r mysqlLocalityRepository) GetLocalityByID(id int) (domain.Locality, error) {
	var locality domain.Locality

	stmt := r.db.QueryRow(queryGetLocalityByID, id)

	err := stmt.Scan()

	if err != nil {
		return domain.Locality{}, fmt.Errorf("o id: %d não foi encontrado", id)
	}

	return locality, nil
}
