package services

import (
	"github.com/gin-gonic/gin"
	"github.com/utkarshkrsingh/goparty/internal/db"
)

type AuthService interface {
	Login(ctx *gin.Context, email, password string) (string, *db.Users, error)
}