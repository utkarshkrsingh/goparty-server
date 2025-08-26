package routehandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not exists",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User is authenticated",
		"user":    user,
	})
}
