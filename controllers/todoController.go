package controllers

import (
	"net/http"
	"todo-app/config"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	config.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}
func GetTodo(c *gin.Context) {
	id := c.Param("id")

	// Declare a variable to hold the result
	var todo models.Todo

	// Fetch the Todo from the database where the ID matches
	if err := config.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		// If there's an error (e.g., record not found), return a 404 error
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found", "error": err.Error()})
		return
	}

	// Return the fetched Todo with status 200 OK
	c.JSON(http.StatusOK, todo)
}
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := config.DB.First(&todo, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.BindJSON(&todo)
	config.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := config.DB.First(&todo, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	config.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
