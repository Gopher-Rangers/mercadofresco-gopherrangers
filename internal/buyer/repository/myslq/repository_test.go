package myslq_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/repository/myslq"
	buyersRepository "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/repository/myslq"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func mockRows() *sqlmock.Rows {
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
	return rows
}

func TestRepositoryGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	rows := mockRows()
	query := "SELECT \\* FROM buyers"

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

	query := "SELECT \\* FROM buyersRepository`"

	mock.ExpectQuery(query).WillReturnRows(rows)

	buyersRepo := myslq.NewRepository(db)

	_, err = buyersRepo.GetAll(context.Background())
	assert.Error(t, err)
}

func TestGetAllFailSelect(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	query := "SELECT \\* FROM buyers"

	mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)

	buyersRepo := myslq.NewRepository(db)

	_, err = buyersRepo.GetAll(context.Background())
	assert.Error(t, err)
}

func TestRepositoryCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		buyer := createBaseData()[0]
		mock.ExpectExec(regexp.QuoteMeta(buyersRepository.SqlStore)).WithArgs(&buyer.CardNumberId, &buyer.FirstName, &buyer.LastName).WillReturnResult(sqlmock.NewResult(1, 1))
		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.Create(context.Background(), buyer)
		assert.NoError(t, err)
		assert.Equal(t, result, buyer)
	})
	t.Run("create_fail_exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(buyersRepository.SqlStore))
		buyer := createBaseData()[0]
		stmt.ExpectExec().WithArgs(&buyer.CardNumberId, &buyer.FirstName,
			&buyer.LastName).WillReturnError(sql.ErrNoRows)
		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.Create(context.Background(), buyer)
		assert.Error(t, err)
		assert.Equal(t, result, domain.Buyer{})
	})
	t.Run("create_fail_zero_rows_affected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		buyer := createBaseData()[0]
		mock.ExpectExec(regexp.QuoteMeta(buyersRepository.SqlStore)).WithArgs(
			&buyer.CardNumberId, &buyer.FirstName, &buyer.LastName).WillReturnResult(
			sqlmock.NewResult(1, 0))

		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.Create(context.Background(), buyer)
		assert.Equal(t, err, fmt.Errorf("error while saving"))
		assert.Equal(t, result, domain.Buyer{})
	})
}

func TestRepositoryGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		buyersData := createBaseData()
		mockBuyers := []domain.Buyer{
			{
				ID:           1,
				CardNumberId: "Card1",
				FirstName:    "Victor",
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
		)

		mock.ExpectQuery(regexp.QuoteMeta(buyersRepository.SqlGetById)).WithArgs(1).WillReturnRows(rows)

		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.GetById(context.Background(), buyersData[0].ID)
		assert.NoError(t, err)
		assert.Equal(t, buyersData[0], result)
	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		errNotFound := fmt.Errorf("buyer with id (10) not founded")

		mock.ExpectQuery(regexp.QuoteMeta(buyersRepository.SqlGetById)).WithArgs(10).WillReturnError(errNotFound)

		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.GetById(context.Background(), 10)
		assert.Equal(t, err, errNotFound)
		assert.Equal(t, result, domain.Buyer{})
	})

	t.Run("find_by_id_fail_exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		prod := createBaseData()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(buyersRepository.SqlGetById))
		stmt.ExpectQuery().WithArgs(prod[0].ID).WillReturnError(sql.ErrNoRows)
		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.GetById(context.Background(), prod[0].ID)
		assert.Equal(t, result, domain.Buyer{})
		assert.Error(t, err)
	})
}

