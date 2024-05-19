package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Todo struct represents a Todo item
type Todo struct {
	Id          string `json:"id"`          // Unique identifier for the Todo
	Name        string `json:"name"`        // Name of the Todo
	Description string `json:"description"` // Description of the Todo
	CreatedAt   string `json:"created_at"`  // Timestamp when the Todo was created
	UpdatedAt   string `json:"updated_at"`  // Timestamp when the Todo was updated
	Completed   bool   `json:"completed"`   // Indicates if the Todo is completed
}

var todos = []Todo{
	{Id: uuid.New().String(), Name: "Learn Go", Description: "Description", CreatedAt: time.Now().String(), UpdatedAt: time.Now().String(), Completed: false},
	{Id: uuid.New().String(), Name: "Build a web server", Description: "Description", CreatedAt: time.Now().String(), UpdatedAt: time.Now().String(), Completed: false},
	{Id: uuid.New().String(), Name: "Implement authentication", Description: "Description", CreatedAt: time.Now().String(), UpdatedAt: time.Now().String(), Completed: false},
	{Id: uuid.New().String(), Name: "Deploy to production", Description: "Description", CreatedAt: time.Now().String(), UpdatedAt: time.Now().String(), Completed: false},
	{Id: uuid.New().String(), Name: "Write documentation", Description: "Description", CreatedAt: time.Now().String(), UpdatedAt: time.Now().String(), Completed: false},
}

func main() {
	r := gin.Default()

	// Endpoint to check if the server is running
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD endpoints for Todo
	r.GET("/api/todos", GetTodos)
	r.GET("/api/todo/:id", GetTodoDetail)
	r.PUT("/api/todo/:id", UpdateTodo)
	r.DELETE("/api/todo/:id", DeleteTodo)
	r.POST("/api/todos", CreateTodo)

	r.Run() // Listen and serve on 0.0.0.0:8080
}

// CreateTodo creates a new Todo item
func CreateTodo(c *gin.Context) {
	var todo Todo
	// Parse request body into Todo struct
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a new unique ID for the new Todo
	todo.Id = uuid.New().String()
	todo.CreatedAt = time.Now().String()
	todo.UpdatedAt = time.Now().String()

	// Append the new Todo to the todos slice
	todos = append(todos, todo)

	// Respond with the created Todo
	c.JSON(http.StatusCreated, todo)
}

// GetTodos retrieves all Todo items
func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

// GetTodoDetail retrieves a specific Todo item by ID
func GetTodoDetail(c *gin.Context) {
	id := c.Param("id")

	// Loop through todos to find the matching Todo by ID
	for _, todo := range todos {
		if todo.Id == id {
			// Respond with the matching Todo
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	// Respond with an error if Todo is not found
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo Not Found"})
}

// UpdateTodo updates an existing Todo item
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Loop through todos to find the matching Todo by ID
	for index, todo := range todos {
		if todo.Id == id {
			// Update the Todo with new values
			updatedTodo.Id = todo.Id                    // Keep the original ID
			updatedTodo.CreatedAt = todo.CreatedAt      // Keep the original CreatedAt
			updatedTodo.UpdatedAt = time.Now().String() // change update at time
			todos[index] = updatedTodo
			// Respond with the updated Todo
			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}

	// Respond with an error if Todo is not found
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo Not Found"})
}

// DeleteTodo deletes a Todo item by ID
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	// Loop through todos to find the matching Todo by ID
	for index, todo := range todos {
		if todo.Id == id {
			// Remove the Todo from the todos slice
			todos = append(todos[:index], todos[index+1:]...)
			// Respond with success message
			c.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted Todo"})
			return
		}
	}

	// Respond with an error if Todo is not found
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo Not Found"})
}
