package section

import (
	"fmt"

	store "github.com/mercadofresco-gopherrangers.git/pkg/store"
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
	UpdateSecID(id, secNum int) (Section, error)
	DeleteSection(id int) error

	LastID() int
	AvailableID() int
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r repository) GetAll() ([]Section, error) {
	var ListSections []Section
	r.db.Read(&ListSections)
	return ListSections, nil
}

func (r repository) GetByID(id int) ([]Section, error) {
	var ListSections []Section
	r.db.Read(&ListSections)
	return ListSections, nil
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

func (r repository) UpdateSecID(id, secNum int) (Section, error) {
	var ListSections []Section
	r.db.Read(&ListSections)

	var flagExist bool
	var index int
	for i := range ListSections {
		if ListSections[i].ID == id {
			index = i
			flagExist = true
		}
		if ListSections[i].SectionNumber == secNum {
			return Section{}, fmt.Errorf("seção com sectionNumber: %d já existe no banco de dados", id)
		}
	}

	if flagExist {
		ListSections[index].SectionNumber = secNum
		r.db.Write(ListSections)
		return ListSections[index], nil
	}
	return Section{}, fmt.Errorf("seção %d não encontrada", id)
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

func (r repository) AvailableID() int {
	var ListSections []Section
	r.db.Read(&ListSections)

	for prevI := range ListSections[:len(ListSections)-1] {
		i := prevI + 1
		if ListSections[i].ID != (ListSections[prevI].ID + 1) {
			id := ListSections[prevI].ID + 1
			return id
		}
	}
	return r.LastID()
}

func (r repository) LastID() int {
	var ListSections []Section
	r.db.Read(&ListSections)

	if len(ListSections) == 0 || ListSections[0].ID != 1 {
		return 1
	}
	return ListSections[len(ListSections)-1].ID + 1
}
