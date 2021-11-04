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
		&User{},
		&Olduser{},
		&Contact{},
		&Sex{},
		&Religion{},
		&Account{},
		&Preorder{},
		&PaymentMethod{},
		&Payment{},
		&DeliveryType{},
		&Return{},
		&Staff{},
		&Product{},
		&ProductType{},
		&Supplier{},
		&ProductStock{},
		&Order{},
		&Status{},
	)
	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	// User Data
	db.Model(&User{}).Create(&User{
		Name:     "Pawarit Praneetponkrang",
		Email:    "prawarit@gmail.com",
		Password: string(password),
	})

	db.Model(&User{}).Create(&User{
		Name:     "Phatcha  Srisuwo",
		Email:    "phatcha@gmail.com",
		Password: string(password),
	})

	db.Model(&User{}).Create(&User{
		Name:     "Narudee Arunno",
		Email:    "narudee@gmail.com",
		Password: string(password),
	})

	db.Model(&User{}).Create(&User{
		Name:     "Nawamin Saengsaikaew",
		Email:    "nawamin@gmail.com",
		Password: string(password),
	})

	db.Model(&User{}).Create(&User{
		Name:     "Patnarin Aiewchoei",
		Email:    "patnarin@gmail.com",
		Password: string(password),
	})

	var prawarit User
	var phatcha User
	var narudee User
	var nawamin User
	var patnarin User
	db.Raw("SELECT * FROM users WHERE email = ?", "prawarit@gmail.com").Scan(&prawarit)
	db.Raw("SELECT * FROM users WHERE email = ?", "phatcha@gmail.com").Scan(&phatcha)
	db.Raw("SELECT * FROM users WHERE email = ?", "narudee@gmail.com").Scan(&narudee)
	db.Raw("SELECT * FROM users WHERE email = ?", "nawamin@gmail.com").Scan(&nawamin)
	db.Raw("SELECT * FROM users WHERE email = ?", "patnarin@gmail.com").Scan(&patnarin)

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
		Type: "จัดส่งถึงบ้าน",
	}
	db.Model(&DeliveryType{}).Create(&type2)

	// Staff data
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

	// Product Data
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

	water600 := Product{
		Name:        "น้ำดื่ม ขนาด 600 ซีซี 12 ขวด",
		ProductType: drink,
	}
	db.Model(&Product{}).Create(&water600)

	// Supplier Data
	db.Model(&Supplier{}).Create(&Supplier{
		Name: "มทส.",
	})

	db.Model(&Supplier{}).Create(&Supplier{
		Name: "นักศึกษา",
	})

	db.Model(&Supplier{}).Create(&Supplier{
		Name: "เกษตกร",
	})

	// Status Data
	status1 := Status{
		Statusorder: "Confirm",
	}
	db.Model(&Status{}).Create(&status1)

	status2 := Status{
		Statusorder: "Not sure",
	}
	db.Model(&Status{}).Create(&status2)

	// ContactType Data
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

	// Account Data
	db.Model(&Account{}).Create(&Account{
		Address:  "Suranaree U. A.Mueng",
		Province: "Nakhonratchasrima",
		Owner:    prawarit,
		Contact:  phone,
		Sex:      male,
		Olduser:  yet,
		Religion: buddhism,
	})

	db.Model(&Account{}).Create(&Account{
		Address:  "Suranaree U. A.Mueng",
		Province: "Nakhonratchasrima",
		Owner:    phatcha,
		Contact:  phone,
		Sex:      female,
		Olduser:  yet,
		Religion: christianity,
	})

	db.Model(&Account{}).Create(&Account{
		Address:  "Suranaree U. A.Mueng",
		Province: "Nakhonratchasrima",
		Owner:    narudee,
		Contact:  phone,
		Sex:      female,
		Olduser:  yet,
		Religion: buddhism,
	})

	db.Model(&Account{}).Create(&Account{
		Address:  "Suranaree U. A.Mueng",
		Province: "Nakhonratchasrima",
		Owner:    nawamin,
		Contact:  post,
		Sex:      male,
		Olduser:  yet,
		Religion: buddhism,
	})

	db.Model(&Account{}).Create(&Account{
		Address:  "Suranaree U. A.Mueng",
		Province: "Nakhonratchasrima",
		Owner:    patnarin,
		Contact:  email,
		Sex:      female,
		Olduser:  yet,
		Religion: christianity,
	})

	// Preorder Data
	pre1ofprawarit := Preorder{
		Amount:        3,
		User:          prawarit,
		Product:       durian,
		PaymentMethod: Method1,
	}
	db.Model(&Preorder{}).Create(&pre1ofprawarit)

	pre2ofprawarit := Preorder{
		Amount:        3,
		User:          prawarit,
		Product:       mango,
		PaymentMethod: Method2,
	}
	db.Model(&Preorder{}).Create(&pre2ofprawarit)

	pre1ofnawamin := Preorder{
		Amount:        4,
		User:          nawamin,
		Product:       durian,
		PaymentMethod: Method3,
	}
	db.Model(&Preorder{}).Create(&pre1ofnawamin)

	pre1ofphatcha := Preorder{
		Amount:        5,
		User:          phatcha,
		Product:       durian,
		PaymentMethod: Method1,
	}
	db.Model(&Preorder{}).Create(&pre1ofphatcha)

	pre2ofphatcha := Preorder{
		Amount:        1,
		User:          phatcha,
		Product:       rambutan,
		PaymentMethod: Method2,
	}
	db.Model(&Preorder{}).Create(&pre2ofphatcha)

	pre3ofphatcha := Preorder{
		Amount:        2,
		User:          phatcha,
		Product:       pork,
		PaymentMethod: Method3,
	}
	db.Model(&Preorder{}).Create(&pre3ofphatcha)

	pre1ofnarudee := Preorder{
		Amount:        8,
		User:          narudee,
		Product:       milk,
		PaymentMethod: Method2,
	}
	db.Model(&Preorder{}).Create(&pre1ofnarudee)

	pre2ofnarudee := Preorder{
		Amount:        5,
		User:          narudee,
		Product:       mango,
		PaymentMethod: Method3,
	}
	db.Model(&Preorder{}).Create(&pre2ofnarudee)

	pre3ofnarudee := Preorder{
		Amount:        2,
		User:          narudee,
		Product:       durian,
		PaymentMethod: Method3,
	}
	db.Model(&Preorder{}).Create(&pre3ofnarudee)

	pre1ofpatnarin := Preorder{
		Amount:        4,
		User:          patnarin,
		Product:       milk,
		PaymentMethod: Method1,
	}
	db.Model(&Preorder{}).Create(&pre1ofpatnarin)

	pre2ofpatnarin := Preorder{
		Amount:        4,
		User:          patnarin,
		Product:       sunflower_sprout,
		PaymentMethod: Method2,
	}
	db.Model(&Preorder{}).Create(&pre2ofpatnarin)

	pre3ofpatnarin := Preorder{
		Amount:        4,
		User:          patnarin,
		Product:       durian,
		PaymentMethod: Method3,
	}
	db.Model(&Preorder{}).Create(&pre3ofpatnarin)

	pre4ofpatnarin := Preorder{
		Amount:        6,
		User:          patnarin,
		Product:       egg_omega,
		PaymentMethod: Method4,
	}
	db.Model(&Preorder{}).Create(&pre4ofpatnarin)

	// Order data
	order1 := Order{
		User:      prawarit,
		Preorder:  pre1ofprawarit,
		Status:    status1,
		OrderTime: time.Now(),
	}
	db.Model(&Order{}).Create(&order1)

	order2 := Order{
		User:      phatcha,
		Preorder:  pre1ofphatcha,
		Status:    status2,
		OrderTime: time.Now(),
	}
	db.Model(&Order{}).Create(&order2)

	order3 := Order{
		User:      phatcha,
		Preorder:  pre1ofphatcha,
		Status:    status2,
		OrderTime: time.Now(),
	}
	db.Model(&Order{}).Create(&order3)

	order4 := Order{
		User:      narudee,
		Preorder:  pre2ofnarudee,
		Status:    status1,
		OrderTime: time.Now(),
	}
	db.Model(&Order{}).Create(&order4)

	order5 := Order{
		User:      narudee,
		Preorder:  pre2ofnarudee,
		Status:    status1,
		OrderTime: time.Now(),
	}
	db.Model(&Order{}).Create(&order5)

	order6 := Order{
		User:      patnarin,
		Preorder:  pre1ofpatnarin,
		Status:    status1,
		OrderTime: time.Now(),
	}
	db.Model(&Order{}).Create(&order6)

	order7 := Order{
		User:      patnarin,
		Preorder:  pre2ofpatnarin,
		Status:    status1,
		OrderTime: time.Now(),
	}
	db.Model(&Order{}).Create(&order7)
}
