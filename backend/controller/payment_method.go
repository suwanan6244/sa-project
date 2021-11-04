package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suwanan6244/sa-project/entity"
)

// POST /Paymentmethods
func CreatePaymentmethod(c *gin.Context) {
	var Paymentmethod entity.PaymentMethod
	if err := c.ShouldBindJSON(&Paymentmethod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Paymentmethod).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Paymentmethod})
}

// GET /Paymentmethod/:id
func GetPaymentmethod(c *gin.Context) {
	var Paymentmethod entity.PaymentMethod
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM payment_methods WHERE id = ?", id).Scan(&Paymentmethod).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Paymentmethod})
}

// GET /Paymentmethods
func ListPaymentmethods(c *gin.Context) {
	var Paymentmethods []entity.PaymentMethod
	if err := entity.DB().Raw("SELECT * FROM payment_methods").Scan(&Paymentmethods).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Paymentmethods})
}

// DELETE /Paymentmethods/:id
func DeletePaymentmethod(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM payment_methods WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Paymentmethod not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Paymentmethods
func UpdatePaymentmethod(c *gin.Context) {
	var Paymentmethod entity.PaymentMethod
	if err := c.ShouldBindJSON(&Paymentmethod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Paymentmethod.ID).First(&Paymentmethod); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Paymentmethod not found"})
		return
	}

	if err := entity.DB().Save(&Paymentmethod).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Paymentmethod})
}
