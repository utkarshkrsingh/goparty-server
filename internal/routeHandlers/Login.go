package routehandlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/utkarshkrsingh/goparty/internal/initializer"
	"github.com/utkarshkrsingh/goparty/internal/repositories"
	"github.com/utkarshkrsingh/goparty/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	// Get the email and password from the req body
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body: " + err.Error(),
		})
		return
	}

	// Look for requested user
	repositoryManager := repositories.PostgresUserRepository{DB: initializer.DB}
	user, err := repositoryManager.FindByEmail(body.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found: " + err.Error(),
		})
		return
	}

	// Compare sent in password with saved user password hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		utils.RespondError(ctx, http.StatusUnauthorized, fmt.Sprintf("Invalid email or password: %v", err.Error()))
		return
	}

	// Generate a JWT token
	jwtManager := utils.NewJWTManager(os.Getenv("JWT_SECRET"), time.Hour*24*30)
	tokenString, err := jwtManager.GenerateToken(user.ID, user.UserName, user.Email)
	if err != nil {
		utils.RespondError(ctx, http.StatusInternalServerError, fmt.Sprintf("Failed to create token: %v", err.Error()))
		return
	}

	// Send it back
	secure := gin.Mode() == gin.ReleaseMode
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "/", "", secure, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login Successful",
		"user": gin.H{
			"id":       user.ID,
			"username": user.UserName,
			"email":    user.Email,
		},
	})
}
