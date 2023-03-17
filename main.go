package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alx9110/payment-calculator/controllers"
	"github.com/alx9110/payment-calculator/ext"
	"github.com/alx9110/payment-calculator/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	var env string
	var production bool = false
	env, exists := os.LookupEnv("PAYMENT_CALCULATOR_ENV")
	if !exists {
		env = "DEVELOPMENT"
	}
	if env == "PRODUCTION" {
		production = true
	}
	viper.AddConfigPath("./")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	viper.ReadInConfig()

	port := viper.Get(fmt.Sprintf("%s.port", env))
	connection_string := viper.Get(fmt.Sprintf("%s.connection_string", env))
	// Routing
	models.ConnectDatabase(connection_string)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			fmt.Sprintf("http://localhost:%s", port),
			fmt.Sprintf("http://0.0.0.0:%s", port),
			fmt.Sprintf("http://127.0.0.1:%s", port),
			fmt.Sprintf("http://pay-calc.ru:%s", port),
		},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	if production {
		// Serve frontend static files
		r.Use(static.Serve("/", static.LocalFile("/srv/www", true)))
	}
	// Setup route group for the API
	api := r.Group("/api")
	{
		// JWT auth
		api.POST("/user/token", controllers.GenerateToken)
		api.POST("/user/create", controllers.RegisterUser)
		// Protected API
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
