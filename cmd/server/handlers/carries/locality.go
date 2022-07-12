package carries

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/domain"
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

	localityIDs := ctx.Query("id")

	splitLocalityIDs := strings.Split(localityIDs, ",")

	results := []domain.Locality{}

	if localityIDs == "" {

		localities, err := l.service.GetAllCarriesLocality()

		if err != nil {
			ctx.JSON(web.DecodeError(http.StatusInternalServerError, "erro ao acessar o banco de dados"))
			return
		}

		ctx.JSON(web.NewResponse(http.StatusOK, localities))
		return

	} else {

		for _, stringId := range splitLocalityIDs {

			id, err := strconv.Atoi(stringId)

			if err != nil {
				ctx.JSON(web.DecodeError(http.StatusBadRequest, "id fornecido é inválido!"))
				return
			}

			locality, err := l.service.GetCarryLocalityByID(id)

			if err != nil {
				ctx.JSON(web.DecodeError(http.StatusNotFound, "a localidade não foi encontrada!"))
				return
			}

			results = append(results, locality)

		}

		ctx.JSON(web.NewResponse(http.StatusOK, results))

	}

}
