package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID             int `json:"id"`
	SectionNumber  int `json:"section_number" binding:"required"`
	CurTemperature int `json:"current_temperature" binding:"required"`
	MinTemperature int `json:"minimum_temperature" binding:"required"`
	CurCapacity    int `json:"current_capacity" binding:"required"`
	MinCapacity    int `json:"minimum_capacity" binding:"required"`
	MaxCapacity    int `json:"maximum_capacity" binding:"required"`
	WareHouseID    int `json:"warehouse_id" binding:"required"`
	ProductTypeID  int `json:"product_type_id" binding:"required"`
}

type Section struct {
	service section.Services
}

func NewSection(p section.Services) Section {
	new := Section{p}
	return new
}

func (p *Section) TokenAuthMiddleware(ctx *gin.Context) {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("Variavel de sistema TOKEN vazia")
	}

	token := ctx.GetHeader("token")
	if token == "" {
		ctx.AbortWithStatusJSON(web.DecodeError(http.StatusUnauthorized, "token vazio"))
		return
	}

	if token != requiredToken {
		ctx.AbortWithStatusJSON(web.DecodeError(http.StatusUnauthorized, "token inválido"))
		return
	}

	ctx.Next()
}

func (p *Section) IdVerificatorMiddleware(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(web.DecodeError(http.StatusBadRequest, "id não é alphanumérico"))
		return
	}

	if 0 > id || id > p.service.LastID() {
		ctx.AbortWithStatusJSON(web.DecodeError(http.StatusBadRequest, "id menor que 0 ou maior que o ultimo id"))
		return
	}

	ctx.Next()
}

func (p *Section) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		prod := p.service.GetAll()
		c.JSON(web.NewResponse(http.StatusOK, prod))
	}
}

func (p *Section) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		prod, err := p.service.GetByID(id)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(web.NewResponse(http.StatusOK, prod))
	}
}

func (p *Section) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		prod, err := p.service.Create(req.SectionNumber, req.CurTemperature, req.MinTemperature,
			req.CurCapacity, req.MinCapacity, req.MaxCapacity, req.WareHouseID, req.ProductTypeID)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusCreated, prod))
	}
}

func (p *Section) UpdateSecID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		id, _ := strconv.Atoi(c.Param("id"))
		prod, err := p.service.UpdateSecID(id, req.SectionNumber)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusOK, prod))
	}
}

func (p *Section) DeleteSection() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		err := p.service.DeleteSection(id)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}

		prod := fmt.Sprintf("O produto %d foi removido", id)
		c.JSON(web.NewResponse(http.StatusOK, prod))
	}
}
