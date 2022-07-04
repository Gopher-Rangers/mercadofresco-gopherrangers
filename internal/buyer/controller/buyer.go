package controller

import (
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/domain"
	"net/http"
	"os"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

const (
	ERROR_BUYER_CARD_NUMBER = "card_number_id is mandatory"
	ERROR_BUYER_FIRST_NAME  = "first_name is mandatory"
	ERROR_BUYER_LAST_NAME   = "last_name is mandatory"
)

type buyerRequest struct {
	ID           int    `json:"id"`
	CardNumberId string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
}

type buyerRequestUpdate struct {
	ID           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type Buyer struct {
	service domain.Service
}

func NewBuyer(s domain.Service) Buyer {
	return Buyer{s}
}

func (Buyer) AuthToken(context *gin.Context) {
	privateToken := os.Getenv("TOKEN")

	providedToken := context.GetHeader("token")

	if providedToken != privateToken {
		context.AbortWithStatusJSON(web.DecodeError(http.StatusUnauthorized, "invalid token"))
		return
	}

	context.Next()
}

func (Buyer) ValidateID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id < 0 {
		ctx.AbortWithStatusJSON(web.DecodeError(http.StatusBadRequest, "Id need to be a valid integer"))
		return
	}

	ctx.Next()
}

func (Buyer) validateBody(req domain.Buyer, c *gin.Context) bool {
	if req.CardNumberId == "" {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, ERROR_BUYER_CARD_NUMBER))
		return false
	}
	if req.FirstName == "" {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, ERROR_BUYER_FIRST_NAME))
		return false
	}
	if req.LastName == "" {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, ERROR_BUYER_LAST_NAME))
		return false
	}
	return true
}

// GetAll ListBuyers godoc
// @Summary List buyers
// @Tags Buyers
// @Description get all buyers
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Failure 401 {object} web.Response "We need token"
// @Failure 404 {object} web.Response
// @Success 200 {object} web.Response
// @Router /api/v1/buyers [GET]
func (b *Buyer) GetAll(c *gin.Context) {
	data, _ := b.service.GetAll(c.Request.Context())

	c.JSON(web.NewResponse(http.StatusOK, data))
}

// GetBuyerById GetBuyer godoc
// @Summary List buyer
// @Tags Buyers
// @Description get a especific buyer by id
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Failure 401 {object} web.Response "We need token"
// @Failure 404 {object} web.Response
// @Success 200 {object} web.Response
// @Router /api/v1/buyers/{id} [GET]
func (b *Buyer) GetBuyerById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	data, err := b.service.GetById(c.Request.Context(), id)

	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusOK, data))
}

// GetAllBuyerPurchaseOrdersById GetPurchaseOrders godoc
// @Summary List buyer
// @Tags Buyers
// @Description Get number of purchase Orders by an ID of a specific buyer
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Failure 401 {object} web.Response "We need token"
// @Failure 404 {object} web.Response
// @Success 200 {object} web.Response
// @Router /api/v1/buyers/{id} [GET]
func (b *Buyer) GetAllBuyerPurchaseOrdersById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	data, err := b.service.GetById(c.Request.Context(), id)

	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusOK, data))
}

// Create CreateBuyer godoc
// @Summary Create buyer
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
func (b *Buyer) Create(c *gin.Context) {

	var req buyerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "validation error",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}
	buyer := domain.Buyer{CardNumberId: req.CardNumberId, FirstName: req.FirstName, LastName: req.LastName}
	newBuyer, err := b.service.Create(c.Request.Context(), buyer)
	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusCreated, newBuyer))
}

// Update UpdateBuyers godoc
// @Summary Update buyer by ID
// @Tags Buyers
// @Description update buyer
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param some_id path int true "Some ID"
// @Param buyer body buyerRequest true "Buyer to update"
// @Failure 401 {object} web.Response "We need token"
// @Failure 400 {object} web.Response "We need ID"
// @Failure 404 {object} web.Response "Can not find ID"
// @Failure 422 {object} web.Response "Missing some mandatory field"
// @Success 200 {object} web.Response
// @Router /api/v1/buyers/{id} [PUT]
func (b *Buyer) Update(c *gin.Context) {
	var req buyerRequestUpdate
	c.Bind(&req)

	req.ID, _ = strconv.Atoi(c.Param("id"))

	if !b.validateBody(domain.Buyer{ID: req.ID, CardNumberId: req.CardNumberId, FirstName: req.FirstName, LastName: req.LastName}, c) {
		return
	}

	newBuyer, err := b.service.Update(c.Request.Context(), domain.Buyer(req))
	if err != nil {
		c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusOK, newBuyer))
}

// Delete DeleteBuyers godoc
// @Summary Delete buyers by ID
// @Tags Buyers
// @Description delete buyer by ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param some_id path int true "Some ID"
// @Failure 401 {object} web.Response "We need token"
// @Failure 400 {object} web.Response "We need ID"
// @Failure 404 {object} web.Response "Can not find ID"
// @Success 204 {object} web.Response
// @Router /api/v1/buyers/{id} [DELETE]
func (b *Buyer) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := b.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	sec := fmt.Sprintf("Buyer with id %d deleted", id)
	c.JSON(web.NewResponse(http.StatusNoContent, sec))
}
