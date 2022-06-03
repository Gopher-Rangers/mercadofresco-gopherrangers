package main

import (
	"log"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/tree/main/cmd/server/handler/products"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("failed to load .env")
	}
	
	gin := gin.Default()

	baseRoute := gin.Group("/api/v1/")
	{
		productsRouterGroup := baseRoute.Group("/products")
		{
			productsRouterGroup.POST("/", productHandler.Save())
			//productsRouterGroup.GET("/", productHandler.GetAll())
			//productsRouterGroup.GET("/:id", productHandler.GetById())
			//productsRouterGroup.PUT("/:id", productHandler.Update())
			//productsRouterGroup.DELETE("/:id", productHandler.Delete())
			//productsRouterGroup.PATCH("/:id", productHandler.PatchNamePrice())
		}
	}
}
