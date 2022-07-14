package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/database"
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/handlers"
	products "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	productrecord "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_record"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func ProductRecord(routerGroup *gin.RouterGroup, productsService products.Service) {
	productRecordRepository := productrecord.NewRepository(database.GetInstance())
	productRecordService := productrecord.NewService(productRecordRepository, productsService)
	productRecordHandler := handler.NewProductRecord(productRecordService)

	productRecordRouterGroupPost := routerGroup.Group("/productRecords")
	{
		productRecordRouterGroupPost.POST("/", productRecordHandler.Store())
	}
	productRecordRouterGroupGet := routerGroup.Group("/products/reportRecords")
	{
		productRecordRouterGroupGet.GET("/", productRecordHandler.Get())
	}
}
