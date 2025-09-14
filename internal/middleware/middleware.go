// Package middleware is for CORS
package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/utkarshkrsingh/goparty/internal/db"
	"github.com/utkarshkrsingh/goparty/internal/initializer"
)

func RequireAuth(ctx *gin.Context) {
	// Get the cookie of request
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil || !token.Valid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check exp
		exp, ok := claims["exp"].(float64)
		if !ok || float64(time.Now().Unix()) > exp {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Find the user with token sub
		var user db.Users
		dbQuery := `SELECT * FROM users WHERE id = $1`
		err := initializer.DB.Get(&user, dbQuery, claims["sub"])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "User not found: " + err.Error(),
			})
			return
		}

		if user.ID == "0" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach to req
		ctx.Set("user", user)

		// continue
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
