package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suwanan6244/sa-project/entity"
)

// POST /preorders
func CreatePreorder(c *gin.Context) {

	var user entity.User
	var preorder entity.Preorder
	var product entity.Product
	var paymentmethod entity.PaymentMethod

	if err := c.ShouldBindJSON(&preorder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", preorder.UserID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", preorder.ProductID).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", preorder.PaymentMethodID).First(&paymentmethod); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paymentmethod not found"})
		return
	}
	// 12: สร้าง Preorder
	pre := entity.Preorder{
		User:          user,
		Product:       product,         // โยงความสัมพันธ์กับ Entity Product
		PaymentMethod: paymentmethod,   // โยงความสัมพันธ์กับ Entity Paymentmethod
		Amount:        preorder.Amount, // ตั้งค่าฟิลด์ Amount
	}

	// 13: บันทึก
	if err := entity.DB().Create(&pre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pre})
}

// GET /preorder/:id
func GetPreorder(c *gin.Context) {
	var preorder entity.Preorder
	id := c.Param("id")
	if err := entity.DB().Preload("User").Preload("Product").Preload("Paymentmethod").Raw("SELECT * FROM preorders WHERE id = ?", id).Find(&preorder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": preorder})
}

// GET /preorders
func ListPreorders(c *gin.Context) {
	var preorders []entity.Preorder
	if err := entity.DB().Preload("User").Preload("Product").Preload("Paymentmethod").Raw("SELECT * FROM preorders").Find(&preorders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": preorders})
}

// DELETE /preorders/:id
func DeletePreorder(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM preorders WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "preorder not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /preorders
func UpdatePreorder(c *gin.Context) {
	var preorder entity.Preorder
	if err := c.ShouldBindJSON(&preorder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", preorder.ID).First(&preorder); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "preorder not found"})
		return
	}

	if err := entity.DB().Save(&preorder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": preorder})
}
