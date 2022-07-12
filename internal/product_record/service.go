package productrecord

import (
	"fmt"
	"time"
	"context"

	products "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
)

const (
	ERROR_INEXISTENT_PRODUCT = "the product id doesn`t exist"
	ERROR_WRONG_LAST_UPDATE_DATE = "the last update date must be greater than the system time"
)

type Service interface {
	Store(ctx context.Context, prod ProductRecord) (ProductRecord, error)
	GetAll(ctx context.Context) ([]ProductRecordGet, error)
	GetById(ctx context.Context, id int) (ProductRecordGet, error)
}

type service struct {
	repository Repository
	productsService products.Service
}

func NewService(r Repository, productsService products.Service) Service {
	return &service{
		repository: r,
		productsService: productsService}
}

func (s *service) checkIfProductExists(ctx context.Context, prod ProductRecord) bool {
	_, err := s.productsService.GetById(ctx, prod.ProductId)
	return err == nil
}

func (s *service) checkDatetime(last_update_time string) (bool, error) {
	currentTime := time.Now()
	loc := currentTime.Location()
	layout := "2006-01-02 15:04:00"
	lastTime, err := time.ParseInLocation(layout, last_update_time, loc)
	if err != nil {
		return false, err
	}
	diff := lastTime.Sub(currentTime)
	return diff > 0, nil
}

func (s *service) Store(ctx context.Context, prod ProductRecord) (ProductRecord, error) {
	if !s.checkIfProductExists(ctx, prod) {
		return ProductRecord{}, fmt.Errorf(ERROR_INEXISTENT_PRODUCT)
	}
	dateTimeOk, err := s.checkDatetime(prod.LastUpdateDate)
	if err != nil {
		return ProductRecord{}, err
	}
	if !dateTimeOk {
		return ProductRecord{}, fmt.Errorf(ERROR_WRONG_LAST_UPDATE_DATE)
	}
	product, err := s.repository.Store(ctx, prod)
	if err != nil {
		return ProductRecord{}, err
	}
	return product, nil
}

func (s *service) GetById(ctx context.Context, id int) (ProductRecordGet, error) {
	ps, err := s.repository.GetById(ctx, id)
	if err != nil {
		return ProductRecordGet{}, err
	}
	return ps, nil
}

func (s *service) GetAll(ctx context.Context) ([]ProductRecordGet, error) {
	ps, _ := s.repository.GetAll(ctx)
	return ps, nil
}
