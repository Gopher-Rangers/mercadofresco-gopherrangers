package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

type requestSeller struct {
	CompanyId   int    `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
}

type Seller struct {
	service seller.Service
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{service: s}
}

func (s *Seller) GetAll(c *gin.Context) {

	sellerList, err := s.service.GetAll()

	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusOK, sellerList))
}

func (s *Seller) GetOne(c *gin.Context) {
	id := c.Param("id")

	idConvertido, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(web.DecodeError(http.StatusInternalServerError, err.Error()))
	}

	oneSeller, err := s.service.GetOne(idConvertido)

	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusOK, oneSeller))
}

func (s *Seller) Update(c *gin.Context) {
	id := c.Param("id")

	idConvertido, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(web.DecodeError(http.StatusInternalServerError, err.Error()))
	}

	var req requestSeller

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, validateFields(req).Error()))
		return
	}

	updateSeller, err := s.service.Update(idConvertido, req.CompanyId, req.CompanyName, req.Address, req.Telephone)

	if err != nil {
		c.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusOK, updateSeller))
}

func (s *Seller) Create(c *gin.Context) {
	var req requestSeller

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, validateFields(req).Error()))
		return
	}

	if err := validateFields(req); err != nil {
		c.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
	}

	newSeller, err := s.service.Create(req.CompanyId, req.CompanyName, req.Address, req.Telephone)

	if err != nil {
		fmt.Println(err)
		c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusCreated, newSeller))
}

func (s *Seller) Delete(c *gin.Context) {
	id := c.Param("id")

	idConvertido, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(web.DecodeError(http.StatusInternalServerError, err.Error()))
	}

	err = s.service.Delete(idConvertido)

	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusNoContent, fmt.Sprintf("the seller %d was removed", idConvertido)))
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
	return nil
}
