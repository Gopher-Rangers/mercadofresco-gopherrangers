package seller_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality"
	localityMock "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality/mocks"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Delete(t *testing.T) {
	t.Run("Se a exclusão for bem-sucedida, o item não aparecerá na lista.", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		var id int = 1

		sellerList := []seller.Seller{{Id: 1, CompanyId: 5, CompanyName: "TestDelete", Address: "BR", Telephone: "5501154545454"},
			{Id: 3, CompanyId: 6, CompanyName: "TestDelete", Address: "BR", Telephone: "5501154545454"},
			{Id: 4, CompanyId: 6, CompanyName: "TestDelete", Address: "BR", Telephone: "5501154545454"}}

		mockRepo.On("GetOne", context.Background(), id).Return(sellerList[0], nil)
		mockRepo.On("Delete", context.Background(), 1).Return(nil)

		service := seller.NewService(mockRepo, mockLocalityRepo)
		err := service.Delete(context.Background(), id)

		assert.Nil(t, err)
	})

	t.Run("Se o elemento a ser removido não existir, retornará null.", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		var id int = 2

		expectedError := fmt.Errorf("the id %d does not exists", id)

		mockRepo.On("GetOne", context.Background(), id).Return(seller.Seller{}, expectedError)

		service := seller.NewService(mockRepo, mockLocalityRepo)
		err := service.Delete(context.Background(), id)

		assert.Equal(t, expectedError, err)
	})

	t.Run("Deve retornar erro ao chamar método delete.", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		var id int = 1

		sellerList := []seller.Seller{{Id: 1, CompanyId: 5, CompanyName: "TestDelete", Address: "BR", Telephone: "5501154545454"},
			{Id: 3, CompanyId: 6, CompanyName: "TestDelete", Address: "BR", Telephone: "5501154545454"},
			{Id: 4, CompanyId: 6, CompanyName: "TestDelete", Address: "BR", Telephone: "5501154545454"}}

		mockRepo.On("GetOne", context.Background(), id).Return(sellerList[0], nil)
		mockRepo.On("Delete", context.Background(), 1).Return(fmt.Errorf("error"))

		service := seller.NewService(mockRepo, mockLocalityRepo)
		err := service.Delete(context.Background(), id)

		assert.Error(t, err)
	})
}

