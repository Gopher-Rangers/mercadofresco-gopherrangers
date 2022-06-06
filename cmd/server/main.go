package main

import (
	"log"
	"os"

	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/docs"
	products "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Mercado Fresco
// @version 1.0
// @description This API Handle Mercado Fresco Sellers, Warehouse, Section, Products, Employees and Buyer
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("failed to load .env")
	}

	gin.SetMode("release")

	server := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	baseRoute := server.Group("/api/v1/")
	{
		productRouterGroup := baseRoute.Group("/products")
		{
			file := store.New(store.FileType, "../../internal/product/products.json")
			prod_rep := products.NewRepository(file)
			prod_service := products.NewService(prod_rep)
			prod := handler.NewProduct(prod_service)

			productRouterGroup.POST("/", prod.Store())
			productRouterGroup.GET("/", prod.GetAll())
			productRouterGroup.GET("/:id", prod.GetById())
			productRouterGroup.PATCH("/:id", prod.Update())
			productRouterGroup.DELETE("/:id", prod.Delete())
		}

		buyerRouterGroup := baseRoute.Group("/buyers")
		{
			buyerHandler := handler.NewBuyerHandler()

			buyerRouterGroup.Use(buyerHandler.AuthToken)

			buyerRouterGroup.GET("/", buyerHandler.GetAll)
			buyerRouterGroup.POST("/", buyerHandler.Create)
			buyerRouterGroup.GET("/:id", buyerHandler.ValidateID, buyerHandler.GetBuyerById)
			buyerRouterGroup.PUT("/:id", buyerHandler.ValidateID, buyerHandler.Update)
			buyerRouterGroup.DELETE("/:id", buyerHandler.ValidateID, buyerHandler.Delete)
		}

		sectionRouterGroup := baseRoute.Group("/sections")
		{
			file := store.New(store.FileType, "../../internal/section/sections.json")
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
			file := store.New(store.FileType, "../../internal/warehouse/warehouses.json")
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
