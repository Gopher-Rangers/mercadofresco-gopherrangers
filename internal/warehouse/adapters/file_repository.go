package adapters

import (
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/usecases"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
)

type fileRepository struct {
	db store.Store
}

func NewFileRepository(db store.Store) usecases.Repository {
	return &fileRepository{db: db}
}

func (r fileRepository) GetAll() []domain.Warehouse {
	var ListWarehouse []domain.Warehouse
	r.db.Read(&ListWarehouse)

	return ListWarehouse
}

func (r fileRepository) GetByID(id int) (domain.Warehouse, error) {
	var ListWarehouse []domain.Warehouse

	r.db.Read(&ListWarehouse)

	for _, warehouse := range ListWarehouse {
		if warehouse.ID == id {
			return warehouse, nil
		}
	}
	return domain.Warehouse{}, fmt.Errorf("o id: %d não foi encontrado", id)
}

func (r *fileRepository) CreateWarehouse(
	code,
	address,
	tel string,
	minCap,
	minTemp int) (domain.Warehouse, error) {

	var ListWarehouse []domain.Warehouse

	err := r.db.Read(&ListWarehouse)

	if err != nil {
		return domain.Warehouse{}, fmt.Errorf("não foi possível ler o arquivo")
	}

	id := r.incrementID()

	w := domain.Warehouse{
		ID:             id,
		WarehouseCode:  code,
		Address:        address,
		Telephone:      tel,
		MinCapacity:    minCap,
		MinTemperature: minTemp,
	}

	ListWarehouse = append(ListWarehouse, w)

	r.db.Write(ListWarehouse)

	return w, nil
}

func (r *fileRepository) UpdatedWarehouseID(id int, code string) (domain.Warehouse, error) {
	var ListWarehouse []domain.Warehouse

	r.db.Read(&ListWarehouse)

	for i := range ListWarehouse {
		if ListWarehouse[i].ID == id {

			ListWarehouse[i].WarehouseCode = code

			r.db.Write(ListWarehouse)
			return ListWarehouse[i], nil
		}
	}
	return domain.Warehouse{}, fmt.Errorf("o id: %d informado não existe", id)
}

func (r *fileRepository) DeleteWarehouse(id int) error {
	var ListWarehouse []domain.Warehouse

	r.db.Read(&ListWarehouse)

	for i := range ListWarehouse {
		if ListWarehouse[i].ID == id {

			ListWarehouse = append(ListWarehouse[:i], ListWarehouse[i+1:]...)

			r.db.Write(ListWarehouse)

			return nil
		}
	}
	return fmt.Errorf("não foi achado warehouse com esse id: %d", id)
}

func (r fileRepository) incrementID() int {
	var ListWarehouse []domain.Warehouse

	r.db.Read(&ListWarehouse)

	return len(ListWarehouse) + 1
}

func (r fileRepository) FindByWarehouseCode(code string) (domain.Warehouse, error) {
	var ListWarehouse []domain.Warehouse

	r.db.Read(&ListWarehouse)

	for _, warehouse := range ListWarehouse {
		if warehouse.WarehouseCode == code {
			return warehouse, nil
		}
	}
	return domain.Warehouse{},
		fmt.Errorf("o warehouse com esse `warehouse_code`: %s não foi encontrado", code)
}
