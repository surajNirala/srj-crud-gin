package responses

import "github.com/gin-gonic/gin"

func ResponseError(code int, status string, message string, payload interface{}) gin.H {
	return gin.H{
		"code":    code,
		"status":  status,
		"message": message,
		"data":    payload,
	}
}

func ResponseSuccess(code int, status string, message string, payload interface{}) gin.H {
	return gin.H{
		"code":    code,
		"status":  status,
		"message": message,
		"data":    payload,
	}
}
