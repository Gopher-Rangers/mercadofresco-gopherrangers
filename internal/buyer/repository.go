package buyer

import (
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
)

type repository struct {
	db store.Store
}

type Repository interface {
	GetAll() ([]Buyer, error)
	Create(id int, cardNumberId string, firstName string, lastName string) (Buyer, error)
	Update(id int, cardNumberId string, firstName string, lastName string) (Buyer, error)
	Delete(id int) error
	GetById(id int) (Buyer, error)
}

func NewRepository() Repository {
	db := store.New(store.FileType, "../../internal/buyer/buyers.json")
	return &repository{db: db}
}

func (r *repository) GetAll() ([]Buyer, error) {
	var buyers []Buyer
	r.db.Read(&buyers)

	if len(buyers) == 0 {
		return make([]Buyer, 0), nil
	}

	return buyers, nil
}

func (r *repository) GetById(id int) (Buyer, error) {
	var buyerList []Buyer
	r.db.Read(&buyerList)

	buyerIndex := -1

	for i := range buyerList {
		if buyerList[i].Id == id {
			buyerIndex = i
		}
	}

	if buyerIndex == -1 {
		return Buyer{}, fmt.Errorf("buyer with id %d not founded", id)
	}
	return buyerList[buyerIndex], nil
}

func (r *repository) Create(id int, cardNumberId string, firstName string, lastName string) (Buyer, error) {
	var buyers []Buyer
	r.db.Read(&buyers)
	newBuyer := Buyer{id, cardNumberId, firstName, lastName}

	buyers = append(buyers, newBuyer)
	if err := r.db.Write(buyers); err != nil {
		return Buyer{}, err
	}
	return newBuyer, nil
}

func (r repository) Update(id int, cardNumberId string, firstName string, lastName string) (Buyer, error) {
	var buyers []Buyer
	r.db.Read(&buyers)

	index := -1
	for i := range buyers {
		if buyers[i].Id == id {
			index = i
		}
	}

	if index != -1 {
		buyers[index] = Buyer{id, cardNumberId, firstName, lastName}
		r.db.Write(buyers)
		return buyers[index], nil
	}
	return Buyer{}, fmt.Errorf("buyer with id: %d not found", id)
}

func (r repository) Delete(id int) error {
	var buyersList []Buyer
	r.db.Read(&buyersList)

	for i := range buyersList {
		if buyersList[i].Id == id {
			buyersList = append(buyersList[:i], buyersList[i+1:]...)
			r.db.Write(buyersList)
			return nil
		}
	}
	return fmt.Errorf("buyer with id : %d not founded", id)
}
