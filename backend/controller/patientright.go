package controller

import (
	"net/http"

	"github.com/pechkr2020/sa-project/entity"
	"github.com/gin-gonic/gin"
)

// GET /patientright/:id
func GetPatientRight(c *gin.Context) {
	var patientright entity.PatientRight
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM patient_rights WHERE id = ?", id).Scan(&patientright).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patientright})
}

// GET /patientrights
func ListPatientRights(c *gin.Context) {
	var patientrights []entity.PatientRight
	if err := entity.DB().Raw("SELECT * FROM patient_rights").Scan(&patientrights).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patientrights})
}

