package handlers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

type sectionRequest struct {
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
	return Section{p}
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

	if id < 0 {
		ctx.AbortWithStatusJSON(web.DecodeError(http.StatusNotFound, "id negativo inválido"))
		return
	}

	ctx.Next()
}

func (p *Section) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sec, _ := p.service.GetAll()
		c.JSON(web.NewResponse(http.StatusOK, sec))
	}
}

func (p *Section) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		sec, err := p.service.GetByID(id)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(web.NewResponse(http.StatusOK, sec))
	}
}

func (p *Section) CreateSection() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req sectionRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		sec, err := p.service.Create(req.SectionNumber, req.CurTemperature, req.MinTemperature,
			req.CurCapacity, req.MinCapacity, req.MaxCapacity, req.WareHouseID, req.ProductTypeID)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusCreated, sec))
	}
}

func (p *Section) UpdateSecID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req sectionRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		id, _ := strconv.Atoi(c.Param("id"))

		sec, err := p.service.UpdateSecID(id, req.SectionNumber)
		if err.Code != 200 {
			c.JSON(web.DecodeError(err.Code, err.Message.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusOK, sec))
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

		c.JSON(web.NewResponse(http.StatusNoContent, ""))
	}
}
