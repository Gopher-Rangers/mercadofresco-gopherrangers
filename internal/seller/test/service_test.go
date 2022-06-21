package test

import (
	"errors"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Delete(t *testing.T) {
	t.Run("Se a exclusão for bem-sucedida, o item não aparecerá na lista.", func(t *testing.T) {
		mockRepo := &mocks.MockRepository{}

		var id int = 1

		sellerList := []seller.Seller{{Id: 1, CompanyId: 5, CompanyName: "TestDelete", Address: "BR", Telephone: "5501154545454"},
			{Id: 3, CompanyId: 6, CompanyName: "TestDelete", Address: "BR", Telephone: "5501154545454"},
			{Id: 4, CompanyId: 6, CompanyName: "TestDelete", Address: "BR", Telephone: "5501154545454"}}

		mockRepo.On("GetOne", id).Return(sellerList[0], nil)
		mockRepo.On("Delete", 1).Return(nil)

		service := seller.NewService(mockRepo)
		err := service.Delete(id)

		assert.Nil(t, err)
	})

	t.Run("Se o elemento a ser removido não existir, retornará null.", func(t *testing.T) {
		mockRepo := &mocks.MockRepository{}

		var id int = 2

		expectedError := fmt.Errorf("the id %d does not exists", id)

		mockRepo.On("GetOne", id).Return(seller.Seller{}, expectedError)
		mockRepo.On("Delete", id).Return(expectedError)

		service := seller.NewService(mockRepo)
		err := service.Delete(id)

		assert.Equal(t, expectedError, err)
	})
}

func TestService_Update(t *testing.T) {
	t.Run("Se os campos forem atualizados com sucesso retornará a informação do elemento atualizado", func(t *testing.T) {

		mockRepo := &mocks.MockRepository{}

		sellerList := []seller.Seller{{Id: 1, CompanyId: 5, CompanyName: "TestUpdate", Address: "BR", Telephone: "5501154545454"},
			{Id: 3, CompanyId: 6, CompanyName: "ServiceSeller", Address: "BR", Telephone: "5501154545454"}}

		expectedResult := seller.Seller{Id: 1, CompanyId: 5, CompanyName: "Meli", Address: "América do Sul", Telephone: "5501154545454"}

		mockRepo.On("GetOne", 1).Return(sellerList[0], nil)
		mockRepo.On("GetAll").Return(sellerList, nil)
		mockRepo.On("Update", expectedResult.CompanyId, expectedResult.CompanyName, expectedResult.Address, expectedResult.Telephone, sellerList[0]).Return(expectedResult, nil)

		service := seller.NewService(mockRepo)
		response, _ := service.Update(1, 5, "Meli", "América do Sul", "5501154545454")

		assert.Equal(t, expectedResult, response)
	})

	t.Run("Se o elemento a ser atualizado não existir, retornar null", func(t *testing.T) {
		mockRepo := &mocks.MockRepository{}

		var id int = 2

		sellerList := []seller.Seller{{Id: 1, CompanyId: 5, CompanyName: "TestUpdate", Address: "BR", Telephone: "5501154545454"},
			{Id: 3, CompanyId: 6, CompanyName: "ServiceSeller", Address: "BR", Telephone: "5501154545454"}}

		expectedResult := seller.Seller{}
		expectedError := fmt.Errorf("the id %d does not exists", id)

		mockRepo.On("GetOne", id).Return(seller.Seller{}, expectedError)
		mockRepo.On("GetAll").Return(sellerList, nil)
		mockRepo.On("Update", expectedResult.CompanyId, expectedResult.CompanyName, expectedResult.Address, expectedResult.Telephone, sellerList[0]).Return(seller.Seller{}, expectedError)

		service := seller.NewService(mockRepo)
		response, err := service.Update(id, 5, "Meli", "América do Sul", "5501154545454")

		assert.Equal(t, expectedResult, response)
		assert.Equal(t, expectedError, err)
	})
}

func TestService_GetOne(t *testing.T) {
	t.Run("Se o elemento procurado por id existir, ele retornará as informações do elemento solicitado", func(t *testing.T) {
		mockrepo := mocks.MockRepository{}

		sellerList := []seller.Seller{{Id: 1, CompanyId: 5, CompanyName: "TestGetOne", Address: "BR", Telephone: "5501154545454"},
			{Id: 3, CompanyId: 5, CompanyName: "ServiceSeller", Address: "BR", Telephone: "5501154545454"}}

		mockrepo.On("GetOne", 1).Return(sellerList[0], nil)

		service := seller.NewService(&mockrepo)
		response1, _ := service.GetOne(sellerList[0].Id)
		assert.Equal(t, sellerList[0], response1)

		mockrepo.On("GetOne", 3).Return(sellerList[1], nil)
		response2, _ := service.GetOne(3)
		assert.Equal(t, sellerList[1], response2)
	})

	t.Run("Se o elemento procurado por id não existir, retorna null", func(t *testing.T) {

		mockRepo := &mocks.MockRepository{}
		var id int = 2

		expectedError := fmt.Errorf("the id %d does not exists", id)

		mockRepo.On("GetOne", id).Return(seller.Seller{}, expectedError)

		service := seller.NewService(mockRepo)
		_, err := service.GetOne(id)

		assert.Equal(t, expectedError, err)
		assert.NotNil(t, err)
	})
}

func TestService_Create(t *testing.T) {
	t.Run("Se contiver os campos necessários, o vendedor será criado", func(t *testing.T) {
		mockRepo := &mocks.MockRepository{}

		expected := seller.Seller{Id: 1, CompanyId: 5, CompanyName: "TestCreate", Address: "BR", Telephone: "5501154545454"}
		input := seller.Seller{CompanyId: 5, CompanyName: "TestCreate", Address: "BR", Telephone: "5501154545454"}

		mockRepo.On("GetAll").Return([]seller.Seller{}, nil)
		mockRepo.On("Create", expected.CompanyId, expected.CompanyName, expected.Address, expected.Telephone).Return(expected, nil)

		service := seller.NewService(mockRepo)
		response, _ := service.Create(input.CompanyId, input.CompanyName, input.Address, input.Telephone)

		assert.Equal(t, expected, response)
	})

	t.Run("Se o cid já existir, o vendedor não pode ser criado", func(t *testing.T) {
		mockRepo := &mocks.MockRepository{}

		sellerList := []seller.Seller{{Id: 1, CompanyId: 5, CompanyName: "TestCreate", Address: "BR", Telephone: "5501154545454"}}
		input := seller.Seller{CompanyId: 5, CompanyName: "TestCreate", Address: "BR", Telephone: "5501154545454"}
		expectedError := errors.New("the cid already exists")

		mockRepo.On("GetAll").Return(sellerList, nil)
		mockRepo.On("Create", input.CompanyId, input.CompanyName, input.Address, input.Telephone).Return(expectedError)

		service := seller.NewService(mockRepo)
		response, err := service.Create(input.CompanyId, input.CompanyName, input.Address, input.Telephone)

		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, seller.Seller{}, response)
	})
}

func TestService_GetAll(t *testing.T) {
	t.Run("Se a lista tiver n elementos, retornará uma quantidade do total de elementos", func(t *testing.T) {

		mockRepository := mocks.NewRepository(t)
		expectedResult := []seller.Seller{
			{Id: 1, CompanyId: 5, CompanyName: "ServiceSeller", Address: "BR", Telephone: "5501154545454"},
			{Id: 2, CompanyId: 6, CompanyName: "ServiceSeller", Address: "BR", Telephone: "5501154545454"},
		}

		mockRepository.On("GetAll").Return(expectedResult, nil)

		service := seller.NewService(mockRepository)
		response, _ := service.GetAll()

		assert.Equal(t, 2, len(response))
		assert.Equal(t, expectedResult, response)
	})
}
