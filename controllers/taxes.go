// controllers/books.go

package controllers

import (
	"net/http"

	"github.com/alx9110/payment-calculator/models"
	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books
func FindTaxes(c *gin.Context) {
	var tax []models.Tax
	models.DB.Find(&tax)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"data": tax,
	})
}

// GET /books/:id
// Find a book
func FindTax(c *gin.Context) {
	var tax models.Tax

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&tax).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tax not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tax})
}

func CreateTax(c *gin.Context) {
	// Validate input
	var input models.CreateTaxInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	Tax := models.Tax{Name: input.Name, Value: input.Value}
	models.DB.Create(&Tax)

	c.JSON(http.StatusOK, gin.H{"data": Tax})
}

// PATCH /books/:id
// Update a book
func UpdateTax(c *gin.Context) {
	// Get model if exist
	var Tax models.Tax
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Tax).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tax not found!"})
		return
	}

	// Validate input
	var input models.UpdateTaxInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&Tax).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": Tax})
}

// DELETE /books/:id
// Delete a book
func DeleteTax(c *gin.Context) {
	// Get model if exist
	var Tax models.Tax
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Tax).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tax not found!"})
		return
	}

	models.DB.Delete(&Tax)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