func TestService_Update(t *testing.T) {
	t.Run("Se os campos forem atualizados com sucesso retornará a informação do elemento atualizado", func(t *testing.T) {

		mockRepo := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		sellerList := []seller.Seller{{Id: 1, CompanyId: 5, CompanyName: "TestUpdate", Address: "BR", Telephone: "5501154545454"},
			{Id: 3, CompanyId: 6, CompanyName: "ServiceSeller", Address: "BR", Telephone: "5501154545454"}}

		localityOne := locality.Locality{Id: 1, LocalityName: "Cecap", ProvinceName: "Gru", CountryName: "SP"}

		expectedResult := seller.Seller{Id: 1, CompanyId: 7, CompanyName: "Meli", Address: "América do Sul", Telephone: "5501154545454", LocalityID: 1}

		mockLocalityRepo.On("GetById", context.Background(), 1).Return(localityOne, nil)
		mockRepo.On("GetOne", context.Background(), 1).Return(sellerList[0], nil)
		mockRepo.On("GetAll", context.Background()).Return(sellerList, nil)
		mockRepo.On("Update", context.Background(), expectedResult.CompanyId, expectedResult.CompanyName, expectedResult.Address,
			expectedResult.Telephone, expectedResult.LocalityID, sellerList[0]).Return(expectedResult, nil)

		service := seller.NewService(mockRepo, mockLocalityRepo)
		response, _ := service.Update(context.Background(), 1, 7, "Meli", "América do Sul", "5501154545454", 1)

		assert.Equal(t, expectedResult, response)
	})

	t.Run("Se o elemento a ser atualizado não existir, retornar null", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		var id int = 2

		expectedResult := seller.Seller{}
		expectedError := fmt.Errorf("the id %d does not exists", id)

		mockRepo.On("GetOne", context.Background(), id).Return(seller.Seller{}, expectedError)

		service := seller.NewService(mockRepo, mockLocalityRepo)
		response, err := service.Update(context.Background(), id, 5, "Meli", "América do Sul", "5501154545454", 1)

		assert.Equal(t, expectedResult, response)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Se o cid já existir, o elemento não poderá ser atualizador.", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		sellerList := []seller.Seller{{Id: 1, CompanyId: 5, CompanyName: "TestUpdate", Address: "BR", Telephone: "5501154545454", LocalityID: 1},
			{Id: 3, CompanyId: 6, CompanyName: "ServiceSeller", Address: "BR", Telephone: "5501154545454", LocalityID: 1}}
		localityOne := locality.Locality{Id: 1, LocalityName: "Cecap", ProvinceName: "Gru", CountryName: "SP"}

		expectedError := errors.New("the cid already exists")

		mockLocalityRepo.On("GetById", context.Background(), 1).Return(localityOne, nil)
		mockRepo.On("GetOne", context.Background(), 1).Return(sellerList[0], nil)
		mockRepo.On("GetAll", context.Background()).Return(sellerList, nil)

		service := seller.NewService(mockRepo, mockLocalityRepo)
		_, err := service.Update(context.Background(), 1, 5, "Meli", "América do Sul", "5501154545454", 1)

		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Deve retornar erro ao chamar método Update", func(t *testing.T) {

		mockRepo := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		sellerList := []seller.Seller{{Id: 1, CompanyId: 5, CompanyName: "TestUpdate", Address: "BR", Telephone: "5501154545454", LocalityID: 1},
			{Id: 3, CompanyId: 6, CompanyName: "ServiceSeller", Address: "BR", Telephone: "5501154545454", LocalityID: 1}}

		expectedResult := seller.Seller{Id: 1, CompanyId: 7, CompanyName: "Meli", Address: "América do Sul", Telephone: "5501154545454", LocalityID: 1}
		localityOne := locality.Locality{Id: 1, LocalityName: "Cecap", ProvinceName: "Gru", CountryName: "SP"}

		mockLocalityRepo.On("GetById", context.Background(), 1).Return(localityOne, nil)
		mockRepo.On("GetOne", context.Background(), 1).Return(sellerList[0], nil)
		mockRepo.On("GetAll", context.Background()).Return(sellerList, nil)
		mockRepo.On("Update", context.Background(), expectedResult.CompanyId, expectedResult.CompanyName, expectedResult.Address, expectedResult.Telephone, expectedResult.LocalityID, sellerList[0]).
			Return(seller.Seller{}, fmt.Errorf("error"))

		service := seller.NewService(mockRepo, mockLocalityRepo)
		_, err := service.Update(context.Background(), 1, 7, "Meli", "América do Sul", "5501154545454", 1)

		assert.Error(t, err)
	})
}

func TestService_GetOne(t *testing.T) {
	t.Run("Se o elemento procurado por id existir, ele retornará as informações do elemento solicitado", func(t *testing.T) {
		mockrepo := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		sellerList := []seller.Seller{{Id: 1, CompanyId: 5, CompanyName: "TestGetOne", Address: "BR", Telephone: "5501154545454", LocalityID: 1},
			{Id: 3, CompanyId: 5, CompanyName: "ServiceSeller", Address: "BR", Telephone: "5501154545454", LocalityID: 1}}

		mockrepo.On("GetOne", context.Background(), 1).Return(sellerList[0], nil)

		service := seller.NewService(mockrepo, mockLocalityRepo)
		response1, _ := service.GetOne(context.Background(), sellerList[0].Id)
		assert.Equal(t, sellerList[0], response1)
	})

	t.Run("Se o elemento procurado por id não existir, retorna null", func(t *testing.T) {

		mockRepo := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		var id int = 2

		expectedError := fmt.Errorf("the id %d does not exists", id)

		mockRepo.On("GetOne", context.Background(), id).Return(seller.Seller{}, expectedError)

		service := seller.NewService(mockRepo, mockLocalityRepo)
		_, err := service.GetOne(context.Background(), id)

		assert.Equal(t, expectedError, err)
		assert.NotNil(t, err)
	})
}

