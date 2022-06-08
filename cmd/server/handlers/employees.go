package handlers

import (
	"net/http"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

type employeeRequest struct {
	ID          int    `json:"id"`
	CardNumber  int    `json:"card_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	WareHouseID int    `json:"warehouse_id"`
}

type Employee struct {
	service employee.Services
}

func NewEmployee(e employee.Services) Employee {
	return Employee{e}
}

func (e *Employee) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req employeeRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		emp, err := e.service.Create(req.CardNumber, req.FirstName, req.LastName,
			req.WareHouseID)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusCreated, emp))
	}
}

func (e Employee) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		employees := e.service.GetAll()
		c.JSON(web.NewResponse(http.StatusOK, employees))
	}
}

func (e Employee) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(web.DecodeError(http.StatusBadRequest, "Id inválido"))
			return
		}
		err = e.service.Delete(id)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusNoContent, "funcionario deletado"))
	}
}

func (e Employee) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(web.DecodeError(http.StatusBadRequest, "Id inválido"))
			return
		}
		employee, err := e.service.GetById(id)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(web.NewResponse(http.StatusAccepted, employee))
	}
}

func (e *Employee) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req employee.Employee
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		id, _ := strconv.Atoi(c.Param("id"))
		employee, err := e.service.Update(req, id)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusOK, employee))
	}
}
