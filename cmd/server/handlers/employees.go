package handlers

import (
	"net/http"

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

func NewEmployee(p employee.Services) Employee {
	return Employee{p}
}

func (p *Employee) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req employeeRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		emp, err := p.service.Create(req.CardNumber, req.FirstName, req.LastName,
			req.WareHouseID)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusConflict, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusCreated, emp))
	}
}

func (p Employee) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		employees := p.service.GetAll()
		c.JSON(web.NewResponse(http.StatusOK, employees))
	}
}
