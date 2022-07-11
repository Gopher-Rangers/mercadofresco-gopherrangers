package adapters

import (
	"database/sql"
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
)

const (
	queryGetCarryLocalityByID = `SELECT ca.locality_id, lo.locality_name, COUNT(*) AS carries_count 
	FROM carriers AS ca
	INNER JOIN localities AS lo
	ON ca.locality_id = lo.id
	WHERE ca.locality_id = ? GROUP BY ca.locality_id`

	queryGetAllCarriesLocality = `SELECT ca.locality_id, lo.locality_name, COUNT(*) AS carries_count 
	FROM carriers AS ca
	INNER JOIN localities AS lo
	ON ca.locality_id = lo.id
	GROUP BY ca.locality_id
	`
)

type mysqlLocalityRepository struct {
	db *sql.DB
}

func NewMySqlLocalityRepository(db *sql.DB) usecases.RepositoryLocality {
	return &mysqlLocalityRepository{db: db}
}

func (r mysqlLocalityRepository) GetCarryLocalityByID(id int) (domain.Locality, error) {
	var locality domain.Locality

	stmt := r.db.QueryRow(queryGetCarryLocalityByID, id)

	err := stmt.Scan(&locality.ID, &locality.Name, &locality.Count)

	if err != nil {
		return domain.Locality{}, fmt.Errorf("o id: %d não foi encontrado", id)
	}

	return locality, nil
}

func (r mysqlLocalityRepository) GetAllCarriesLocality() ([]domain.Locality, error) {
	rows, err := r.db.Query(queryGetAllCarriesLocality)

	if err != nil {
		fmt.Println(err)
		return []domain.Locality{}, err

	}

	defer rows.Close()

	localities := []domain.Locality{}

	for rows.Next() {
		locality := domain.Locality{}

		rows.Scan(&locality.ID, &locality.Name, &locality.Count)

		localities = append(localities, locality)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return []domain.Locality{}, err
	}

	return localities, nil
}
