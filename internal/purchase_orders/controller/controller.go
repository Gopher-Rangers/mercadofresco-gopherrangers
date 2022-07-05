package controller

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PurchaseOrdersCreate struct {
	ID              int    `json:"id"`
	OrderNumber     string `json:"order_number"`
	OrderDate       string `json:"order_date"`
	TrackingCode    string `json:"tracking_code"`
	BuyerId         int    `json:"buyer_id"`
	ProductRecordId int    `json:"product_record_id"`
	OrderStatusId   int    `json:"order_status_id"`
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
// @Description store a new buyer
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param buyer body buyerRequest true "Product to store"
// @Failure 401 {object} web.Response "We need token"
// @Failure 404 {object} web.Response
// @Failure 422 {object} web.Response "Missing some mandatory field"
// @Success 201 {object} web.Response
// @Router /api/v1/buyers [POST]
func (b *PurchaseOrders) Create(c *gin.Context) {

	var req PurchaseOrdersCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "validation error",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}
	purchaseOrder := domain.PurchaseOrders{OrderNumber: req.OrderNumber, OrderDate: req.OrderDate, TrackingCode: req.TrackingCode,
		BuyerId: req.BuyerId, ProductRecordId: req.ProductRecordId, OrderStatusId: req.OrderStatusId}
	newPurchaseOrder, err := b.service.Create(c.Request.Context(), purchaseOrder)
	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusCreated, newPurchaseOrder))
}

// GetPurchaseOrderById GetPurchaseOrder godoc
// @Summary List buyer
// @Tags Buyers
// @Description get a especific purchase order by id
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Failure 401 {object} web.Response "We need token"
// @Failure 404 {object} web.Response
// @Success 200 {object} web.Response
// @Router /api/v1/buyers/{id} [GET]
func (b *PurchaseOrders) GetPurchaseOrderById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	data, err := b.service.GetById(c.Request.Context(), id)

	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusOK, data))
}
