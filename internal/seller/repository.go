package seller

import (
	"errors"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
)

type Repository interface {
	GetOne(id int) (Seller, error)
	GetAll() ([]Seller, error)
	Create(cid int, companyName, address, telephone string) (Seller, error)
	Update(id, cid int, companyName, address, telephone string) (Seller, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r repository) GetOne(id int) (Seller, error) {
	var sellerList []Seller

	err := r.db.Read(&sellerList)

	if err != nil {
		fmt.Println("error reading file", err)
	}

	for _, seller := range sellerList {
		if seller.Id == id {
			return seller, nil
		}
	}
	return Seller{}, fmt.Errorf("the id %d does not exists", id)
}

func (r *repository) GetAll() ([]Seller, error) {
	var sellerList []Seller

	err := r.db.Read(&sellerList)

	if err != nil {
		fmt.Println("error reading file", err)
	}

	if len(sellerList) == 0 {
		sellerList = make([]Seller, 0)
	}
	return sellerList, err
}

func (r *repository) Create(cid int, companyName, address, telephone string) (Seller, error) {
	var sellerList []Seller

	err := r.db.Read(&sellerList)

	if err != nil {
		fmt.Println("error reading file", err)
		return Seller{}, err
	}

	for i := range sellerList {
		if sellerList[i].CompanyId == cid {
			return Seller{}, errors.New("the cid already exists")
		}
	}

	newSeller := Seller{CompanyId: cid, CompanyName: companyName, Address: address, Telephone: telephone}
	newSeller = r.generateId(&newSeller)
	sellerList = append(sellerList, newSeller)

	if err := r.db.Write(sellerList); err != nil {
		return Seller{}, err
	}

	return newSeller, nil
}

func (r *repository) Update(id, cid int, companyName, address, telephone string) (Seller, error) {

	var sellerList []Seller

	err := r.db.Read(&sellerList)
	if err != nil {
		return Seller{}, err
	}

	updateSeller, err := r.GetOne(id)

	if err != nil {
		return Seller{}, err
	}

	err = r.checkIfCidExists(cid, updateSeller)

	if err != nil {
		return Seller{}, err
	}

	updateSeller.CompanyId = cid
	updateSeller.CompanyName = companyName
	updateSeller.Address = address
	updateSeller.Telephone = telephone

	for i := range sellerList {
		if sellerList[i].Id == updateSeller.Id {
			sellerList[i] = updateSeller
		}
	}

	if err := r.db.Write(sellerList); err != nil {
		return Seller{}, err
	}

	return updateSeller, nil
}

func (r *repository) Delete(id int) error {
	var sellerList []Seller

	err := r.db.Read(&sellerList)
	if err != nil {
		return err
	}

	var index int
	seller, err := r.GetOne(id)

	if err != nil {
		return err
	}

	for i := range sellerList {
		if sellerList[i].Id == seller.Id {
			index = i
		}
	}

	sellerList = append(sellerList[:index], sellerList[index+1:]...)
	if err := r.db.Write(sellerList); err != nil {
		return err
	}
	return nil
}

func (r repository) generateId(newSeller *Seller) Seller {
	var sellerList []Seller

	err := r.db.Read(&sellerList)

	if err != nil {
		fmt.Println("error reading file", err)
	}

	if len(sellerList) == 0 {
		newSeller.Id = 1
		return *newSeller
	}

	lastSeller := len(sellerList) - 1
	newSeller.Id = lastSeller + 1
	return *newSeller
}

func (r repository) checkIfCidExists(cid int, seller Seller) error {
	var sellerList []Seller

	err := r.db.Read(&sellerList)

	if err != nil {
		fmt.Println("error reading file", err)
	}

	for i := range sellerList {
		if sellerList[i].CompanyId == cid && sellerList[i].Id != seller.Id {
			return errors.New("the cid already exists")
		}
	}
	return nil
}
