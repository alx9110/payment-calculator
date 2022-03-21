package main

import (
	"github.com/gin-gonic/gin"

	"app/controllers"
	"app/models"
)

func main() {
	// Routing
	models.ConnectDatabase()
	r := gin.Default()
	r.GET("/", controllers.FindProducts)
	r.GET("/:id", controllers.FindProduct)
	r.PATCH("/:id", controllers.UpdateProduct)
	r.POST("/", controllers.CreateProduct)
	r.DELETE("/:id", controllers.DeleteProduct)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
