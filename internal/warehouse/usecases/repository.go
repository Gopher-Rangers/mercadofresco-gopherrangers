package usecases

import "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/domain"

type Repository interface {
	GetAll() []domain.Warehouse
	GetByID(id int) (domain.Warehouse, error)
	CreateWarehouse(
		code,
		address,
		tel string,
		localityID int) (domain.Warehouse, error)
	UpdatedWarehouseID(id int, code string) (domain.Warehouse, error)
	DeleteWarehouse(id int) error
	FindByWarehouseCode(code string) (domain.Warehouse, error)
}
