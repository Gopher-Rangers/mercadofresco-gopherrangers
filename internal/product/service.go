package products

type Service interface {
	GetAll() ([]Product, error)
	Store(prod Product) (Product, error)
	Update(prod Product) (Product, error)
	UpdateDescription(id int, description string) (Product, error)
	Delete(id int) (error)
}
