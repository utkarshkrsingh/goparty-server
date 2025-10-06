package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utkarshkrsingh/goparty/internal/utils"
)

// CreateRoom handles the request to the "/create-room" endpoint
func createRoom(c *gin.Context) {
	roomCode, err := utils.GenerateCode()
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
