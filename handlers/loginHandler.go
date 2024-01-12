package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/srj-crud-gin/config"
	"github.com/srj-crud-gin/models"
	"github.com/srj-crud-gin/responses"
)

func Login(c *gin.Context) {
	DB := config.DB
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		res := responses.ResponseError(500, "error", "Invalid Data Request.", nil)
		c.JSON(500, res)
		return
	}
	email := user.Email
	password := user.Password
	if !responses.ValidateEmail(email) {
		res := responses.ResponseError(400, "error", "Invalid Email Address.", nil)
		c.JSON(400, res)
		return
	}
	if !responses.ValidatePassword(password) {
		res := responses.ResponseError(400, "error", "Password must be least 8 characters.", nil)
		c.JSON(400, res)
		return
	}
	userDetail := DB.Where("email = ?", email).First(&user)
	if userDetail.Error != nil || userDetail.RowsAffected == 0 {
		res := responses.ResponseError(404, "error", "User not found.", nil)
		c.JSON(404, res)
		return
	}
	if !responses.WrongPassword(user.Password, password) {
		res := responses.ResponseError(400, "error", "Invalid Email/Password.", nil)
		c.JSON(400, res)
		return
	}
	res := responses.ResponseSuccess(200, "success", "Login Successfully.", user)
	c.JSON(200, res)
	return
}
