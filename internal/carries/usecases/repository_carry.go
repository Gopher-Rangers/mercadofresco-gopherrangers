package usecases

import "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carries/domain"

type RepositoryCarry interface {
	CreateCarry(carry domain.Carry) (domain.Carry, error)
	GetCaryPerLocality(id int) (domain.Carry, error)
	GetCarryByCid(cid string) (domain.Carry, error)
}
