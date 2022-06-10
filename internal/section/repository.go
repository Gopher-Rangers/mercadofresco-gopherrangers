package section

import (
	"fmt"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
)

type Section struct {
	ID             int `json:"id"`
	SectionNumber  int `json:"section_number"`
	CurTemperature int `json:"current_temperature"`
	MinTemperature int `json:"minimum_temperature"`
	CurCapacity    int `json:"current_capacity"`
	MinCapacity    int `json:"minimum_capacity"`
	MaxCapacity    int `json:"maximum_capacity"`
	WareHouseID    int `json:"warehouse_id"`
	ProductTypeID  int `json:"product_type_id"`
}

type Repository interface {
	GetAll() []Section
	GetByID(id int) (Section, error)
	Create(id, secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID int) (Section, error)
	UpdateSecID(id, secNum int) (Section, CodeError)
	DeleteSection(id int) error
}

type CodeError struct {
	Code    int
	Message error
}

func (c CodeError) Error() error {
	return c.Message
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r repository) GetAll() []Section {
	var ListSections []Section
	r.db.Read(&ListSections)
	return ListSections
}

func (r repository) GetByID(id int) (Section, error) {
	var ListSections []Section
	r.db.Read(&ListSections)
	for i := range ListSections {
		if ListSections[i].ID == id {
			return ListSections[i], nil
		}
	}

	return Section{}, fmt.Errorf("seção %d não encontrada", id)
}

func (r repository) Create(id, secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID int) (Section, error) {
	var ListSections []Section
	r.db.Read(&ListSections)

	p := Section{id, secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID}

	for i := range ListSections {
		if ListSections[i].ID+1 == id {
			post := make([]Section, len(ListSections[i+1:]))
			copy(post, ListSections[i+1:])

			ListSections = append(ListSections[:i+1], p)
			ListSections = append(ListSections, post...)
			break
		}
	}

	if id == 1 {
		sec := []Section{p}
		ListSections = append(sec, ListSections...)
	}
	r.db.Write(ListSections)
	return p, nil
}

func (r repository) UpdateSecID(id, secNum int) (Section, CodeError) {
	var ListSections []Section
	r.db.Read(&ListSections)

	for i := range ListSections {
		if ListSections[i].ID == id {
			ListSections[i].SectionNumber = secNum
			r.db.Write(ListSections)
			return ListSections[i], CodeError{0, nil}
		}
	}

	return Section{}, CodeError{404, fmt.Errorf("seção %d não encontrada", id)}
}

func (r repository) DeleteSection(id int) error {
	var ListSections []Section
	r.db.Read(&ListSections)

	for i := range ListSections {
		if ListSections[i].ID == id {
			ListSections = append(ListSections[:i], ListSections[i+1:]...)
			r.db.Write(ListSections)
			return nil
		}
	}
	return fmt.Errorf("seção %d não encontrada", id)
}
