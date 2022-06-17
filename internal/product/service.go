package products

import (
	"fmt"
)

const (
	ERROR_UNIQUE_PRODUCT_CODE = "the product code must be unique"
)

type Service interface {
	Store(prod Product) (Product, error)
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Update(prod Product, id int) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) checkProductTypeId(prod Product) bool {
	ps, _ := s.GetAll()
	for i := range ps {
		if ps[i].ProductCode == prod.ProductCode && ps[i].ID != prod.ID {
			return false
		}
	}
	return true
}

func (s *service) Store(prod Product) (Product, error) {
	if !s.checkProductTypeId(prod) {
		return Product{}, fmt.Errorf(ERROR_UNIQUE_PRODUCT_CODE)
	}
	lastId, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}
	lastId++
	product, err := s.repository.Store(prod, lastId)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) GetById(id int) (Product, error) {
	ps, err := s.repository.GetById(id)
	if err != nil {
		return Product{}, err
	}
	return ps, nil
}

func (s *service) Update(prod Product, id int) (Product, error) {
	if !s.checkProductTypeId(prod) {
		return Product{}, fmt.Errorf(ERROR_UNIQUE_PRODUCT_CODE)
	}
	product, err := s.repository.Update(prod, id)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
