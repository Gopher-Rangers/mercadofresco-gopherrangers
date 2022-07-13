package handlers

import (
	"net/http"

	inboundorders "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/inbound_orders"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

type inboundOrderRequest struct {
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeId     int    `json:"employee_id"`
	ProductBatchId int    `json:"product_batch_id"`
	WarehouseId    int    `json:"warehouse_id"`
}

type InboundOrder struct {
	service inboundorders.Services
}

func NewInboundOrder(io inboundorders.Services) InboundOrder {
	return InboundOrder{io}
}

func (io *InboundOrder) checkBody(req inboundOrderRequest, c *gin.Context) bool {
	if req.OrderDate == "" || req.OrderNumber == "" {
		c.JSON(web.DecodeError(
			http.StatusUnprocessableEntity,
			ERROR_ALLMANDATORY))
		return false
	}
	if req.WarehouseId == 0 || req.EmployeeId == 0 || req.ProductBatchId == 0 {
		c.JSON(web.DecodeError(
			http.StatusUnprocessableEntity,
			ERROR_ALLMANDATORY))
		return false
	}

	return true
}

func (io *InboundOrder) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req inboundOrderRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}
		if !io.checkBody(req, c) {
			return
		}

		inboundOrder, err := io.service.Create(req.OrderDate, req.OrderNumber, req.EmployeeId, req.ProductBatchId, req.WarehouseId)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusCreated, inboundOrder))
	}
}
