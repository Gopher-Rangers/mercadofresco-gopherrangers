package controller

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

	var req domain.PurchaseOrders
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
