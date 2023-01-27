package controllers

import (
	"net/http"

	"github.com/alx9110/payment-calculator/ext"
	"github.com/alx9110/payment-calculator/models"
	"github.com/gin-gonic/gin"
)

func FindTaxes(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	email, validation_error := ext.ValidateToken(tokenString)
	if validation_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validation_error.Error()})
		return
	}

	var tax []models.Tax
	models.DB.Where("Email = ?", email).Find(&tax)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"data": tax,
	})
}

func FindTax(c *gin.Context) {
	var tax models.Tax
	if err := models.DB.Where("ID = ?", c.Param("id")).First(&tax).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tax not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tax})
}

func CreateTax(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	email, validation_error := ext.ValidateToken(tokenString)
	if validation_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validation_error.Error()})
		return
	}

	var input models.CreateTaxInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Tax := models.Tax{
		HotPrice:     input.HotPrice,
		ColdPrice:    input.ColdPrice,
		EnergyPrice:  input.EnergyPrice,
		DrenagePrice: input.DrenagePrice,
		Email:        email,
	}

	models.DB.Create(&Tax)

	c.JSON(http.StatusOK, gin.H{"data": Tax})
}

func UpdateTax(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	email, validation_error := ext.ValidateToken(tokenString)
	if validation_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validation_error.Error()})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	var json models.UpdateTaxInput
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var tax models.Tax
	if err := models.DB.Where("id = ? AND Email = ?", c.Param("id"), email).First(&tax).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tax not found!"})
		return
	}
	models.DB.Model(&tax).Updates(json)

	c.JSON(http.StatusOK, gin.H{"data": tax})
}

func DeleteTax(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	email, validation_error := ext.ValidateToken(tokenString)
	if validation_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validation_error.Error()})
		return
	}

	var Tax models.Tax
	if err := models.DB.Where("id = ? AND Email = ?", c.Param("id"), email).First(&Tax).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tax not found!"})
		return
	}

	models.DB.Delete(&Tax)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
