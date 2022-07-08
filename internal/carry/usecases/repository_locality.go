package usecases

import "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carries/domain"

type RepositoryLocality interface {
	GetLocalityByID(id int) (domain.Locality, error)
}
