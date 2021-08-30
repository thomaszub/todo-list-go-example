package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

var todos = make([]todo, 0)

var nextId = 0

func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func writeTodo(c *gin.Context) {
	var newTodo todo
	if err := c.Bind(&newTodo); err != nil {
		return
	}
	newTodo.Id = nextId
	nextId++
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func main() {
	r := gin.Default()
	api := r.Group("/api")
	api.GET("/todos", getAllTodos)
	api.POST("/todos", writeTodo)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
