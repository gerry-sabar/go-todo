package service

import (
	"go-todo/domain"
)

type UserServiceImpl struct {
	UserRepository domain.UserRepository
}

func NewUserServiceImpl(userRepository domain.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (s *UserServiceImpl) Create(data domain.User) (domain.User, error) {
	user := domain.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	id, err := s.UserRepository.Create(data)
	if err != nil {
		return domain.User{}, err
	}

	user.Id = id
	return user, nil
}

func (s *UserServiceImpl) GetAll() ([]domain.User, error) {
	users, err := s.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserServiceImpl) GetById(id int) (domain.User, error) {
	user, err := s.UserRepository.GetById(id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *UserServiceImpl) Update(id int, data domain.User) error {
	_, err := s.UserRepository.GetById(id)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	err = s.UserRepository.Update(id, user)
	return err
}

func (s *UserServiceImpl) Delete(id int) error {
	_, err := s.UserRepository.GetById(id)
	if err != nil {
		return err
	}

	err = s.UserRepository.Delete(id)
	return err
}
