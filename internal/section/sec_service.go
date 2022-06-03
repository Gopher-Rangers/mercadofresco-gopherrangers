package section

type Services interface {
	GetAll() []Section
	GetByID(id int) (Section, error)
	Create(secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID int) (Section, error)
	UpdateSecID(id, secNum int) (Section, error)
	DeleteSection(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Services {
	s := service{r}
	return &s
}

func (s service) LastID() int {
	return s.repository.LastID()
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
	id := s.repository.AvailableID()
	ps, err := s.repository.Create(id, secNum, curTemp, minTemp, curCap, minCap, maxCap, wareID, typeID)
	if err != nil {
		return Section{}, err
	}
	return ps, nil
}

func (s *service) UpdateSecID(id, secNum int) (Section, error) {
	ps, err := s.repository.UpdateSecID(id, secNum)
	if err != nil {
		return Section{}, err
	}

	return ps, nil
}

func (s *service) DeleteProduct(id int) error {
	err := s.repository.DeleteSection(id)
	if err != nil {
		return err
	}
	return nil
}
