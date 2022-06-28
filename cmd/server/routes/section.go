package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/database"
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/section"
	"github.com/gin-gonic/gin"
)

func Sections(routerGroup *gin.RouterGroup) {
	sectionRouterGroup := routerGroup.Group("/sections")
	{
		sec_rep := section.NewRepository(database.GetInstance())
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
