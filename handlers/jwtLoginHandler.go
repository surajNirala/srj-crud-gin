package handlers

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/srj-crud-gin/config"
	middlware "github.com/srj-crud-gin/middleware"
	"github.com/srj-crud-gin/models"
	"github.com/srj-crud-gin/responses"
)

var jwtKey = []byte("your-secret-key")

func JwtLogin(c *gin.Context) {
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
		c.Abort()
		return
	}
	if !responses.ValidatePassword(password) {
		res := responses.ResponseError(400, "error", "Password must be least 8 characters.", nil)
		c.JSON(400, res)
		c.Abort()
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
		c.Abort()
		return
	}
	// expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
	expirationTime := time.Now().Add(time.Minute * 3)
	claims := &models.Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		res := responses.ResponseError(500, "error", "Error Creating token.", nil)
		c.JSON(500, res)
		return
	}
	payloadToken := map[string]string{
		"token": signedToken,
	}
	res := responses.ResponseSuccess(200, "success", "Login Successfully.", payloadToken)
	c.JSON(200, res)
	return
}

func JwtLogout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		res := responses.ResponseError(400, "error", "Token not provided.", nil)
		c.JSON(400, res)
		c.Abort()
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	middlware.BlacklistedTokensMux.Lock()
	defer middlware.BlacklistedTokensMux.Unlock()
	middlware.BlacklistedTokens[tokenString] = true
	res := responses.ResponseSuccess(200, "success", "Logout Successfully.", nil)
	c.JSON(200, res)
	c.Abort()
	return
}
