package service

import "go-todo/domain"

type UserService interface {
	Create(data domain.User) (domain.User, error)
	GetById(id int) (domain.User, error)
	Update(id int, data domain.User) error
	GetAll() ([]domain.User, error)
	Delete(id int) error
}
