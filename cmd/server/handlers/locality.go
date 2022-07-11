package handlers

import (
	"net/http"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

type Locality struct {
	service usecases.ServiceLocality
}

func NewLocality(l usecases.ServiceLocality) Locality {
	return Locality{l}
}

func (l Locality) GetCarryLocality(ctx *gin.Context) {

	localityID, _ := strconv.Atoi(ctx.Query("id"))

	if localityID == 0 {
		localities, err := l.service.GetAllCarriesLocality()

		if err != nil {
			ctx.JSON(web.DecodeError(http.StatusBadRequest, "erro ao acessar o banco de dados"))
			return
		}

		ctx.JSON(web.NewResponse(http.StatusOK, localities))

	} else {

		locality, err := l.service.GetCarryLocalityByID(localityID)

		if err != nil {
			ctx.JSON(web.DecodeError(http.StatusNotFound, "a localidade não foi encontrada!"))
			return
		}

		ctx.JSON(web.NewResponse(http.StatusOK, locality))

	}

}
