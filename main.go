package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var todos = []Todo{}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.POST("/todos", createTodo)
	router.PUT("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)

	router.Run("localhost:8080")
}

// getTodos responds with the list of all todos as JSON.
func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

// createTodo adds a todo from JSON received in the request body.
func createTodo(c *gin.Context) {
	var newTodo Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	newTodo.ID = uuid.New().String() // Generate a unique ID for the new todo.
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// updateTodo updates the completed status of a todo given its ID.
func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo Todo

	if err := c.BindJSON(&updatedTodo); err != nil {
		return
	}

	for i, t := range todos {
		if t.ID == id {
			todos[i].IsCompleted = updatedTodo.IsCompleted
			c.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

// deleteTodo deletes a todo given its ID.
func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "todo deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}
