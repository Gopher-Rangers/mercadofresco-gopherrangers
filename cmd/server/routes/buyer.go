package routes

import (
	"database/sql"
	"fmt"
	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func Buyers(routerGroup *gin.RouterGroup) {

	productsDatabase := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
	)
	conn, err := sql.Open("mysql", productsDatabase)
	if err != nil {
		log.Fatal("failed: ", err.Error())
	}

	repo := buyer.NewRepository(conn)
	buyersService := buyer.NewService(repo)
	buyerHandler := handler.NewBuyer(buyersService)

	buyerRouterGroup := routerGroup.Group("/buyers")
	{

		buyerRouterGroup.Use(buyerHandler.AuthToken)

		buyerRouterGroup.GET("/", buyerHandler.GetAll)
		buyerRouterGroup.POST("/", buyerHandler.Create)
		buyerRouterGroup.GET("/:id", buyerHandler.ValidateID, buyerHandler.GetBuyerById)
		buyerRouterGroup.PUT("/:id", buyerHandler.ValidateID, buyerHandler.Update)
		buyerRouterGroup.DELETE("/:id", buyerHandler.ValidateID, buyerHandler.Delete)
	}
}
