package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/database"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/handlers/validation"
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/controller"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/repository/myslq"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Buyers(routerGroup *gin.RouterGroup) {

	repo := myslq.NewRepository(database.GetInstance())
	buyersService := service.NewService(repo)
	buyerHandler := handler.NewBuyer(buyersService)

	buyerRouterGroup := routerGroup.Group("/buyers")
	{

		buyerRouterGroup.GET("/", buyerHandler.GetAll)
		buyerRouterGroup.POST("/", buyerHandler.Create)
		buyerRouterGroup.GET("/:id", validation.ValidateID, buyerHandler.GetBuyerById)
		buyerRouterGroup.PUT("/:id", validation.ValidateID, buyerHandler.Update)
		buyerRouterGroup.DELETE("/:id", validation.ValidateID, buyerHandler.Delete)
		buyerRouterGroup.GET("/report-purchase-orders", buyerHandler.ReportPurchaseOrdersByBuyer)
	}
}
