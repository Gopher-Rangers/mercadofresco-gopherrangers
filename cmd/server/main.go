package main

import (
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode("release")
	_ = godotenv.Load("./.env")

	server := gin.Default()

	baseRoute := server.Group("/api/v1/")
	{
		// sectionRouterGroupproductsRouterGroup := baseRoute.Group("/products")
		// {
		// 	productsRouterGroup.POST("/", productHandler.Save())
		// 	productsRouterGroup.GET("/", productHandler.GetAll())
		// 	productsRouterGroup.GET("/:id", productHandler.GetById())
		// 	productsRouterGroup.PUT("/:id", productHandler.Update())
		// 	productsRouterGroup.DELETE("/:id", productHandler.Delete())
		// 	productsRouterGroup.PATCH("/:id", productHandler.PatchNamePrice())
		// }

		sectionRouterGroup := baseRoute.Group("/sections")
		{
			file := store.New(store.FileType, "./internal/section/sections.json")
			sec_rep := section.NewRepository(file)
			sec_service := section.NewService(sec_rep)
			sec_p := handler.NewSection(sec_service)

			sectionRouterGroup.Use(sec_p.TokenAuthMiddleware)

			sectionRouterGroup.GET("/", sec_p.GetAll())
			sectionRouterGroup.POST("/", sec_p.CreateProduct())
			sectionRouterGroup.GET("/:id", sec_p.IdVerificatorMiddleware, sec_p.GetByID())
			sectionRouterGroup.PATCH("/:id", sec_p.IdVerificatorMiddleware, sec_p.UpdateSecID())
			sectionRouterGroup.DELETE("/:id", sec_p.IdVerificatorMiddleware, sec_p.DeleteSection())
		}

		warehouseRouterGroup := baseRoute.Group("/warehouses")
		{
			file := store.New(store.FileType, "./internal/warehouse/warehouses.json")
			warehouseRep := warehouse.NewRepository(file)
			warehouseService := warehouse.NewService(warehouseRep)
			warehouse := handler.NewWarehouse(warehouseService)

			warehouseRouterGroup.GET("/", warehouse.GetAll)
			warehouseRouterGroup.GET("/:id", warehouse.GetByID)
			warehouseRouterGroup.POST("/", warehouse.CreateWarehouse)
			warehouseRouterGroup.PATCH("/:id", warehouse.UpdatedWarehouseID)
			warehouseRouterGroup.DELETE("/:id", warehouse.DeleteWarehouse)
		}
	}
	server.Run()
}
