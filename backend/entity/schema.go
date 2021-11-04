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

	Orders    []Order    `gorm:"foreignKey:UserID"`
	Returns   []Return   `gorm:"foreignKey:OwnerID"`
	Preorders []Preorder `gorm:"foreignKey:UserID"`
	Accounts  []Account  `gorm:"foreignKey:OwnerID"`
}

type Staff struct {
	gorm.Model
	Name          string
	Email         string `gorm:"uniqueIndex"`
	Password      string
	ProductStocks []ProductStock `gorm:"foreignKey:StaffID"`
	Returns       []Return       `gorm:"foreignKey:StaffID"`
}

type Contact struct {
	gorm.Model
	Ctype    string
	Accounts []Account `gorm:"foreignKey:ContactID"`
}
type Sex struct {
	gorm.Model
	Stype    string
	Accounts []Account `gorm:"foreignKey:SexID"`
}

type Olduser struct {
	gorm.Model
	Otype    string
	Accounts []Account `gorm:"foreignKey:OlduserID"`
}
type Religion struct {
	gorm.Model
	Rtype    string
	Accounts []Account `gorm:"foreignKey:ReligionID"`
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

	ProductTypeID *uint
	ProductType   ProductType    `gorm:"references:id"`
	ProductStocks []ProductStock `gorm:"foreignKey:ProductID"`
}

type Status struct {
	gorm.Model
	Statusorder string

	Orders []Order `gorm:"foreignKey:StatusID"`
}

type PaymentMethod struct {
	gorm.Model
	Method  string
	Payment []Payment `gorm:"foreignKey:PaymentMethodID"`
}

type DeliveryType struct {
	gorm.Model
	Type    string
	Payment []Payment `gorm:"foreignKey:DeliveryTypeID"`
}

//ระบบย่อย ระบบสมาชิก
type Account struct {
	gorm.Model
	Address    string
	Province   string
	OwnerID    *uint
	Owner      User `gorm:"references:id"`
	ContactID  *uint
	Contact    Contact `gorm:"references:id"`
	SexID      *uint
	Sex        Sex `gorm:"references:id"`
	OlduserID  *uint
	Olduser    Olduser `gorm:"references:id"`
	ReligionID *uint
	Religion   Religion `gorm:"references:id"`
}

//ระบบย่อย ระบบสั่งจองสินค้า
type Preorder struct {
	gorm.Model
	Amount int

	UserID *uint
	User   User `gorm:"references:id"`

	ProductID *uint
	Product   Product `gorm:"references:id"`

	PaymentMethodID *uint
	PaymentMethod   PaymentMethod `gorm:"references:id"`
}

//ระบบย่อย ระบบสั่งซื้อสินค้า
type Order struct {
	gorm.Model
	OrderTime time.Time

	UserID *uint
	User   User `gorm:"references:id"`

	PreorderID *uint
	Preorder   Preorder `gorm:"references:id"`

	StatusID *uint
	Status   Status `gorm:"references:id"`

	Payment []Payment `gorm:"foreignKey:OrderID"`
	Returns []Return  `gorm:"foreignKey:OrderID"`
}

//ระบบย่อย ระบบคลังสินค้า
type ProductStock struct {
	gorm.Model

	ProductID *uint
	Product   Product `gorm:"references:id"`

	Price  int
	Amount int

	SupplierID *uint
	Supplier   Supplier `gorm:"references:id"`

	StaffID *uint
	Staff   Staff `gorm:"references:id"`

	ProductTime time.Time
}

//ระบบย่อย ระบบขอคืนสินค้า
type Return struct {
	gorm.Model
	OwnerID *uint
	Owner   User `gorm:"references:id"`

	OrderID *uint
	Order   Order `gorm:"references:id"`

	StaffID *uint
	Staff   Staff `gorm:"references:id"`

	Reason     string
	Returndate time.Time
}

//ระบบย่อย ระบบการชำระเงิน
type Payment struct {
	gorm.Model
	Phone       string
	Price       float32
	PaymentTime time.Time

	OrderID *uint
	Order   Order `gorm:"references:id"`

	PaymentMethodID *uint
	PaymentMethod   PaymentMethod `gorm:"references:id"`

	DeliveryTypeID *uint
	DeliveryType   DeliveryType `gorm:"references:id"`
}
