package main

import (
	"fmt"
	"log"
	"os"

	"todo-app/config"
	"todo-app/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Example usage of environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	jwtSecret := os.Getenv("JWT_SECRET")

	fmt.Println("Database User:", dbUser)
	fmt.Println("Database Password:", dbPassword)
	fmt.Println("Database Name:", dbName)
	fmt.Println("JWT Secret:", jwtSecret)

	config.ConnectDatabase()

	// Initialize the router
	router := routes.SetupRouter()

	// Start the server on port 8080
	router.Run(":8080") // This will start the server on http://localhost:8080
}

// Define routes
// r.GET("/", func(c *gin.Context) {
// var token, err1 = utils.GenerateToken(123456)
// if err1 != nil {
// 	return
// }
// fmt.Println(token)
// fmt.Println(utils.ValidateToken(token))
// config.ConnectDatabase()
// var user []models.User

// // Query the "Test" table directly
// err := config.DB.Find(&user).Error
// if err != nil {
// 	c.JSON(500, gin.H{"message": "Error fetching data", "error": err})
// 	return
// }

// // Print the results to the console
// fmt.Println(user)
// c.JSON(200, gin.H{
// 	"message": "Welcome to the Todo App API!",
// 	"name":    config.DB.Name(),
// 	"table":   "User",
// })

// })

// Initialize Gin and start the server
// r := gin.Default()

// r.GET("/status", func(c *gin.Context) {
// c.JSON(200, gin.H{
// 	"status": "OK",
// })
// })
