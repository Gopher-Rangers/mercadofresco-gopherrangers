package section

import (
	"database/sql"
	"fmt"
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
	GetAll() ([]Section, error)
	GetByID(id int) (Section, error)
	Create(secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID int) (Section, error)
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
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) GetAll() ([]Section, error) {
	var sections []Section

	rows, err := r.db.Query(sqlGetAll)
	if err != nil {
		return sections, err
	}

	defer rows.Close()

	for rows.Next() {
		var sec Section

		err := rows.Scan(&sec.ID, &sec.SectionNumber, &sec.CurTemperature, &sec.MinTemperature,
			&sec.CurCapacity, &sec.MinCapacity, &sec.MaxCapacity, &sec.WareHouseID, &sec.ProductTypeID)
		if err != nil {
			return sections, err
		}

		sections = append(sections, sec)
	}

	return sections, nil
}

func (r repository) GetByID(id int) (Section, error) {
	var sec Section

	rows, err := r.db.Query(sqlGetById, id)
	if err != nil {
		return Section{}, fmt.Errorf("seção %d não encontrada", id)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&sec.ID, &sec.SectionNumber, &sec.CurTemperature, &sec.MinTemperature,
			&sec.CurCapacity, &sec.MinCapacity, &sec.MaxCapacity, &sec.WareHouseID, &sec.ProductTypeID)
		if err != nil {
			return Section{}, err
		}
	}

	return sec, nil
}

func (r repository) Create(secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID int) (Section, error) {
	res, err := r.db.Exec(sqlStore, secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID)
	if err != nil {
		return Section{}, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return Section{}, err
	}

	sec := Section{int(lastID), secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID}
	return sec, nil
}

func (r repository) UpdateSecID(id, secNum int) (Section, CodeError) {
	_, err := r.db.Exec(sqlUpdateSecID, secNum, id)
	if err != nil {
		return Section{}, CodeError{500, err}
	}

	sec, _ := r.GetByID(id)
	return sec, CodeError{200, nil}
}

func (r repository) DeleteSection(id int) error {
	_, err := r.db.Exec(sqlDelete, id)
	if err != nil {
		return err
	}

	return nil
}
