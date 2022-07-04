package myslq_test

import (
	"context"
	"database/sql"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/repository/myslq"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mockBuyers := []domain.Buyer{
		{
			ID:           1,
			CardNumberId: "Card1",
			FirstName:    "Victor",
			LastName:     "Beltramini",
		},
		{
			ID:           2,
			CardNumberId: "Card2",
			FirstName:    "Hugo",
			LastName:     "Beltramini",
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name",
	}).AddRow(
		mockBuyers[0].ID,
		mockBuyers[0].CardNumberId,
		mockBuyers[0].FirstName,
		mockBuyers[0].LastName,
	).AddRow(
		mockBuyers[1].ID,
		mockBuyers[1].CardNumberId,
		mockBuyers[1].FirstName,
		mockBuyers[1].LastName,
	)

	query := "SELECT \\* FROM `mercado-fresco`.`buyers`"

	mock.ExpectQuery(query).WillReturnRows(rows)

	buyersRepo := myslq.NewRepository(db)

	result, err := buyersRepo.GetAll(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, result[0].FirstName, "Victor")
	assert.Equal(t, result[1].FirstName, "Hugo")
}

func TestGetAllFailScan(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name",
	}).AddRow("", "", "", "")

	query := "SELECT \\* FROM `mercado-fresco`.`buyers`"

	mock.ExpectQuery(query).WillReturnRows(rows)

	buyersRepo := myslq.NewRepository(db)

	_, err = buyersRepo.GetAll(context.Background())
	assert.Error(t, err)
}

func TestGetAllFailSelect(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	query := "SELECT \\* FROM `mercado-fresco`.`buyers`"

	mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)

	buyersRepo := myslq.NewRepository(db)

	_, err = buyersRepo.GetAll(context.Background())
	assert.Error(t, err)
}

//func TestRepositoryGetByIdAll(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	assert.NoError(t, err)
//	defer db.Close()
//
//	mockBuyers := []buyer.Buyer{
//		{
//			ID:           1,
//			CardNumberId: "Card1",
//			FirstName:    "Victor",
//			LastName:     "Beltramini",
//		},
//	}
//
//	rows := sqlmock.NewRows([]string{
//		"id", "card_number_id", "first_name", "last_name",
//	}).AddRow(
//		mockBuyers[0].ID,
//		mockBuyers[0].CardNumberId,
//		mockBuyers[0].FirstName,
//		mockBuyers[0].LastName,
//	)
//
//	query := "SELECT * FROM `mercado-fresco`.`buyers` WHERE id=?"
//
//	mock.ExpectQuery(query).WillReturnRows(rows)
//
//	buyersRepo := buyer.NewRepository(db)
//
//	result, err := buyersRepo.GetById(context.Background(), 1)
//	assert.NoError(t, err)
//
//	fmt.Println(result)
//
//	assert.NotNil(t, result)
//
//	assert.Equal(t, result.FirstName, "Victor")
//}
