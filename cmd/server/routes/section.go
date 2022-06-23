package routes

import (
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/store"
	"github.com/gin-gonic/gin"
)

func Sections(routerGroup *gin.RouterGroup) {
	sectionRouterGroup := routerGroup.Group("/sections")
	{
		file := store.New(store.FileType, "../../internal/section/sections.json")
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
