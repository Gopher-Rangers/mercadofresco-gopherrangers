package main

import "github.com/gin-gonic/gin"

func main() {
	gin := gin.Default()

	baseRoute := gin.Group("/api/v1/")
	{
		productsRouterGroup := baseRoute.Group("/products")
		{
			productsRouterGroup.POST("/", productHandler.Save())
			productsRouterGroup.GET("/", productHandler.GetAll())
			productsRouterGroup.GET("/:id", productHandler.GetById())
			productsRouterGroup.PUT("/:id", productHandler.Update())
			productsRouterGroup.DELETE("/:id", productHandler.Delete())
			productsRouterGroup.PATCH("/:id", productHandler.PatchNamePrice())
		}
	}

}
