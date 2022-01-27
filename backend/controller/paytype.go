package controller

import (
	"net/http"

	"github.com/pechkr2020/sa-project/entity"
	"github.com/gin-gonic/gin"
)

// GET /patientright/:id
func GetPaytype(c *gin.Context) {
	var paytype entity.Paytype
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM paytypes WHERE id = ?", id).Scan(&paytype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paytype})
}

// GET /patientrights
func ListPaytypes(c *gin.Context) {
	var paytypes []entity.Paytype
	if err := entity.DB().Raw("SELECT * FROM paytypes").Scan(&paytypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paytypes})
}

