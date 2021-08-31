package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	serv := NewTodosService()
	con := NewTodosController(&serv)
	con.RegisterAtGroup(api)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
