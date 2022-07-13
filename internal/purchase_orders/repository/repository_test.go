package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/domain"
	purchaseOrdersRepo "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/repository"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestDBRepositoryGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		purchasesData := createBaseData()
		mockBuyers := []domain.PurchaseOrders{
			{
				ID:              1,
				OrderNumber:     "Order1",
				OrderDate:       "2008-11-11",
				TrackingCode:    "1",
				BuyerId:         1,
				ProductRecordId: 1,
				OrderStatusId:   1,
			},
		}
		rows := sqlmock.NewRows([]string{
			"id",
			"order_number",
			"order_date",
			"tracking_code",
			"buyer_id",
			"product_record_id",
			"order_status_id",
		}).AddRow(
			mockBuyers[0].ID,
			mockBuyers[0].OrderNumber,
			mockBuyers[0].OrderDate,
			mockBuyers[0].TrackingCode,
			mockBuyers[0].BuyerId,
			mockBuyers[0].ProductRecordId,
			mockBuyers[0].OrderStatusId,
		)

		mock.ExpectQuery(regexp.QuoteMeta(purchaseOrdersRepo.SqlGetById)).WithArgs(1).WillReturnRows(rows)

		buyersRepo := purchaseOrdersRepo.NewRepository(db)
		result, err := buyersRepo.GetById(context.Background(), purchasesData[0].ID)
		assert.NoError(t, err)
		assert.Equal(t, purchasesData[0], result)
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		errNotFound := fmt.Errorf("purchase order with id (10) not founded")

		mock.ExpectQuery(regexp.QuoteMeta(purchaseOrdersRepo.SqlGetById)).WithArgs(10).WillReturnError(errNotFound)

		purchaseRepository := purchaseOrdersRepo.NewRepository(db)
		result, err := purchaseRepository.GetById(context.Background(), 10)
		assert.Equal(t, err, errNotFound)
		assert.Equal(t, result, domain.PurchaseOrders{})
	})
	t.Run("find_by_id_fail_exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		purchases := createBaseData()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(purchaseOrdersRepo.SqlGetById))
		stmt.ExpectQuery().WithArgs(purchases[0].ID).WillReturnError(sql.ErrNoRows)
		buyersRepo := purchaseOrdersRepo.NewRepository(db)
		result, err := buyersRepo.GetById(context.Background(), purchases[0].ID)
		assert.Equal(t, result, domain.PurchaseOrders{})
		assert.Error(t, err)
	})
}

func TestRepositoryCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		purchaseOrder := createBaseData()[0]
		mock.ExpectExec(regexp.QuoteMeta(purchaseOrdersRepo.SqlCreate)).WithArgs(&purchaseOrder.OrderNumber, &purchaseOrder.OrderDate,
			&purchaseOrder.TrackingCode, &purchaseOrder.BuyerId, &purchaseOrder.ProductRecordId, &purchaseOrder.OrderStatusId).WillReturnResult(sqlmock.NewResult(1, 1))
		repo := purchaseOrdersRepo.NewRepository(db)
		result, err := repo.Create(context.Background(), purchaseOrder)
		assert.NoError(t, err)
		assert.Equal(t, result, purchaseOrder)
	})
	t.Run("create_fail_last_id", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		purchaseOrder := createBaseData()[0]
		mock.ExpectExec(regexp.QuoteMeta(purchaseOrdersRepo.SqlCreate)).WithArgs(&purchaseOrder.OrderNumber, &purchaseOrder.OrderDate,
			&purchaseOrder.TrackingCode, &purchaseOrder.BuyerId, &purchaseOrder.ProductRecordId, &purchaseOrder.OrderStatusId).WillReturnResult(sqlmock.NewResult(0, 1))
		repo := purchaseOrdersRepo.NewRepository(db)
		result, err := repo.Create(context.Background(), purchaseOrder)
		assert.Error(t, err)
		assert.Equal(t, result, domain.PurchaseOrders{})
	})
	t.Run("create_fail_exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		purchaseOrder := createBaseData()[0]
		mock.ExpectExec(regexp.QuoteMeta(purchaseOrdersRepo.SqlCreate)).WithArgs(&purchaseOrder.OrderNumber, &purchaseOrder.OrderDate,
			&purchaseOrder.TrackingCode, &purchaseOrder.BuyerId, &purchaseOrder.ProductRecordId, &purchaseOrder.OrderStatusId).WillReturnError(sql.ErrNoRows)
		repo := purchaseOrdersRepo.NewRepository(db)
		result, err := repo.Create(context.Background(), purchaseOrder)
		assert.Error(t, err)
		assert.Equal(t, result, domain.PurchaseOrders{})
	})
	t.Run("create_fail_zero_rows_affected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		purchaseOrder := createBaseData()[0]
		mock.ExpectExec(regexp.QuoteMeta(purchaseOrdersRepo.SqlCreate)).WithArgs(&purchaseOrder.OrderNumber, &purchaseOrder.OrderDate,
			&purchaseOrder.TrackingCode, &purchaseOrder.BuyerId, &purchaseOrder.ProductRecordId, &purchaseOrder.OrderStatusId).WillReturnResult(sqlmock.NewResult(1, 0))

		repo := purchaseOrdersRepo.NewRepository(db)
		result, err := repo.Create(context.Background(), purchaseOrder)
		assert.Equal(t, err, fmt.Errorf("error while saving"))
		assert.Equal(t, result, domain.PurchaseOrders{})
	})
}