func TestService_Create(t *testing.T) {
	t.Run("Se contiver os campos necessários, o vendedor será criado", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		expected := seller.Seller{Id: 1, CompanyId: 5, CompanyName: "TestCreate", Address: "BR", Telephone: "5501154545454", LocalityID: 1}
		input := seller.Seller{CompanyId: 5, CompanyName: "TestCreate", Address: "BR", Telephone: "5501154545454", LocalityID: 1}
		localityOne := locality.Locality{Id: 1, LocalityName: "Cecap", ProvinceName: "Gru", CountryName: "SP"}

		mockLocalityRepo.On("GetById", context.Background(), 1).Return(localityOne, nil)
		mockRepo.On("GetAll", context.Background()).Return([]seller.Seller{}, nil)
		mockRepo.On("Create", context.Background(), expected.CompanyId, expected.CompanyName, expected.Address, expected.Telephone, localityOne.Id).Return(expected, nil)

		service := seller.NewService(mockRepo, mockLocalityRepo)
		response, _ := service.Create(context.Background(), input.CompanyId, input.CompanyName, input.Address, input.Telephone, localityOne.Id)

		assert.Equal(t, expected, response)
	})

	t.Run("Se o cid já existir, o vendedor não pode ser criado", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		sellerList := []seller.Seller{
			{Id: 1, CompanyId: 5, CompanyName: "ServiceSeller", Address: "BR", Telephone: "5501154545454", LocalityID: 1},
			{Id: 2, CompanyId: 6, CompanyName: "ServiceSeller", Address: "BR", Telephone: "5501154545454", LocalityID: 1},
		}

		localityOne := locality.Locality{Id: 1, LocalityName: "Cecap", ProvinceName: "Gru", CountryName: "SP"}
		input := seller.Seller{CompanyId: 5, CompanyName: "TestCreate", Address: "BR", Telephone: "5501154545454", LocalityID: 1}
		expectedError := errors.New("the cid already exists")

		mockLocalityRepo.On("GetById", context.Background(), 1).Return(localityOne, nil)
		mockRepo.On("GetAll", context.Background()).Return(sellerList, nil)

		service := seller.NewService(mockRepo, mockLocalityRepo)
		_, err := service.Create(context.Background(), input.CompanyId, input.CompanyName, input.Address, input.Telephone, input.LocalityID)

		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Deve retornar erro ao chamar método create", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		input := seller.Seller{CompanyId: 7, CompanyName: "TestCreate", Address: "BR", Telephone: "5501154545454", LocalityID: 1}
		localityOne := locality.Locality{Id: 1, LocalityName: "Cecap", ProvinceName: "Gru", CountryName: "SP"}

		mockLocalityRepo.On("GetById", context.Background(), 1).Return(localityOne, nil)
		mockRepo.On("GetAll", context.Background()).Return([]seller.Seller{}, nil)
		mockRepo.On("Create", context.Background(), input.CompanyId, input.CompanyName, input.Address, input.Telephone, input.LocalityID).
			Return(seller.Seller{}, fmt.Errorf("error"))

		service := seller.NewService(mockRepo, mockLocalityRepo)
		_, err := service.Create(context.Background(), input.CompanyId, input.CompanyName, input.Address, input.Telephone, localityOne.Id)

		assert.NotNil(t, err)
	})

	t.Run("Deve retornar erro ao executar findByCid", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		input := seller.Seller{CompanyId: 5, CompanyName: "TestCreate", Address: "BR", Telephone: "5501154545454", LocalityID: 1}

		mockLocalityRepo.On("GetById", context.Background(), 1).Return(locality.Locality{}, nil)
		mockRepo.On("GetAll", context.Background()).Return([]seller.Seller{}, fmt.Errorf("error"))

		service := seller.NewService(mockRepo, mockLocalityRepo)
		_, err := service.Create(context.Background(), input.CompanyId, input.CompanyName, input.Address, input.Telephone, input.LocalityID)

		assert.Error(t, err)
	})
}

func TestService_GetAll(t *testing.T) {
	t.Run("Se a lista tiver n elementos, retornará uma quantidade do total de elementos", func(t *testing.T) {

		mockRepository := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		expectedResult := []seller.Seller{
			{Id: 1, CompanyId: 5, CompanyName: "ServiceSeller", Address: "BR", Telephone: "5501154545454", LocalityID: 1},
			{Id: 2, CompanyId: 6, CompanyName: "ServiceSeller", Address: "BR", Telephone: "5501154545454", LocalityID: 2},
		}

		mockRepository.On("GetAll", context.Background()).Return(expectedResult, nil)

		service := seller.NewService(mockRepository, mockLocalityRepo)
		response, _ := service.GetAll(context.Background())

		assert.Equal(t, 2, len(response))
		assert.Equal(t, expectedResult, response)
	})

	t.Run("Deve retornar erro", func(t *testing.T) {

		mockRepository := mocks.NewRepository(t)
		mockLocalityRepo := localityMock.NewRepository(t)

		expectedError := errors.New("erro ao inicializar a lista")

		mockRepository.On("GetAll", context.Background()).Return([]seller.Seller{}, expectedError)

		service := seller.NewService(mockRepository, mockLocalityRepo)
		_, err := service.GetAll(context.Background())

		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
}
