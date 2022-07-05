package productbatch

import (
	"database/sql"
	"fmt"
)

type ProductBatch struct {
	ID              int    `json:"id"`
	BatchNumber     int    `json:"batch_number"`
	CurQuantity     int    `json:"current_quantity"`
	CurTemperature  int    `json:"current_temperature"`
	DueDate         string `json:"due_date"`
	InitialQuantity int    `json:"initial_quantity"`
	ManufactDate    string `json:"manufacturing_date"`
	ManufactHour    int    `json:"manufacturing_hour"`
	MinTemperature  int    `json:"minimum_temperature"`
	ProductTypeID   int    `json:"product_id"`
	SectionID       int    `json:"section_id"`
}

type Report struct {
	SecID     int `json:"section_id"`
	SecNum    int `json:"section_number"`
	ProdCount int `json:"products_count"`
}

type Repository interface {
	Create(pb ProductBatch) (ProductBatch, error)
	Report() ([]Report, error)
	ReportByID(id int) (Report, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) Create(pb ProductBatch) (ProductBatch, error) {
	res, err := r.db.Exec(sqlCreateBatch, pb.BatchNumber, pb.CurQuantity, pb.CurTemperature, pb.DueDate,
		pb.InitialQuantity, pb.ManufactDate, pb.ManufactHour, pb.MinTemperature, pb.ProductTypeID, pb.SectionID)
	if err != nil {
		return ProductBatch{}, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected <= 0 {
		return ProductBatch{}, fmt.Errorf("rows not affected")
	}

	lastID, _ := res.LastInsertId()

	pb.ID = int(lastID)
	return pb, nil
}

func (r repository) Report() ([]Report, error) {
	var rep []Report

	rows, err := r.db.Query(sqlReportBatchAll)
	if err != nil {
		return rep, err
	}

	defer rows.Close()

	for rows.Next() {
		var row Report

		err = rows.Scan(&row.SecID, &row.SecNum, &row.ProdCount)
		if err != nil {
			return []Report{}, err
		}

		rep = append(rep, row)
	}

	return rep, nil
}

func (r repository) ReportByID(id int) (Report, error) {
	var rep Report

	rows := r.db.QueryRow(sqlReportBatchByID, id)

	err := rows.Scan(&rep.SecID, &rep.SecNum, &rep.ProdCount)
	if err != nil {
		return Report{}, err
	}

	return rep, nil
}
