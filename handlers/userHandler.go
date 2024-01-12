package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/srj-crud-gin/config"
	"github.com/srj-crud-gin/models"
	"github.com/srj-crud-gin/responses"
	"github.com/srj-crud-gin/transforms"
)

func GetAllUser(c *gin.Context) {
	DB := config.DB
	var users []models.User
	DB.Find(&users)
	var transformedUsers []transforms.UserResponse
	for _, user := range users {
		transformedUsers = append(transformedUsers, transforms.TransformUser(user))
	}
	res := responses.ResponseSuccess(200, "success", "User list.", transformedUsers)
	c.JSON(200, res)
}

func CreateUser(c *gin.Context) {
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
	hashedPassword, err := responses.HashPassword(user.Password)
	if err != nil {
		res := responses.ResponseError(400, "error", "Error hasing password.", nil)
		c.JSON(400, res)
		return
	}
	// Check for duplicate email
	if !responses.IsEmailUnique(DB, user.Email) {
		res := responses.ResponseError(400, "error", "Email already exists.", nil)
		c.JSON(400, res)
		return
	}
	user.Password = hashedPassword
	result := DB.Create(&user)
	if result.Error != nil {
		res := responses.ResponseError(400, "error", "Internal Server error.", nil)
		c.JSON(400, res)
		return
	}
	res := responses.ResponseSuccess(200, "success", "User created successfully", &user)
	c.JSON(200, res)
}

func GetUser(c *gin.Context) {
	DB := config.DB
	var user models.User
	id := c.Params.ByName("id")
	userDetail := DB.Where("id = ?", id).First(&user)
	if userDetail.Error != nil || userDetail.RowsAffected == 0 {
		res := responses.ResponseError(404, "error", "User not found.", nil)
		c.JSON(404, res)
		return
	}
	data := transforms.TransformUser(user)
	res := responses.ResponseError(200, "error", "Fetch User Details.", data)
	c.JSON(200, res)
}

func UpdateUser(c *gin.Context) {
	DB := config.DB
	var user models.User
	var updatedUser models.User
	userID := c.Params.ByName("id")
	userDetail := DB.Find(&user, userID)
	if userDetail.Error != nil || userDetail.RowsAffected == 0 {
		res := responses.ResponseError(404, "error", "User not found.", nil)
		c.JSON(404, res)
		return
	}
	err := c.ShouldBindJSON(&updatedUser)
	if err != nil {
		res := responses.ResponseError(500, "error", "Invalid Data Request.", nil)
		c.JSON(500, res)
		return
	}
	var existingUser models.User
	if DB.Where("email = ? AND id != ?", updatedUser.Email, userID).First(&existingUser).RowsAffected > 0 {
		res := responses.ResponseError(400, "error", "Email already exists.", nil)
		c.JSON(400, res)
		return
	}
	// Update user details
	user.FirstName = updatedUser.FirstName
	user.LastName = updatedUser.LastName
	user.Email = updatedUser.Email
	// Save the updated user to the database
	DB.Save(&user)
	data := transforms.TransformUser(user)
	res := responses.ResponseSuccess(200, "success", "User Updated successfully", data)
	c.JSON(200, res)
}

func DeleteUser(c *gin.Context) {
	DB := config.DB
	var user models.User
	user_id := c.Params.ByName("id")
	userDetail := DB.Delete(&user, user_id)
	if userDetail.Error != nil || userDetail.RowsAffected == 0 {
		res := responses.ResponseError(404, "error", "User not found.", nil)
		c.JSON(404, res)
		return
	}
	res := responses.ResponseSuccess(200, "success", "User Delete successfully", nil)
	c.JSON(200, res)
}
