package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/srj-crud-gin/handlers/web"
)

func WebRoutes(router *gin.Engine) {
	router.GET("/", web.Index)
	router.GET("user-list", web.UserList)
}
