package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/internal/produto"
	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name        string  `json:"name"`
	ProductType string  `json:"type"`
	Count       int     `json:"count"`
	Price       float64 `json:"price"`
}

type Product struct {
	service produto.Services
}

func NewProduct(p produto.Services) Product {
	new := Product{p}
	return new
}

func (p *Product) TokenAuthMiddleware(ctx *gin.Context) {
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

func (p *Product) IdVerificatorMiddleware(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(web.DecodeError(http.StatusBadRequest, "id não é alphanumérico"))
		return
	}

	if 0 > id || id > p.service.LastID() {
		ctx.AbortWithStatusJSON(web.DecodeError(http.StatusBadRequest, "id fora do limite"))
		return
	}

	ctx.Next()
}

func (p *Product) GetAll(ctx *gin.Context) {
	prod, err := p.service.GetAll()
	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(web.NewResponse(http.StatusOK, prod))
}

func (p *Product) CreateProduct(ctx *gin.Context) {
	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
		return
	}

	if req.Name == "" {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'nome' é obrigatório"))
		return
	}

	if req.ProductType == "" {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'type' é obrigatório"))
		return
	}

	if req.Count <= 0 {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'count' não pode ser menor que 1"))
		return
	}

	if req.Price < 0 {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'price' não pode ser menor que 0"))
		return
	}

	prod, err := p.service.Create(req.Name, req.ProductType, req.Count, req.Price)
	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	ctx.JSON(web.NewResponse(http.StatusOK, prod))
}

func (p *Product) Update(ctx *gin.Context) {
	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
		return
	}

	if req.Name == "" {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'nome' é obrigatório"))
		return
	}

	if req.ProductType == "" {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'type' é obrigatório"))
		return
	}

	if req.Count <= 0 {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'count' não pode ser menor que 1"))
		return
	}

	if req.Price < 0 {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'price' não pode ser menor que 0"))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	prod, err := p.service.Update(id, req.Name, req.ProductType, req.Count, req.Price)
	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	ctx.JSON(web.NewResponse(http.StatusOK, prod))
}

func (p *Product) UpdateName(ctx *gin.Context) {
	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
		return
	}

	if req.Name == "" {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'nome' é obrigatório"))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	prod, err := p.service.UpdateName(id, req.Name)
	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	ctx.JSON(web.NewResponse(http.StatusOK, prod))
}

func (p *Product) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := p.service.DeleteProduct(id)
	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	prod := fmt.Sprintf("O produto %d foi removido", id)
	ctx.JSON(web.NewResponse(http.StatusOK, prod))
}
