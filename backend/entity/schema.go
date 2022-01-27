package entity


import (

  "time"

  "gorm.io/gorm"

)


type Patient struct {

  gorm.Model

  ID_card		string

  Firstname		string

  Lastname		string

  Birthdate		time.Time

  Age 			uint8

  Allergy		string

  Underlying_Disease	string

  Gender 		string

  Recorder		string

  Examinations		[]Examination		`gorm:"foreignKey:PatientID"`


}

type Examination struct {

	gorm.Model

	TreatmentTime 	time.Time

	Treatment		string

	Treatment_cost	uint

	Medicine_cost	uint

	PatientID		*uint

	Patient			Patient		`gorm:"references:id"`

	DoctorId		uint

	Clinic			uint

	Disease			uint8

	Medicine		uint8

	BillItems 		[]BillItem	`gorm:"foreignKey:ExaminationID"`


}

type Cashier struct {

	gorm.Model

	Name		string

	Email		string		`gorm:"uniqueIndex"`

	Password	string

	Bills		[]Bill		`gorm:"foreignKey:CashierID"`
}

type PatientRight struct {

	gorm.Model

	Name 		string

	Discount	uint

	Bills		[]Bill		`gorm:"foreignKey:PatientRightID"`
}

type Bill struct {

	gorm.Model

	PatientRightID	*uint

	PatientRight	PatientRight	`gorm:"references:id"`

	BillTime	time.Time

	Total 			uint

	Telephone 		string

	PaytypeID		*uint

	Paytype			Paytype 	`gorm:"references:id"`

	CashierID		*uint

	Cashier		Cashier	`gorm:"references:id"`

	BillItems	[]BillItem  	`gorm:"foreignKey:BillID; constraint:OnDelete:CASCADE"`
	

}

type BillItem struct {

	gorm.Model 

	ExaminationID 	*uint

	Examination 	Examination		`gorm:"references:id"`

	BillID 			*uint
	Bill			Bill			`gorm:"references:id"`

}


type Paytype struct {

	gorm.Model

	Type 		string

	Bills		[]Bill		`gorm:"foreignKey:PaytypeID"`


}
