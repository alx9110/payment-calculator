// controllers/books.go

package controllers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books
func FindProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(200, gin.H{
		"message": products,
	})
}

// GET /books/:id
// Find a book
func FindProduct(c *gin.Context) { // Get model if exist
	var product models.Product

	if err := models.DB.Where("ID = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProduct(c *gin.Context) {
	// Validate input
	var input models.CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	product := models.Product{Code: input.Code, Price: input.Price}
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// PATCH /books/:id
// Update a book
func UpdateProduct(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DELETE /books/:id
// Delete a book
func DeleteProduct(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
