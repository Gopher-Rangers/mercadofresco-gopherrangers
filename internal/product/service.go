package products

import (
	"context"
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"
)

const (
	ERROR_INEXISTENT_SELLER = "the seller id doesn`t exist"
	ERROR_INEXISTENT_PRODUCT_TYPE = "the product type id doesn`t exist"
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
	sellerService seller.Service
}

func NewService(r Repository, sellerService seller.Service) Service {
	return &service{
		repository: r,
		sellerService: sellerService}
}

func (s *service) checkIfSellerExists(ctx context.Context, prod Product) bool {
	_, err := s.sellerService.GetOne(ctx, prod.ID)
	return err == nil
}

func (s *service) Store(ctx context.Context, prod Product) (Product, error) {
	if !s.repository.CheckProductType(ctx, prod.ProductTypeId) {
		return Product{}, fmt.Errorf(ERROR_INEXISTENT_PRODUCT_TYPE)
	}
	if !s.checkIfSellerExists(ctx, prod) {
		return Product{}, fmt.Errorf(ERROR_INEXISTENT_SELLER)
	}
	if !s.repository.CheckProductCode(ctx, prod.ID, prod.ProductCode) {
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

func (s *service) Update(ctx context.Context, prod Product, id int) (
	Product, error) {
	if !s.repository.CheckProductType(ctx, prod.ProductTypeId) {
		return Product{}, fmt.Errorf(ERROR_INEXISTENT_PRODUCT_TYPE)
	}
	if !s.checkIfSellerExists(ctx, prod) {
		return Product{}, fmt.Errorf(ERROR_INEXISTENT_SELLER)
	}
	if !s.repository.CheckProductCode(ctx, prod.ID, prod.ProductCode) {
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
