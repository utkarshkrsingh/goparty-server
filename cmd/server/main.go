package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/utkarshkrsingh/goparty/internal/initializer"
	"github.com/utkarshkrsingh/goparty/internal/middleware"
	routehandlers "github.com/utkarshkrsingh/goparty/internal/routeHandlers"
)

func init() {
	initializer.EnvVariables()
	initializer.ConnectDB()
}

func main() {
	router := gin.Default()
	defer initializer.DB.Close()

	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // frontend origin
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/create-room", middleware.RequireAuth, routehandlers.CreateRoom)
	router.POST("/signup", routehandlers.Signup)
	router.POST("/login", routehandlers.Login)
	router.POST("/logout", routehandlers.Logout)
	router.GET("/validate", middleware.RequireAuth, routehandlers.Validate)

	// Health check
	router.GET("/ping-db", func(ctx *gin.Context) {
		if err := initializer.DB.Ping(); err != nil {
			ctx.JSON(500, gin.H{"status": "db down", "error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"status": "db up"})
	})

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting Server on Port: %s", os.Getenv("APP_PORT"))
	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
