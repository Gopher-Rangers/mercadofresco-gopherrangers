package products_test

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"context"

	products "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func mockRowsArray() *sqlmock.Rows {
	prod := createProductsArray()
	rows := sqlmock.NewRows([]string{
		"id", "product_code", "description",
		"width", "height", "length", "net_weight", "expiration_rate",
		"recommended_freezing_temperature", "freezing_rate",
		"product_type_id", "seller_id"}).AddRow(
		prod[0].ID, prod[0].ProductCode, prod[0].Description, prod[0].Width,
		prod[0].Height, prod[0].Length, prod[0].NetWeight,
		prod[0].ExpirationRate, prod[0].RecommendedFreezingTemperature,
		prod[0].FreezingRate, prod[0].ProductTypeId, prod[0].SellerId).AddRow(
		prod[1].ID, prod[1].ProductCode, prod[1].Description, prod[1].Width,
		prod[1].Height, prod[1].Length, prod[1].NetWeight,
		prod[1].ExpirationRate, prod[1].RecommendedFreezingTemperature,
		prod[1].FreezingRate, prod[1].ProductTypeId, prod[1].SellerId)
	return rows
}

func mockRow() *sqlmock.Rows {
	prod := createProductsArray()
	rows := sqlmock.NewRows([]string{
		"id", "product_code", "description",
		"width", "height", "length", "net_weight", "expiration_rate",
		"recommended_freezing_temperature", "freezing_rate",
		"product_type_id", "seller_id"}).AddRow(
		prod[0].ID, prod[0].ProductCode, prod[0].Description, prod[0].Width,
		prod[0].Height, prod[0].Length, prod[0].NetWeight,
		prod[0].ExpirationRate, prod[0].RecommendedFreezingTemperature,
		prod[0].FreezingRate, prod[0].ProductTypeId, prod[0].SellerId)
	return rows
}

func TestDBRepositoryStore(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.STORE))
		prod := createProductsArray()[0]
		stmt.ExpectExec().WithArgs(&prod.ProductCode, &prod.Description,
			&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
			&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
			&prod.FreezingRate, &prod.ProductTypeId,
			&prod.SellerId).WillReturnResult(sqlmock.NewResult(1, 1))
		productsRepo := products.NewRepository(db)
		result, err := productsRepo.Store(context.Background(), prod)
		assert.NoError(t, err)
		assert.Equal(t, result, prod)
	})
	t.Run("create_fail_exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.STORE))
		prod := createProductsArray()[0]
		stmt.ExpectExec().WithArgs(&prod.ProductCode, &prod.Description,
			&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
			&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
			&prod.FreezingRate, &prod.ProductTypeId,
			&prod.SellerId).WillReturnError(sql.ErrNoRows)
		productsRepo := products.NewRepository(db)
		result, err := productsRepo.Store(context.Background(), prod)
		assert.Error(t, err)
		assert.Equal(t, result, products.Product{})
	})
	t.Run("create_fail_zero_rows_affected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.STORE))
		prod := createProductsArray()[0]
		stmt.ExpectExec().WithArgs(&prod.ProductCode, &prod.Description,
			&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
			&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
			&prod.FreezingRate, &prod.ProductTypeId,
			&prod.SellerId).WillReturnResult(sqlmock.NewResult(1, 0))
		productsRepo := products.NewRepository(db)
		result, err := productsRepo.Store(context.Background(), prod)
		assert.Equal(t, err, fmt.Errorf("fail to save"))
		assert.Equal(t, result, products.Product{})
	})
}

func TestDBRepositoryGetAll(t *testing.T) {
	t.Run("find_all_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		prod := createProductsArray()
		rows := mockRowsArray()
		mock.ExpectQuery(regexp.QuoteMeta(products.GETALL)).WillReturnRows(rows)
		productsRepo := products.NewRepository(db)
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
			"id", "product_code", "description",
			"width", "height", "length", "net_weight", "expiration_rate",
			"recommended_freezing_temperature", "freezing_rate",
			"product_type_id", "seller_id"}).AddRow(
			"", "", "", "", "", "", "", "", "", "", "", "")
		mock.ExpectQuery(regexp.QuoteMeta(products.GETALL)).WillReturnRows(rows)
		productsRepo := products.NewRepository(db)
		prod, err := productsRepo.GetAll(context.Background())
		assert.Equal(t, prod, []products.Product(nil))
		assert.Error(t, err)
	})
	t.Run("find_all_fail_select", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		mock.ExpectQuery(regexp.QuoteMeta(
			products.GETALL)).WillReturnError(sql.ErrNoRows)
		productsRepo := products.NewRepository(db)
		prod, err := productsRepo.GetAll(context.Background())
		assert.Equal(t, prod, []products.Product(nil))
		assert.Error(t, err)
	})
}

func TestDBRepositoryGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		prod := createProductsArray()
		rows := mockRow()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.GETBYID))
		stmt.ExpectQuery().WithArgs(prod[0].ID).WillReturnRows(rows)
		productsRepo := products.NewRepository(db)
		result, err := productsRepo.GetById(context.Background(), prod[0].ID)
		assert.NoError(t, err)
		assert.Equal(t, result, prod[0])
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.GETBYID))
		errNotFound := fmt.Errorf("product 999 not found")
		stmt.ExpectQuery().WithArgs(999).WillReturnError(errNotFound)
		productsRepo := products.NewRepository(db)
		result, err := productsRepo.GetById(context.Background(), 999)
		assert.Equal(t, err, errNotFound)
		assert.Equal(t, result, products.Product{})
	})
	t.Run("find_by_id_fail_exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		prod := createProductsArray()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.GETBYID))
		stmt.ExpectQuery().WithArgs(prod[0].ID).WillReturnError(sql.ErrNoRows)
		productsRepo := products.NewRepository(db)
		result, err := productsRepo.GetById(context.Background(), prod[0].ID)
		assert.Equal(t, result, products.Product{})
		assert.Error(t, err)
	})
}

func TestDBRepositoryUpdate(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		prod := createProductsArray()[0]
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.UPDATE))
		stmt.ExpectExec().WithArgs(&prod.ProductCode, &prod.Description,
			&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
			&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
			&prod.FreezingRate, &prod.ProductTypeId,
			&prod.SellerId, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		productsRepo := products.NewRepository(db)
		result, err := productsRepo.Update(context.Background(), prod, 1)
		assert.NoError(t, err)
		assert.Equal(t, result, prod)
	})
	t.Run("update_non_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		prod := createProductsArray()[0]
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.UPDATE))
		stmt.ExpectExec().WithArgs(&prod.ProductCode, &prod.Description,
			&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
			&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
			&prod.FreezingRate, &prod.ProductTypeId,
			&prod.SellerId, 1).WillReturnResult(sqlmock.NewResult(1, 0))
		productsRepo := products.NewRepository(db)
		result, err := productsRepo.Update(context.Background(), prod, 1)
		assert.Equal(t, err, fmt.Errorf("product 1 not found"))
		assert.Equal(t, result, products.Product{})
	})
	t.Run("update_fail_exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		prod := createProductsArray()[0]
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.UPDATE))
		stmt.ExpectExec().WithArgs(&prod.ProductCode, &prod.Description,
			&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
			&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
			&prod.FreezingRate, &prod.ProductTypeId,
			&prod.SellerId, 1).WillReturnError(sql.ErrNoRows)
		productsRepo := products.NewRepository(db)
		result, err := productsRepo.Update(context.Background(), prod, 1)
		assert.Equal(t, result, products.Product{})
		assert.Error(t, err)
	})
}

func TestDBRepositoryDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.DELETE))
		stmt.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
		productsRepo := products.NewRepository(db)
		err = productsRepo.Delete(context.Background(), 1)
		assert.NoError(t, err)
	})
	t.Run("delete_non_existent", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.DELETE))
		stmt.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 0))
		productsRepo := products.NewRepository(db)
		err = productsRepo.Delete(context.Background(), 1)
		assert.Equal(t, err, fmt.Errorf("product 1 not found"))
	})
	t.Run("delete_fail_exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.DELETE))
		stmt.ExpectExec().WithArgs(1).WillReturnError(sql.ErrNoRows)
		productsRepo := products.NewRepository(db)
		err = productsRepo.Delete(context.Background(), 1)
		assert.Error(t, err)
	})
}
