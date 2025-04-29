package main

import (
	"user-service/database"
	"user-service/handlers"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggingMiddleware())

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	user := r.Group("/api/user", middleware.AuthMiddleware())
	{
		user.GET("/profile", handlers.GetProfile)
		user.PUT("/profile", handlers.UpdateProfile)
	}

	r.GET("/api/users/:id", middleware.AuthMiddleware(), handlers.GetUserByID)

	admin := r.Group("/api/users", middleware.AuthMiddleware(), middleware.RequireRole("admin"))
	{
		admin.GET("/", handlers.GetAllUsers)
		admin.POST("/", handlers.CreateUser)
		admin.PUT("/:id", handlers.UpdateUser)
		admin.DELETE("/:id", handlers.DeleteUser)
	}

	r.Run(":8001")
}
