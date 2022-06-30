package seller

import (
	"errors"
)

type Service interface {
	GetOne(id int) (Seller, error)
	GetAll() ([]Seller, error)
	Create(cid int, companyName, address, telephone string) (Seller, error)
	Update(id, cid int, companyName, address, telephone string) (Seller, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]Seller, error) {
	sellerList, err := s.repository.GetAll()

	if err != nil {
		return sellerList, err
	}
	return sellerList, nil
}

func (s *service) Create(cid int, companyName, address, telephone string) (Seller, error) {

	err := s.findByCid(cid)

	if err != nil {
		return Seller{}, err
	}

	newSeller, err := s.repository.Create(cid, companyName, address, telephone)

	if err != nil {
		return Seller{}, err
	}
	return newSeller, nil
}

func (s *service) Update(id, cid int, companyName, address, telephone string) (Seller, error) {
	oneSeller, err := s.GetOne(id)

	if err != nil {
		return Seller{}, err
	}

	sellerList, err := s.GetAll()

	for i := range sellerList {
		if sellerList[i].CompanyId == cid && sellerList[i].Id != oneSeller.Id {
			return Seller{}, errors.New("this cid already exists and and belongs to another company")
		}
	}

	updateSeller, err := s.repository.Update(cid, companyName, address, telephone, oneSeller)

	if err != nil {
		return Seller{}, err
	}
	return updateSeller, nil
}

func (s *service) GetOne(id int) (Seller, error) {
	oneSeller, err := s.repository.GetOne(id)

	if err != nil {
		return Seller{}, err
	}
	return oneSeller, nil
}

func (s *service) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (s service) findByCid(cid int) error {
	var sellerList []Seller

	sellerList, err := s.GetAll()

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
