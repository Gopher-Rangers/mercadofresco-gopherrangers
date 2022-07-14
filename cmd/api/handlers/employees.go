package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"
	inboundorders "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/inbound_orders"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	ERROR_UNIQUE_ID    = "não insira id"
	ERROR_ALLMANDATORY = "preencha todos os campos"
	ERROR_NOTFOUND     = "funcionario nao encontrado"
)

type employeeRequest struct {
	ID          int    `json:"id"`
	CardNumber  int    `json:"card_number_id" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	WareHouseID int    `json:"warehouse_id" validate:"required"`
}

type Employee struct {
	employeeService     employee.Services
	inboundOrderService inboundorders.Services
}

func NewEmployee(e employee.Services, io inboundorders.Services) Employee {
	return Employee{employeeService: e, inboundOrderService: io}
}

func (e *Employee) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var validate *validator.Validate = validator.New()
		var req employeeRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}
		errValidate := validate.Struct(req)
		if errValidate != nil {
			if _, ok := errValidate.(*validator.InvalidValidationError); ok {
				c.JSON(web.DecodeError(http.StatusNotFound, errValidate.Error()))
				return
			}
			for _, errValidate := range errValidate.(validator.ValidationErrors) {
				if errValidate != nil {
					s := fmt.Sprintf("%s é obrigatório", errValidate.Field())
					c.JSON(web.DecodeError(http.StatusUnprocessableEntity, s))
					return
				}
			}
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
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(web.DecodeError(http.StatusBadRequest, ERROR_ID))
			return
		}
		var validate *validator.Validate = validator.New()
		var req employee.Employee
		if err := c.Bind(&req); err != nil {
			c.JSON(web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
			return
		}
		req.ID = id
		errValidate := validate.Struct(req)
		if errValidate != nil {
			if _, ok := errValidate.(*validator.InvalidValidationError); ok {
				c.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
				return
			}
			for _, errValidate := range errValidate.(validator.ValidationErrors) {
				if errValidate != nil {
					s := fmt.Sprintf("%s is mandatory", errValidate.Field())
					c.JSON(web.DecodeError(http.StatusUnprocessableEntity, s))
					return
				}
			}
		}
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
