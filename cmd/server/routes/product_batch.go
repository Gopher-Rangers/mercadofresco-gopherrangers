package routes

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/database"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers/product_batches"
	productbatch "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product_batch"
	"github.com/gin-gonic/gin"
)

func ProductBatches(routerGroup *gin.RouterGroup) {
	pb_rep := productbatch.NewRepository(database.GetInstance())
	pb_service := productbatch.NewService(pb_rep)
	productBatch := product_batches.NewProductBatch(pb_service)

	routerGroup.POST("productBatches/", productBatch.Create())
	routerGroup.GET("sections/reportProducts", productBatch.GetByID())
}
