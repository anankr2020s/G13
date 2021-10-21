package controller

import (
	"net/http"

	"github.com/pechkr2020/sa-project/entity"
	"github.com/gin-gonic/gin"
)

// GET /patients
// List all patient
func ListPatient(c *gin.Context) {
	var patients []entity.Patient
	if err := entity.DB().Raw("SELECT * FROM patients").Scan(&patients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patients})
}

// GET /patient/:id
// Get patient by id
func GetPatient(c *gin.Context) {
	var patient entity.Patient
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM patients WHERE id = ?", id).Scan(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patient})
}

