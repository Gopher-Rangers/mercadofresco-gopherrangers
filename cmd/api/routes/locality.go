package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/database"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality"
	"github.com/gin-gonic/gin"
)

func Localities(routerGroup *gin.RouterGroup) locality.Service {
	localityRepository := locality.NewMariaDBRepository(database.GetInstance())
	localityService := locality.NewService(localityRepository)
	localityController := handlers.NewLocality(localityService)

	localityRouterGroup := routerGroup.Group("/localities")
	{
		localityRouterGroup.GET("/", localityController.GetAll)
		localityRouterGroup.GET("/reportSellers", localityController.ReportSellers)
		localityRouterGroup.POST("/", localityController.Create)
	}

	return localityService
}
