package usecases

import "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carries/domain"

type ServiceLocality interface {
	GetLocalityByID(id int) (domain.Locality, error)
}

type serviceLocality struct {
	repository RepositoryLocality
}

func NewServiceLocality(r RepositoryLocality) ServiceLocality {
	return &serviceLocality{r}
}

func (s serviceLocality) GetLocalityByID(id int) (domain.Locality, error) {
	locality, _ := s.repository.GetLocalityByID(id)

	return locality, nil
}
