// controllers/books.go

package controllers

import (
	"net/http"

	"github.com/alx9110/payment-calculator/ext"
	"github.com/alx9110/payment-calculator/models"
	"github.com/gin-gonic/gin"
)

// GET /records
// Get all records
func FindRecords(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	email, validation_error := ext.ValidateToken(tokenString)
	if validation_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validation_error.Error()})
		return
	}

	var records []models.Record
	models.DB.Where("Email = ?", email).Find(&records)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"data": records,
	})
}

// GET /records/:id
// Find a record
func FindRecord(c *gin.Context) {
	var product models.Record

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// POST /records/
// Create a record
func CreateRecord(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	tokenString := c.GetHeader("Authorization")
	email, validation_error := ext.ValidateToken(tokenString)
	if validation_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validation_error.Error()})
		return
	}

	// Validate input
	var input models.CreateRecordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tax models.Tax
	if err := models.DB.Last(&tax).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tax not found!"})
		return
	}
	var previous_record models.Record
	if err := models.DB.Where("Email = ?", email).Last(&previous_record).Error; err != nil {
		previous_record := models.Record{
			HotValue:     input.HotValue,
			HotCost:      0,
			ColdValue:    input.ColdValue,
			ColdCost:     0,
			EnergyValue:  input.EnergyValue,
			EnergyCost:   0,
			DrenageValue: input.HotValue + input.ColdValue,
			DrenageCost:  0,
			Email:        email,
		}
		models.DB.Create(&previous_record)
		return
	}
	models.DB.Last(&previous_record)
	record := models.Record{
		HotValue:     input.HotValue,
		HotCost:      ext.CalcCost(tax.HotPrice, input.HotValue, previous_record.HotValue),
		ColdValue:    input.ColdValue,
		ColdCost:     ext.CalcCost(tax.ColdPrice, input.ColdValue, previous_record.ColdValue),
		EnergyValue:  input.EnergyValue,
		EnergyCost:   ext.CalcCost(tax.EnergyPrice, input.EnergyValue, previous_record.EnergyValue),
		DrenageValue: input.HotValue + input.ColdValue,
		DrenageCost:  ext.CalcCost(tax.DrenagePrice, input.HotValue+input.ColdValue, previous_record.DrenageValue),
		Email:        email,
	}
	models.DB.Create(&record)

	c.JSON(http.StatusOK, gin.H{"data": record})
}

// PATCH /records/:id
// Update a record
func UpdateRecord(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var record models.Record
	if err := models.DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateRecordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&record).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": record})
}

// DELETE /records/:id
// Delete a record
func DeleteRecord(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	email, validation_error := ext.ValidateToken(tokenString)
	if validation_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validation_error.Error()})
		return
	}

	var record models.Record
	if err := models.DB.Where("id = ? AND Email = ?", c.Param("id"), email).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&record)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
