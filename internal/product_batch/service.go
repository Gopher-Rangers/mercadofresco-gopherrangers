package productbatch

type Services interface {
	GetByID(id int) (ProductBatch, error)
	Create(pb ProductBatch) (ProductBatch, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Services {
	s := service{r}
	return &s
}

func (r service) GetByID(id int) (ProductBatch, error) {
	return ProductBatch{}, nil
}

func (r service) Create(pb ProductBatch) (ProductBatch, error) {
	return ProductBatch{}, nil
}
