package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	// 1 user เป็นเจ้าของได้หลาย order
	Orders []Order `gorm:"foreignKey:UserID"`
	// 1 user เป็นเจ้าของได้หลาย return
	Returns []Return `gorm:"foreignKey:OwnerID"`
}

type Order struct {
	gorm.Model
	// UserID ทำหน้าที่เป็น FK
	UserID     *uint
	User       User `gorm:"references:id"`
	PreorderID int
	StatusID   int
	Ordertime  time.Time
	Returns    []Return `gorm:"foreignKey:OrderID"`
}

type Staff struct {
	gorm.Model
	Name          string
	Email         string `gorm:"uniqueIndex"`
	Password      string
	ProductStocks []ProductStock `gorm:"foreignKey:StaffID"`
	Returns       []Return       `gorm:"foreignKey:StaffID"`
}

type Supplier struct {
	gorm.Model
	Name          string
	ProductStocks []ProductStock `gorm:"foreignKey:SupplierID"`
}

type ProductType struct {
	gorm.Model
	Ptype    string
	Products []Product `gorm:"foreignKey:ProductTypeID"`
}

type Product struct {
	gorm.Model
	Name string

	// ProductTypeID ทำหน้าที่เป็น FK
	ProductTypeID *uint
	ProductType   ProductType    `gorm:"references:id"`
	ProductStocks []ProductStock `gorm:"foreignKey:ProductID"`
}

type ProductStock struct {
	gorm.Model

	Price  int
	Amount int
	// ProductID ทำหน้าที่เป็น FK
	ProductID *uint
	Product   Product `gorm:"references:id"`

	// SupplierID  ทำหน้าที่เป็น FK
	SupplierID *uint
	Supplier   Supplier `gorm:"references:id"`

	// StaffID ทำหน้าที่เป็น FK
	StaffID *uint
	Staff   Staff `gorm:"references:id"`

	ProductTime time.Time
}

type Return struct {
	gorm.Model
	// OwnerID ทำหน้าที่เป็น FK
	OwnerID *uint
	Owner   User `gorm:"references:id"`

	// OderID ทำหน้าที่เป็น FK
	OrderID *uint
	Order   Order `gorm:"references:id"`

	// StaffID ทำหน้าที่เป็น FK
	StaffID *uint
	Staff   Staff `gorm:"references:id"`

	Reason     string
	Returndate time.Time
}
