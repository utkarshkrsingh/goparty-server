// Package routes handles all the incoming request.
package routes

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/utkarshkrsingh/goparty/internal/db"
	"github.com/utkarshkrsingh/goparty/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func login(ctx *gin.Context) {
	// Get the email and password from the req body
	var body = db.Users{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, fmt.Sprintf("Unable to read body: %v", err.Error()))
		return
	}

	// Look for requested user
	user, err := body.FindByMail()
	if err != nil {
		utils.RespondError(ctx, http.StatusNotFound, fmt.Sprintf("User not found: %v", err.Error()))
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

func signup(ctx *gin.Context) {
	// Get the username, email and password from the request body
	var body = db.Users{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, fmt.Sprintf("Failed to read body: %v", err.Error()))
		return
	}
	fmt.Printf("Data: %v, %v and %v", body.UserName, body.Email, body.Password)

	if body.UserName == "" || body.Email == "" || body.Password == "" {
		utils.RespondError(ctx, http.StatusBadRequest, "Username, Email and Password are required")
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, fmt.Sprintf("Failed to hash password: %v", err.Error()))
		return
	}

	// Create the user with hashed password
	body.Password = string(hash)
	err = body.CreateUser()
	if err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, fmt.Sprintf("Failed to create user: %v", err.Error()))
		return
	}

	// Respond
	ctx.JSON(http.StatusOK, gin.H{
		"status": "signup successful",
	})
}

func logout(ctx *gin.Context) {
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

func validate(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		utils.RespondError(ctx, http.StatusNotFound, "User not exists")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User is authenticated",
		"user":    user,
	})
}
