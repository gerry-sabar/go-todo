package service

import "go-todo/domain"

type TodoService interface {
	Create(data domain.Todo) (domain.Todo, error)
	GetById(id int) (domain.Todo, error)
	Update(id int, data domain.Todo) error
	GetAll() ([]domain.Todo, error)
	Delete(id int) error
}
