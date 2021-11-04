package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suwanan6244/sa-project/entity"
)

func CreateReturnod(c *gin.Context) {

	var returns entity.Return
	var user entity.User
	var order entity.Order
	var staff entity.Staff

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร returnod
	if err := c.ShouldBindJSON(&returns); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", returns.OwnerID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// 10: ค้นหา order ด้วย id
	if tx := entity.DB().Where("id = ?", returns.OrderID).First(&order); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})
		return
	}

	// 11: ค้นหา staff ด้วย id
	if tx := entity.DB().Where("id = ?", returns.StaffID).First(&staff); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "staff not found"})
		return
	}
	// 12: สร้าง Returnod
	rd := entity.Return{
		Owner:      user,               // โยงความสัมพันธ์กับ Entity User
		Order:      order,              // โยงความสัมพันธ์กับ Entity Order
		Staff:      staff,              // โยงความสัมพันธ์กับ Entity Staff
		Reason:     returns.Reason,     // ตั้งค่าฟิลด์ Reason
		Returndate: returns.Returndate, // ตั้งค่าฟิลด์ Returndate
	}

	// 13: บันทึก
	if err := entity.DB().Create(&rd).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rd})
}

// GET /returnod/:id
func GetReturn(c *gin.Context) {
	var returns entity.Return
	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Preload("Order").Preload("Staff").Raw("SELECT * FROM returns WHERE id = ?", id).Find(&returns).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": returns})
}

// GET /return_ods
func ListReturn(c *gin.Context) {
	var returns []entity.Return
	if err := entity.DB().Preload("Owner").Preload("Order").Preload("Staff").Raw("SELECT * FROM returns ").Find(&returns).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": returns})
}

// GET /return_ods
func ListReturns(c *gin.Context) {
	var returns []entity.Return
	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Preload("Order").Preload("Staff").Raw("SELECT * FROM returns where owner_id = ? ", id).Find(&returns).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": returns})
}

// DELETE /return_ods/:id
func DeleteReturn(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM returns WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "return not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /return_ods
func UpdateReturn(c *gin.Context) {
	var returnod entity.Return
	if err := c.ShouldBindJSON(&returnod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", returnod.ID).First(&returnod); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "return not found"})
		return
	}

	if err := entity.DB().Save(&returnod).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": returnod})
}
