package usecases

import "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"

type RepositoryLocality interface {
	GetCarryLocalityByID(id int) (domain.Locality, error)
	GetAllCarriesLocality() ([]domain.Locality, error)
}
