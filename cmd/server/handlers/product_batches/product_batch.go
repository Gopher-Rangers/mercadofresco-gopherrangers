package product_batches

import (
	"net/http"
	"strconv"

	productbatch "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_batch"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductBatch struct {
	service productbatch.Services
}

func NewProductBatch(p productbatch.Services) ProductBatch {
	return ProductBatch{p}
}

func (p *ProductBatch) GetByID() gin.HandlerFunc {
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

func (p *ProductBatch) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req productbatch.ProductBatch
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		sec, err := p.service.Create(req)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusCreated, sec))
	}
}
