package routes

import (
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
	"github.com/gin-gonic/gin"
)

func Buyers(routerGroup *gin.RouterGroup) {

	storeBuyers := store.New(store.FileType, "../../internal/buyer/buyers.json")
	repo := buyer.NewRepository(storeBuyers)
	buyersService := buyer.NewService(repo)
	buyerHandler := handler.NewBuyerHandler(buyersService)

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
