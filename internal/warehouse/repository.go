package warehouse

import (
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
)

type Warehouse struct {
	ID             int    `json:"id"`
	WarehouseCode  string `json:"warehouse_code" binding:"required"`
	Address        string `json:"address"`
	Telephone      string `json:"telephone"`
	MinCapacity    int    `json:"minimun_capacity"`
	MinTemperature int    `json:"minimun_temperature"`
}

type Repository interface {
	GetAll() []Warehouse
	GetByID(id int) (Warehouse, error)
	CreateWarehouse(
		id int,
		code,
		address,
		tel string,
		minCap,
		minTemp int) (Warehouse, error)
	UpdatedWarehouseID(id int, code string) (Warehouse, error)
	DeleteWarehouse(id int) error
	IncrementID() int
	FindByWarehouseCode(code string) (Warehouse, error)
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
	return Warehouse{}, fmt.Errorf("o id: %d não foi encontrado", id)
}

func (r *repository) CreateWarehouse(
	id int,
	code,
	address,
	tel string,
	minCap,
	minTemp int) (Warehouse, error) {

	var ListWarehouse []Warehouse

	r.db.Read(&ListWarehouse)

	w := Warehouse{id, code, address, tel, minCap, minTemp}

	ListWarehouse = append(ListWarehouse, w)

	r.db.Write(ListWarehouse)

	return w, nil
}

func (r *repository) UpdatedWarehouseID(id int, code string) (Warehouse, error) {
	var ListWarehouse []Warehouse

	r.db.Read(&ListWarehouse)

	for i := range ListWarehouse {
		if ListWarehouse[i].ID == id {
			ListWarehouse[i].WarehouseCode = code
			r.db.Write(ListWarehouse)
			return ListWarehouse[i], nil
		}
	}
	return Warehouse{}, fmt.Errorf("o id: %d informado não existe", id)
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
	return fmt.Errorf("não foi achado warehouse com esse id: %d", id)
}

func (r repository) IncrementID() int {
	var ListWarehouse []Warehouse
	r.db.Read(&ListWarehouse)

	return len(ListWarehouse) + 1
}

func (r repository) FindByWarehouseCode(code string) (Warehouse, error) {
	var ListWarehouse []Warehouse

	r.db.Read(&ListWarehouse)

	for _, warehouse := range ListWarehouse {
		if warehouse.WarehouseCode == code {
			return warehouse, nil
		}
	}
	return Warehouse{},
		fmt.Errorf("o warehouse com esse `warehouse_code`: %s não foi encontrado", code)
}
