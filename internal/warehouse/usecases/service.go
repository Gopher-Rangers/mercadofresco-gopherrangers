package usecases

import (
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/domain"
)

type Service interface {
	GetAll() []domain.Warehouse
	GetByID(id int) (domain.Warehouse, error)
	CreateWarehouse(
		code,
		address,
		tel string,
		minCap,
		minTemp int) (domain.Warehouse, error)
	UpdatedWarehouseID(id int, code string) (domain.Warehouse, error)
	DeleteWarehouse(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s service) GetAll() []domain.Warehouse {
	return s.repository.GetAll()
}

func (s service) GetByID(id int) (domain.Warehouse, error) {

	warehouse, err := s.repository.GetByID(id)

	if err != nil {
		return domain.Warehouse{}, err
	}

	return warehouse, nil

}

func (s service) CreateWarehouse(code, address, tel string, minCap, minTemp int) (domain.Warehouse, error) {

	_, err := s.repository.FindByWarehouseCode(code)

	if err == nil {
		return domain.Warehouse{}, fmt.Errorf("o `warehouse_code` já está em uso")
	}

	warehouse, err := s.repository.CreateWarehouse(code, address, tel, minCap, minTemp)

	if err != nil {
		return domain.Warehouse{}, err
	}

	return warehouse, nil
}

func (s service) UpdatedWarehouseID(id int, code string) (domain.Warehouse, error) {
	_, err := s.repository.FindByWarehouseCode(code)

	if err == nil {
		return domain.Warehouse{}, fmt.Errorf("o `warehouse_code` já está em uso")
	}

	warehouse, err := s.repository.UpdatedWarehouseID(id, code)

	if err != nil {
		return domain.Warehouse{}, err
	}

	return warehouse, nil
}

func (s service) DeleteWarehouse(id int) error {
	err := s.repository.DeleteWarehouse(id)

	if err != nil {
		return err
	}

	return nil
}
