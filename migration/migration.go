package migration

import (
	"github.com/srj-crud-gin/config"
	"github.com/srj-crud-gin/models"
)

func DatabaseUp() {
	DB := config.DB
	DB.AutoMigrate(&models.Todo{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
}
