package service_test

import (
	"context"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/domain/mocks"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		newService := service.NewService(mockRepository)
		purchasesData := createBaseData()
		mockRepository.On("GetById", ctx, 1).Return(purchasesData[0], nil)
		purchaseData, err := newService.GetById(ctx, 1)
		assert.Nil(t, err)
		assert.Equal(t, purchaseData, purchasesData[0])
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		serv := service.NewService(mockRepository)
		mockRepository.On("GetById", ctx, 10).Return(domain.PurchaseOrders{}, fmt.Errorf("purchase order with id %d not founded", 10))
		foundedBuyer, err := serv.GetById(ctx, 10)
		assert.Equal(t, fmt.Errorf("purchase order with id %d not founded", 10), err)
		assert.Equal(t, foundedBuyer, domain.PurchaseOrders{})
	})
}

func TestCreate(t *testing.T) {
	t.Run("create_conflict", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		newService := service.NewService(mockRepository)
		expected := domain.PurchaseOrders{
			ID:              1,
			OrderNumber:     "Order1",
			OrderDate:       "2008-11-11",
			TrackingCode:    "1",
			BuyerId:         1,
			ProductRecordId: 1,
			OrderStatusId:   1,
		}
		mockRepository.On("ValidadeOrderNumber", expected.OrderNumber).Return(false, nil)
		_, err := newService.Create(ctx, expected)
		assert.Equal(t, fmt.Errorf("order number: Order1 already exist"), err)
	})
	t.Run("create_conflict_error", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		newService := service.NewService(mockRepository)
		expected := domain.PurchaseOrders{
			ID:              1,
			OrderNumber:     "Order1",
			OrderDate:       "2008-11-11",
			TrackingCode:    "1",
			BuyerId:         1,
			ProductRecordId: 1,
			OrderStatusId:   1,
		}
		mockRepository.On("ValidadeOrderNumber", expected.OrderNumber).Return(false, fmt.Errorf("error"))
		_, err := newService.Create(ctx, expected)
		assert.Equal(t, fmt.Errorf("error"), err)
	})
	t.Run("create_ok", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		newService := service.NewService(mockRepository)
		//baseData := createBaseData()
		expected := domain.PurchaseOrders{
			ID:              1,
			OrderNumber:     "Order1",
			OrderDate:       "2008-11-11",
			TrackingCode:    "1",
			BuyerId:         1,
			ProductRecordId: 1,
			OrderStatusId:   1,
		}
		mockRepository.On("ValidadeOrderNumber", expected.OrderNumber).Return(true, nil)
		expected.ID = 1
		mockRepository.On("Create", ctx, expected).Return(expected, nil)
		newPurchase, err := newService.Create(ctx, expected)
		assert.Nil(t, err)
		assert.Equal(t, expected, newPurchase)
	})
	t.Run("create_error", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		newService := service.NewService(mockRepository)
		//baseData := createBaseData()
		expected := domain.PurchaseOrders{
			ID:              1,
			OrderNumber:     "Order1",
			OrderDate:       "2008-11-11",
			TrackingCode:    "1",
			BuyerId:         1,
			ProductRecordId: 1,
			OrderStatusId:   1,
		}
		mockRepository.On("ValidadeOrderNumber", expected.OrderNumber).Return(true, nil)
		expected.ID = 1
		mockRepository.On("Create", ctx, expected).Return(domain.PurchaseOrders{}, fmt.Errorf("error"))
		newPurchase, err := newService.Create(ctx, expected)
		assert.Error(t, err)
		assert.Equal(t, newPurchase, domain.PurchaseOrders{})
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
