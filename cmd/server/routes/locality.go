package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/database"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/adapters"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
	"github.com/gin-gonic/gin"
)

func Locality(routerGroup *gin.RouterGroup) {

	localityRepository := adapters.NewMySqlLocalityRepository(database.GetInstance())
	localityService := usecases.NewServiceLocality(localityRepository)
	localityHandler := handlers.NewLocality(localityService)

	localityRouterGroup := routerGroup.Group("/localities/reportCarries")
	{

		localityRouterGroup.GET("/", localityHandler.GetCarryLocality)
	}
}
