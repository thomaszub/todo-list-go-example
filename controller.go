package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	group.POST("/todos", t.addTodo)
}

func (t *TodosController) getAllTodos(c *gin.Context) {
	todos := t.service.GetAllTodos()
	var responses []TodoResponse
	for _, todo := range todos {
		responses = append(responses, mapTodoToResponse(&todo))
	}
	c.JSON(http.StatusOK, responses)
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

func mapTodoToResponse(todo *Todo) TodoResponse {
	return TodoResponse{
		Id:   todo.Id,
		Name: todo.Name,
	}
}
