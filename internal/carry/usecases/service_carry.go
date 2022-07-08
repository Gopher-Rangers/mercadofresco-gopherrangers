package usecases

import "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carries/domain"

type ServiceCarry interface {
	CreateCarry(carry domain.Carry) (domain.Carry, error)
	GetCaryPerLocality(id int) (domain.Carry, error)
	GetCarryByCid(cid string) (domain.Carry, error)
}

type serviceCarry struct {
	repository RepositoryCarry
}

func NewServiceCarry(r ServiceCarry) ServiceCarry {
	return &service{r}
}
