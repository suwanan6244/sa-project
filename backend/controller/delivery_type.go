package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suwanan6244/sa-project/entity"
)

// POST /Deliverytypes
func CreateDeliverytype(c *gin.Context) {
	var Deliverytype entity.DeliveryType
	if err := c.ShouldBindJSON(&Deliverytype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Deliverytype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Deliverytype})
}

// GET /Deliverytype/:id
func GetDeliverytype(c *gin.Context) {
	var Deliverytype entity.DeliveryType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM delivery_types WHERE id = ?", id).Scan(&Deliverytype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Deliverytype})
}

// GET /Deliverytypes
func ListDeliverytypes(c *gin.Context) {
	var Deliverytypes []entity.DeliveryType
	if err := entity.DB().Raw("SELECT * FROM delivery_types").Scan(&Deliverytypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Deliverytypes})
}

// DELETE /Deliverytypes/:id
func DeleteDeliverytype(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM delivery_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Deliverytype not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Deliverytypes
func UpdateDeliverytype(c *gin.Context) {
	var Deliverytype entity.DeliveryType
	if err := c.ShouldBindJSON(&Deliverytype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Deliverytype.ID).First(&Deliverytype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Deliverytype not found"})
		return
	}

	if err := entity.DB().Save(&Deliverytype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Deliverytype})
}
