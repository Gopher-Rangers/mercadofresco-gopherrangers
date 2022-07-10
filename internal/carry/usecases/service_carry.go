package usecases

import (
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
)

type ServiceCarry interface {
	CreateCarry(carry domain.Carry) (domain.Carry, error)
	GetCaryPerLocality(id int) (domain.Carry, error)
}

type serviceCarry struct {
	repository RepositoryCarry
}

func NewServiceCarry(r RepositoryCarry) ServiceCarry {
	return &serviceCarry{r}
}

func (s *serviceCarry) CreateCarry(carry domain.Carry) (domain.Carry, error) {

	_, err := s.repository.GetCarryByCid(carry.Cid)

	if err == nil {
		return domain.Carry{}, fmt.Errorf("o `cid` já está em uso")
	}

	carry, err = s.repository.CreateCarry(carry)

	if err != nil {
		return domain.Carry{}, err
	}

	return carry, nil

}
func (s serviceCarry) GetCaryPerLocality(id int) (domain.Carry, error) {
	return domain.Carry{}, nil
}
