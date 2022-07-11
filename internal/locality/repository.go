package locality

import (
	"context"
	"database/sql"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Locality, error)
	GetById(ctx context.Context, id int) (Locality, error)
	Create(ctx context.Context, zipCode, localityName, provinceName, countryName string) (Locality, error)
	ReportSellers(ctx context.Context, id int) (ReportSeller, error)
}

const (
	GET_REPORT_SELLER = "SELECT l.id, l.locality_name, COUNT(sellers.id) FROM localities l LEFT JOIN sellers ON l.id=sellers.locality_id WHERE l.id = ?"
	INSERT            = "INSERT INTO localities (zip_code, locality_name, province_name, country_name) VALUES (?,?,?,?)"
	GETALL            = "SELECT * FROM localities"
	GETBYID           = "SELECT * FROM localities WHERE id = ?"
)

type mariaDBRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) Repository {
	return &mariaDBRepository{db: db}
}

func (m mariaDBRepository) ReportSellers(ctx context.Context, id int) (ReportSeller, error) {
	var reportSeller ReportSeller

	rows, err := m.db.QueryContext(ctx, GET_REPORT_SELLER, id)

	if err != nil {
		return ReportSeller{}, err
	}

	defer rows.Close()

	for rows.Next() {

		rows.Scan(&reportSeller.LocalityID, &reportSeller.LocalityName, &reportSeller.SellersCount)

		if err != nil {
			return ReportSeller{}, err
		}
	}
	return reportSeller, nil
}

func (m mariaDBRepository) Create(ctx context.Context, zipCode, localityName, provinceName, countryName string) (Locality, error) {

	locality := Locality{ZipCode: zipCode, LocalityName: localityName, ProvinceName: provinceName, CountryName: countryName}

	res, err := m.db.ExecContext(ctx, INSERT, &locality.ZipCode, &locality.LocalityName, &locality.ProvinceName, &locality.CountryName)

	if err != nil {
		return Locality{}, err
	}

	lastID, err := res.LastInsertId()

	if err != nil {
		return Locality{}, err
	}

	locality.Id = int(lastID)

	return locality, nil
}

func (m mariaDBRepository) GetAll(ctx context.Context) ([]Locality, error) {
	var localityList []Locality

	rows, err := m.db.QueryContext(ctx, GETALL)

	if err != nil {
		return localityList, err
	}

	defer rows.Close()

	for rows.Next() {
		var locality Locality

		err = rows.Scan(&locality.Id, &locality.ZipCode, &locality.LocalityName, &locality.ProvinceName, &locality.CountryName)

		if err != nil {
			return localityList, err
		}

		localityList = append(localityList, locality)
	}

	return localityList, err
}

func (m mariaDBRepository) GetById(ctx context.Context, id int) (Locality, error) {
	var locality Locality

	rows, err := m.db.QueryContext(ctx, GETBYID, id)

	if err != nil {
		return locality, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&locality.Id, &locality.ZipCode, &locality.LocalityName, &locality.ProvinceName, &locality.CountryName)

		if err != nil {
			return locality, err
		}

	}

	if err := rows.Err(); err != nil {
		return locality, err
	}

	return locality, nil
}
