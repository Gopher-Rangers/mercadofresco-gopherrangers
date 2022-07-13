package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/database"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/api/handlers/validation"
	purchaseOrdersHandler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/controller"
	purchaseOrdersRepo "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/repository"
	purchaseOrdersService "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/purchase_orders/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func PurchaseOrders(routerGroup *gin.RouterGroup) {

	repo := purchaseOrdersRepo.NewRepository(database.GetInstance())
	service := purchaseOrdersService.NewService(repo)
	handler := purchaseOrdersHandler.NewPurchaseOrder(service)

	purchaseOrderGroup := routerGroup.Group("/purchase-orders")
	{
		purchaseOrderGroup.POST("/", handler.Create)
		purchaseOrderGroup.GET("/:id", validation.ValidateID, handler.GetPurchaseOrderById)
	}
}
