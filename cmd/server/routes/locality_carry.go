package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/database"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers/carries"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/adapters"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
	"github.com/gin-gonic/gin"
)

func LocalityCarry(routerGroup *gin.RouterGroup) {

	localityRepository := adapters.NewMySqlLocalityRepository(database.GetInstance())
	localityService := usecases.NewServiceLocality(localityRepository)
	localityHandler := carries.NewLocality(localityService)

	localityRouterGroup := routerGroup.Group("/localities/reportCarries")
	{

		localityRouterGroup.GET("/", localityHandler.GetCarryLocality)

	}
}
