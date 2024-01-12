package middlware

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/srj-crud-gin/responses"
)

var (
	BlacklistedTokens    = make(map[string]bool)
	BlacklistedTokensMux sync.RWMutex
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		// fmt.Println("Token: ", tokenString)
		if tokenString == "" {
			res := responses.ResponseSuccess(401, "error", "Unauthorized.", nil)
			c.JSON(401, res)
			c.Abort()
			return
		}
		tokenString = tokenString[len("Bearer "):]
		// Check if the token is blacklisted
		BlacklistedTokensMux.RLock()
		defer BlacklistedTokensMux.RUnlock()
		if _, blacklisted := BlacklistedTokens[tokenString]; blacklisted {
			res := responses.ResponseSuccess(401, "error", "Token is blacklisted.", nil)
			c.JSON(401, res)
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil
		})
		if err != nil || !token.Valid {
			res := responses.ResponseSuccess(401, "error", "Invalid Token.", nil)
			c.JSON(401, res)
			c.Abort()
			return
		}
		fmt.Println("Error: ", err)
		c.Next()
	}
}
