package controller

import (
	"net/http"

	"github.com/suwanan6244/sa-project/entity"

	"github.com/gin-gonic/gin"
)

// POST /contacts
func CreateContact(c *gin.Context) {
	var contact entity.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": contact})
}

// GET /contact/:id
func GetContact(c *gin.Context) {
	var contact entity.Contact

	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM contacts WHERE id = ?", id).Find(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contact})
}

// GET /contacts
func ListContacts(c *gin.Context) {
	var contacts []entity.Contact
	if err := entity.DB().Raw("SELECT * FROM contacts").Scan(&contacts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contacts})
}

// DELETE /contacts/:id
func DeleteContact(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM contacts WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "contact not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /contacts
func UpdateContact(c *gin.Context) {
	var contact entity.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", contact.ID).First(&contact); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "contact not found"})
		return
	}

	if err := entity.DB().Save(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contact})
}
