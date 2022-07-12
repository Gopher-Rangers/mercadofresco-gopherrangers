package usecases

import "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"

type RepositoryCarry interface {
	CreateCarry(carry domain.Carry) (domain.Carry, error)
	GetCarryByCid(cid string) (domain.Carry, error)
}
