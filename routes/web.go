package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/srj-crud-gin/handlers"
)

func WebRoutes(router *gin.Engine) {
	// router.GET("/", web.Index)
	// router.GET("user-list", web.UserList)

	router.GET("/", handlers.Index)
	// router.GET("/user-login", handlers.UserLogin)
	// router.POST("/user-login-post", handlers.UserLoginPost)

	router.GET("/add-user", handlers.AddUser)
	router.POST("/add-user-post", handlers.AddUserPost)
	router.GET("/:user_id/edit", handlers.EditUserInfo)
	router.POST("/:user_id/update-user", handlers.UpdateUserInfo)

	router.POST("/delete", handlers.DeletePost)
	// router.GET("/generate-pdf", handlers.GeneratePDF)
	// router.GET("/generate-html-to-pdf", handlers.GenerateHTMLTOPDF)
	// router.GET("/fake-store", handlers.FakeStore)
	// router.GET("/export-excel", handlers.ExportExcel)
	// router.POST("/import-csv", handlers.ImportCSV)

	// router.GET("/{user_id}/send-email", handlers.SendEmail)
	// router.GET("/user-logout", handlers.UserLogout)
}
