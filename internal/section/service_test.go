package section_test

import (
	"errors"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section/mocks"
	"github.com/stretchr/testify/assert"
)

func createSectionArray() []section.Section {
	var sec []section.Section
	sec1 := section.Section{
		ID:             1,
		SectionNumber:  40,
		CurTemperature: 90,
		MinTemperature: 10,
		CurCapacity:    20,
		MinCapacity:    10,
		MaxCapacity:    999,
		WareHouseID:    9876,
		ProductTypeID:  7659,
	}

	sec2 := section.Section{
		ID:             2,
		SectionNumber:  20,
		CurTemperature: 32,
		MinTemperature: 15,
		CurCapacity:    100,
		MinCapacity:    20,
		MaxCapacity:    500,
		WareHouseID:    9876,
		ProductTypeID:  3747,
	}

	sec = append(sec, sec1, sec2)
	return sec
}

func TestCreate(t *testing.T) {
	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)

	t.Run("create_ok", func(t *testing.T) {
		secs := createSectionArray()
		exp := secs[0]
		exp.ID = 3
		exp.SectionNumber = 50

		mockRepository.On("GetAll").Return(secs, nil)
		mockRepository.On("Create", exp.SectionNumber, exp.CurCapacity, exp.MinTemperature,
			exp.CurCapacity, exp.MinCapacity, exp.MaxCapacity, exp.WareHouseID, exp.ProductTypeID).Return(exp, nil)

		prod, err := service.Create(exp.SectionNumber, exp.CurCapacity, exp.MinTemperature,
			exp.CurCapacity, exp.MinCapacity, exp.MaxCapacity, exp.WareHouseID, exp.ProductTypeID)
		assert.Nil(t, err)
		assert.Equal(t, exp, prod)
	})

	t.Run("create_conflict", func(t *testing.T) {
		secs := createSectionArray()
		exp := secs[0]

		mockRepository.On("GetAll").Return(secs, nil)
		prod, err := service.Create(exp.SectionNumber, exp.CurCapacity, exp.MinTemperature,
			exp.CurCapacity, exp.MinCapacity, exp.MaxCapacity, exp.WareHouseID, exp.ProductTypeID)
		assert.Equal(t, section.Section{}, prod)
		assert.Equal(t, errors.New("seção com sectionNumber: 40 já existe no banco de dados"), err)
	})
}

func TestGetAll(t *testing.T) {
	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)
	exp := createSectionArray()

	t.Run("find_all", func(t *testing.T) {
		mockRepository.On("GetAll").Return(exp, nil)
		sections, _ := service.GetAll()
		assert.Equal(t, exp, sections)
	})
}

func TestGetByID(t *testing.T) {
	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)
	secs := createSectionArray()
	exp := secs[0]

	t.Run("find_by_id_non_existent", func(t *testing.T) {
		mockRepository.On("GetByID", 10).Return(section.Section{}, errors.New("seção 10 não encontrada"))
		sec, err := service.GetByID(10)
		assert.Equal(t, errors.New("seção 10 não encontrada"), err)
		assert.Equal(t, section.Section{}, sec)
	})

	t.Run("find_by_id_existent", func(t *testing.T) {
		mockRepository.On("GetByID", 1).Return(exp, nil)
		sec, err := service.GetByID(1)
		assert.Nil(t, err)
		assert.Equal(t, exp, sec)
	})
}

func TestUpdateSecID(t *testing.T) {
	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)
	secs := createSectionArray()
	exp := section.Section{
		ID:             2,
		SectionNumber:  572836456385,
		CurTemperature: 32,
		MinTemperature: 15,
		CurCapacity:    100,
		MinCapacity:    20,
		MaxCapacity:    500,
		WareHouseID:    9876,
		ProductTypeID:  3747,
	}

	mockRepository.On("GetAll").Return(secs, nil)

	t.Run("update_existent", func(t *testing.T) {
		mockRepository.On("UpdateSecID", 2, 572836456385).Return(exp, section.CodeError{})
		sec, err := service.UpdateSecID(2, 572836456385)
		assert.Equal(t, section.CodeError{}, err)
		assert.Equal(t, exp, sec)
	})

	t.Run("update_conflict", func(t *testing.T) {
		sec, err := service.UpdateSecID(2, 20)
		assert.Equal(t, section.CodeError{409, errors.New("seção com section_number: 20 já existe no banco de dados")}, err)
		assert.Equal(t, section.Section{}, sec)
	})

	t.Run("update_non_existent", func(t *testing.T) {
		mockRepository.On("UpdateSecID", 99, 99).Return(section.Section{},
			section.CodeError{404, errors.New("seção 99 não encontrada")})
		sec, err := service.UpdateSecID(99, 99)
		assert.Equal(t, section.CodeError{404, errors.New("seção 99 não encontrada")}, err)
		assert.Equal(t, section.Section{}, sec)
	})
}

func TestDelete(t *testing.T) {
	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)

	t.Run("delete_non_existent", func(t *testing.T) {
		mockRepository.On("DeleteSection", 99).Return(errors.New("seção 99 não encontrada"))
		err := service.DeleteSection(99)
		assert.NotNil(t, err)
	})

	t.Run("delete_ok", func(t *testing.T) {
		mockRepository.On("DeleteSection", 1).Return(nil)
		err := service.DeleteSection(1)
		assert.Nil(t, err)
	})
}
