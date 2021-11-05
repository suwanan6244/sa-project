package controller

import (
	"net/http"

	"github.com/suwanan6244/sa-project/entity"

	"github.com/gin-gonic/gin"
)

// POST /accounts
func CreateAccount(c *gin.Context) {

	var account entity.Account
	var sex entity.Sex
	var olduser entity.Olduser
	var contact entity.Contact
	var religion entity.Religion

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// : ค้นหา sex ด้วย id
	if tx := entity.DB().Where("id = ?", account.SexID).First(&sex); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sex not found"})
		return
	}

	// : ค้นหา contact ด้วย id
	if tx := entity.DB().Where("id = ?", account.ContactID).First(&contact); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "contact not found"})
		return
	}

	// : ค้นหา olduser ด้วย id
	if tx := entity.DB().Where("id = ?", account.OlduserID).First(&olduser); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "olduser not found"})
		return
	}
	// : ค้นหา sex ด้วย id
	if tx := entity.DB().Where("id = ?", account.SexID).First(&sex); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sex not found"})
		return
	}

	// : ค้นหา religion ด้วย id
	if tx := entity.DB().Where("id = ?", account.ReligionID).First(&religion); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "religion not found"})
		return
	}
	// : สร้าง account
	wv := entity.Account{

		Address:  account.Address,  // ตั้งค่าฟิลด์ Address
		Province: account.Province, // ตั้งค่าฟิลด์ Province
		Sex:      sex,              // โยงความสัมพันธ์กับ Entity sex
		Contact:  contact,          // โยงความสัมพันธ์กับ Entity contact
		Olduser:  olduser,          // โยงความสัมพันธ์กับ Entity olduser
		Religion: religion,         // โยงความสัมพันธ์กับ Entity religion

	}

	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /account/:id
func GetAccount(c *gin.Context) {
	var account entity.Account
	id := c.Param("id")
	if err := entity.DB().Preload("Sex").Preload("Contact").Preload("Olduser").Preload("Religion").Raw("SELECT * FROM accounts WHERE id = ?", id).Find(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": account})
}

// GET /accounts
func ListAccounts(c *gin.Context) {
	var accounts []entity.Account
	if err := entity.DB().Preload("Sex").Preload("Contact").Preload("Olduser").Preload("Religion").Raw("SELECT * FROM accounts").Find(&accounts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": accounts})
}

// DELETE /accounts/:id
func DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM accounts WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /accounts
func UpdateAccount(c *gin.Context) {
	var account entity.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", account.ID).First(&account); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account not found"})
		return
	}

	if err := entity.DB().Save(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": account})
}
