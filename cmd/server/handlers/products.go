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
	ERROR_TOKEN = "token inválido"
	ERROR_ID = "id inválido"
)

type requestProduct struct {
	ID int `json:"id"`
	ProductCode string `json:"product_code"`
	Description string `json:"description"`
	Width float64 `json:"width"`
	Height float64 `json:"height"`
	Length float64 `json:"length"`
	NetWeight float64 `json:"net_weight"`
	ExpirationRate string `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate float64 `json:"freezing_rate"`
	ProductTypeId int `json:"product_type_id"`
	SellerId int `json:"seller_id"`
}

type requestDescription struct {
	Description string `json:"description"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{service: p}
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
		p, err := prod.service.Store(req)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(web.NewResponse(http.StatusOK, p))
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

func (prod *Product) UpdatePut() gin.HandlerFunc {
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
		if req.ProductCode == "" {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_PRODUCT_CODE))
			return
		}
		if req.Description == "" {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_DESCRIPTION))
			return
		}
		if req.Width == 0 {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_WIDTH))
			return
		}
		if req.Height == 0 {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_HEIGHT))
			return
		}
		if req.Length == 0 {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_LENGTH))
			return
		}
		if req.NetWeight == 0 {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_NET_WEIGHT))
			return
		}
		if req.ExpirationRate == "" {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_EXPIRATIONN_RATE))
			return
		}
		if req.RecommendedFreezingTemperature == 0 {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_RECOM_FREEZING_TEMP))
			return
		}
		if req.FreezingRate == 0 {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_FREEZING_RATE))
			return
		}
		if req.ProductTypeId == 0 {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_PRODUCT_TYPE_ID))
			return
		}
		p, err := prod.service.UpdatePut(req, int(id))
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
		c.JSON(web.NewResponse(http.StatusOK, p))
	}
	return fn
}
