package routehandlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/utkarshkrsingh/goparty/internal/db"
	"github.com/utkarshkrsingh/goparty/internal/initializer"
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
	var user db.Users
	dbQuery := `SELECT * FROM users WHERE email = $1`
	if err := initializer.DB.Get(&user, dbQuery, body.Email); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found: " + err.Error(),
		})
		return
	}

	// Compare sent in password with saved user password hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password: " + err.Error(),
		})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"username": user.UserName,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret-key
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server misconfigured",
		})
		return
	}

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create token: " + err.Error(),
		})
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
