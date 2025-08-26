package routehandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	// Overwrite the cookie and expire it
	ctx.SetCookie(
		"Authorization", // cookie name
		"",              // value
		-1,              // maxAge (in seconds) → -1 means delete immediately
		"/",             // path
		"",              // domain (empty = current domain)
		true,            // secure (true = HTTPS only; set false if dev on localhost HTTP)
		true,            // httpOnly (not accessible to JS)
	)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
