package productrecord_test

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	productrecord "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_record"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func mockRow() *sqlmock.Rows {
	prod := createProductRecordGetArray()
	rows := sqlmock.NewRows([]string{
		"product_id", "description", "records_count"}).AddRow(
		prod[0].ProductId, prod[0].Description, prod[0].RecordsCount)
	return rows
}

func mockRowsArray() *sqlmock.Rows {
	prod := createProductRecordGetArray()
	rows := sqlmock.NewRows([]string{
		"product_id", "description", "records_count"}).AddRow(
		prod[0].ProductId, prod[0].Description, prod[0].RecordsCount).AddRow(
		prod[1].ProductId, prod[1].Description, prod[1].RecordsCount)
	return rows
}

func TestRepositoryStore(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(productrecord.STORE))
		prod := createProductRecordArray()[0]
		stmt.ExpectExec().WithArgs(&prod.LastUpdateDate, &prod.PurchasePrice,
			&prod.SalePrice, &prod.ProductId).WillReturnResult(
			sqlmock.NewResult(1, 1))
		productsRepo := productrecord.NewRepository(db)
		result, err := productsRepo.Store(context.Background(), prod)
		assert.NoError(t, err)
		assert.Equal(t, result, prod)
	})
	t.Run("create_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(productrecord.STORE))
		prod := createProductRecordArray()[0]
		stmt.ExpectExec().WithArgs(&prod.LastUpdateDate, &prod.PurchasePrice,
			&prod.SalePrice, &prod.ProductId).WillReturnResult(
			sqlmock.NewResult(1, 1))
		productsRepo := productrecord.NewRepository(db)
		result, err := productsRepo.Store(context.Background(), prod)
		assert.NoError(t, err)
		assert.Equal(t, result, prod)
	})
	t.Run("create_prepare_fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		errPrepare := fmt.Errorf("fail to preprare")
		mock.ExpectPrepare(regexp.QuoteMeta(
			productrecord.STORE)).WillReturnError(errPrepare)
		prod := createProductRecordArray()[0]
		productsRepo := productrecord.NewRepository(db)
		result, err := productsRepo.Store(context.Background(), prod)
		assert.Equal(t, err, errPrepare)
		assert.Equal(t, result, productrecord.ProductRecord{})
	})
	t.Run("create_fail_exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(productrecord.STORE))
		prod := createProductRecordArray()[0]
		stmt.ExpectExec().WithArgs(&prod.LastUpdateDate, &prod.PurchasePrice,
			&prod.SalePrice, &prod.ProductId).WillReturnError(sql.ErrNoRows)
		productsRepo := productrecord.NewRepository(db)
		result, err := productsRepo.Store(context.Background(), prod)
		assert.Error(t, err)
		assert.Equal(t, result, productrecord.ProductRecord{})
	})
	t.Run("create_fail_zero_rows_affected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(productrecord.STORE))
		prod := createProductRecordArray()[0]
		stmt.ExpectExec().WithArgs(&prod.LastUpdateDate, &prod.PurchasePrice,
			&prod.SalePrice, &prod.ProductId).WillReturnResult(
			sqlmock.NewResult(1, 0))
		productsRepo := productrecord.NewRepository(db)
		result, err := productsRepo.Store(context.Background(), prod)
		assert.Equal(t, err, fmt.Errorf("fail to save"))
		assert.Equal(t, result, productrecord.ProductRecord{})
	})
}

func TestRepositoryGetAll(t *testing.T) {
	t.Run("find_all_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		prod := createProductRecordGetArray()
		rows := mockRowsArray()
		mock.ExpectQuery(regexp.QuoteMeta(
			productrecord.GETALL)).WillReturnRows(rows)
		productsRepo := productrecord.NewRepository(db)
		result, err := productsRepo.GetAll(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, result[0], prod[0])
		assert.Equal(t, result[1], prod[1])
	})
	t.Run("find_all_fail_scan", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		rows := sqlmock.NewRows([]string{
			"product_id", "description", "records_count"}).AddRow(
			"", "", "")
		mock.ExpectQuery(regexp.QuoteMeta(
			productrecord.GETALL)).WillReturnRows(rows)
		productsRepo := productrecord.NewRepository(db)
		prod, err := productsRepo.GetAll(context.Background())
		assert.Equal(t, prod, []productrecord.ProductRecordGet(nil))
		assert.Error(t, err)
	})
	t.Run("find_all_fail_select", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(regexp.QuoteMeta(
			productrecord.GETALL)).WillReturnError(sql.ErrNoRows)
		productsRepo := productrecord.NewRepository(db)
		prod, err := productsRepo.GetAll(context.Background())
		assert.Equal(t, prod, []productrecord.ProductRecordGet(nil))
		assert.Error(t, err)
	})
}

func TestRepositoryGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		prod := createProductArray()
		prodGet := createProductRecordGetArray()
		rows := mockRow()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(productrecord.GETBYID))
		stmt.ExpectQuery().WithArgs(prod[0].ID).WillReturnRows(rows)
		productsRepo := productrecord.NewRepository(db)
		result, err := productsRepo.GetById(context.Background(), prod[0].ID)
		assert.NoError(t, err)
		assert.Equal(t, result, prodGet[0])
	})
	t.Run("find_by_id_prepare_fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		prod := createProductArray()
		errPrepare := fmt.Errorf("fail to preprare")
		mock.ExpectPrepare(regexp.QuoteMeta(
			productrecord.GETBYID)).WillReturnError(errPrepare)
		productsRepo := productrecord.NewRepository(db)
		result, err := productsRepo.GetById(context.Background(), prod[0].ID)
		assert.Equal(t, err, errPrepare)
		assert.Equal(t, result, productrecord.ProductRecordGet{})
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(productrecord.GETBYID))
		errNotFound := fmt.Errorf("product record 999 not found")
		stmt.ExpectQuery().WithArgs(999).WillReturnError(errNotFound)
		productsRepo := productrecord.NewRepository(db)
		result, err := productsRepo.GetById(context.Background(), 999)
		assert.Equal(t, err, errNotFound)
		assert.Equal(t, result, productrecord.ProductRecordGet{})
	})
	t.Run("find_by_id_fail_exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		prod := createProductArray()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(productrecord.GETBYID))
		stmt.ExpectQuery().WithArgs(prod[0].ID).WillReturnError(sql.ErrNoRows)
		productsRepo := productrecord.NewRepository(db)
		result, err := productsRepo.GetById(context.Background(), prod[0].ID)
		assert.Equal(t, result, productrecord.ProductRecordGet{})
		assert.Error(t, err)
	})
}
