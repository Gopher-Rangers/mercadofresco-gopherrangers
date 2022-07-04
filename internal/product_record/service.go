package productrecord

import (
	"fmt"
	"time"

	products "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
)

const (
	ERROR_INEXISTENT_PRODUCT = "the product id doesn`t exist"
	ERROR_WRONG_LAST_UPDATE_DATE = "the last update date must be greater than the system time"
)

type Service interface {
	Store(prod ProductRecord) (ProductRecord, error)
	GetAll() ([]ProductRecordGet, error)
	GetById(id int) (ProductRecordGet, error)
	GetAllProductRecords() ([]ProductRecord, error)
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

func (s *service) checkIfProductExists(prod ProductRecord) bool {
	_, err := s.productsService.GetById(prod.ProductId)
	return err == nil
}

func (s *service) checkDatetime(last_update_time string) bool {
	currentTime := time.Now()
	loc := currentTime.Location()
	layout := "2006-01-02 15:04:00"
	lastTime, err := time.ParseInLocation(layout, last_update_time, loc)
	if err != nil {
		fmt.Println(err)
	}
	diff := lastTime.Sub(currentTime)
	return diff > 0
}

func (s *service) Store(prod ProductRecord) (ProductRecord, error) {
	if !s.checkIfProductExists(prod) {
		return ProductRecord{}, fmt.Errorf(ERROR_INEXISTENT_PRODUCT)
	}
	if !s.checkDatetime(prod.LastUpdateDate) {
		return ProductRecord{}, fmt.Errorf(ERROR_WRONG_LAST_UPDATE_DATE)
	}
	product, err := s.repository.Store(prod)
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

func (s *service) GetAllProductRecords() ([]ProductRecord, error) {
	ps, _ := s.repository.GetAllProductRecords()
	return ps, nil
}
