package products

import (
	"fmt"
	"context"
)

const (
	ERROR_UNIQUE_PRODUCT_CODE = "the product code must be unique"
)

type Service interface {
	Store(ctx context.Context, prod Product) (Product, error)
	GetAll(ctx context.Context) ([]Product, error)
	GetById(ctx context.Context, id int) (Product, error)
	Update(ctx context.Context, prod Product, id int) (Product, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Store(ctx context.Context, prod Product) (Product, error) {
	if !s.repository.CheckProductCode(prod.ID, prod.ProductCode) {
		return Product{}, fmt.Errorf(ERROR_UNIQUE_PRODUCT_CODE)
	}
	product, err := s.repository.Store(ctx, prod)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) GetAll(ctx context.Context) ([]Product, error) {
	ps, _ := s.repository.GetAll(ctx)
	return ps, nil
}

func (s *service) GetById(ctx context.Context, id int) (Product, error) {
	ps, err := s.repository.GetById(ctx, id)
	if err != nil {
		return Product{}, err
	}
	return ps, nil
}

func (s *service) Update(ctx context.Context, prod Product, id int) (Product, error) {
	if !s.repository.CheckProductCode(prod.ID, prod.ProductCode) {
		return Product{}, fmt.Errorf(ERROR_UNIQUE_PRODUCT_CODE)
	}
	product, err := s.repository.Update(ctx, prod, id)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
