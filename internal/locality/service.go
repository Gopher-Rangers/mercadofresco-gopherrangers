package locality

import (
	"context"
	"fmt"
)

type Service interface {
	GetAll(ctx context.Context) ([]Locality, error)
	GetById(ctx context.Context, id int) (Locality, error)
	ReportSellers(ctx context.Context, id int) (ReportSeller, error)
	Create(ctx context.Context, id int, localityName, provinceName, countryName string) (Locality, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s service) ReportSellers(ctx context.Context, localityId int) (ReportSeller, error) {
	var reportSeller ReportSeller

	locality, err := s.repository.GetById(ctx, localityId)

	if err != nil {
		return ReportSeller{}, err
	}

	reportSeller, err = s.repository.ReportSellers(ctx, locality.Id)
	if err != nil {
		return ReportSeller{}, err
	}

	return reportSeller, nil
}

func (s service) Create(ctx context.Context, id int, localityName, provinceName, countryName string) (Locality, error) {

	exists, err := s.cepIdExists(ctx, id)

	if err != nil {
		return Locality{}, err
	}

	if exists {
		return Locality{}, err
	}

	newLocality, err := s.repository.Create(ctx, id, localityName, provinceName, countryName)

	if err != nil {
		return Locality{}, err
	}
	return newLocality, nil
}

func (s service) GetAll(ctx context.Context) ([]Locality, error) {

	localityList, err := s.repository.GetAll(ctx)

	if err != nil {
		return localityList, err
	}

	return localityList, err
}

func (s service) GetById(ctx context.Context, id int) (Locality, error) {

	locality, err := s.repository.GetById(ctx, id)

	if err != nil {
		return locality, err
	}

	return locality, nil
}

func (s service) cepIdExists(ctx context.Context, id int) (bool, error) {

	localities, err := s.GetAll(ctx)

	if err != nil {
		return false, err
	}

	for i := range localities {
		if localities[i].Id == id {
			return true, fmt.Errorf("id already exists")
		}
	}
	return false, nil
}
