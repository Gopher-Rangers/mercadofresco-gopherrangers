package handler

import (
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type requestSeller struct {
	Cid         int    `json:"company_id" binding:"required"`
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
		c.JSON(http.StatusNotFound, gin.H{
			"error": "não há vendedores",
		})
		return
	}
	c.JSON(http.StatusOK, sellerList)
}

func (s *Seller) GetOne(c *gin.Context) {
	id := c.Param("id")

	idConvertido, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	oneSeller, err := s.service.GetOne(idConvertido)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, oneSeller)
}

func (s *Seller) Update(c *gin.Context) {
	id := c.Param("id")

	idConvertido, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	var req requestSeller

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	updateSeller, err := s.service.Update(idConvertido, req.CompanyName, req.Address, req.Telephone)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updateSeller)
}

func (s *Seller) Create(c *gin.Context) {
	var req requestSeller

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	newSeller, err := s.service.Create(req.Cid, req.CompanyName, req.Address, req.Telephone)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusConflict, fmt.Sprintf("error: %s", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, newSeller)
}

func (s *Seller) Delete(c *gin.Context) {
	id := c.Param("id")

	idConvertido, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = s.service.Delete(idConvertido)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("o vendedor %d foi removido", idConvertido))
}
