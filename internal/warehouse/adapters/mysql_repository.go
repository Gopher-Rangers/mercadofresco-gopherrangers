package adapters

import (
	"database/sql"
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/usecases"
)

type mysqlRepository struct {
	db *sql.DB
}

const (
	queryGetAll = "SELECT * FROM warehouse"

	queryGetByID = "SELECT * FROM warehouse WHERE id=?"

	queryCreateWarehouse = "INSERT INTO warehouse (warehouse_code, address, telephone, locality_id) VALUES (?, ?, ?, ?)"

	queryFindByWarehouseCode = "SELECT * FROM warehouse WHERE warehouse_code=?"

	queryUpdateWarehouse = "UPDATE warehouse SET warehouse_code=? WHERE id=?"

	queryDeleteWarehouse = "DELETE FROM warehouse WHERE id=?"
)

func NewMySqlRepository(db *sql.DB) usecases.Repository {
	return &mysqlRepository{db: db}
}

func (r mysqlRepository) GetAll() []domain.Warehouse {

	rows, err := r.db.Query(queryGetAll)

	if err != nil {
		return []domain.Warehouse{}
	}

	defer rows.Close()

	warehouses := []domain.Warehouse{}

	for rows.Next() {
		w := domain.Warehouse{}
		rows.Scan(&w.ID, &w.WarehouseCode, &w.Address, &w.Telephone, &w.LocalityID)
		warehouses = append(warehouses, w)
	}

	if err = rows.Err(); err != nil {
		return []domain.Warehouse{}
	}

	return warehouses
}

func (r mysqlRepository) GetByID(id int) (domain.Warehouse, error) {
	var warehouse domain.Warehouse

	stmt := r.db.QueryRow(queryGetByID, id)

	err := stmt.Scan(&warehouse.ID, &warehouse.WarehouseCode, &warehouse.Address, &warehouse.Telephone, &warehouse.LocalityID)

	if err != nil {
		return domain.Warehouse{}, fmt.Errorf("o id: %d não foi encontrado", id)
	}

	return warehouse, nil
}

func (r *mysqlRepository) CreateWarehouse(
	code,
	address,
	tel string,
	localityID int) (domain.Warehouse, error) {

	stmt, err := r.db.Prepare(queryCreateWarehouse)

	if err != nil {
		return domain.Warehouse{}, fmt.Errorf("erro ao preparar a query")
	}

	defer stmt.Close()

	result, err := stmt.Exec(code, address, tel, localityID)

	if err != nil {
		return domain.Warehouse{}, fmt.Errorf("erro ao executar a query")
	}

	id, err := result.LastInsertId()

	if err != nil {
		return domain.Warehouse{}, fmt.Errorf("falha ao obter o id no banco de dados")
	}

	return domain.Warehouse{
		ID:            int(id),
		WarehouseCode: code,
		Address:       address,
		Telephone:     tel,
		LocalityID:    localityID,
	}, nil

}

func (r *mysqlRepository) UpdatedWarehouseID(id int, code string) (domain.Warehouse, error) {
	warehouse, err := r.GetByID(id)

	if err != nil {
		return domain.Warehouse{}, fmt.Errorf("id inexistente no banco de dados")
	}

	stmt, err := r.db.Prepare(queryUpdateWarehouse)

	if err != nil {
		return domain.Warehouse{}, fmt.Errorf("erro ao preparar a query")
	}

	defer stmt.Close()

	_, err = stmt.Exec(code, id)

	if err != nil {
		return domain.Warehouse{}, fmt.Errorf("erro ao executar a query")
	}

	return domain.Warehouse{
		ID:            id,
		WarehouseCode: code,
		Address:       warehouse.Address,
		Telephone:     warehouse.Telephone,
		LocalityID:    warehouse.LocalityID,
	}, nil
}

func (r *mysqlRepository) DeleteWarehouse(id int) error {
	stmt, err := r.db.Prepare(queryDeleteWarehouse)

	if err != nil {
		return fmt.Errorf("erro ao preparar a query: %v", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return fmt.Errorf("erro ao executar query: %v", err)
	}

	return nil
}

func (r mysqlRepository) FindByWarehouseCode(code string) (domain.Warehouse, error) {
	var warehouse domain.Warehouse

	stmt := r.db.QueryRow(queryFindByWarehouseCode, code)

	err := stmt.Scan(&warehouse.ID, &warehouse.WarehouseCode, &warehouse.Address, &warehouse.Telephone, &warehouse.LocalityID)

	if err != nil {
		return domain.Warehouse{}, fmt.Errorf("o warehouse com esse `warehouse_code`: %s não foi encontrado", code)
	}

	return warehouse, nil

}