func TestValidadeOrderNumber(t *testing.T) {
	t.Run("test_valid_order_number", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		purchasesData := createBaseData()
		rows := sqlmock.NewRows([]string{
			"result",
		}).AddRow(
			0,
		)

		mock.ExpectQuery(regexp.QuoteMeta(purchaseOrdersRepo.SqlExistsOrderNumber)).WithArgs(&purchasesData[0].OrderNumber).WillReturnRows(rows)

		repo := purchaseOrdersRepo.NewRepository(db)
		result, err := repo.ValidadeOrderNumber(purchasesData[0].OrderNumber)
		assert.NoError(t, err)
		assert.Equal(t, false, result)
	})
	t.Run("test_invalid_order_number", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		purchasesData := createBaseData()
		rows := sqlmock.NewRows([]string{
			"result",
		}).AddRow(
			1,
		)

		mock.ExpectQuery(regexp.QuoteMeta(purchaseOrdersRepo.SqlExistsOrderNumber)).WithArgs(&purchasesData[0].OrderNumber).WillReturnRows(rows)

		repo := purchaseOrdersRepo.NewRepository(db)
		result, err := repo.ValidadeOrderNumber(purchasesData[0].OrderNumber)
		assert.NoError(t, err)
		assert.Equal(t, true, result)
	})
	t.Run("test_validate_order_number_query_err", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		purchasesData := createBaseData()

		mock.ExpectQuery(regexp.QuoteMeta(purchaseOrdersRepo.SqlExistsOrderNumber)).WithArgs(&purchasesData[0].OrderNumber).WillReturnError(sql.ErrNoRows)

		repo := purchaseOrdersRepo.NewRepository(db)
		result, err := repo.ValidadeOrderNumber(purchasesData[0].OrderNumber)
		assert.Error(t, err)
		assert.Equal(t, false, result)
	})

}

func createBaseData() []domain.PurchaseOrders {
	var purchases []domain.PurchaseOrders
	purchaseOne := domain.PurchaseOrders{
		ID:              1,
		OrderNumber:     "Order1",
		OrderDate:       "2008-11-11",
		TrackingCode:    "1",
		BuyerId:         1,
		ProductRecordId: 1,
		OrderStatusId:   1,
	}
	purchaseTwo := domain.PurchaseOrders{
		ID:              1,
		OrderNumber:     "Order1",
		OrderDate:       "2008-11-11",
		TrackingCode:    "1",
		BuyerId:         1,
		ProductRecordId: 1,
		OrderStatusId:   1,
	}
	purchases = append(purchases, purchaseOne, purchaseTwo)
	return purchases
}
