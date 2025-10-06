// Package utils handles utilities working as backbones
package utils

import (
	"github.com/gin-gonic/gin"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func RespondError(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, APIError{Code: status, Message: message})
}

