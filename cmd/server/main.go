package main

import (
	"log"

	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("failed to load .env")
	}

	gin.SetMode("release")

	server := gin.Default()

	baseRoute := server.Group("/api/v1/")
	{
		productRouterGroup := baseRoute.Group("/products")
		{
			file := store.New(store.FileType, "../../internal/product/products.json")
			prod_rep := products.NewRepository(file)
			prod_service := products.NewService(prod_rep)
			prod := handler.NewProduct(prod_service)

			//productRouterGroup.Use(prod.TokenAuthMiddleware)

			productRouterGroup.POST("/", prod.Store())
			productRouterGroup.GET("/", prod.GetAll())
			productRouterGroup.GET("/:id", prod.GetById())
			productRouterGroup.PUT("/:id", prod.UpdatePut())
			//productRouterGroup.PATCH("/:id", prod.IdVerificatorMiddleware, prod.UpdatePatch())
			productRouterGroup.DELETE("/:id", prod.Delete())
		}

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
	}
	server.Run()
}
