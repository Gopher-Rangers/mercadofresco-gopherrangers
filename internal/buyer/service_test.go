package buyer_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/mocks"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		mockRepository.On("Delete", ctx, 25735482).Return(nil)
		service := buyer.NewService(mockRepository)
		err := service.Delete(ctx, 25735482)
		assert.Nil(t, err)
	})
	t.Run("delete_non_existent", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		mockRepository.On("Delete", ctx, 9).Return(fmt.Errorf("buyer with id : %d not founded", 9))

		service := buyer.NewService(mockRepository)
		err := service.Delete(ctx, 9)

		assert.Equal(t, fmt.Errorf("buyer with id : %d not founded", 9), err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		service := buyer.NewService(mockRepository)
		mockedBuyers := createBaseData()
		mockRepository.On("GetAll", ctx).Return(mockedBuyers, nil)
		buyersFromTest, err := service.GetAll(ctx)
		assert.Nil(t, err)
		assert.Equal(t, buyersFromTest, mockedBuyers)
		assert.Equal(t, len(buyersFromTest), len(mockedBuyers))
	})
	t.Run("find_all_error", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		service := buyer.NewService(mockRepository)
		mockRepository.On("GetAll", ctx).Return(nil, fmt.Errorf("Could not getAll"))
		_, err := service.GetAll(ctx)
		assert.Equal(t, err, fmt.Errorf("Could not getAll"))
	})
}

func createBaseData() []buyer.Buyer {
	var buyers []buyer.Buyer
	buyerOne := buyer.Buyer{
		ID:           25735482,
		CardNumberId: "Card1",
		FirstName:    "Victor",
		LastName:     "Beltramini",
	}
	buyerTwo := buyer.Buyer{
		ID:           25735582,
		CardNumberId: "Card2",
		FirstName:    "Victor",
		LastName:     "Beltramini",
	}
	buyers = append(buyers, buyerOne, buyerTwo)
	return buyers
}

func TestGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		service := buyer.NewService(mockRepository)
		mockedBuyers := createBaseData()
		mockRepository.On("GetById", ctx, 25735482).Return(mockedBuyers[0], nil)
		foundedBuyer, err := service.GetById(ctx, 25735482)
		assert.Nil(t, err)
		assert.Equal(t, foundedBuyer, mockedBuyers[0])
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		service := buyer.NewService(mockRepository)
		mockRepository.On("GetById", ctx, 25735481).Return(buyer.Buyer{}, fmt.Errorf("buyer with id %d not founded", 25735481))
		foundedBuyer, err := service.GetById(ctx, 25735481)
		assert.Equal(t, fmt.Errorf("buyer with id %d not founded", 25735481), err)
		assert.Equal(t, foundedBuyer, buyer.Buyer{})
	})
}

func TestCreate(t *testing.T) {
	t.Run("create_conflict", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		service := buyer.NewService(mockRepository)
		baseData := createBaseData()
		expected := buyer.Buyer{
			CardNumberId: "Card3",
			FirstName:    "Victor",
			LastName:     "Beltramini",
		}
		mockRepository.On("GetAll", ctx).Return(baseData, nil)
		expected.ID = 25735482
		mockRepository.On("Create", ctx, expected).Return(buyer.Buyer{}, fmt.Errorf("buyer with card_number_id %s already exists", expected.CardNumberId))
		newBuyer, err := service.Create(ctx, expected)
		assert.Equal(t, err, fmt.Errorf("buyer with card_number_id %s already exists", expected.CardNumberId))
		assert.Equal(t, buyer.Buyer{}, newBuyer)
	})
	t.Run("create_conflict_service", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		service := buyer.NewService(mockRepository)
		buyersData := createBaseData()
		expected := buyer.Buyer{
			ID:           25735482,
			CardNumberId: "Card2",
			FirstName:    "Victor",
			LastName:     "Beltramini",
		}
		mockRepository.On("GetAll", ctx).Return(buyersData, nil)
		_, err := service.Create(ctx, expected)
		assert.Equal(t, fmt.Errorf("buyer with card_number_id %s already exists", expected.CardNumberId), err)
	})
	t.Run("create_ok", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		service := buyer.NewService(mockRepository)
		baseData := createBaseData()
		expected := buyer.Buyer{
			CardNumberId: "Card3",
			FirstName:    "Victor",
			LastName:     "Beltramini",
		}
		mockRepository.On("GetAll", ctx).Return(baseData, nil)
		expected.ID = 25735482
		mockRepository.On("Create", ctx, expected).Return(expected, nil)
		newBuyer, err := service.Create(ctx, expected)
		assert.Nil(t, err)
		assert.Equal(t, expected.CardNumberId, newBuyer.CardNumberId)
		assert.Equal(t, expected.FirstName, newBuyer.FirstName)
		assert.Equal(t, expected.LastName, newBuyer.LastName)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update_existent", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		service := buyer.NewService(mockRepository)
		buyersData := createBaseData()
		expected := buyer.Buyer{
			ID:           25735482,
			CardNumberId: "Card77",
			FirstName:    "Victor",
			LastName:     "Beltramini",
		}
		mockRepository.On("GetAll", ctx).Return(buyersData, nil)
		mockRepository.On("Update", ctx, expected).Return(expected, nil)
		prod, err := service.Update(ctx, expected)
		assert.Nil(t, err)
		assert.Equal(t, prod, expected)
	})
	t.Run("update_non_existent", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		service := buyer.NewService(mockRepository)
		buyersData := createBaseData()
		expected := buyer.Buyer{
			ID:           22735482,
			CardNumberId: "Card77",
			FirstName:    "Victor",
			LastName:     "Beltramini",
		}
		mockRepository.On("GetAll", ctx).Return(buyersData, nil)
		mockRepository.On("Update", ctx, expected).Return(buyer.Buyer{}, fmt.Errorf("buyer with id: %d not found", expected.ID))
		prod, err := service.Update(ctx, expected)
		assert.Equal(t, fmt.Errorf("buyer with id: %d not found", expected.ID), err)
		assert.Equal(t, prod, buyer.Buyer{})
	})
	t.Run("update_conflict", func(t *testing.T) {
		ctx := context.Background()
		mockRepository := mocks.NewRepository(t)
		service := buyer.NewService(mockRepository)
		buyersData := createBaseData()
		expected := buyer.Buyer{
			ID:           25735482,
			CardNumberId: "Card2",
			FirstName:    "Victor",
			LastName:     "Beltramini",
		}
		mockRepository.On("GetAll", ctx).Return(buyersData, nil)
		_, err := service.Update(ctx, expected)
		assert.Equal(t, fmt.Errorf("buyer with card_number_id %s already exists", expected.CardNumberId), err)
	})
}