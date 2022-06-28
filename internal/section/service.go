package section

import (
	"errors"
	"fmt"
)

type Services interface {
	GetAll() ([]Section, error)
	GetByID(id int) (Section, error)
	Create(secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID int) (Section, error)
	UpdateSecID(id, secNum int) (Section, CodeError)
	DeleteSection(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Services {
	s := service{r}
	return &s
}

func (s *service) GetAll() ([]Section, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return ps, err
	}
	return ps, nil
}

func (s *service) GetByID(id int) (Section, error) {
	ps, err := s.repository.GetByID(id)
	if err != nil {
		return Section{}, err
	}
	return ps, nil
}

func (s *service) Create(secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID int) (Section, error) {
	ListSections, err := s.repository.GetAll()
	if err != nil {
		return Section{}, err
	}

	for i := range ListSections {
		if ListSections[i].SectionNumber == secNum {
			return Section{}, fmt.Errorf("seção com sectionNumber: %d já existe no banco de dados", secNum)
		}
	}

	ps, err := s.repository.Create(secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID)
	if err != nil {
		return Section{}, err
	}
	return ps, nil
}

func (s *service) UpdateSecID(id, secNum int) (Section, CodeError) {
	ListSections, err := s.repository.GetAll()
	if err != nil {
		return Section{}, CodeError{500, errors.New("internal server error")}
	}

	for i := range ListSections {
		if ListSections[i].SectionNumber == secNum {
			return Section{}, CodeError{409,
				fmt.Errorf("seção com section_number: %d já existe no banco de dados", secNum)}
		}
	}

	ps, erro := s.repository.UpdateSecID(id, secNum)
	if erro.Code != 0 {
		return Section{}, erro
	}

	return ps, CodeError{0, nil}
}

func (s *service) DeleteSection(id int) error {
	err := s.repository.DeleteSection(id)
	if err != nil {
		return err
	}
	return nil
}
