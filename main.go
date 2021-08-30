package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type todo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var todos = make([]todo, 0)

func getAllTodos(w http.ResponseWriter, _ *http.Request) {
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		log.Fatal("Error encoding response body: ", err)
	}
}

func writeTodo(_ http.ResponseWriter, r *http.Request) {
	var newTodo todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		log.Fatal("Error decoding request body: ", err)
		return
	}
	todos = append(todos, newTodo)
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/todos", getAllTodos).Methods(http.MethodGet)
	api.HandleFunc("/todos", writeTodo).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}
