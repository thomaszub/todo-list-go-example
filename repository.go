package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(dbname string) (*TodoRepository, error) {
	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&Todo{}); err != nil {
		return nil, err
	}
	return &TodoRepository{db: db}, nil
}

func (r *TodoRepository) GetAllTodos() []Todo {
	todos := make([]Todo, 0)
	r.db.Find(&todos)
	return todos
}

func (r *TodoRepository) GetTodoById(id uint) (*Todo, bool) {
	todo := Todo{}
	result := r.db.Find(&todo, id)
	if result.RowsAffected == 0 {
		return nil, false
	}
	return &todo, true
}

func (r *TodoRepository) AddTodo(name string) Todo {
	todo := Todo{Name: name}
	r.db.Create(&todo)
	return todo
}

func (r *TodoRepository) DeleteTodo(id uint) bool {
	result := r.db.Delete(&Todo{}, id)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}
