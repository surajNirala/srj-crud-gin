package main

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/srj-crud-gin/config"
	"github.com/srj-crud-gin/handlers"
	middlware "github.com/srj-crud-gin/middleware"
	"github.com/srj-crud-gin/models"
)

var (
	db                   *gorm.DB
	blacklistedTokens    = make(map[string]bool)
	blacklistedTokensMux sync.RWMutex
)

func main() {
	// Connect to the database
	config.InitDB()
	DB := config.DB

	DB.AutoMigrate(&models.Todo{})
	// Set up Gin router
	router := gin.Default()
	// Define API routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome",
		})
	})
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

	// Start the Gin server
	router.Run(":8081")
}

/* func initDB() {
	var err error
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_crud?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&Todo{})
} */
