package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/database"
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	employees "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"
	inboundOrders "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/inbound_orders"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Employees(routerGroup *gin.RouterGroup) {
	db := database.GetInstance()
	employeesRepository := employees.NewRepository(db)

	inboundOrderRepository := inboundOrders.NewRepository(db)
	inboundOrderService := inboundOrders.NewService(inboundOrderRepository)

	employeesService := employees.NewService(employeesRepository)
	employeesHandler := handler.NewEmployee(employeesService, inboundOrderService)

	employeesRouterGroup := routerGroup.Group("/employees")
	{
		employeesRouterGroup.POST("/", employeesHandler.Create())
		employeesRouterGroup.GET("/", employeesHandler.GetAll())
		employeesRouterGroup.GET("/:id", employeesHandler.GetById())
		employeesRouterGroup.GET("/:id/reportInboundOrders", employeesHandler.GetOrderCount())

		employeesRouterGroup.PATCH("/:id", employeesHandler.Update())
		employeesRouterGroup.DELETE("/:id", employeesHandler.Delete())
	}
}
