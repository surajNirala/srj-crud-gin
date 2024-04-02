package web

import (
	"github.com/gin-gonic/gin"
	"github.com/srj-crud-gin/config"
	"github.com/srj-crud-gin/models"
)

func UserList(c *gin.Context) {
	DB := config.DB
	var users []models.User
	DB.Find(&users)
	c.HTML(200, "users/list.html", gin.H{
		"Title": "User List",
		"Users": users,
	})
}
