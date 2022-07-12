package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/database"
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	products "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	seller "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Products(routerGroup *gin.RouterGroup, sellerService seller.Service) products.Service {
	productsRepository := products.NewRepository(database.GetInstance())
	productsService := products.NewService(productsRepository, sellerService)
	productsHandler := handler.NewProduct(productsService)

	productsRouterGroup := routerGroup.Group("/products")
	{
		productsRouterGroup.POST("/", productsHandler.Store())
		productsRouterGroup.GET("/", productsHandler.GetAll())
		productsRouterGroup.GET("/:id", productsHandler.GetById())
		productsRouterGroup.PATCH("/:id", productsHandler.Update())
		productsRouterGroup.DELETE("/:id", productsHandler.Delete())
	}
	return productsService
}
