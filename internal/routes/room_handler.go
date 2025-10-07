package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utkarshkrsingh/goparty/internal/utils"
)

// CreateRoom handles the request to the "/create-room" endpoint
func createRoom(ctx *gin.Context) {
	roomCode, err := utils.GenerateCode()
	if err != nil {
		utils.RespondError(ctx, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"roomCode": roomCode,
	})
}
