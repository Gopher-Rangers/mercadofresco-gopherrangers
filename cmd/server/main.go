package main

import (
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section"
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
			section := handler.NewSection(sec_service)

			sectionRouterGroup.Use(section.TokenAuthMiddleware)

			sectionRouterGroup.GET("/", section.GetAll())
			sectionRouterGroup.POST("/", section.CreateSection())
			sectionRouterGroup.GET("/:id", section.IdVerificatorMiddleware, section.GetByID())
			sectionRouterGroup.PATCH("/:id", section.IdVerificatorMiddleware, section.UpdateSecID())
			sectionRouterGroup.DELETE("/:id", section.IdVerificatorMiddleware, section.DeleteSection())
		}
	}
	server.Run()
}
