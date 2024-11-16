package routes

import (
	"todo-app/controllers"
	"todo-app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes
	r.Group("/todos", middlewares.JWTAuth()).
		POST("/", controllers.CreateTodo).
		GET("/", controllers.GetTodos).
		GET("/:id", controllers.GetTodo).
		PUT("/:id", controllers.UpdateTodo).
		DELETE("/:id", controllers.DeleteTodo)

	return r
}
