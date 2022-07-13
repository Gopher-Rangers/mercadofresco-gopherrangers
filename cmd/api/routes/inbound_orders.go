package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/database"
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/handlers"
	io "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/inbound_orders"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func InboundOrders(routerGroup *gin.RouterGroup) {

	inboundOrdersRepository := io.NewRepository(database.GetInstance())

	inboundOrdersService := io.NewService(inboundOrdersRepository)
	inboundOrdersHandler := handler.NewInboundOrder(inboundOrdersService)

	inboundOrdersRouterGroup := routerGroup.Group("/inboundOrders")
	{
		inboundOrdersRouterGroup.POST("/", inboundOrdersHandler.Create())
	}
}
