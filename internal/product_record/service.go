package productrecord

import (
	"fmt"
)

const (
	ERROR_UNIQUE_PRODUCT_ID = "the product id must be unique"
)

type Service interface {
	Store(prod ProductRecord) (ProductRecord, error)
	GetAll() ([]ProductRecordGet, error)
	GetById(id int) (ProductRecordGet, error)
	Get() ([]ProductRecord, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) checkProductId(prod ProductRecord) bool {
	ps, _ := s.Get()
	for i := range ps {
		if ps[i].ProductId == prod.ProductId && ps[i].ID != prod.ID {
			return false
		}
	}
	return true
}

func (s *service) Store(prod ProductRecord) (ProductRecord, error) {
	if !s.checkProductId(prod) {
		return ProductRecord{}, fmt.Errorf(ERROR_UNIQUE_PRODUCT_ID)
	}
	lastId, err := s.repository.LastId()
	if err != nil {
		return ProductRecord{}, err
	}
	lastId++
	product, err := s.repository.Store(prod, lastId)
	if err != nil {
		return ProductRecord{}, err
	}
	return product, nil
}

func (s *service) GetById(id int) (ProductRecordGet, error) {
	ps, err := s.repository.GetById(id)
	if err != nil {
		return ProductRecordGet{}, err
	}
	return ps, nil
}

func (s *service) GetAll() ([]ProductRecordGet, error) {
	ps, _ := s.repository.GetAll()
	return ps, nil
}

func (s *service) Get() ([]ProductRecord, error) {
	ps, _ := s.repository.Get()
	return ps, nil
}
