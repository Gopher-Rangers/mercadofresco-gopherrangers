package handlers

import (
	"net/http"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

type requestWarehouse struct {
	ID              int    `json:"id"`
	Warehouse_code  string `json:"warehouse_code" binding:"required"`
	Address         string `json:"address"`
	Telephone       string `json:"telephone"`
	Min_Capacity    int    `json:"minimun_capacity"`
	Min_Temperature int    `json:"minimun_temperature"`
}

type Warehouse struct {
	service warehouse.Service
}

func NewWarehouse(w warehouse.Service) Warehouse {
	return Warehouse{w}
}

func (w Warehouse) GetAll(c *gin.Context) {
	warehouse := w.service.GetAll()

	c.JSON(web.NewResponse(http.StatusOK, warehouse))
}

func (w Warehouse) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(web.DecodeError(http.StatusBadRequest, "O id passado não é um número!"))
		return
	}

	warehouse, err := w.service.GetByID(id)

	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, "O warehouse não foi encontrado!"))
		return
	}

	c.JSON(web.NewResponse(http.StatusOK, warehouse))
}

func (w Warehouse) CreateWarehouse(c *gin.Context) {
	var req requestWarehouse

	if err := c.Bind(&req); err != nil {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	warehouse, err := w.service.CreateWarehouse(req.ID, req.Warehouse_code, req.Address,
		req.Telephone, req.Min_Capacity, req.Min_Temperature)

	if err != nil {
		c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusCreated, warehouse))
}

func (w Warehouse) UpdatedWarehouseID(c *gin.Context) {
	var req requestWarehouse

	if err := c.Bind(&req); err != nil {
		c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(web.DecodeError(http.StatusBadRequest, "O id passado não é um número!"))
		return
	}

	warehouse, err := w.service.UpdatedWarehouseID(id, req.Warehouse_code)

	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusOK, warehouse))
}

func (w Warehouse) DeleteWarehouse(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(web.DecodeError(http.StatusBadRequest, "O id passado não é um número!"))
		return
	}

	err = w.service.DeleteWarehouse(id)

	if err != nil {
		c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(web.NewResponse(http.StatusNoContent, "O warehouse foi removido com sucesso! "))
}
