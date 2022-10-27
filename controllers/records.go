// controllers/books.go

package controllers

import (
	"net/http"

	"github.com/alx9110/payment-calculator/models"
	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books
func FindRecords(c *gin.Context) {
	var products []models.Record
	models.DB.Find(&products)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"data": products,
	})
}

// GET /books/:id
// Find a book
func FindRecord(c *gin.Context) { // Get model if exist
	var product models.Record

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateRecord(c *gin.Context) {
	// Validate input
	var input models.CreateRecordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	var tax models.Tax
	if err := models.DB.Where("Name = ?", input.Name).Last(&tax).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tax not found!"})
		return
	}
	var previous_record models.Record
	if err := models.DB.Where("Name = ?", input.Name).Last(&previous_record).Error; err != nil {
		previous_record := models.Record{Name: input.Name, Value: input.Value, Cost: 0}
		models.DB.Create(&previous_record)
	}
	models.DB.Where("Name = ?", input.Name).Last(&previous_record)
	record := models.Record{Name: input.Name, Value: input.Value, Cost: tax.Value * (input.Value - previous_record.Value)}
	models.DB.Create(&record)

	c.JSON(http.StatusOK, gin.H{"data": record})
}

// PATCH /books/:id
// Update a book
func UpdateRecord(c *gin.Context) {
	// Get model if exist
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

// DELETE /books/:id
// Delete a book
func DeleteProduct(c *gin.Context) {
	// Get model if exist
	var record models.Record
	if err := models.DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&record)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
