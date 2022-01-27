package main

import (

	"github.com/pechkr2020/sa-project/controller"

	"github.com/pechkr2020/sa-project/entity"

	"github.com/pechkr2020/sa-project/middlewares"
  
	"github.com/gin-gonic/gin"
  
  )
  
  
  func main() {
  
	entity.SetupDatabase()
  
  
	r := gin.Default()

	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		
		protected := api.Use(middlewares.Authorizes())
		{

			//Patient Routes

			protected.GET("/patients", controller.ListPatient)
			protected.GET("/patient/:id", controller.GetPatient)

			//Examinations Routes
			protected.GET("/examinations", controller.ListExaminations)
			protected.GET("/examination/:id", controller.GetExamination)

			//PatientRight Routes
			protected.GET("/patientrights", controller.ListPatientRights)
			protected.GET("/patientright/:id", controller.GetPatientRight)

			//Paytype Routes
			protected.GET("/paytypes", controller.ListPaytypes)
			protected.GET("/paytype/:id", controller.GetPaytype)

			//Crashier Routes
			protected.GET("/cashiers", controller.ListCashiers)
			protected.GET("/cashier/:id", controller.GetCashier)

			//Bill Routes
			protected.GET("/bills", controller.ListBills)
			protected.GET("/bill/:id", controller.GetBill)
			protected.POST("/bills", controller.CreateBill)
			protected.PATCH("/bills", controller.UpdateBill)
			protected.DELETE("/bills/:id", controller.DeleteBill)
		}

	}

	// Authentication Routes
	r.POST("/login", controller.Login)
	
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