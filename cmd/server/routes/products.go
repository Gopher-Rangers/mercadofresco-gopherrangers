package routes

import (
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
	"github.com/gin-gonic/gin"
)

func Products(routerGroup *gin.RouterGroup) {

	productsStore := store.New(store.FileType, "../../internal/product/products.json")
	productsRepository := products.NewRepository(productsStore)
	productsService := products.NewService(productsRepository)
	productsHandler := handler.NewProduct(productsService)

	productRouterGroup := routerGroup.Group("/products")
	{
		productRouterGroup.POST("/", productsHandler.Store())
		productRouterGroup.GET("/", productsHandler.GetAll())
		productRouterGroup.GET("/:id", productsHandler.GetById())
		productRouterGroup.PATCH("/:id", productsHandler.Update())
		productRouterGroup.DELETE("/:id", productsHandler.Delete())
	}
}
