package main

import (

	"github.com/pechkr2020/sa-project/controller"

	"github.com/pechkr2020/sa-project/entity"
  
	"github.com/gin-gonic/gin"
  
  )
  
  
  func main() {
  
	entity.SetupDatabase()
  
  
	r := gin.Default()

	r.Use(CORSMiddleware())
  

	//Patient Routes

	r.GET("/patients", controller.ListPatient)
	r.GET("/patient/:id", controller.GetPatient)

	//Examinations Routes
	r.GET("/examinations", controller.ListExaminations)
	r.GET("/examination/:id", controller.GetExamination)

	//PatientRight Routes
	r.GET("/patientrights", controller.ListPatientRights)
	r.GET("/patientright/:id", controller.GetPatientRight)

	//Crashier Routes
	r.GET("/cashiers", controller.ListCashiers)
	r.GET("/cashier/:id", controller.GetCashier)

	//Bill Routes
	r.GET("/bills", controller.ListBills)
	r.GET("/bill/:id", controller.GetBill)
	r.POST("/bills", controller.CreateBill)
	r.PATCH("/bills", controller.UpdateBill)
	r.DELETE("/bills/:id", controller.DeleteBill)
	
	r.Run()
  
  }

  func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}