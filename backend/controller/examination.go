package controller

import (
	"net/http"

	"github.com/pechkr2020/sa-project/entity"
	"github.com/gin-gonic/gin"
)

// GET /examination
// List all examination
func ListExaminations(c *gin.Context) {
	var examinations []entity.Examination
	if err := entity.DB().Raw("SELECT * FROM examinations ").Scan(&examinations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": examinations})
}

// GET /examination/:id
// Get examination by id
func GetExamination(c *gin.Context) {
	var examination []entity.Examination
	id := c.Param("id")
	if err := entity.DB().Preload("Patient").Raw("SELECT * FROM examinations WHERE patient_id = ?", id).Find(&examination).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": examination})
}




// PATCH /examinations
func UpdateExamination(c *gin.Context) {
	var examination entity.Examination
	if err := c.ShouldBindJSON(&examination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", examination.ID).First(&examination); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "examination not found"})
		return
	}

	if err := entity.DB().Save(&examination).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": examination})
}

