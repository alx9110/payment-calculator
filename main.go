package main

import (
	"time"

	"github.com/alx9110/payment-calculator/controllers"
	"github.com/alx9110/payment-calculator/ext"
	"github.com/alx9110/payment-calculator/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Routing
	models.ConnectDatabase()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Serve frontend static files
	// r.Use(static.Serve("/", static.LocalFile("web/routing-app/dist/routing-app", true)))
	// Setup route group for the API
	api := r.Group("/api")
	{
		api.POST("/user/token", controllers.GenerateToken)
		api.POST("/user/create", controllers.RegisterUser)
		secured := api.Use(ext.Auth())
		// Records
		secured.GET("/records/", controllers.FindRecords)
		secured.GET("/records/:id", controllers.FindRecord)
		secured.PATCH("/records/:id", controllers.UpdateRecord)
		secured.POST("/records/", controllers.CreateRecord)
		secured.DELETE("/records/:id", controllers.DeleteRecord)

		// Taxes
		secured.GET("/taxes/", controllers.FindTaxes)
		secured.GET("/taxes/:id", controllers.FindTax)
		secured.PATCH("/taxes/:id", controllers.UpdateTax)
		secured.POST("/taxes/", controllers.CreateTax)
		secured.DELETE("/taxes/:id", controllers.DeleteTax)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
