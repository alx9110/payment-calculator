package main

import (
	"github.com/alx9110/payment-calculator/controllers"
	"github.com/alx9110/payment-calculator/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Routing
	models.ConnectDatabase()
	r := gin.Default()
	// Serve frontend static files
	// r.Use(static.Serve("/", static.LocalFile("./web", true)))
	// Setup route group for the API
	r.SetTrustedProxies([]string{"0.0.0.0"})

	api := r.Group("/api")
	{
		// Records
		api.GET("/records/", controllers.FindRecords)
		api.GET("/records/:id", controllers.FindRecord)
		api.PATCH("/records/:id", controllers.UpdateRecord)
		api.POST("/records/", controllers.CreateRecord)
		api.DELETE("/records/:id", controllers.DeleteProduct)

		// Taxes
		api.GET("/taxes/", controllers.FindTaxes)
		api.GET("/taxes/:id", controllers.FindTax)
		api.PATCH("/taxes/:id", controllers.UpdateTax)
		api.POST("/taxes/", controllers.CreateTax)
		api.DELETE("/taxes/:id", controllers.DeleteTax)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
