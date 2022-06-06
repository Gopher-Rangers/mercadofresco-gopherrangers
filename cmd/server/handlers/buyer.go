package handler

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

func (p *Buyer) AuthToken(context *gin.Context) {
	privateToken := os.Getenv("TOKEN")

	providedToken := context.GetHeader("token")

	if providedToken != privateToken {
		context.AbortWithStatusJSON(web.DecodeError(http.StatusUnauthorized, "invalid token"))
		return
	}

	context.Next()
}

func (p *Buyer) ValidateID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id < 0 {
		ctx.AbortWithStatusJSON(web.DecodeError(http.StatusBadRequest, "Id need to be a valid integer"))
		return
	}

	ctx.Next()
}

func (b *Buyer) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		data, err := b.service.GetAll()

		if err != nil {
			c.JSON(web.DecodeError(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusOK, data))
	}
}

func (b *Buyer) GetBuyerById() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, _ := strconv.Atoi(c.Param("id"))

		data, err := b.service.GetById(id)

		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusOK, data))
	}
}

func (p *Buyer) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req buyerRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		newBuyer, err := p.service.Create(req.CardNumberId, req.FirstName, req.LastName)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusCreated, newBuyer))
	}
}

func (p *Buyer) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req buyerRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		id, _ := strconv.Atoi(c.Param("id"))

		newBuyer, err := p.service.Update(id, req.CardNumberId, req.FirstName, req.LastName)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusCreated, newBuyer))
	}
}

func (p *Buyer) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req buyerRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		id, _ := strconv.Atoi(c.Param("id"))

		err := p.service.Delete(id)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}

		sec := fmt.Sprintf("Buyer with id %d deleted", id)
		c.JSON(web.NewResponse(http.StatusNoContent, sec))
	}
}
