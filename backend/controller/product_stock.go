package controller

import (
	"github.com/suwanan6244/sa-project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /product_stocks
func CreateProductStock(c *gin.Context) {

	var productstock entity.ProductStock
	var product entity.Product
	var supplier entity.Supplier
	var staff entity.Staff

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร productstock
	if err := c.ShouldBindJSON(&productstock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// : ค้นหา product ด้วย id
	if tx := entity.DB().Where("id = ?", productstock.ProductID).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	// : ค้นหา supplier ด้วย id
	if tx := entity.DB().Where("id = ?", productstock.SupplierID).First(&supplier); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "supplier not found"})
		return
	}

	// : ค้นหา staff ด้วย id
	if tx := entity.DB().Where("id = ?", productstock.StaffID).First(&staff); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "staff not found"})
		return
	}

	// : สร้าง productstock
	ps := entity.ProductStock{
		Product:     product,                  // โยงความสัมพันธ์กับ Entity product
		Supplier:    supplier,                 // โยงความสัมพันธ์กับ Entity Supplier
		Staff:       staff,                    // โยงความสัมพันธ์กับ Entity Staff
		ProductTime: productstock.ProductTime, // ตั้งค่าฟิลด์ productTime
		Price:       productstock.Price,       // ตั้งค่าฟิลด์ Price
		Amount:      productstock.Amount,      // ตั้งค่าฟิลด์ Amount
	}

	// : บันทึก
	if err := entity.DB().Create(&ps).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ps})
}

// GET /productstock/:id
func GetProductStock(c *gin.Context) {
	var productstock entity.ProductStock
	id := c.Param("id")
	if err := entity.DB().Preload("Product").Preload("Supplier").Preload("Staff").Raw("SELECT * FROM product_stocks WHERE id = ?", id).Find(&productstock).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": productstock})
}

// GET /product_stocks
func ListProductStocks(c *gin.Context) {
	var productstocks []entity.ProductStock
	if err := entity.DB().Preload("Product").Preload("Supplier").Preload("Staff").Raw("SELECT * FROM product_stocks").Find(&productstocks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": productstocks})
}

// DELETE /product_stocks/:id
func DeleteProductStock(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM product_stocks WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "productstock not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /product_stocks
func UpdateProductStock(c *gin.Context) {
	var productstock entity.ProductStock
	if err := c.ShouldBindJSON(&productstock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", productstock.ID).First(&productstock); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "productstock not found"})
		return
	}
	if err := entity.DB().Save(&productstock).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": productstock})
}
