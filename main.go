package main

import (
	"time"

	"github.com/alx9110/payment-calculator/controllers"
	"github.com/alx9110/payment-calculator/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Routing
	models.ConnectDatabase()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Serve frontend static files
	r.Use(static.Serve("/", static.LocalFile("web/routing-app/dist/routing-app", true)))
	// Setup route group for the API
	api := r.Group("/api")
	{
		// Records
		api.GET("/records/", controllers.FindRecords)
		api.GET("/records/:id", controllers.FindRecord)
		api.PATCH("/records/:id", controllers.UpdateRecord)
		api.POST("/records/", controllers.CreateRecord)
		api.DELETE("/records/:id", controllers.DeleteRecord)

		// Taxes
		api.GET("/taxes/", controllers.FindTaxes)
		api.GET("/taxes/:id", controllers.FindTax)
		api.PATCH("/taxes/:id", controllers.UpdateTax)
		api.POST("/taxes/", controllers.CreateTax)
		api.DELETE("/taxes/:id", controllers.DeleteTax)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
