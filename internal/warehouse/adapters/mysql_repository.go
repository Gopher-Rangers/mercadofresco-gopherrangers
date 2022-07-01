package adapters

import (
	"database/sql"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/usecases"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMySqlRepository(db *sql.DB) usecases.Repository {
	return &mysqlRepository{db: db}
}

func (r mysqlRepository) GetAll() []domain.Warehouse {
	const query = `SELECT * FROM warehouse`

	rows, err := r.db.Query(query)

	if err != nil {
		return []domain.Warehouse{}
	}

	defer rows.Close()

	warehouses := []domain.Warehouse{}

	for rows.Next() {
		w := domain.Warehouse{}
		rows.Scan(&w.ID, &w.Address, &w.MinCapacity, &w.MinTemperature, &w.Telephone, &w.WarehouseCode)
		warehouses = append(warehouses, w)
	}

	if err = rows.Err(); err != nil {
		return []domain.Warehouse{}
	}

	return warehouses
}

func (r mysqlRepository) GetByID(id int) (domain.Warehouse, error) {

}

func (r *mysqlRepository) CreateWarehouse(
	id int,
	code,
	address,
	tel string,
	minCap,
	minTemp int) (domain.Warehouse, error) {

}

func (r *mysqlRepository) UpdatedWarehouseID(id int, code string) (domain.Warehouse, error) {

}

func (r *mysqlRepository) DeleteWarehouse(id int) error {

}

func (r mysqlRepository) IncrementID() int {

}

func (r mysqlRepository) FindByWarehouseCode(code string) (domain.Warehouse, error) {

}
