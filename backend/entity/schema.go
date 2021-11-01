package entity

import (
	"time"

	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	Name          string
	Email         string `gorm:"uniqueIndex"`
	Password      string
	ProductStocks []ProductStock `gorm:"foreignKey:StaffID"`
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
