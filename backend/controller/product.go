package controller

import (
	"github.com/suwanan6244/sa-project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /products
func CreateProduct(c *gin.Context) {
	var product entity.Product
	var producttype entity.ProductType

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 11: ค้นหา producttype ด้วย id
	if tx := entity.DB().Where("id = ?", product.ProductTypeID).First(&producttype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "producttype not found"})
		return
	}

	// : สร้าง producttype
	p := entity.Product{
		ProductType: producttype, // โยงความสัมพันธ์กับ Entity ProductType
	}

	if err := entity.DB().Create(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": p})
}

// GET /product/:id
func GetProduct(c *gin.Context) {
	var product entity.Product
	id := c.Param("id")
	if err := entity.DB().Preload("ProductType").Raw("SELECT * FROM products WHERE id = ?", id).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GET /products
func ListProducts(c *gin.Context) {
	var products []entity.Product
	if err := entity.DB().Preload("ProductType").Raw("SELECT * FROM products").Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// DELETE /products/:id
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM products WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /products
func UpdateProduct(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", product.ID).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	if err := entity.DB().Save(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}
