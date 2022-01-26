package controller

import (
	"github.com/suwanan6244/sa-project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /suppliers
func CreateSupplier(c *gin.Context) {
	var supplier entity.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&supplier).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": supplier})
}

// GET /supplier/:id
func GetSupplier(c *gin.Context) {
	var supplier entity.Supplier
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM suppliers WHERE id = ?", id).Scan(&supplier).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": supplier})
}

// GET /suppliers
func ListSuppliers(c *gin.Context) {
	var suppliers []entity.Supplier
	if err := entity.DB().Raw("SELECT * FROM suppliers").Scan(&suppliers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": suppliers})
}

// DELETE /suppliers/:id
func DeleteSupplier(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM suppliers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "supplier not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /suppliers
func UpdateSupplier(c *gin.Context) {
	var supplier entity.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", supplier.ID).First(&supplier); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "supplier not found"})
		return
	}

	if err := entity.DB().Save(&supplier).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": supplier})
}

// ของตัวเอง
