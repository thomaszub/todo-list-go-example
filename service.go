package main

type Todo struct {
	Id   int
	Name string
}

type TodosService struct {
	nextId int
	todos  []Todo
}

func NewTodosService() TodosService {
	return TodosService{
		nextId: 0,
		todos:  make([]Todo, 0),
	}
}

func (s *TodosService) GetAllTodos() []Todo {
	return s.todos
}

func (s *TodosService) AddTodo(name string) Todo {
	newTodo := Todo{
		Id:   s.nextId,
		Name: name,
	}
	s.nextId++
	s.todos = append(s.todos, newTodo)
	return newTodo
}
