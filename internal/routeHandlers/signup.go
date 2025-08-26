package routehandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utkarshkrsingh/goparty/internal/initializer"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	// Get the username, email and password from the request body
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body: " + err.Error(),
		})
		return
	}

	if body.Username == "" || body.Email == "" || body.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Username, Email and Password are required",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password: " + err.Error(),
		})
		return
	}

	// Create the user with hashed password
	body.Password = string(hash)
	dbQuery := `INSERT INTO users (username, email, password_hash)
    VALUES ($1, $2, $3)`
	_, err = initializer.DB.Exec(dbQuery, body.Username, body.Email, body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user: " + err.Error(),
		})
		return
	}

	// Respond
	ctx.JSON(http.StatusOK, gin.H{
		"status": "signup successful",
	})
}
