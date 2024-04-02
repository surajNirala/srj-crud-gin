package handlers

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/srj-crud-gin/config"
	"github.com/srj-crud-gin/models"
)

func Index(c *gin.Context) {
	param := c.Query("type")
	paramSearch := c.Query("search")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	totalPages := GetTotalPages(pageSize)
	users, _ := FetchUsers(page, pageSize, paramSearch)
	if param != "" && param == "json" {
		c.JSON(200, gin.H{
			"Param": param,
			// "Title": "User List",
			"Users": users,
		})
		return
	}
	session := sessions.Default(c)
	success, _ := session.Get("success").(string)
	session.Delete("success")
	session.Save()
	c.HTML(200, "list.html", gin.H{
		"Param":       param,
		"Title":       "User List",
		"Users":       users,
		"TotalPage":   totalPages,
		"Page":        page,
		"PageSize":    pageSize,
		"Success":     success,
		"paramSearch": paramSearch,
		"PrevPage":    page - 1,
		"NextPage":    page + 1,
		"Pages":       generatePageNumbers(page, totalPages),
		// "DangerMessage": dangerMessage,
	})
}

// Function to generate page numbers for pagination
func generatePageNumbers(currentPage, totalPages int) []int {
	pages := make([]int, totalPages)
	for i := 0; i < totalPages; i++ {
		pages[i] = i + 1
	}
	return pages
}
func GetTotalPages(pageSize int) int {
	// DB := config.DB
	var count int64 = 100
	// DB.Model(&models.User{}).Count(&count)
	return int((count + int64(pageSize) - 1) / int64(pageSize))
}

