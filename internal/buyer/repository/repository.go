package buyer

import (
	"fmt"
	"github.com/google/uuid"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
)

type repository struct {
	db store.Store
}

type Repository interface {
	GetAll() ([]Buyer, error)
	Create(buyer Buyer) (Buyer, error)
	Update(buyer Buyer) (Buyer, error)
	Delete(id int) error
	GetValidId() int
	GetById(id int) (Buyer, error)
}

func NewRepository(db store.Store) Repository {
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

func (r *repository) Create(buyer Buyer) (Buyer, error) {
	var buyers []Buyer
	r.db.Read(&buyers)

	buyers = append(buyers, buyer)
	if err := r.db.Write(buyers); err != nil {
		return Buyer{}, err
	}
	return buyer, nil
}

func (r repository) Update(buyer Buyer) (Buyer, error) {
	var buyers []Buyer
	r.db.Read(&buyers)

	index := -1
	for i := range buyers {
		if buyers[i].Id == buyer.Id {
			index = i
		}
	}

	if index != -1 {
		buyers[index] = buyer
		r.db.Write(buyers)
		return buyers[index], nil
	}
	return Buyer{}, fmt.Errorf("buyer with id: %d not found", buyer.Id)
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

func (r repository) GetValidId() int {
	generatedId := int(uuid.New().ID())

	entities, _ := r.GetAll()

	for i := 0; i < len(entities); i++ {
		if entities[i].Id == generatedId {
			generatedId = int(uuid.New().ID())
			i = 0
		}
	}

	return generatedId
}
