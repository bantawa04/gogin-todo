package controllers

import (
	"net/http"
	"todo-api/db"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

// ListTodos is a controller function that handles HTTP requests to retrieve a list of todos from the database
func ListTodos(c *gin.Context) {
	// Declare a variable to hold the list of todos
	var todo []models.Todo

	// Connect to the database using the ConnectDatabase function from the db package
	db := db.ConnectDatabase()

	// Retrieve all todos from the database and store them in the todo variable
	db.Find(&todo)

	// Return the list of todos as a JSON response with a HTTP status code of 200 (OK)
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// GetTodo is a controller function that handles HTTP requests to retrieve a single todo from the database
func GetTodo(c *gin.Context) {
	// Get the id of the todo from the request URL parameters
	id := c.Param("id")

	// Declare a variable to hold the single todo
	var todo models.Todo

	// Connect to the database using the ConnectDatabase function from the db package
	db := db.ConnectDatabase()

	// Retrieve the todo with the specified id from the database and store it in the todo variable
	db.Where("id = ?", id).Find(&todo)

	// Return the single todo as a JSON response with a HTTP status code of 200 (OK)
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func CreateTodo(c *gin.Context) {
	var input models.CreateTodo
	// Parse the request body into the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create a new todo with the input data
	todo := models.Todo{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
	}
	db := db.ConnectDatabase()
	// Insert the new todo into the database
	if err := db.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the new todo as a JSON response
	c.JSON(http.StatusOK, gin.H{"data": todo})

}

func UpdateTodo(c *gin.Context) {
	// Get the ID of the todo from the request URL parameters
	id := c.Param("id")

	// Declare a variable to hold the input data
	var input models.UpdateTodo

	// Parse the request body and bind the data to the input variable
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Connect to the database
	db := db.ConnectDatabase()

	// Find the todo with the specified ID in the database
	var todo models.Todo
	if err := db.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// Update the todo status based on the input data
	todo.Status = input.Status

	// Save the updated todo to the database
	if err := db.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the updated todo as a JSON response with a HTTP status code of 200 (OK)
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTodo(c *gin.Context) {
	// Get the ID of the todo to be deleted from the URL parameters
	id := c.Param("id")

	// Connect to the database
	db := db.ConnectDatabase()

	// Find the todo with the specified ID in the database
	var todo models.Todo
	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// Delete the todo from the database
	if err := db.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message as a JSON response with a HTTP status code of 200 (OK)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
