package routes

import (
	"database/sql"
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Buyers(databaseConection *sql.DB, routerGroup *gin.RouterGroup) {

	repo := buyer.NewRepository(databaseConection)
	buyersService := buyer.NewService(repo)
	buyerHandler := handler.NewBuyer(buyersService)

	buyerRouterGroup := routerGroup.Group("/buyers")
	{

		buyerRouterGroup.Use(buyerHandler.AuthToken)

		buyerRouterGroup.GET("/", buyerHandler.GetAll)
		buyerRouterGroup.POST("/", buyerHandler.Create)
		buyerRouterGroup.GET("/:id", buyerHandler.ValidateID, buyerHandler.GetBuyerById)
		buyerRouterGroup.PUT("/:id", buyerHandler.ValidateID, buyerHandler.Update)
		buyerRouterGroup.DELETE("/:id", buyerHandler.ValidateID, buyerHandler.Delete)
	}
}
