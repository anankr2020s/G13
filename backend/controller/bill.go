package controller

import (
	"net/http"

	

	"github.com/pechkr2020/sa-project/entity"
	"github.com/gin-gonic/gin"
)

// POST /bills
func CreateBill(c *gin.Context) {

	
	var patientright entity.PatientRight
	var cashier entity.Cashier
	var bill entity.Bill
	var paytype entity.Paytype
	
	
	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	

	// 10: ค้นหา patientright ด้วย id
	if tx := entity.DB().Where("id = ?", bill.PatientRightID).First(&patientright); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patientright not found"})
		return
	}

	// 11: ค้นหา crashier ด้วย id
	if tx := entity.DB().Where("id = ?", bill.CashierID).First(&cashier); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cashier not found"})
		return
	}

	
		
	// 12: สร้าง Bill
	bl := entity.Bill{       
		PatientRight:       patientright,
		Paytype : paytype,                 // โยงความสัมพันธ์กับ Entity PatientRight
		Cashier:    cashier,               // โยงความสัมพันธ์กับ Entity Crashier
		BillTime: bill.BillTime, // ตั้งค่าฟิลด์ BillTime
		
	}

	// 13: บันทึก
	if err := entity.DB().Create(&bl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	

	var items []entity.BillItem

	for _,item := range bill.BillItems{

		var exams entity.Examination

		if tx := entity.DB().Where("id = ?", item.ExaminationID).First(&exams); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "patientright not found"})
			return
		}

		i := entity.BillItem{

			Examination : exams,
			Bill : bl,
		}

		items = append(items, i)
	}

	// 13: บันทึก
	if err := entity.DB().Create(&items).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity.DB().Preload("BillItems").Preload("BillItems.Examination").Raw("SELECT * FROM bills WHERE id = ?", bl.ID).Find(&bl)

	c.JSON(http.StatusOK, gin.H{"data": bl})

}

// GET /bill/:id
func GetBill(c *gin.Context) {
	var bill entity.Bill
	id := c.Param("id")
	if err := entity.DB().Preload("Examination").Preload("PatientRight").Preload("Cashier").Raw("SELECT * FROM bills WHERE id = ?", id).Find(&bill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bill})
}

// GET /bills
func ListBills(c *gin.Context) {
	var bills []entity.Bill
	if err := entity.DB().Preload("BillItem").Preload("PatientRight").Preload("Cashier").Raw("SELECT * FROM bills").Find(&bills).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bills})
}

// DELETE /bills/:id
func DeleteBill(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM bills WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bill not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /bills
func UpdateBill(c *gin.Context) {
	var bill entity.Bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", bill.ID).First(&bill); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bill not found"})
		return
	}

	if err := entity.DB().Save(&bill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bill})
}
