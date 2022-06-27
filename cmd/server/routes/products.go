package routes

import (
	"database/sql"

	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	products "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Products(databaseConection *sql.DB, routerGroup *gin.RouterGroup) {
	productsRepository := products.NewDBRepository(databaseConection)
	productsService := products.NewService(productsRepository)
	productsHandler := handler.NewProduct(productsService)

	productsRouterGroup := routerGroup.Group("/products")
	{
		productsRouterGroup.POST("/", productsHandler.Store())
		productsRouterGroup.GET("/", productsHandler.GetAll())
		productsRouterGroup.GET("/:id", productsHandler.GetById())
		productsRouterGroup.PATCH("/:id", productsHandler.Update())
		productsRouterGroup.DELETE("/:id", productsHandler.Delete())
	}
}
