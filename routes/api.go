package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/srj-crud-gin/handlers"
	middlware "github.com/srj-crud-gin/middleware"
)

func ApiRoutes(router *gin.Engine) {
	router.GET("/todos", handlers.GetTodos)
	router.GET("/todos-deleted", handlers.GetTodosDeleted)
	router.GET("/todos/:id", handlers.GetTodo)
	router.POST("/todos", handlers.CreateTodo)
	router.PUT("/todos/:id", handlers.UpdateTodo)
	router.DELETE("/todos/:id", handlers.DeleteTodo)

	router.GET("/users", handlers.GetAllUser)
	router.POST("/users", handlers.CreateUser)
	router.GET("/users/:id", handlers.GetUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)

	/*****************Start Login Handler****************/
	router.POST("/login", handlers.Login)
	/*****************End Login Handler****************/

	/*****************Start JWT Login Handler****************/
	router.POST("/jwt-login", handlers.JwtLogin)
	// Logout route
	router.POST("/logout", handlers.JwtLogout)
	/*****************End JWT Login Handler****************/

	/*****************Start Product Handler****************/
	authGroup := router.Group("/auth")
	authGroup.Use(middlware.AuthMiddleware())
	authGroup.GET("/products", handlers.GetAllProduct)
	/*****************End Product Handler****************/
}
