package inboundorders_test

import (
	"fmt"
	"testing"

	inboundorders "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/inbound_orders"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/inbound_orders/mocks"
	"github.com/stretchr/testify/assert"
)

func createInboundOrdersArray() []inboundorders.InboundOrder {
	var ios []inboundorders.InboundOrder
	inboundorder1 := inboundorders.InboundOrder{
		ID:             1,
		OrderDate:      "2022-04-04",
		OrderNumber:    "order#1",
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}
	inboundorder2 := inboundorders.InboundOrder{
		ID:             1,
		OrderDate:      "2022-04-04",
		OrderNumber:    "order#1",
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}
	ios = append(ios, inboundorder1, inboundorder2)
	return ios
}

func TestCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := inboundorders.NewService(mockRepository)
		expected := inboundorders.InboundOrder{
			OrderDate:      "2022-04-04",
			OrderNumber:    "order#1",
			EmployeeId:     1,
			ProductBatchId: 1,
			WarehouseId:    1,
		}

		mockRepository.On("Create", expected.OrderDate, expected.OrderNumber, expected.EmployeeId,
			expected.ProductBatchId, expected.WarehouseId).Return(expected, nil)
		io, err := service.Create(expected.OrderDate, expected.OrderNumber, expected.EmployeeId, expected.ProductBatchId, expected.WarehouseId)
		assert.Nil(t, err)
		assert.Equal(t, expected, io)
	})
	t.Run("create_error", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := inboundorders.NewService(mockRepository)

		employeeWrong := inboundorders.InboundOrder{
			OrderDate:      "2022-04-04",
			OrderNumber:    "order#1",
			EmployeeId:     100,
			ProductBatchId: 1,
			WarehouseId:    1,
		}

		mockRepository.On("Create", employeeWrong.OrderDate, employeeWrong.OrderNumber, employeeWrong.EmployeeId, employeeWrong.ProductBatchId, employeeWrong.WarehouseId).Return(employeeWrong, fmt.Errorf("funcionario nao existe"))
		_, err := service.Create(employeeWrong.OrderDate, employeeWrong.OrderNumber, employeeWrong.EmployeeId, employeeWrong.ProductBatchId, employeeWrong.WarehouseId)
		assert.Equal(t, err, fmt.Errorf("funcionario nao existe"))
	})
}

func TestGetCounterByEmployee(t *testing.T) {
	t.Run("get_counter_ok", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		service := inboundorders.NewService(mockRepository)
		expected := 2

		mockRepository.On("GetCountByEmployee", 2).Return(2, nil)
		counter := service.GetCounterByEmployee(2)

		assert.Equal(t, expected, counter)
	})
}
