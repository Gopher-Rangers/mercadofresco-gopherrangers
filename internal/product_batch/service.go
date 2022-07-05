package productbatch

type Services interface {
	Report() ([]Report, error)
	ReportByID(id int) (Report, error)
	Create(pb ProductBatch) (ProductBatch, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Services {
	s := service{r}
	return &s
}

func (s service) Create(pb ProductBatch) (ProductBatch, error) {
	ps, err := s.repository.Create(pb)
	if err != nil {
		return ProductBatch{}, err
	}
	return ps, nil
}

func (s service) Report() ([]Report, error) {
	pb, err := s.repository.Report()
	if err != nil {
		return []Report{}, err
	}
	return pb, nil
}

func (s service) ReportByID(id int) (Report, error) {
	pb, err := s.repository.ReportByID(id)
	if err != nil {
		return Report{}, err
	}
	return pb, nil
}
