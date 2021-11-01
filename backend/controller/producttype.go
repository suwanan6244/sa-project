package controller

import (
	"github.com/suwanan6244/sa-project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /producttypes
func CreateProductType(c *gin.Context) {
	var producttype entity.ProductType
	if err := c.ShouldBindJSON(&producttype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&producttype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": producttype})
}

// GET /producttype/:id
func GetProductType(c *gin.Context) {
	var producttype entity.ProductType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM product_types WHERE id = ?", id).Scan(&producttype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": producttype})
}

// GET /producttypes
func ListProductTypes(c *gin.Context) {
	var producttypes []entity.ProductType
	if err := entity.DB().Raw("SELECT * FROM product_types").Scan(&producttypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": producttypes})
}

// DELETE /producttypes/:id
func DeleteProductType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM product_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "producttype not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /producttypes
func UpdateProductType(c *gin.Context) {
	var producttype entity.ProductType
	if err := c.ShouldBindJSON(&producttype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", producttype.ID).First(&producttype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "producttype not found"})
		return
	}

	if err := entity.DB().Save(&producttype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": producttype})
}
