package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	ERROR_PRODUCT_CODE = "ProductCode is mandatory"
	ERROR_DESCRIPTION = "Description is mandatory"
	ERROR_WIDTH = "Width is mandatory"
	ERROR_HEIGHT = "Height is mandatory"
	ERROR_LENGTH = "Length is mandatory"
	ERROR_NET_WEIGHT  = "NetWeight is mandatory"
	ERROR_EXPIRATIONN_RATE = "ExpirationRate is mandatory"
	ERROR_RECOM_FREEZING_TEMP = "RecommendedFreezingTemperature is mandatory"
	ERROR_FREEZING_RATE = "FreezingRate is mandatory"
	ERROR_PRODUCT_TYPE_ID = "ProductTypeId is mandatory"
	ERROR_TOKEN = "invalid token"
	ERROR_ID = "invalid id"
	ERROR_UNIQUE_PRODUCT_CODE = "the product code must be unique"
)

type requestProduct struct {
	ProductCode string `json:"product_code"`
	Description string `json:"description"`
	Width float64 `json:"width"`
	Height float64 `json:"height"`
	Length  float64 `json:"length"`
	NetWeight float64 `json:"net_weight"`
	ExpirationRate string  `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate float64 `json:"freezing_rate"`
	ProductTypeId int `json:"product_type_id"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{service: p}
}

func NewRequestProduct() requestProduct {
	p := requestProduct{}
	return p
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body requestProduct true "Product to store"
// @Failure 401 {object} web.Response "We need token"
// @Failure 404 {object} web.Response
// @Failure 422 {object} web.Response "Missing some mandatory field"
// @Success 201 {object} web.Response
// @Router /api/v1/products [POST]
func (prod *Product) Store() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(web.DecodeError(http.StatusUnauthorized, ERROR_TOKEN))
			return
		}
		var validate *validator.Validate = validator.New()
		var req products.Product
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		errValidate := validate.Struct(req)
		if errValidate != nil {
			if _, ok := errValidate.(*validator.InvalidValidationError); ok {
				c.JSON(web.DecodeError(http.StatusNotFound, errValidate.Error()))
				return
			}
			for _, errValidate := range errValidate.(validator.ValidationErrors) {
				if errValidate != nil {
					s := fmt.Sprintf("%s is mandatory", errValidate.Field())
					c.JSON(web.DecodeError(http.StatusUnprocessableEntity, s))
					return
				}
			}
		}
		p, err := prod.service.Store(req)
		if err != nil {
			if err.Error() == ERROR_UNIQUE_PRODUCT_CODE {
				c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
				return
			}
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(web.NewResponse(http.StatusCreated, p))
	}
	return fn
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Failure 401 {object} web.Response "We need token"
// @Failure 404 {object} web.Response
// @Success 200 {object} web.Response
// @Router /api/v1/products [GET]
func (prod *Product) GetAll() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(web.DecodeError(http.StatusUnauthorized, ERROR_TOKEN))
			return
		}
		p, _ := prod.service.GetAll()
		c.JSON(web.NewResponse(http.StatusOK, p))
	}
	return fn
}

// ListProductsById godoc
// @Summary List products by ID
// @Tags Products
// @Description list products by ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param some_id path int true "Some ID"
// @Failure 401 {object} web.Response "We need token"
// @Failure 400 {object} web.Response "We need ID"
// @Failure 404 {object} web.Response "Can not find ID"
// @Success 200 {object} web.Response
// @Router /api/v1/products/{some_id} [GET]
func (prod *Product) GetById() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
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

// UpdateProducts godoc
// @Summary Update products by ID
// @Tags Products
// @Description update products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param some_id path int true "Some ID"
// @Param product body requestProduct true "Product to update"
// @Failure 401 {object} web.Response "We need token"
// @Failure 400 {object} web.Response "We need ID"
// @Failure 404 {object} web.Response "Can not find ID"
// @Failure 422 {object} web.Response "Missing some mandatory field"
// @Success 200 {object} web.Response
// @Router /api/v1/products/{some_id} [PATCH]
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
		var validate *validator.Validate = validator.New()
		var req products.Product
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
			return
		}
		req.ID = id
		errValidate := validate.Struct(req)
		if errValidate != nil {
			if _, ok := errValidate.(*validator.InvalidValidationError); ok {
				c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
				return
			}
			for _, errValidate := range errValidate.(validator.ValidationErrors) {
				if errValidate != nil {
					s := fmt.Sprintf("%s is mandatory", errValidate.Field())
					c.JSON(web.DecodeError(http.StatusUnprocessableEntity, s))
					return
				}
			}
		}
		p, err := prod.service.Update(req, int(id))
		if err != nil {
			if err.Error() == ERROR_UNIQUE_PRODUCT_CODE {
				c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
				return
			}
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(web.NewResponse(http.StatusOK, p))
	}
	return fn
}

// DeleteProducts godoc
// @Summary Delete products by ID
// @Tags Products
// @Description delete products by ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param some_id path int true "Some ID"
// @Failure 401 {object} web.Response "We need token"
// @Failure 400 {object} web.Response "We need ID"
// @Failure 404 {object} web.Response "Can not find ID"
// @Success 204 {object} web.Response
// @Router /api/v1/products/{some_id} [DELETE]
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
