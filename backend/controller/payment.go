package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suwanan6244/sa-project/entity"
)

// POST /payments
func CreatePayment(c *gin.Context) {

	var Payment entity.Payment
	var PaymentMethod entity.PaymentMethod
	var DeliveryType entity.DeliveryType
	var Order entity.Order

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Payment
	if err := c.ShouldBindJSON(&Payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เช็คการบันทึก payment ถ้ามีการชำระเงินหมายเลขคำสั่งซื้อที่ซ้ำกัน ให้ return การชำระเงินนั้นออกไป
	if tx := entity.DB().Table("payments").Where("order_id = ?", Payment.OrderID).First(&Payment); tx.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"duplicatepayment": Payment})
		return
	}

	// 9: ค้นหา Order ด้วย id
	if tx := entity.DB().Where("id = ?", Payment.OrderID).First(&Order); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	// 10: ค้นหา PaymentMethod ด้วย id
	if tx := entity.DB().Where("id = ?", Payment.PaymentMethodID).First(&PaymentMethod); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PaymentMethod not found"})
		return
	}

	// 11: ค้นหา DeliveryType ด้วย id
	if tx := entity.DB().Where("id = ?", Payment.DeliveryTypeID).First(&DeliveryType); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DeliveryType not found"})
		return
	}
	// 12: สร้าง Payment
	pm := entity.Payment{
		PaymentMethod: PaymentMethod,       // โยงความสัมพันธ์กับ Entity PaymentMethod
		Order:         Order,               // โยงความสัมพันธ์กับ Entity Order
		DeliveryType:  DeliveryType,        // โยงความสัมพันธ์กับ Entity DeliveryType
		PaymentTime:   Payment.PaymentTime, // ตั้งค่าฟิลด์ PaymentTime
		Phone:         Payment.Phone,
	}

	// 13: บันทึก
	if err := entity.DB().Create(&pm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pm})
}

// GET /Payment/:id
func GetPayment(c *gin.Context) {
	var Payment entity.Payment
	id := c.Param("id")
	if err := entity.DB().Preload("PaymentMethod").Preload("DeliveryType").Preload("Order").Raw("SELECT * FROM payments WHERE id = ?", id).Find(&Payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Payment})
}

// GET /payments
func ListPayments(c *gin.Context) {
	var Payments []entity.Payment
	if err := entity.DB().Preload("PaymentMethod").Preload("DeliveryType").Preload("Order").Raw("SELECT * FROM payments").Find(&Payments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Payments})
}

// DELETE /payments/:id
func DeletePayment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM payments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /payments
func UpdatePayment(c *gin.Context) {
	var Payment entity.Payment
	if err := c.ShouldBindJSON(&Payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Payment.ID).First(&Payment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payment not found"})
		return
	}

	if err := entity.DB().Save(&Payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Payment})
}
