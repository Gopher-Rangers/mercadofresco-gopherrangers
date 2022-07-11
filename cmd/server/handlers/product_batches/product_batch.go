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

func (p *ProductBatch) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req productbatch.ProductBatch
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		pb, err := p.service.Create(ctx, req)
		if err != nil {
			ctx.JSON(web.DecodeError(http.StatusConflict, err.Error()))
			return
		}

		ctx.JSON(web.NewResponse(http.StatusCreated, pb))
	}
}

func (p *ProductBatch) Report() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var repID interface{}
		var err error

		id, _ := strconv.Atoi(ctx.Query("id"))
		if id == 0 {
			repID, err = p.service.Report(ctx)

			if err != nil {
				ctx.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
				return
			}

		} else {
			repID, err = p.service.ReportByID(ctx, id)

			if err != nil {
				ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
				return
			}
		}

		ctx.JSON(web.NewResponse(http.StatusOK, repID))
	}
}
