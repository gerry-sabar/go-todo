package service

import (
	"go-todo/domain"
)

type TodoServiceImpl struct {
	TodoRepository domain.TodoRepository
}

func NewTodoServiceImpl(todoRepository domain.TodoRepository) *TodoServiceImpl {
	return &TodoServiceImpl{TodoRepository: todoRepository}
}

func (s *TodoServiceImpl) Create(data domain.Todo) (domain.Todo, error) {
	todo := domain.Todo{
		Title:  data.Title,
		Status: data.Status,
	}

	id, err := s.TodoRepository.Create(data)
	if err != nil {
		return domain.Todo{}, err
	}

	todo.Id = id
	return todo, nil
}

func (s *TodoServiceImpl) GetAll() ([]domain.Todo, error) {
	todos, err := s.TodoRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoServiceImpl) GetById(id int) (domain.Todo, error) {
	todo, err := s.TodoRepository.GetById(id)
	if err != nil {
		return domain.Todo{}, err
	}

	return todo, nil
}

func (s *TodoServiceImpl) Update(id int, data domain.Todo) error {
	_, err := s.TodoRepository.GetById(id)
	if err != nil {
		return err
	}

	todo := domain.Todo{
		Title:  data.Title,
		Status: data.Status,
	}

	err = s.TodoRepository.Update(id, todo)
	return err
}

func (s *TodoServiceImpl) Delete(id int) error {
	_, err := s.TodoRepository.GetById(id)
	if err != nil {
		return err
	}

	err = s.TodoRepository.Delete(id)
	return err
}
