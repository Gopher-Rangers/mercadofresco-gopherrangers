package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/database"
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"
	"github.com/gin-gonic/gin"
)

func Sellers(routerGroup *gin.RouterGroup, localityService locality.Service) seller.Service {

	sellerRepository := seller.NewMariaDBRepository(database.GetInstance())
	sellerService := seller.NewService(sellerRepository, localityService)
	sellerController := handler.NewSeller(sellerService)

	sellerRouterGroup := routerGroup.Group("/sellers")
	{
		sellerRouterGroup.GET("/", sellerController.GetAll)
		sellerRouterGroup.GET("/:id", sellerController.GetOne)
		sellerRouterGroup.PUT("/:id", sellerController.Update)
		sellerRouterGroup.POST("/", sellerController.Create)
		sellerRouterGroup.DELETE("/:id", sellerController.Delete)
	}
	return sellerService
}
