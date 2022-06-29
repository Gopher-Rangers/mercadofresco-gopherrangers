package productbatch

import (
	"database/sql"
	"time"
)

type ProductBatch struct {
	ID              int       `json:"id"`
	BatchNumber     string    `json:"batch_number"`
	CurQuantity     int       `json:"current_quantity"`
	CurTemperature  float64   `json:"current_temperature"`
	MinTemperature  float64   `json:"minimum_temperature"`
	DueDate         time.Time `json:"due_date"`
	InitialQuantity int       `json:"initial_quantity"`
	ManufactDate    time.Time `json:"manufacturing_date"`
	ManufactHour    time.Time `json:"manufacturing_hour"`
	ProductTypeID   int       `json:"product_id"`
	SectionID       int       `json:"section_id"`
}

type Repository interface {
	GetByID(id int) (ProductBatch, error)
	Create(pb ProductBatch) (ProductBatch, error)
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

func (r repository) GetByID(id int) (ProductBatch, error) {
	return ProductBatch{}, nil
}

func (r repository) Create(pb ProductBatch) (ProductBatch, error) {
	return ProductBatch{}, nil
}
