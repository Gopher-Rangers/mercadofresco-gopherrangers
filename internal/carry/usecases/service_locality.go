package usecases

import "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"

type ServiceLocality interface {
	GetCarryLocalityByID(id int) (domain.Locality, error)
	GetAllCarriesLocalityByID() ([]domain.Locality, error)
}

type serviceLocality struct {
	repository RepositoryLocality
}

func NewServiceLocality(r RepositoryLocality) ServiceLocality {
	return &serviceLocality{r}
}

func (s serviceLocality) GetCarryLocalityByID(id int) (domain.Locality, error) {
	locality, _ := s.repository.GetCarryLocalityByID(id)

	return locality, nil
}

func (s serviceLocality) GetAllCarriesLocalityByID() ([]domain.Locality, error) {
	localities, _ := s.repository.GetAllCarriesLocalityByID()

	return localities, nil
}
