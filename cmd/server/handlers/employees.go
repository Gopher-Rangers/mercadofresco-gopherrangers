package handlers

import (
	"net/http"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"
	inboundorders "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/inbound_orders"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

const (
	ERROR_UNIQUE_ID    = "não insira id"
	ERROR_ALLMANDATORY = "preencha todos os campos"
	ERROR_NOTFOUND     = "funcionario nao encontrado"
)

type employeeRequest struct {
	ID          int    `json:"id"`
	CardNumber  int    `json:"card_number_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	WareHouseID int    `json:"warehouse_id"`
}

type Employee struct {
	employeeService     employee.Services
	inboundOrderService inboundorders.Services
}

func NewEmployee(e employee.Services, io inboundorders.Services) Employee {
	return Employee{employeeService: e, inboundOrderService: io}
}

func (emp *Employee) checkBody(req employeeRequest, c *gin.Context) bool {
	employees, _ := emp.employeeService.GetAll()
	for i := range employees {
		if employees[i].ID == req.ID || req.ID != 0 {
			c.JSON(web.DecodeError(
				http.StatusUnprocessableEntity,
				ERROR_UNIQUE_ID))
			return false
		}
	}
	if req.CardNumber == 0 {
		c.JSON(web.DecodeError(
			http.StatusUnprocessableEntity,
			ERROR_ALLMANDATORY))
		return false
	}
	if req.FirstName == "" {
		c.JSON(web.DecodeError(
			http.StatusUnprocessableEntity,
			ERROR_ALLMANDATORY))
		return false
	}
	if req.LastName == "" {
		c.JSON(web.DecodeError(
			http.StatusUnprocessableEntity,
			ERROR_ALLMANDATORY))
		return false
	}
	if req.WareHouseID == 0 {
		c.JSON(web.DecodeError(
			http.StatusUnprocessableEntity,
			ERROR_ALLMANDATORY))
		return false
	}

	return true
}

func (e *Employee) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req employeeRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}
		if !e.checkBody(req, c) {
			return
		}

		emp, err := e.employeeService.Create(req.CardNumber, req.FirstName, req.LastName,
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
		employees, _ := e.employeeService.GetAll()
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
		err = e.employeeService.Delete(id)
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
		employee, err := e.employeeService.GetById(id)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(web.NewResponse(http.StatusOK, employee))
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
		employee, err := e.employeeService.Update(req, id)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(web.NewResponse(http.StatusOK, employee))
	}
}

func (e Employee) GetOrderCount() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(web.DecodeError(http.StatusBadRequest, "Id inválido"))
			return
		}

		count := e.inboundOrderService.GetCounterByEmployee(id)

		employee, err := e.employeeService.GetCount(id, count)
		if err != nil {
			c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusOK, employee))
	}
}
