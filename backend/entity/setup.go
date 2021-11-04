package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(

		// สมาชิก
		&User{},
		&Olduser{},
		&Contact{},
		&Sex{},
		&Religion{},
		&Account{},

		// ระบบสั่งจองสินค้า
		&Preorder{},

		// ระบบสั่งสินค้า
		&Order{},
		&Status{},

		// ระบบจ่ายเงิน
		&PaymentMethod{},
		&Payment{},
		&DeliveryType{},

		// ระบบขอคืนสินค้า
		&Return{},

		// ระบบคลังสินค้า
		&Staff{},
		&Product{},
		&ProductType{},
		&Supplier{},
		&ProductStock{},
	)
	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	// ข้อมูล user
	db.Model(&User{}).Create(&User{
		Name:     "Narudee",
		Email:    "narudee@gmail.com",
		Password: string(password),
	})

	db.Model(&User{}).Create(&User{
		Name:     "Phatcha",
		Email:    "phatcha@gmail.com",
		Password: string(password),
	})

	db.Model(&User{}).Create(&User{
		Name:     "Patnarin",
		Email:    "patnarin@gmail.com",
		Password: string(password),
	})

	db.Model(&User{}).Create(&User{
		Name:     "name",
		Email:    "name@example.com",
		Password: string(password),
	})

	var narudee User
	var phatcha User
	var patnarin User
	var name User
	db.Raw("SELECT * FROM users WHERE email = ?", "narudee@gmail.com").Scan(&narudee)
	db.Raw("SELECT * FROM users WHERE email = ?", "phatcha@gmail.com").Scan(&phatcha)
	db.Raw("SELECT * FROM users WHERE email = ?", "patnarin@gmail.com").Scan(&patnarin)
	db.Raw("SELECT * FROM users WHERE email = ?", "name@example.com").Scan(&name)

	// ระบบชำระเงิน
	// PaymentMethod Data
	Method1 := PaymentMethod{
		Method: "Bank",
	}
	db.Model(&PaymentMethod{}).Create(&Method1)

	Method2 := PaymentMethod{
		Method: "Promtpay",
	}
	db.Model(&PaymentMethod{}).Create(&Method2)

	Method3 := PaymentMethod{
		Method: "จ่ายสดหน้าร้าน",
	}
	db.Model(&PaymentMethod{}).Create(&Method3)

	Method4 := PaymentMethod{
		Method: "เก็บเงินปลายทาง",
	}
	db.Model(&PaymentMethod{}).Create(&Method4)

	// DeliveryType Data
	type1 := DeliveryType{
		Type: "รับสินค้าที่ร้าน",
	}
	db.Model(&DeliveryType{}).Create(&type1)

	type2 := DeliveryType{
		Type: "จัดส่งตามที่อยู่ในระบบ",
	}
	db.Model(&DeliveryType{}).Create(&type2)

	// ระบบคลังสินค้า
	// staff data
	var suwanan Staff
	var kanyanat Staff
	var sirilak Staff
	db.Raw("SELECT * FROM staffs WHERE email = ?", "suwanan@gmail.com").Scan(&suwanan)
	db.Raw("SELECT * FROM staffs WHERE email = ?", "kanyanat@gmail.com").Scan(&kanyanat)
	db.Raw("SELECT * FROM staffs WHERE email = ?", "sirilak@gmail.com").Scan(&sirilak)

	db.Model(&Staff{}).Create(&Staff{
		Name:     "Suwanan",
		Email:    "suwanan@gmail.com",
		Password: string(password),
	})

	db.Model(&Staff{}).Create(&Staff{
		Name:     "Kanyanat",
		Email:    "kanyanat@gmail.com",
		Password: string(password),
	})

	db.Model(&Staff{}).Create(&Staff{
		Name:     "Sirilak",
		Email:    "sirilak@gmail.com",
		Password: string(password),
	})

	// ProductType Data
	drink := ProductType{
		Ptype: "เครื่องดื่ม",
	}
	db.Model(&ProductType{}).Create(&drink)

	vegetable := ProductType{
		Ptype: "ผัก",
	}
	db.Model(&ProductType{}).Create(&vegetable)

	fruit := ProductType{
		Ptype: "ผลไม้",
	}
	db.Model(&ProductType{}).Create(&fruit)

	meat := ProductType{
		Ptype: "เนื้อสัตว์",
	}
	db.Model(&ProductType{}).Create(&meat)

	another := ProductType{
		Ptype: "อื่นๆ",
	}
	db.Model(&ProductType{}).Create(&another)

	// --- Product Data
	mango := Product{
		Name:        "มะม่วงน้ำดอกไม้",
		ProductType: fruit,
	}
	db.Model(&Product{}).Create(&mango)

	maprang := Product{
		Name:        "มะปราง",
		ProductType: fruit,
	}
	db.Model(&Product{}).Create(&maprang)

	rambutan := Product{
		Name:        "เงาะ",
		ProductType: fruit,
	}
	db.Model(&Product{}).Create(&rambutan)

	durian := Product{
		Name:        "ทุเรียน",
		ProductType: fruit,
	}
	db.Model(&Product{}).Create(&durian)

	egg := Product{
		Name:        "ไข่ไก่ No.0",
		ProductType: another,
	}
	db.Model(&Product{}).Create(&egg)

	egg_omega := Product{
		Name:        "ไข่ไก่ OMEGA-3",
		ProductType: another,
	}
	db.Model(&Product{}).Create(&egg_omega)

	salad := Product{
		Name:        "ผักสลัด",
		ProductType: vegetable,
	}
	db.Model(&Product{}).Create(&salad)

	sunflower_sprout := Product{
		Name:        "ต้นอ่อนทานตะวัน",
		ProductType: vegetable,
	}
	db.Model(&Product{}).Create(&sunflower_sprout)

	milk := Product{
		Name:        "นมสดพาสเจอร์ไรส์ รสจืด ขนาด 5 ลิตร",
		ProductType: drink,
	}
	db.Model(&Product{}).Create(&milk)

	milk_banana := Product{
		Name:        "นมรสกล้วย 150 ml",
		ProductType: drink,
	}
	db.Model(&Product{}).Create(&milk_banana)

	milk_lychee := Product{
		Name:        "นมรสลิ้นจี่ 150 ml",
		ProductType: drink,
	}
	db.Model(&Product{}).Create(&milk_lychee)

	milk_melon := Product{
		Name:        "นมรสเมล่อน 150 ml",
		ProductType: drink,
	}
	db.Model(&Product{}).Create(&milk_melon)

	pork := Product{
		Name:        "เนื้อหมู",
		ProductType: meat,
	}
	db.Model(&Product{}).Create(&pork)

	turkey := Product{
		Name:        "ไก่งวง 3-5 กิโลกรัม",
		ProductType: meat,
	}
	db.Model(&Product{}).Create(&turkey)

	water350 := Product{
		Name:        "น้ำดื่ม ขนาด 350 ซีซี 12 ขวด",
		ProductType: drink,
	}
	db.Model(&Product{}).Create(&water350)

	water600 := Product{
		Name:        "น้ำดื่ม ขนาด 600 ซีซี 12 ขวด",
		ProductType: drink,
	}
	db.Model(&Product{}).Create(&water600)

	// --- Supplier Data
	db.Model(&Supplier{}).Create(&Supplier{
		Name: "มทส.",
	})

	db.Model(&Supplier{}).Create(&Supplier{
		Name: "นักศึกษา",
	})

	db.Model(&Supplier{}).Create(&Supplier{
		Name: "มทส.",
	})

	// ระบบสั่งจองสินค้า
	preorder1 := Preorder{
		Amount:        2,
		User:          narudee,
		Product:       durian,
		PaymentMethod: Method1,
	}
	db.Model(&Preorder{}).Create(&preorder1)

	preorder2 := Preorder{
		Amount:        1,
		User:          phatcha,
		Product:       mango,
		PaymentMethod: Method2,
	}

	preorder3 := Preorder{
		Amount:        1,
		User:          narudee,
		Product:       turkey,
		PaymentMethod: Method1,
	}
	db.Model(&Preorder{}).Create(&preorder3)

	preorder4 := Preorder{
		Amount:        2,
		User:          phatcha,
		Product:       milk_banana,
		PaymentMethod: Method2,
	}
	db.Model(&Preorder{}).Create(&preorder4)

	// ระบบสั่งสินค้า
	// Status Data
	status1 := Status{
		Statusorder: "Confirm",
	}
	db.Model(&Status{}).Create(&status1)

	status2 := Status{
		Statusorder: "not sure",
	}
	db.Model(&Status{}).Create(&status2)

	// Order data
	order1 := Order{
		User:      narudee,
		Preorder:  preorder1,
		Status:    status1,
		OrderTime: time.Now(),
	}

	db.Model(&Order{}).Create(&order1)

	order2 := Order{
		User:      phatcha,
		Preorder:  preorder2,
		Status:    status1,
		OrderTime: time.Now(),
	}
	db.Model(&Order{}).Create(&order2)

	order3 := Order{
		User:      narudee,
		Preorder:  preorder3,
		Status:    status1,
		OrderTime: time.Now(),
	}

	db.Model(&Order{}).Create(&order3)

	order4 := Order{
		User:      phatcha,
		Preorder:  preorder4,
		Status:    status1,
		OrderTime: time.Now(),
	}
	db.Model(&Order{}).Create(&order4)

	// ระบบสมาชิก
	// --- ContactType Data
	email := Contact{
		Ctype: "Email",
	}
	db.Model(&Contact{}).Create(&email)

	phone := Contact{
		Ctype: "Phone",
	}
	db.Model(&Contact{}).Create(&phone)

	post := Contact{
		Ctype: "Post",
	}
	db.Model(&Contact{}).Create(&post)

	// Sex Data
	male := Sex{
		Stype: "Male",
	}
	db.Model(&Sex{}).Create(&male)

	female := Sex{
		Stype: "Female",
	}
	db.Model(&Sex{}).Create(&female)

	// Olduser Data
	ever := Olduser{
		Otype: "Ever",
	}
	db.Model(&Olduser{}).Create(&ever)

	yet := Olduser{
		Otype: "Yet",
	}
	db.Model(&Olduser{}).Create(&yet)

	// Religion Data
	buddhism := Religion{
		Rtype: "Buddhism",
	}
	db.Model(&Religion{}).Create(&buddhism)

	christianity := Religion{
		Rtype: "Christianity",
	}
	db.Model(&Religion{}).Create(&christianity)

	islam := Religion{
		Rtype: "Islam",
	}
	db.Model(&Religion{}).Create(&islam)

	Another := Religion{
		Rtype: "Another",
	}
	db.Model(&Religion{}).Create(&Another)

	// Account 1
	db.Model(&Account{}).Create(&Account{
		Address:  "10 M.22 T.Suranaree A.Mueng",
		Province: "Nakhonratchasrima",
		Owner:    name,
		Contact:  phone,
		Sex:      female,
		Olduser:  yet,
		Religion: christianity,
	})

	// Account 2
	db.Model(&Account{}).Create(&Account{
		Address:  "Suranaree U. A.Mueng",
		Province: "Nakhonratchasrima",
		Owner:    phatcha,
		Contact:  phone,
		Sex:      female,
		Olduser:  yet,
		Religion: buddhism,
	})

	// Account 3
	db.Model(&Account{}).Create(&Account{
		Address:  "Suranaree U. A.Mueng",
		Province: "Nakhonratchasrima",
		Owner:    patnarin,
		Contact:  email,
		Sex:      female,
		Olduser:  yet,
		Religion: buddhism,
	})

	// Account 3
	db.Model(&Account{}).Create(&Account{
		Address:  "Suranaree U. A.Mueng",
		Province: "Nakhonratchasrima",
		Owner:    narudee,
		Contact:  post,
		Sex:      female,
		Olduser:  yet,
		Religion: buddhism,
	})

}
