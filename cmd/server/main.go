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

	gin := gin.Default()

	db := store.New(store.FileType, "./internal/section/sections.json")

	sec_rep := section.NewRepository(db)
	sec_service := section.NewService(sec_rep)
	sec_p := handler.NewSection(sec_service)

	baseRoute := gin.Group("/api/v1/")
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
			sectionRouterGroup.Use(sec_p.TokenAuthMiddleware)

			sectionRouterGroup.GET("/", sec_p.GetAll())
			sectionRouterGroup.POST("/", sec_p.CreateProduct())
			sectionRouterGroup.GET("/:id", sec_p.IdVerificatorMiddleware, sec_p.GetByID())
			sectionRouterGroup.PATCH("/:id", sec_p.IdVerificatorMiddleware, sec_p.UpdateSecID())
			sectionRouterGroup.DELETE("/:id", sec_p.IdVerificatorMiddleware, sec_p.DeleteSection())
		}
	}
	gin.Run()
}
