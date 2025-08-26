// Package routehandlers handles the request comming to each endpoint
package routehandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	roomcode "github.com/utkarshkrsingh/goparty/internal/roomCode"
)

// CreateRoom handles the request to the "/create-room" endpoint
func CreateRoom(c *gin.Context) {
	roomCode, err := roomcode.GenerateCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Code: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"roomCode": roomCode,
	})
}
