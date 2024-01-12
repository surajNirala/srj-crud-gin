package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/srj-crud-gin/config"
	"github.com/srj-crud-gin/models"
	"github.com/srj-crud-gin/responses"
	"github.com/srj-crud-gin/transforms"
)

func GetAllProduct(c *gin.Context) {
	DB := config.DB
	var products []models.Product
	_ = DB.Preload("User").Order("created_at desc").Find(&products)
	var transformProducts []transforms.ProductResponse
	for _, product := range products {
		transformProducts = append(transformProducts, transforms.TransformProduct(product))
	}
	// responses.ResponseSuccess(w, http.StatusOK, "Fetched Product List.", transformProducts)
	res := responses.ResponseSuccess(200, "success", "Product list.", transformProducts)
	c.JSON(200, res)
}
