package seller

import (
	"context"
	"errors"
	"fmt"
)

type Service interface {
	GetOne(ctx context.Context, id int) (Seller, error)
	GetAll(ctx context.Context) ([]Seller, error)
	Create(ctx context.Context, cid int, companyName, address, telephone string, localityID int) (Seller, error)
	Update(ctx context.Context, id, cid int, companyName, address, telephone string) (Seller, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
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

	if err := s.existsLocality(ctx, localityID); err != nil {
		return Seller{}, err
	}

	err := s.findByCid(ctx, cid)

	if err != nil {
		return Seller{}, err
	}

	newSeller, err := s.repository.Create(ctx, cid, companyName, address, telephone, localityID)

	if err != nil {
		return Seller{}, err
	}
	return newSeller, nil
}

func (s *service) Update(ctx context.Context, id, cid int, companyName, address, telephone string) (Seller, error) {
	oneSeller, err := s.GetOne(ctx, id)

	if err != nil {
		return Seller{}, err
	}

	err = s.findByCid(ctx, cid)

	if err != nil {
		return Seller{}, err
	}

	updateSeller, err := s.repository.Update(ctx, cid, companyName, address, telephone, oneSeller)

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

func (s service) existsLocality(ctx context.Context, localityID int) error {
	var sellerList []Seller

	sellerList, err := s.GetAll(ctx)

	if err != nil {
		return err
	}

	for i := range sellerList {
		if sellerList[i].LocalityID == localityID {
			return nil
		}
	}
	return fmt.Errorf("locality_id does not exists")
}

func (s service) findByCid(ctx context.Context, cid int) error {
	var sellerList []Seller

	sellerList, err := s.GetAll(ctx)

	if err != nil {
		return err
	}

	for i := range sellerList {
		if sellerList[i].CompanyId == cid {
			return errors.New("the cid already exists")
		}
	}
	return nil
}
