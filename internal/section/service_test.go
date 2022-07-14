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

	t.Run("create_fail_getall", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := section.NewService(mockRepository)
		mockRepository.On("GetAll").Return([]section.Section{}, errors.New("rows not affected"))

		secs := createSectionArray()
		exp := secs[0]

		prod, err := service.Create(exp.SectionNumber, exp.CurCapacity, exp.MinTemperature,
			exp.CurCapacity, exp.MinCapacity, exp.MaxCapacity, exp.WareHouseID, exp.ProductTypeID)

		exp.SectionNumber = 50
		assert.Equal(t, section.Section{}, prod)
		assert.Equal(t, errors.New("internal server error"), err)
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

	t.Run("find_fail_getall", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := section.NewService(mockRepository)
		mockRepository.On("GetAll").Return([]section.Section{}, errors.New("rows not affected"))

		prod, err := service.GetAll()

		assert.Equal(t, []section.Section{}, prod)
		assert.Equal(t, errors.New("internal server error"), err)
	})
}

func TestGetByID(t *testing.T) {
	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)
	secs := createSectionArray()
	exp := secs[0]

	mockRepository.On("GetAll").Return(secs, nil)

	t.Run("find_by_id_non_existent", func(t *testing.T) {
		sec, err := service.GetByID(10)
		assert.Equal(t, errors.New("seção com id: 10 não existe no banco de dados"), err)
		assert.Equal(t, section.Section{}, sec)
	})

	t.Run("find_by_id_existent", func(t *testing.T) {
		mockRepository.On("GetByID", 1).Return(exp, nil)
		sec, err := service.GetByID(1)
		assert.Nil(t, err)
		assert.Equal(t, exp, sec)
	})

	t.Run("find_by_id_fail_getall", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := section.NewService(mockRepository)
		mockRepository.On("GetAll").Return([]section.Section{}, errors.New("rows not affected"))

		prod, err := service.GetByID(1)

		assert.Equal(t, section.Section{}, prod)
		assert.Equal(t, errors.New("internal server error"), err)
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
		mockRepository.On("UpdateSecID", 2, 572836456385).Return(exp, section.CodeError{200, nil})
		sec, err := service.UpdateSecID(2, 572836456385)
		assert.Equal(t, section.CodeError{200, nil}, err)
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

	t.Run("update_fail_getall", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := section.NewService(mockRepository)
		mockRepository.On("GetAll").Return([]section.Section{}, errors.New("rows not affected"))

		prod, err := service.UpdateSecID(2, 572836456385)

		assert.Equal(t, section.Section{}, prod)
		assert.Equal(t, section.CodeError{500, errors.New("internal server error")}, err)
	})
}

func TestDelete(t *testing.T) {
	mockRepository := mocks.NewRepository(t)
	service := section.NewService(mockRepository)

	secs := createSectionArray()
	mockRepository.On("GetAll").Return(secs, nil)

	t.Run("delete_non_existent", func(t *testing.T) {
		err := service.DeleteSection(99)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("seção com id: 99 não existe no banco de dados"), err)
	})

	t.Run("delete_ok", func(t *testing.T) {
		mockRepository.On("DeleteSection", 1).Return(nil)
		err := service.DeleteSection(1)
		assert.Nil(t, err)
	})

	t.Run("delete_fail_scan", func(t *testing.T) {
		mockRepository = mocks.NewRepository(t)
		service = section.NewService(mockRepository)
		mockRepository.On("GetAll").Return(secs, nil)
		mockRepository.On("DeleteSection", 1).Return(errors.New("rows not affected"))
		err := service.DeleteSection(1)

		assert.NotNil(t, err)
		assert.Equal(t, errors.New("rows not affected"), err)
	})

	t.Run("delete_fail_getall", func(t *testing.T) {
		mockRepository = mocks.NewRepository(t)
		service = section.NewService(mockRepository)
		mockRepository.On("GetAll").Return([]section.Section{}, errors.New("rows not affected"))

		err := service.DeleteSection(1)

		assert.NotNil(t, err)
		assert.Equal(t, errors.New("internal server error"), err)
	})
}
