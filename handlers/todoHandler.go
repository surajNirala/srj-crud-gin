package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/srj-crud-gin/config"
	"github.com/srj-crud-gin/models"
)

func GetTodos(c *gin.Context) {
	DB := config.DB
	var todos []models.Todo
	DB.Find(&todos)
	c.JSON(200, todos)
}

func GetTodosDeleted(c *gin.Context) {
	DB := config.DB
	var todos []models.Todo
	DB.Unscoped().Find(&todos)
	c.JSON(200, todos)
}

func GetTodo(c *gin.Context) {
	DB := config.DB
	id := c.Params.ByName("id")
	var todo models.Todo
	if err := DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, todo)
}

func CreateTodo(c *gin.Context) {
	DB := config.DB
	var todo models.Todo
	c.BindJSON(&todo)
	DB.Create(&todo)
	c.JSON(200, todo)
}

func UpdateTodo(c *gin.Context) {
	DB := config.DB
	id := c.Params.ByName("id")
	var todo models.Todo
	if err := DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.BindJSON(&todo)
	DB.Save(&todo)
	c.JSON(200, todo)
}

func DeleteTodo(c *gin.Context) {
	DB := config.DB
	id := c.Params.ByName("id")
	var todo models.Todo
	d := DB.Where("id = ?", id).Delete(&todo)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
