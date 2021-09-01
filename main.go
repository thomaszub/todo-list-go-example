package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	repo, err := NewTodoRepository("todo.db")
	if err != nil {
		panic(err.Error())
	}
	serv := NewTodosService(repo)
	con := NewTodosController(&serv)
	con.RegisterAtGroup(api)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
