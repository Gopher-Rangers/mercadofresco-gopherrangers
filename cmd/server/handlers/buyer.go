package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

type buyerRequest struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
}

type Buyer struct {
	service buyer.Service
}

func NewBuyerHandler() Buyer {
	return Buyer{buyer.NewService()}
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

// ListBuyers godoc
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

	data, err := b.service.GetAll()

	if err != nil {
		c.JSON(web.DecodeError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusOK, data))
}

// GetBuyer godoc
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

	data, err := b.service.GetById(id)

	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusOK, data))
}

// CreateBuyer godoc
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
	if err := c.Bind(&req); err != nil {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	newBuyer, err := b.service.Create(req.CardNumberId, req.FirstName, req.LastName)
	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusCreated, newBuyer))
}

// UpdateBuyers godoc
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
	var req buyerRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	newBuyer, err := b.service.Update(id, req.CardNumberId, req.FirstName, req.LastName)
	if err != nil {
		c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusCreated, newBuyer))
}

// DeleteBuyers godoc
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
	var req buyerRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	err := b.service.Delete(id)
	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	sec := fmt.Sprintf("Buyer with id %d deleted", id)
	c.JSON(web.NewResponse(http.StatusNoContent, sec))
}
