package seller

import (
	"context"
	"errors"
	"fmt"
	l "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality"
)

type Service interface {
	GetOne(ctx context.Context, id int) (Seller, error)
	GetAll(ctx context.Context) ([]Seller, error)
	Create(ctx context.Context, cid int, companyName, address, telephone string, localityID int) (Seller, error)
	Update(ctx context.Context, id, cid int, companyName, address, telephone string, localityID int) (Seller, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository   Repository
	localityRepo l.Repository
}

func NewService(r Repository, lr l.Repository) Service {
	return &service{
		repository:   r,
		localityRepo: lr,
	}
}

func (s *service) GetAll(ctx context.Context) ([]Seller, error) {
	sellerList, err := s.repository.GetAll(ctx)

	if err != nil {
		return sellerList, err
	}
	return sellerList, nil
}

func (s *service) Create(ctx context.Context, cid int, companyName, address, telephone string, localityID int) (Seller, error) {
	locality, err := s.localityRepo.GetById(ctx, localityID)

	if err != nil {
		return Seller{}, fmt.Errorf("locality_id does not exists")
	}

	err = s.findByCid(ctx, cid, Seller{})

	if err != nil {
		return Seller{}, err
	}

	newSeller, err := s.repository.Create(ctx, cid, companyName, address, telephone, locality.Id)

	if err != nil {
		return Seller{}, err
	}
	return newSeller, nil
}

func (s *service) Update(ctx context.Context, id, cid int, companyName, address, telephone string, localityID int) (Seller, error) {
	oneSeller, err := s.GetOne(ctx, id)

	if err != nil {
		return Seller{}, err
	}

	locality, err := s.localityRepo.GetById(ctx, localityID)

	if err != nil {
		return Seller{}, fmt.Errorf("locality_id does not exists")
	}

	err = s.findByCid(ctx, cid, oneSeller)

	if err != nil {
		return Seller{}, err
	}

	updateSeller, err := s.repository.Update(ctx, cid, companyName, address, telephone, locality.Id, oneSeller)

	if err != nil {
		return Seller{}, err
	}
	return updateSeller, nil
}

func (s *service) GetOne(ctx context.Context, id int) (Seller, error) {
	oneSeller, err := s.repository.GetOne(ctx, id)

	if err != nil {
		return Seller{}, err
	}
	return oneSeller, nil
}

func (s *service) Delete(ctx context.Context, id int) error {

	seller, err := s.GetOne(ctx, id)

	if err != nil {
		return err
	}

	if err := s.repository.Delete(ctx, seller.Id); err != nil {
		return err
	}

	return nil
}

func (s service) findByCid(ctx context.Context, cid int, seller Seller) error {
	var sellerList []Seller

	sellerList, err := s.GetAll(ctx)

	if err != nil {
		return err
	}

	for i := range sellerList {
		if sellerList[i].CompanyId == cid && sellerList[i].Id != seller.Id {
			return errors.New("the cid already exists")
		}
	}

	return nil
}
