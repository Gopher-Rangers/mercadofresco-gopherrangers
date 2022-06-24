package domain

type Buyer struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type Repository interface {
	GetAll() ([]Buyer, error)
	Create(buyer Buyer) (Buyer, error)
	Update(buyer Buyer) (Buyer, error)
	Delete(id int) error
	GetValidId() int
	GetById(id int) (Buyer, error)
}

type Service interface {
	GetAll() ([]Buyer, error)
	Create(buyer Buyer) (Buyer, error)
	Update(buyer Buyer) (Buyer, error)
	Delete(id int) error
	GetById(id int) (Buyer, error)
}
