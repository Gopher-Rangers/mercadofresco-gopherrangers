package seller

import "fmt"

type Service interface {
	GetOne(id int) (Seller, error)
	GetAll() ([]Seller, error)
	Create(cid int, companyName, address, telephone string) (Seller, error)
	Update(id int, companyName, address, telephone string) (Seller, error)
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

	fmt.Println("seller", sellerList)
	if err != nil {
		return sellerList, err
	}
	fmt.Println("seller", err)
	return sellerList, nil
}

func (s *service) Create(cid int, companyName, address, telephone string) (Seller, error) {
	newSeller, err := s.repository.Create(cid, companyName, address, telephone)

	if err != nil {
		return Seller{}, err
	}

	return newSeller, nil
}

func (s *service) Update(id int, companyName, address, telephone string) (Seller, error) {
	updateSeller, err := s.repository.Update(id, companyName, address, telephone)

	if err != nil {
		return Seller{}, err
	}
	return updateSeller, nil
}

func (s *service) GetOne(id int) (Seller, error) {
	oneSeller, err := s.repository.GetOne(id)

	if err != nil {
		fmt.Println(err.Error())
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
