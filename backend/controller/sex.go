package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suwanan6244/sa-project/entity"
)

// POST /sexs
func CreateSex(c *gin.Context) {
	var sex entity.Sex
	if err := c.ShouldBindJSON(&sex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&sex).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sex})
}

// GET /sex/:id
func GetSex(c *gin.Context) {
	var sex entity.Sex
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM sexs WHERE id = ?", id).Scan(&sex).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sex})
}

// GET /sexs
func ListSexs(c *gin.Context) {
	var sexs []entity.Sex
	if err := entity.DB().Raw("SELECT * FROM sexs").Scan(&sexs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sexs})
}

// DELETE /sexs/:id
func DeleteSex(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM sexs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sex not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /sexs
func UpdateSex(c *gin.Context) {
	var sex entity.Sex
	if err := c.ShouldBindJSON(&sex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", sex.ID).First(&sex); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sex not found"})
		return
	}

	if err := entity.DB().Save(&sex).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sex})
}
