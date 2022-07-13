package main

import (
	"os"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/routes"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/go-sql-driver/mysql"
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

	godotenv.Load(".env")

	gin.SetMode("release")

	server := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	baseRoute := server.Group("/api/v1/")
	{
		localityService := routes.Localities(baseRoute)
		sellerService := routes.Sellers(baseRoute, localityService)
		productsService := routes.Products(baseRoute, sellerService)

		routes.ProductRecord(baseRoute, productsService)

		routes.Buyers(baseRoute)

		routes.PurchaseOrders(baseRoute)

		routes.Sections(baseRoute)

		routes.ProductBatches(baseRoute)

		routes.Employees(baseRoute)

		routes.InboundOrders(baseRoute)

		routes.Carry(baseRoute)

		routes.LocalityCarry(baseRoute)

		routes.Warehouses(baseRoute)
	}
	server.Run()
}
