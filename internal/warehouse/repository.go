package warehouse

import (
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
)

type Warehouse struct {
	ID              int
	Warehouse_code  string
	Address         string
	Telephone       string
	Min_Capacity    int
	Min_Temperature int
}

type Repository interface {
	GetAll() []Warehouse
	GetByID(id int) (Warehouse, error)
	CreateWarehouse(id int, code, address, tel string, minCap, minTemp int) (Warehouse, error)
	UpdatedWarehouseID(id int, code string) (Warehouse, error)
	DeleteWarehouse(id int) error
	IncrementID() int
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r repository) GetAll() []Warehouse {
	var ListWarehouse []Warehouse
	r.db.Read(&ListWarehouse)

	return ListWarehouse
}

func (r repository) GetByID(id int) (Warehouse, error) {
	var ListWarehouse []Warehouse
	r.db.Read(&ListWarehouse)

	for _, warehouse := range ListWarehouse {
		if warehouse.ID == id {
			return warehouse, nil
		}
	}
	return Warehouse{}, fmt.Errorf("O %d não foi encontrado", id)
}

func (r *repository) CreateWarehouse(id int, code, address, tel string, minCap, minTemp int) (Warehouse, error) {
	var ListWarehouse []Warehouse
	r.db.Read(&ListWarehouse)

	w := Warehouse{id, code, address, tel, minCap, minTemp}

	for _, warehouse := range ListWarehouse {
		if warehouse.Warehouse_code == code {
			return Warehouse{}, fmt.Errorf("O %s já existe no banco de dados", code)
		}
	}

	ListWarehouse = append(ListWarehouse, w)

	r.db.Write(ListWarehouse)

	return w, nil
}

func (r *repository) UpdatedWarehouseID(id int, code string) (Warehouse, error) {
	var ListWarehouse []Warehouse
	r.db.Read(&ListWarehouse)

	for i := range ListWarehouse {
		if ListWarehouse[i].ID == id {
			ListWarehouse[i].Warehouse_code = code
			r.db.Write(ListWarehouse)
			return ListWarehouse[i], nil
		}
	}
	return Warehouse{}, fmt.Errorf("O %d informado não existe", id)
}

func (r *repository) DeleteWarehouse(id int) error {
	var ListWarehouse []Warehouse
	r.db.Read(&ListWarehouse)

	for i := range ListWarehouse {
		if ListWarehouse[i].ID == id {
			ListWarehouse = append(ListWarehouse[:i], ListWarehouse[i+1:]...)
			r.db.Write(ListWarehouse)
			return nil
		}
	}
	return fmt.Errorf("Não foi achado warehouse com esse %d", id)
}

func (r repository) IncrementID() int {
	var ListWarehouse []Warehouse
	r.db.Read(&ListWarehouse)

	return len(ListWarehouse)
}
