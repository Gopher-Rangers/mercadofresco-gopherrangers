package products

import (
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
)

type Product struct {
	ID int `json:"id"`
	ProductCode string `json:"product_code"`
	Description string `json:"description"`
	Width float64 `json:"width"`
	Height float64 `json:"height"`
	Length float64 `json:"length"`
	NetWeight float64 `json:"net_weight"`
	ExpirationRate string `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate float64 `json:"freezing_rate"`
	ProductTypeTd int `json:"product_type_id"`
	SellerId int `json:"seller_id"`
}

type Repository interface {
	LastID() (int, error)
	Store(prod Product, id int) (Product, error)
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	//UpdatePut(prod Product) (Product, error)
	//UpdatePatch(prod Product) (Product, error)
	//Delete(id int) (error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps) - 1].ID, nil
}

func (r *repository) Store(prod Product, id int) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	prod.ID = id
	ps = append(ps, prod)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return prod, nil
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *repository) GetById(id int) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	for i:= range ps {
		if ps[i].ID == id {
			return ps[i], nil
		}
	}
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}

/*
func (r *repository) UpdatePut(prod Product) (Product, error) {

}

func (r *repository) UpdatePatch(prod Product) (Product, error) {

}

func (r *repository) Delete(id int) (error) {

}
*/
