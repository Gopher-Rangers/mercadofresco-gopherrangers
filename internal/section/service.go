package section

import "fmt"

type Services interface {
	GetAll() []Section
	GetByID(id int) (Section, error)
	Create(secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID int) (Section, error)
	UpdateSecID(id, secNum int) (Section, CodeError)
	DeleteSection(id int) error
	LastID() int
}

type service struct {
	repository Repository
}

func NewService(r Repository) Services {
	s := service{r}
	return &s
}

func (s *service) GetAll() []Section {
	ps := s.repository.GetAll()
	return ps
}

func (s *service) GetByID(id int) (Section, error) {
	ps, err := s.repository.GetByID(id)
	if err != nil {
		return Section{}, err
	}
	return ps, nil
}

func (s *service) Create(secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID int) (Section, error) {
	ListSections := s.repository.GetAll()
	for i := range ListSections {
		if ListSections[i].SectionNumber == secNum {
			return Section{}, fmt.Errorf("seção com sectionNumber: %d já existe no banco de dados", secNum)
		}
	}

	id := s.AvailableID()
	ps, err := s.repository.Create(id, secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID)
	if err != nil {
		return Section{}, err
	}
	return ps, nil
}

func (s *service) UpdateSecID(id, secNum int) (Section, CodeError) {
	ListSections := s.repository.GetAll()
	for i := range ListSections {
		if ListSections[i].SectionNumber == secNum {
			return Section{}, CodeError{409,
				fmt.Errorf("seção com section_number: %d já existe no banco de dados", secNum)}
		}
	}

	ps, err := s.repository.UpdateSecID(id, secNum)
	if err.Code != 0 {
		return Section{}, err
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

func (s *service) AvailableID() int {
	ListSections := s.repository.GetAll()

	if len(ListSections) == 0 || ListSections[0].ID != 1 {
		return 1
	}

	for prevI := range ListSections[:len(ListSections)-1] {
		i := prevI + 1
		if ListSections[i].ID != (ListSections[prevI].ID + 1) {
			id := ListSections[prevI].ID + 1
			return id
		}
	}
	return s.LastID()
}

func (s *service) LastID() int {
	ListSections := s.repository.GetAll()

	if len(ListSections) == 0 {
		return 1
	}

	return ListSections[len(ListSections)-1].ID + 1
}
