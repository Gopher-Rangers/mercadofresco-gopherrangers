package controller

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	ERROR_PURCHASE_ID_NOT_FOUNDED = "purchase order with id (%d) not founded"
)

type PurchaseOrdersCreate struct {
	ID              int    `json:"id"`
	OrderNumber     string `json:"order_number" binding:"required"`
	OrderDate       string `json:"order_date" binding:"required"`
	TrackingCode    string `json:"tracking_code" binding:"required"`
	BuyerId         int    `json:"buyer_id" binding:"required"`
	ProductRecordId int    `json:"product_record_id" binding:"required"`
	OrderStatusId   int    `json:"order_status_id" binding:"required"`
}

type PurchaseOrders struct {
	service domain.Service
}

func NewPurchaseOrder(r domain.Service) PurchaseOrders {
	return PurchaseOrders{r}
}

// Create CreatePurchaseOrder godoc
// @Summary Create PurchaseOrder
// @Tags Buyers
// @Description store a new purchase order
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param buyer body buyerRequest true "Purchase Order to store"
// @Failure 401 {object} web.Response "We need token"
// @Failure 404 {object} web.Response
// @Failure 409 {object} web.Response "Conflict"
// @Failure 422 {object} web.Response "Missing some mandatory field"
// @Success 201 {object} web.Response
// @Router /api/v1/purchase-orders [POST]
func (b *PurchaseOrders) Create(c *gin.Context) {

	var req PurchaseOrdersCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, "invalid body"))
		return
	}
	purchaseOrder := domain.PurchaseOrders{OrderNumber: req.OrderNumber, OrderDate: req.OrderDate, TrackingCode: req.TrackingCode,
		BuyerId: req.BuyerId, ProductRecordId: req.ProductRecordId, OrderStatusId: req.OrderStatusId}

	newPurchaseOrder, err := b.service.Create(c.Request.Context(), purchaseOrder)

	if err != nil {
		if err.Error() == domain.ERROR_UNIQUE_ORDER_NUMBER {
			c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
			return
		}
		if err.Error() == "Error 1062: Duplicate entry 'order229' for key 'purchase_orders.UNIQUE_ORDER_NUMBER'" {
			c.JSON(web.DecodeError(http.StatusConflict, domain.ERROR_UNIQUE_ORDER_NUMBER))
			return
		}
		c.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusCreated, newPurchaseOrder))
}

// GetPurchaseOrderById GetPurchaseOrder godoc
// @Summary List buyer
// @Tags Buyers
// @Description get a specific purchase order by id
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Failure 401 {object} web.Response "We need token"
// @Failure 404 {object} web.Response
// @Success 200 {object} web.Response
// @Router /api/v1/purchase-orders/{id} [GET]
func (b *PurchaseOrders) GetPurchaseOrderById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	data, err := b.service.GetById(c.Request.Context(), id)

	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusOK, data))
}