func FetchUsers(page int, pageSize int, searchQuery string) ([]models.User, error) {
	DB := config.DB
	var users []models.User
	offset := (page - 1) * pageSize
	result := DB.Offset(offset).Limit(pageSize).Order("id desc")
	if searchQuery != "" {
		result = result.Where("first_name LIKE ? OR email LIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}
	result.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	// DB.Order("id desc").Find(&users)
	return users, nil
}

func AddUser(c *gin.Context) {
	session := sessions.Default(c)
	errors, ok := session.Get("errors").([]string)
	if !ok {
		errors = []string{} // Initialize empty errors if not found
	}
	// Clear errors from session
	session.Delete("errors")
	session.Save()
	c.HTML(200, "add-user.html", gin.H{
		"Title":  "Create User",
		"Errors": errors,
	})
}

func AddUserPost(c *gin.Context) {
	DB := config.DB
	var user models.User
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	phone := c.PostForm("phone")
	// formData := FormData{}
	errors := make([]string, 0)
	// var errors []string
	if first_name == "" {
		errors = append(errors, "First name cannot be empty")
	}
	if last_name == "" {
		errors = append(errors, "Last name cannot be empty")
	}
	if phone == "" {
		errors = append(errors, "Phone cannot be empty")
	}
	age := c.PostForm("age")
	if age == "" {
		errors = append(errors, "Age cannot be empty")
	}
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		errors = append(errors, "Invalid age string")
	}
	email := c.PostForm("email")
	exist := DB.Where("email = ?", email).First(&user)
	if exist.RowsAffected > 0 {
		errors = append(errors, "Email Already Exist.")
	}
	file, err := c.FormFile("file")
	if err != nil {
		errors = append(errors, "Error getting Image file")
	}
	if len(errors) > 0 {
		// Store errors in session
		session := sessions.Default(c)
		session.Set("errors", errors)
		session.Save()
		// Redirect to the form page
		c.Redirect(http.StatusSeeOther, "/add-user")
		return
	}
	// defer file.Close()
	user.FirstName = first_name
	user.LastName = last_name
	user.Email = email
	user.Age = uint(ageInt)
	user.Phone = phone
	uploadDir := "uploaded/files/"
	filename := generateUniqueFilename1(file.Filename)
	fullPath := filepath.Join(uploadDir, filename)
	// Save the file to the full path
	if err := c.SaveUploadedFile(file, fullPath); err != nil {
		c.JSON(500, gin.H{
			"error":   err.Error(),
			"code":    http.StatusInternalServerError,
			"message": "Error While  Saving File.",
		})
		return
	}
	user.Image = &filename
	result := DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error":   result.Error,
			"code":    http.StatusInternalServerError,
			"message": "Internal Server error.",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("success", "Record Created Successfully.")
	session.Save()
	c.Redirect(http.StatusSeeOther, "/")
	return
}

func generateUniqueFilename1(originalFilename string) string {
	// Get the current timestamp
	currentTime := time.Now()

	// Generate a random identifier (e.g., random number or hash)
	randomIdentifier := rand.Intn(1000) // Adjust the range as needed

	// Create a unique filename using a combination of timestamp and random identifier
	uniqueFilename := fmt.Sprintf("%d_%d_%s", currentTime.Unix(), randomIdentifier, originalFilename)

	// Optionally, sanitize the filename (replace spaces, special characters, etc.)
	// uniqueFilename = sanitizeFilename(uniqueFilename)

	return uniqueFilename
}

func EditUserInfo(c *gin.Context) {
	DB := config.DB
	var user models.User
	user_id := c.Param("user_id")
	// fmt.Println("user_id", user_id)
	result := DB.First(&user, user_id)
	if result.Error != nil {
		c.JSON(409, gin.H{
			"error":   result.Error,
			"code":    http.StatusConflict,
			"message": "User not found.",
		})
		return
	}
	session := sessions.Default(c)
	errors, ok := session.Get("errors").([]string)
	if !ok {
		errors = []string{} // Initialize empty errors if not found
	}
	// Clear errors from session
	session.Delete("errors")
	session.Save()
	c.HTML(200, "edit-user.html", gin.H{
		"Title":  "Update User",
		"User":   user,
		"Errors": errors,
	})
}

func DeletePost(c *gin.Context) {
	DB := config.DB
	var user models.User
	user_id := c.PostForm("user_id")
	// Delete with additional conditions
	result := DB.Where("id = ?", user_id).Delete(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error":   result.Error,
			"code":    http.StatusInternalServerError,
			"message": "Internal Server error.",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("success", "User Delete Successfully.")
	session.Save()
	c.Redirect(http.StatusSeeOther, "/")
}

// isEmpty checks if the uploaded file is empty
func isEmpty(file *multipart.FileHeader) bool {
	return file == nil || file.Size == 0
}
func UpdateUserInfo(c *gin.Context) {
	DB := config.DB
	var user models.User
	param_user_id := c.Param("user_id")
	user_id := c.PostForm("user_id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	phone := c.PostForm("phone")
	age := c.PostForm("age")
	email := c.PostForm("email")
	// formData := FormData{}
	errors := make([]string, 0)
	if first_name == "" {
		errors = append(errors, "First name cannot be empty")
	}
	if last_name == "" {
		errors = append(errors, "Last name cannot be empty")
	}
	if user_id == "" {
		errors = append(errors, "User Found Something wrong.")
	}
	if phone == "" {
		errors = append(errors, "Phone cannot be empty")
	}
	if age == "" {
		errors = append(errors, "Age cannot be empty")
	}
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		errors = append(errors, "Invalid age string")
	}
	found := DB.Where("id = ?", user_id).First(&user)
	if found.Error != nil {
		c.JSON(409, gin.H{
			"error":   found.Error,
			"code":    http.StatusConflict,
			"message": "User not found.",
		})
		return
	}
	exist := DB.Where("email = ?", email).First(&user)
	if exist.RowsAffected > 0 {
		user_idInt, _ := strconv.Atoi(user_id)
		if int(user.ID) != user_idInt {
			errors = append(errors, "Email Already Exist.")
		}
	}
	if len(errors) > 0 {
		// Store errors in session
		session := sessions.Default(c)
		session.Set("errors", errors)
		session.Save()
		// Redirect to the form page
		c.Redirect(http.StatusSeeOther, "/"+param_user_id+"/edit")
		return
	}
	// defer file.Close()
	user.FirstName = first_name
	user.LastName = last_name
	user.Email = email
	user.Age = uint(ageInt)
	user.Phone = phone
	file, _ := c.FormFile("file")
	// Check if the file is empty
	if !isEmpty(file) {
		uploadDir := "uploaded/files/"
		filename := generateUniqueFilename1(file.Filename)
		fullPath := filepath.Join(uploadDir, filename)
		// Save the file to the full path
		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			c.JSON(500, gin.H{
				"error":   err.Error(),
				"code":    http.StatusInternalServerError,
				"message": "Error While  Saving File.",
			})
			return
		}
		user.Image = &filename
	}
	result := DB.Save(&user)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"error":   result.Error,
			"code":    http.StatusInternalServerError,
			"message": "Internal Server error.",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("success", "Record Updated Successfully.")
	session.Save()
	c.Redirect(http.StatusSeeOther, "/")
	// return
}
