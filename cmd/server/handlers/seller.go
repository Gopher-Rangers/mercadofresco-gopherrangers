package handlers

import (
	"errors"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"
	"net/http"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

const (
	ERR_UNIQUE_CID_VALUE          = "the cid already exists"
	ERR_LOCALITY_NON_EXISTS_VALUE = "locality_id does not exists"
)

type requestSeller struct {
	CompanyId   int    `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityID  int    `json:"locality_id" binding:"required"`
}

type Seller struct {
	service seller.Service
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{service: s}
}

func (s *Seller) GetAll(ctx *gin.Context) {

	sellerList, err := s.service.GetAll(ctx)

	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}
	ctx.JSON(web.NewResponse(http.StatusOK, sellerList))
}

func (s *Seller) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")

	idConvertido, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusInternalServerError, err.Error()))
		return
	}

	oneSeller, err := s.service.GetOne(ctx, idConvertido)

	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	ctx.JSON(web.NewResponse(http.StatusOK, oneSeller))
}

func (s *Seller) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	idConvertido, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusInternalServerError, err.Error()))
		return
	}

	_, err = s.service.GetOne(ctx, idConvertido)

	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	var req requestSeller

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(web.DecodeError(http.StatusUnprocessableEntity, validateFields(req).Error()))
		return
	}

	updateSeller, err := s.service.Update(ctx, idConvertido, req.CompanyId, req.CompanyName, req.Address, req.Telephone)

	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	ctx.JSON(web.NewResponse(http.StatusOK, updateSeller))
}

func (s *Seller) Create(ctx *gin.Context) {
	var req requestSeller

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(web.DecodeError(http.StatusUnprocessableEntity, validateFields(req).Error()))
		return
	}

	newSeller, err := s.service.Create(ctx, req.CompanyId, req.CompanyName, req.Address, req.Telephone, req.LocalityID)

	if err != nil {
		switch err.Error() {
		case ERR_UNIQUE_CID_VALUE:
			ctx.JSON(web.DecodeError(http.StatusConflict, err.Error()))
			return

		case ERR_LOCALITY_NON_EXISTS_VALUE:
			ctx.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
			return
		}
	}

	ctx.JSON(web.NewResponse(http.StatusCreated, newSeller))
}

func (s *Seller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	idConvertido, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusInternalServerError, err.Error()))
		return
	}

	err = s.service.Delete(ctx, idConvertido)

	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}
	ctx.JSON(web.NewResponse(http.StatusNoContent, fmt.Sprintf("the seller %d was removed", idConvertido)))
}

func validateFields(req requestSeller) error {
	if req.CompanyId == 0 {
		return errors.New("field cid is required")
	}

	if req.CompanyName == "" {
		return errors.New("field company_name is required")
	}

	if req.Address == "" {
		return errors.New("field address is required")
	}

	if req.Telephone == "" {
		return errors.New("field telephone is required")
	}

	if req.LocalityID == 0 {
		return errors.New("locality_id is required")
	}
	return nil
}
