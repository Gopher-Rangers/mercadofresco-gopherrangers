package handlers

import (
	"errors"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	ERR_UNIQUE_ZIPCODE_VALUE = "zip_code already exists"
)

type requestLocality struct {
	ZipCode      string `json:"zip_code" binding:"required"`
	LocalityName string `json:"locality_name" binding:"required"`
	ProvinceName string `json:"province_name" binding:"required"`
	CountryName  string `json:"country_name" binding:"required"`
}

type Locality struct {
	service locality.Service
}

func NewLocality(s locality.Service) *Locality {
	return &Locality{service: s}
}

func (l *Locality) ReportSellers(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")

	if !ok {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "missing parameter url"))
		return
	}

	idConvertido, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusInternalServerError, err.Error()))
		return
	}

	reportSeller, err := l.service.ReportSellers(ctx, idConvertido)

	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
		return
	}

	ctx.JSON(web.NewResponse(http.StatusOK, reportSeller))
}

func (l *Locality) Create(ctx *gin.Context) {
	var req requestLocality

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(web.DecodeError(http.StatusUnprocessableEntity, validateLocalityFields(req).Error()))
		return
	}

	newLocality, err := l.service.Create(ctx, req.ZipCode, req.LocalityName, req.ProvinceName, req.CountryName)

	if err != nil {
		switch err.Error() {
		case ERR_UNIQUE_ZIPCODE_VALUE:
			ctx.JSON(web.DecodeError(http.StatusConflict, err.Error()))
			return
		default:
			ctx.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
			return
		}
	}

	ctx.JSON(web.NewResponse(http.StatusCreated, newLocality))
}

func (l *Locality) GetAll(ctx *gin.Context) {

	localityList, err := l.service.GetAll(ctx)

	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	ctx.JSON(web.NewResponse(http.StatusOK, localityList))
}

func validateLocalityFields(req requestLocality) error {

	if req.ZipCode == "" {
		return errors.New("invalid input in field zip_code")
	}
	if req.LocalityName == "" {
		return errors.New("invalid input in field locality_name")
	}
	if req.ProvinceName == "" {
		return errors.New("invalid input in field province_name")
	}

	if req.CountryName == "" {
		return errors.New("invalid input in field country_name")
	}
	return nil
}
