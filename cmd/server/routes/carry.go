package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/database"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/adapters"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
	"github.com/gin-gonic/gin"
)

func Carry(routerGroup *gin.RouterGroup) {

	carryRepository := adapters.NewMySqlCarryRepository(database.GetInstance())
	carryService := usecases.NewServiceCarry(carryRepository)
	carryHandler := handlers.NewCarry(carryService)

	carryRouterGroup := routerGroup.Group("/carries")
	{

		carryRouterGroup.POST("/", carryHandler.CreateCarry)
	}
}
