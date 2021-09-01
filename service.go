package main

type TodosService struct {
	repo *TodoRepository
}

func NewTodosService(repo *TodoRepository) TodosService {
	return TodosService{
		repo: repo,
	}
}

func (s *TodosService) GetAllTodos() []Todo {
	return s.repo.GetAllTodos()
}

func (s *TodosService) GetTodo(id uint) (*Todo, bool) {
	return s.repo.GetTodoById(id)
}

func (s *TodosService) AddTodo(name string) Todo {
	return s.repo.AddTodo(name)
}

func (s *TodosService) DeleteTodo(id uint) bool {
	return s.repo.DeleteTodo(id)
}
