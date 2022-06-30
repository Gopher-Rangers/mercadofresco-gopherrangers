package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	productrecord "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_record"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	ERROR_LAST_UPDATE_DATE  = "LastUpdateDate is mandatory"
	ERROR_PURCHASE_PRICE    = "PurchasePrice is mandatory"
	ERROR_SALE_PRICE        = "SalePrice is mandatory"
	ERROR_PRODUCT_ID        = "ProductId is mandatory"
	ERROR_UNIQUE_PRODUCT_ID = "the product id must be unique"
)

type requestProductRecord struct {
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float64 `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductId      int     `json:"product_id"`
}

type ProductRecord struct {
	service productrecord.Service
}

func NewProductRecord(p productrecord.Service) *ProductRecord {
	return &ProductRecord{service: p}
}

func NewRequestProductRecord() requestProductRecord {
	p := requestProductRecord{}
	return p
}

// StoreProducts godoc
// @Summary Store product records
// @Tags Product Records
// @Description store product records
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body requestProductRecord true "Product Record to store"
// @Failure 401 {object} web.Response "We need token"
// @Failure 404 {object} web.Response
// @Failure 422 {object} web.Response "Missing some mandatory field"
// @Success 201 {object} web.Response
// @Router /api/v1/productRecords [POST]
func (prod *ProductRecord) Store() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(web.DecodeError(http.StatusUnauthorized, ERROR_TOKEN))
			return
		}
		var validate *validator.Validate = validator.New()
		var req productrecord.ProductRecord
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

// ListProductRecordsById godoc
// @Summary List product records by ID
// @Tags Products Records
// @Description list products records by ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "ID"
// @Failure 401 {object} web.Response "We need token"
// @Failure 400 {object} web.Response "We need ID"
// @Failure 404 {object} web.Response "Can not find ID"
// @Success 200 {object} web.Response
// @Router /api/v1/productRecords/{id} [GET]
func (prod *ProductRecord) GetById() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(web.DecodeError(http.StatusUnauthorized, ERROR_TOKEN))
			return
		}
		idStr := c.Query("id")
		idNum, err := strconv.Atoi(idStr)
		if err != nil && idStr != "" {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_ID))
			return
		}
		if idStr == "" {
			p, err := prod.service.GetAll()
			if err != nil {
				c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
				return
			}
			c.JSON(web.NewResponse(http.StatusOK, p))
		} else {
			p, err := prod.service.GetById(idNum)
			if err != nil {
				c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
				return
			}
			c.JSON(web.NewResponse(http.StatusOK, p))
		}
	}
	return fn
}
