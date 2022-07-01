package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/database"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/adapters"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse/usecases"

	// "github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
	"github.com/gin-gonic/gin"
)

func Warehouses(routerGroup *gin.RouterGroup) {

	warehouseRouterGroup := routerGroup.Group("/warehouses")

	{
		// file := store.New(store.FileType, "../../internal/warehouse/warehouses.json")
		warehouseRep := adapters.NewMySqlRepository(database.GetInstance())
		warehouseService := usecases.NewService(warehouseRep)
		warehouse := handlers.NewWarehouse(warehouseService)

		warehouseRouterGroup.GET("/", warehouse.GetAll)
		warehouseRouterGroup.GET("/:id", warehouse.GetByID)
		warehouseRouterGroup.POST("/", warehouse.CreateWarehouse)
		warehouseRouterGroup.PATCH("/:id", warehouse.UpdatedWarehouseID)
		warehouseRouterGroup.DELETE("/:id", warehouse.DeleteWarehouse)
	}

}
