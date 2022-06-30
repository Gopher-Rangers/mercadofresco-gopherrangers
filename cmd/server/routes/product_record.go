package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/database"
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	productrecord "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_record"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func ProductRecord(routerGroup *gin.RouterGroup) {
	productRecordRepository := productrecord.NewRepository(database.GetInstance())
	productRecordService := productrecord.NewService(productRecordRepository)
	productRecordHandler := handler.NewProductRecord(productRecordService)

	productRecordRouterGroup := routerGroup.Group("/productRecords")
	{
		productRecordRouterGroup.POST("/", productRecordHandler.Store())
		productRecordRouterGroup.GET("/", productRecordHandler.GetById())
	}
}
