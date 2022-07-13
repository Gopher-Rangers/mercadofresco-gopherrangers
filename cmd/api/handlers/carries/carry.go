package carries

import (
	"net/http"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

type Carry struct {
	service usecases.ServiceCarry
}

func NewCarry(c usecases.ServiceCarry) Carry {
	return Carry{c}
}

func (c Carry) CreateCarry(ctx *gin.Context) {
	var req domain.Carry

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(web.DecodeError(http.StatusUnprocessableEntity, "o campo `cid ` é obrigatório"))
		return
	}

	carry, err := c.service.CreateCarry(req)

	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusConflict, err.Error()))
		return
	}

	ctx.JSON(web.NewResponse(http.StatusCreated, carry))

}