func TestGetBuyerOrdersById(t *testing.T) {
	t.Run("find_by_id_existent_with_purchase_orders", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		buyersData := createBaseDataWithPurchaseOrders()
		mockBuyers := []domain.BuyerTotalOrders{
			{
				ID:                  1,
				CardNumberId:        "Card1",
				FirstName:           "Victor",
				LastName:            "Beltramini",
				PurchaseOrdersCount: 1,
			},
		}
		rows := sqlmock.NewRows([]string{
			"id", "card_number_id", "first_name", "last_name", "purchase_orders_count",
		}).AddRow(
			mockBuyers[0].ID,
			mockBuyers[0].CardNumberId,
			mockBuyers[0].FirstName,
			mockBuyers[0].LastName,
			mockBuyers[0].PurchaseOrdersCount,
		)

		mock.ExpectQuery(regexp.QuoteMeta(buyersRepository.SqlBuyerWithOrdersById)).WithArgs(1).WillReturnRows(rows)

		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.GetBuyerOrdersById(context.Background(), buyersData[0].ID)
		assert.NoError(t, err)
		assert.Equal(t, buyersData[0], result)
	})
	t.Run("find_by_id_non_existent_with_purchase_orders", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		errNotFound := fmt.Errorf("buyer with id (10) not founded")

		mock.ExpectQuery(regexp.QuoteMeta(buyersRepository.SqlBuyerWithOrdersById)).WithArgs(10).WillReturnError(errNotFound)

		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.GetBuyerOrdersById(context.Background(), 10)
		assert.Equal(t, err, errNotFound)
		assert.Equal(t, result, domain.BuyerTotalOrders{})
	})
	t.Run("find_by_id_fail_exec_with_purchase_orders", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		prod := createBaseData()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(buyersRepository.SqlBuyerWithOrdersById))
		stmt.ExpectQuery().WithArgs(prod[0].ID).WillReturnError(sql.ErrNoRows)
		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.GetBuyerOrdersById(context.Background(), prod[0].ID)
		assert.Equal(t, result, domain.BuyerTotalOrders{})
		assert.Error(t, err)
	})
}

func TestGetBuyerOrders(t *testing.T) {
	t.Run("get_all_existent_with_purchase_orders", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		buyersData := createBaseDataWithPurchaseOrders()

		rows := sqlmock.NewRows([]string{
			"id", "card_number_id", "first_name", "last_name", "purchase_orders_count",
		}).AddRow(
			buyersData[0].ID,
			buyersData[0].CardNumberId,
			buyersData[0].FirstName,
			buyersData[0].LastName,
			buyersData[0].PurchaseOrdersCount,
		).AddRow(
			buyersData[1].ID,
			buyersData[1].CardNumberId,
			buyersData[1].FirstName,
			buyersData[1].LastName,
			buyersData[1].PurchaseOrdersCount,
		)

		mock.ExpectQuery(regexp.QuoteMeta(buyersRepository.SqlBuyersWithOrders)).WillReturnRows(rows)

		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.GetBuyerTotalOrders(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, buyersData[0], result[0])
		assert.Equal(t, buyersData[1], result[1])
	})
	t.Run("get_all_error_with_purchase_orders", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		errNotFound := fmt.Errorf("err")

		mock.ExpectQuery(regexp.QuoteMeta(buyersRepository.SqlBuyersWithOrders)).WillReturnError(errNotFound)

		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.GetBuyerTotalOrders(context.Background())
		assert.Equal(t, err, errNotFound)
		assert.Nil(t, result)
	})
	t.Run("get_all_fail_exec_with_purchase_orders", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		//order := createBaseData()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(buyersRepository.SqlBuyersWithOrders))
		stmt.ExpectQuery().WillReturnError(sql.ErrNoRows)
		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.GetBuyerTotalOrders(context.Background())
		assert.Nil(t, result)
		assert.Error(t, err)
	})
}

func TestRepositoryUpdate(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		buyer := createBaseData()[0]

		mock.ExpectExec(regexp.QuoteMeta(buyersRepository.SqlUpdate)).WithArgs(&buyer.CardNumberId, &buyer.FirstName, &buyer.LastName, 1).WillReturnResult(sqlmock.NewResult(1, 1))

		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.Update(context.Background(), buyer)
		assert.NoError(t, err)
		assert.Equal(t, result, buyer)
	})
	t.Run("update_non_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		buyer := createBaseData()[0]

		mock.ExpectExec(regexp.QuoteMeta(buyersRepository.SqlUpdate)).WithArgs(&buyer.CardNumberId, &buyer.FirstName, &buyer.LastName, &buyer.ID).WillReturnResult(sqlmock.NewResult(1, 0))

		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.Update(context.Background(), buyer)
		assert.Equal(t, err, fmt.Errorf("buyer wiht id (1) not founded"))
		assert.Equal(t, domain.Buyer{}, result)
	})
	t.Run("update_fail_exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		buyer := createBaseData()[0]

		mock.ExpectQuery(regexp.QuoteMeta(buyersRepository.SqlGetById)).WithArgs(&buyer.CardNumberId, &buyer.FirstName,
			&buyer.LastName, &buyer.ID).WillReturnError(sql.ErrNoRows)

		buyersRepo := buyersRepository.NewRepository(db)
		result, err := buyersRepo.Update(context.Background(), buyer)
		assert.Equal(t, result, domain.Buyer{})
		assert.Error(t, err)
	})
}

func TestRepositoryDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectExec(regexp.QuoteMeta(buyersRepository.SqlDelete)).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

		buyersRepo := buyersRepository.NewRepository(db)
		err = buyersRepo.Delete(context.Background(), 1)
		assert.NoError(t, err)
	})
	t.Run("delete_non_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(buyersRepository.SqlDelete)).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 0))

		buyersRepo := buyersRepository.NewRepository(db)
		err = buyersRepo.Delete(context.Background(), 1)
		assert.Equal(t, err, fmt.Errorf("buyer with id (1) not founded"))
	})
	t.Run("delete_fail_exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(buyersRepository.SqlDelete))
		stmt.ExpectExec().WithArgs(1).WillReturnError(sql.ErrNoRows)
		buyersRepo := buyersRepository.NewRepository(db)
		err = buyersRepo.Delete(context.Background(), 1)
		assert.Error(t, err)
	})
}

func createBaseData() []domain.Buyer {
	var buyers []domain.Buyer
	buyerOne := domain.Buyer{
		ID:           1,
		CardNumberId: "Card1",
		FirstName:    "Victor",
		LastName:     "Beltramini",
	}
	buyerTwo := domain.Buyer{
		ID:           2,
		CardNumberId: "Card2",
		FirstName:    "Victor",
		LastName:     "Beltramini",
	}
	buyers = append(buyers, buyerOne, buyerTwo)
	return buyers
}
func createBaseDataWithPurchaseOrders() []domain.BuyerTotalOrders {
	var buyers []domain.BuyerTotalOrders
	buyerOne := domain.BuyerTotalOrders{
		ID:                  1,
		CardNumberId:        "Card1",
		FirstName:           "Victor",
		LastName:            "Beltramini",
		PurchaseOrdersCount: 1,
	}
	buyerTwo := domain.BuyerTotalOrders{
		ID:                  2,
		CardNumberId:        "Card2",
		FirstName:           "Victor",
		LastName:            "Beltramini",
		PurchaseOrdersCount: 1,
	}
	buyers = append(buyers, buyerOne, buyerTwo)
	return buyers
}

func TestValidadeOrderNumber(t *testing.T) {
	t.Run("test_valid_order_number", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		buyersData := createBaseData()
		rows := sqlmock.NewRows([]string{
			"",
		}).AddRow(
			1,
		)
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(buyersRepository.SqlUniqueCardNumberId))
		stmt.ExpectQuery().WithArgs(buyersData[0].ID, buyersData[0].CardNumberId).WillReturnRows(rows)

		repo := buyersRepository.NewRepository(db)
		result, err := repo.ValidateCardNumberId(context.Background(), buyersData[0].ID, buyersData[0].CardNumberId)
		assert.NoError(t, err)
		assert.Equal(t, false, result)
	})
	t.Run("test_invalid_order_number", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		buyersData := createBaseData()
		rows := sqlmock.NewRows([]string{
			"result",
		}).AddRow(
			0,
		)

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(buyersRepository.SqlUniqueCardNumberId))
		stmt.ExpectQuery().WithArgs(buyersData[0].ID, buyersData[0].CardNumberId).WillReturnRows(rows)

		repo := buyersRepository.NewRepository(db)
		result, err := repo.ValidateCardNumberId(context.Background(), buyersData[0].ID, buyersData[0].CardNumberId)
		assert.NoError(t, err)
		assert.Equal(t, true, result)
	})
	t.Run("test_validate_order_number_query_err", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		buyersData := createBaseData()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(buyersRepository.SqlUniqueCardNumberId))
		stmt.ExpectQuery().WithArgs(buyersData[0].ID, buyersData[0].CardNumberId).WillReturnError(sql.ErrNoRows)

		repo := buyersRepository.NewRepository(db)
		result, err := repo.ValidateCardNumberId(context.Background(), buyersData[0].ID, buyersData[0].CardNumberId)
		assert.Equal(t, true, result)
	})

}
