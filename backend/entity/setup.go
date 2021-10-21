package entity

import (
	

	"time"

	"gorm.io/gorm"

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

	database.AutoMigrate(&Patient{}, &Examination{}, &PatientRight{}, &Cashier{}, &Bill{})

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
	
	
	//Crashier Data
	Cra1 := Cashier{
		Name: "Somsom",
		Password: "somsom1234",
	}
	db.Model(&Cashier{}).Create(&Cra1)

	//PatientRight Data
	Pr1 := PatientRight{
		Name: "บัตรทอง",
		Discount: 300,
	}
	db.Model(&PatientRight{}).Create(&Pr1)

	Pr2 := PatientRight{
		Name: "บัตรนักศึกษา",
		Discount: 250,
	}
	db.Model(&PatientRight{}).Create(&Pr2)

	Pr3 := PatientRight{
		Name: "บัตรผู้สูงอายุ",
		Discount: 500,
	}
	db.Model(&PatientRight{}).Create(&Pr3)
	




}	
