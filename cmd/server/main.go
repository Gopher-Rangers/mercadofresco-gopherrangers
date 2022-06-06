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
