package entity

import (
	

	"time"

	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db
}

func SetupDatabase(){

	database, err := gorm.Open(sqlite.Open("sa-Bill.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")
	}

	database.AutoMigrate(&Patient{}, &Examination{}, &PatientRight{}, &Cashier{}, &Bill{}, &BillItem{}, &Paytype{})

	db = database

	db.Model(&Patient{}).Create(&Patient{
			ID_card: "1234129a",
			Firstname: "Anan",
			Lastname:	"Krasen",
			Birthdate:	time.Now(),
			Age:		21,
			Allergy:	"-",
			Underlying_Disease:	"-",
			Gender:		"Male",
			Recorder:	"Phuwadon",

	})
	db.Model(&Patient{}).Create(&Patient{
		ID_card: "2399235b",
		Firstname: "Phum",
		Lastname:	"Chai",
		Birthdate:	time.Now(),
		Age:		21,
		Allergy:	"-",
		Underlying_Disease:	"-",
		Gender:		"Male",
		Recorder:	"Phuwadon",
})

	var anan Patient
	var phum Patient
	db.Raw("SELECT * FROM patients WHERE id_card = ?","1234129a").Scan(&anan)
	db.Raw("SELECT * FROM patients WHERE id_card = ?","2399235b").Scan(&phum)



	Ex1 := Examination{
		TreatmentTime: 	time.Now(),

		Treatment:		"ถอนฟัน",

		Treatment_cost:	250,

		Medicine_cost:	50,

		Patient:		anan,

		DoctorId:		2,

		Clinic:			4,

		Disease:		2,

		Medicine:		2,
	}
	db.Model(&Examination{}).Create(&Ex1)
	

	Ex2 := Examination{
		TreatmentTime: 	time.Now(),

		Treatment:		"ทานยา",

		Treatment_cost:	50,

		Medicine_cost:	50,

		Patient:		phum,

		DoctorId:		2,

		Clinic:			4,

		Disease:		1,

		Medicine:		2,
	}
	db.Model(&Examination{}).Create(&Ex2)
	
	password, err := bcrypt.GenerateFromPassword([]byte("1234"), 14)
	password1, err := bcrypt.GenerateFromPassword([]byte("1235"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("2020"), 14)
	
	//Crashier Data
	Cra1 := Cashier{
		Name: "อนันต์",
		Email: "anan1234@gmail.com",
		Password: string(password),
	}
	db.Model(&Cashier{}).Create(&Cra1)

	Cra2 := Cashier{
		Name: "ภูวดล",
		Email: "phu123@gmail.com",
		Password: string(password1),
	}
	db.Model(&Cashier{}).Create(&Cra2)

	Cra3 := Cashier{
		Name: "ภูมิชัย",
		Email: "phumchai123@gmail.com",
		Password: string(password2),
	}
	db.Model(&Cashier{}).Create(&Cra3)

	//PatientRight Data
	Pr1 := PatientRight{
		Name: "สิทธิ์สุขภาพถ้วนหน้า",
		Discount: 80,
	}
	db.Model(&PatientRight{}).Create(&Pr1)

	Pr2 := PatientRight{
		Name: "สิทธิ์นักศึกษา",
		Discount: 50,
	}
	db.Model(&PatientRight{}).Create(&Pr2)

	//Paytype Data

	Pt1 := Paytype{
		Type : "เงินสด",
	}
	db.Model(&Paytype{}).Create(&Pt1)

	Pt2 := Paytype{
		Type : "บัตรเครดิต",
	}
	db.Model(&Paytype{}).Create(&Pt2)


	
	    

}	
