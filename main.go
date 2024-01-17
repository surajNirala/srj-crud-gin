package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/srj-crud-gin/config"
	"github.com/srj-crud-gin/migration"
	"github.com/srj-crud-gin/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	//gin-routes-snippet
	config.InitDB()
	migration.DatabaseUp()
	router := gin.Default()
	// Set the HTML templates folder
	router.LoadHTMLGlob("templates/users/*")
	routes.WebRoutes(router)
	routes.ApiRoutes(router)
	// Start the Gin server
	server_port := os.Getenv("SERVER_PORT")
	router.Run(server_port)
}

/* func initDB() {
	var err error
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_crud?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&Todo{})
} */
