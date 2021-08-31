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
	Id   int    `json:"id"`
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)

	}
	todo, found := t.service.GetTodo(id)
	if !found {
		c.JSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Todo item with id %d not found", id)},
		)
		return
	}
	c.JSON(http.StatusOK, mapTodoToResponse(todo))
}

func mapTodoToResponse(todo *Todo) TodoResponse {
	return TodoResponse{
		Id:   todo.Id,
		Name: todo.Name,
	}
}
