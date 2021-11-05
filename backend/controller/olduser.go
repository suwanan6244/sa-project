package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suwanan6244/sa-project/entity"
)

// POST /oldusers
func CreateOlduser(c *gin.Context) {
	var olduser entity.Olduser
	if err := c.ShouldBindJSON(&olduser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&olduser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": olduser})
}

// GET /olduser/:id
func GetOlduser(c *gin.Context) {
	var olduser entity.Olduser
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM oldusers WHERE id = ?", id).Scan(&olduser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": olduser})
}

// GET /oldusers
func ListOldusers(c *gin.Context) {
	var oldusers []entity.Olduser
	if err := entity.DB().Raw("SELECT * FROM oldusers").Scan(&oldusers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": oldusers})
}

// DELETE /oldusers/:id
func DeleteOlduser(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM oldusers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "olduser not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /oldusers
func UpdateOlduser(c *gin.Context) {
	var olduser entity.Olduser
	if err := c.ShouldBindJSON(&olduser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", olduser.ID).First(&olduser); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "olduser not found"})
		return
	}

	if err := entity.DB().Save(&olduser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": olduser})
}
