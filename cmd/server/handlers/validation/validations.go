package validation

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func AuthToken(context *gin.Context) {
	privateToken := os.Getenv("TOKEN")

	providedToken := context.GetHeader("token")

	if providedToken != privateToken {
		context.AbortWithStatusJSON(web.DecodeError(http.StatusUnauthorized, "invalid token"))
		return
	}

	context.Next()
}

func ValidateID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id < 0 {
		ctx.AbortWithStatusJSON(web.DecodeError(http.StatusBadRequest, "Id need to be a valid integer"))
		return
	}

	ctx.Next()
}
