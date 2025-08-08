package routes

import (
	"net/http"

	"github.com/ByGultekin-tak/elelden-backend/internal/config"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config) {
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Register endpoint - Coming soon"})
			})
			auth.POST("/login", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Login endpoint - Coming soon"})
			})
		}

		// User routes
		users := v1.Group("/users")
		{
			users.GET("/profile", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "User profile endpoint - Coming soon"})
			})
		}

		// Listings routes
		listings := v1.Group("/listings")
		{
			listings.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Get listings endpoint - Coming soon"})
			})
			listings.POST("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Create listing endpoint - Coming soon"})
			})
		}

		// Categories routes
		categories := v1.Group("/categories")
		{
			categories.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Get categories endpoint - Coming soon"})
			})
		}
	}
}
