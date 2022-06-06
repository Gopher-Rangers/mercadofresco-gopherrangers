package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

const (
	ERROR_PRODUCT_CODE = "product_code is mandatory"
	ERROR_DESCRIPTION = "description is mandatory"
	ERROR_WIDTH = "Width is mandatory"
	ERROR_HEIGHT = "height is mandatory"
	ERROR_LENGTH = "length is mandatory"
	ERROR_NET_WEIGHT = "net_weight is mandatory"
	ERROR_EXPIRATIONN_RATE = "expiration_rate is mandatory"
	ERROR_RECOM_FREEZING_TEMP = "recommended_freezing_temperature is mandatory"
	ERROR_FREEZING_RATE = "freezing_rate is mandatory"
	ERROR_PRODUCT_TYPE_ID = "product_type_id is mandatory"
	ERROR_TOKEN = "ivalid token"
	ERROR_ID = "invalid id"
	ERROR_UNIQUE_PRODUCT_CODE = "the product code must be unique"
)

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{service: p}
}

func (prod *Product) checkBody(req products.Product, c *gin.Context) bool {
	ps, _ := prod.service.GetAll()
	for i :=range ps {
		if ps[i].ProductCode == req.ProductCode {
			c.JSON(web.DecodeError(
				http.StatusUnprocessableEntity,
				ERROR_UNIQUE_PRODUCT_CODE))
			return false
		}
	}
	if req.ProductCode == "" {
		c.JSON(web.DecodeError(
			http.StatusUnprocessableEntity,
			ERROR_PRODUCT_CODE))
		return false
	}
	if req.Description == "" {
		c.JSON(web.DecodeError(
			http.StatusUnprocessableEntity,
			ERROR_DESCRIPTION))
		return false
	}
	if req.Width == 0 {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, ERROR_WIDTH))
		return false
	}
	if req.Height == 0 {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, ERROR_HEIGHT))
		return false
	}
	if req.Length == 0 {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, ERROR_LENGTH))
		return false
	}
	if req.NetWeight == 0 {
		c.JSON(web.DecodeError(
			http.StatusUnprocessableEntity,
			ERROR_NET_WEIGHT))
		return false
	}
	if req.ExpirationRate == "" {
		c.JSON(web.DecodeError(
			http.StatusUnprocessableEntity,
			ERROR_EXPIRATIONN_RATE))
		return false
	}
	if req.RecommendedFreezingTemperature == 0 {
		c.JSON(web.DecodeError(
			http.StatusUnprocessableEntity,
			ERROR_RECOM_FREEZING_TEMP))
		return false
	}
	if req.FreezingRate == 0 {
		c.JSON(web.DecodeError(
			http.StatusUnprocessableEntity,
			ERROR_FREEZING_RATE))
		return false
	}
	if req.ProductTypeId == 0 {
		c.JSON(web.DecodeError(
			http.StatusUnprocessableEntity,
			ERROR_PRODUCT_TYPE_ID))
		return false
	}
	return true
}

func (prod *Product) Store() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(web.DecodeError(http.StatusUnauthorized, ERROR_TOKEN))
			return
		}
		var req products.Product
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		if !prod.checkBody(req, c) {
			return
		}
		p, err := prod.service.Store(req)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(web.NewResponse(http.StatusCreated, p))
	}
	return fn
}

func (prod *Product) GetAll() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN"){
			c.JSON(web.DecodeError(http.StatusUnauthorized, ERROR_TOKEN))
			return
		}
		p, err := prod.service.GetAll()
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(web.NewResponse(http.StatusOK, p))
	}
	return fn
}

func (prod *Product) GetById() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN"){
			c.JSON(web.DecodeError(http.StatusUnauthorized, ERROR_TOKEN))
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_ID))
			return
		}
		p, err := prod.service.GetById(id)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(web.NewResponse(http.StatusOK, p))
	}
	return fn
}

func (prod *Product) Update() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(web.DecodeError(http.StatusUnauthorized, ERROR_TOKEN))
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_ID))
			return
		}
		var req products.Product
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
			return
		}
		req.ID = id
		if !prod.checkBody(req, c) {
			return
		}
		p, err := prod.service.Update(req, int(id))
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(web.NewResponse(http.StatusOK, p))
	}
	return fn
}

func (prod *Product) Delete() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(web.DecodeError(http.StatusUnauthorized, ERROR_TOKEN))
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_ID))
			return
		}
		err = prod.service.Delete(int(id))
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		p := fmt.Sprintf("o produto %d foi removido", id)
		c.JSON(web.NewResponse(http.StatusNoContent, p))
	}
	return fn
}
