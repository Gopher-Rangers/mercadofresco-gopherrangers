package main

import (
	"log"
	"os"

	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/routes"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/docs"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/seller"
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

	//gin.SetMode("release")

	server := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	baseRoute := server.Group("/api/v1/")
	{
		routes.Products(baseRoute)

		routes.ProductRecord(baseRoute)

		routes.Buyers(baseRoute)

		routes.Sections(baseRoute)

		sellerRouterGroup := baseRoute.Group("/sellers")
		{
			file := store.New(store.FileType, "../../internal/seller/seller.json")
			sellerRepository := seller.NewRepository(file)
			sellerService := seller.NewService(sellerRepository)
			sellerController := handler.NewSeller(sellerService)

			sellerRouterGroup.GET("/", sellerController.GetAll)
			sellerRouterGroup.GET("/:id", sellerController.GetOne)
			sellerRouterGroup.PUT("/:id", sellerController.Update)
			sellerRouterGroup.POST("/", sellerController.Create)
			sellerRouterGroup.DELETE("/:id", sellerController.Delete)
		}

		employeeRouterGroup := baseRoute.Group("/employees")
		{
			file := store.New(store.FileType, "../../internal/employee/employees.json")
			employee_rep := employee.NewRepository(file)
			employee_service := employee.NewService(employee_rep)
			employee := handler.NewEmployee(employee_service)

			employeeRouterGroup.GET("/", employee.GetAll())
			employeeRouterGroup.POST("/", employee.Create())
			employeeRouterGroup.GET("/:id", employee.GetById())
			employeeRouterGroup.PATCH("/:id", employee.Update())
			employeeRouterGroup.DELETE("/:id", employee.Delete())
		}

		routes.Warehouses(baseRoute)
	}
	server.Run()
}
