package main

import (
	_ "Webhook/docs" // Import the documentation generated by swag
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // Correct import
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Webhook API
// @version 1.0
// @description This is a basic API created with Go and Swagger to receive Webhooks.
// @host localhost:8080
// @BasePath /
// @accept json
// @produce json

// @Summary Receives a webhook with data in JSON format
// @Description Receives data from a POST webhook and processes it
// @Tags Webhook
// @Param webhook body map[string]interface{} true "Webhook Data"
// @Success 200 {string} string "Webhook received successfully"
// @Failure 400 {string} string "Error in the webhook"
// @Router /webhook [post]
func main() {
	r := gin.Default()

	// Swagger UI route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Webhook receiving route
	r.POST("/webhook", func(c *gin.Context) {
		var json map[string]interface{}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"message": "Error in the webhook"})
			return
		}
		fmt.Println("Webhook data received:", json)
		c.JSON(200, gin.H{"message": "Webhook received successfully"})
	})

	// Start the server
	r.Run(":8080")
}
