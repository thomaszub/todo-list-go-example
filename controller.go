package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AddTodoRequest struct {
	Name string `json:"name" binding:"required"`
}

type TodoResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type TodosController struct {
	service *TodosService
}

func NewTodosController(service *TodosService) TodosController {
	return TodosController{
		service: service,
	}
}

func (t *TodosController) RegisterAtGroup(group *gin.RouterGroup) {
	group.GET("/todos", t.getAllTodos)
	group.GET("/todos/:id", t.getTodo)
	group.DELETE("/todos/:id", t.deleteTodo)
	group.POST("/todos", t.addTodo)
}

func (t *TodosController) addTodo(c *gin.Context) {
	var request AddTodoRequest
	if err := c.Bind(&request); err != nil {
		return
	}
	newTodo := t.service.AddTodo(request.Name)
	response := mapTodoToResponse(&newTodo)
	c.JSON(http.StatusCreated, response)
}

func (t *TodosController) getAllTodos(c *gin.Context) {
	todos := t.service.GetAllTodos()
	responses := make([]TodoResponse, 0)
	for _, todo := range todos {
		responses = append(responses, mapTodoToResponse(&todo))
	}
	c.JSON(http.StatusOK, responses)
}

func (t *TodosController) getTodo(c *gin.Context) {
	requestId := c.Param("id")
	id, err := strconv.Atoi(requestId)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("%s is not a valid id", requestId)},
		)
		return
	}
	todo, found := t.service.GetTodo(uint(id))
	if !found {
		c.JSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Todo item with id %d not found", id)},
		)
		return
	}
	c.JSON(http.StatusOK, mapTodoToResponse(todo))
}

func (t *TodosController) deleteTodo(c *gin.Context) {
	requestId := c.Param("id")
	id, err := strconv.Atoi(requestId)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("%s is not a valid id", requestId)},
		)
		return
	}
	deleted := t.service.DeleteTodo(uint(id))
	if !deleted {
		c.JSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Todo item with id %d not found", id)},
		)
		return
	}
	c.Status(http.StatusNoContent)
}

func mapTodoToResponse(todo *Todo) TodoResponse {
	return TodoResponse{
		Id:   todo.Model.ID,
		Name: todo.Name,
	}
}
