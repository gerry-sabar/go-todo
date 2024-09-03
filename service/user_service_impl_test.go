package service

import (
	"errors"
	domain "go-todo/domain"
	"time"

	domainTest "go-todo/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllFail(t *testing.T) {
	mocking := domainTest.NewUserRepository(t)
	mocking.On("GetAll").Return(nil, errors.New("failed to fetch data"))
	service := UserServiceImpl{
		UserRepository: mocking,
	}

	todos, err := service.GetAll()
	assert.Error(t, err)
	assert.Equal(t, todos, []domain.User(nil))

}

func TestGetAllSuccess(t *testing.T) {
	currentTime := time.Now()
	expected := []domain.User{
		{
			Id:        1,
			Name:      "first user",
			Email:     "first@user.com",
			Password:  "test-password",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
		{
			Id:        2,
			Name:      "second user",
			Email:     "second@user.com",
			Password:  "test-password",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
	}

	mocking := domainTest.NewUserRepository(t)
	mocking.On("GetAll").Return(expected, nil)
	service := UserServiceImpl{
		UserRepository: mocking,
	}
	result, err := service.GetAll()
	assert.Equal(t, expected, result)
	assert.NoError(t, err)
}

func TestCreateSuccess(t *testing.T) {
	user := domain.User{
		Id:       1,
		Name:     "test user",
		Email:    "test@user.com",
		Password: "test-password",
	}

	mocking := domainTest.NewUserRepository(t)
	mocking.On("Create", user).Return(int64(1), nil)
	service := UserServiceImpl{
		UserRepository: mocking,
	}
	result, err := service.Create(user)
	assert.Equal(t, user, result)
	assert.NoError(t, err)

}

func TestCreateFail(t *testing.T) {
	mocking := domainTest.NewUserRepository(t)
	mocking.On("Create", domain.User{}).Return(int64(0), errors.New("failed to create user"))
	service := UserServiceImpl{
		UserRepository: mocking,
	}

	result, err := service.Create(domain.User{})
	assert.Equal(t, domain.User{}, result)
	assert.Error(t, err)
}

func TestUpdateSuccess(t *testing.T) {
	currentTime := time.Now()
	user := domain.User{
		Id:        1,
		Name:      "test user",
		Email:     "test@user.com",
		Password:  "test-password",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	update := domain.User{
		Name:     "edit user",
		Email:    "edit@user.com",
		Password: "edit-password",
	}

	mocking := domainTest.NewUserRepository(t)
	mocking.On("GetById", 1).Return(user, nil)
	mocking.On("Update", 1, update).Return(nil)
	service := UserServiceImpl{
		UserRepository: mocking,
	}

	err := service.Update(1, update)
	assert.NoError(t, err)
}

func TestUpdateNotFound(t *testing.T) {
	mocking := domainTest.NewUserRepository(t)
	mocking.On("GetById", 1).Return(domain.User{}, errors.New("user is not found"))
	update := domain.User{
		Name:     "edit user",
		Email:    "edit@user.com",
		Password: "edit-password",
	}
	service := UserServiceImpl{
		UserRepository: mocking,
	}

	err := service.Update(1, update)
	assert.Error(t, err)
}

func TestUpdateFail(t *testing.T) {
	currentTime := time.Now()
	user := domain.User{
		Id:        1,
		Name:      "test user",
		Email:     "test@user.com",
		Password:  "test-password",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	update := domain.User{
		Name:     "edit user",
		Email:    "edit@user.com",
		Password: "edit-password",
	}

	mocking := domainTest.NewUserRepository(t)
	mocking.On("GetById", 1).Return(user, nil)
	mocking.On("Update", 1, update).Return(errors.New("failed to update user"))
	service := UserServiceImpl{
		UserRepository: mocking,
	}

	err := service.Update(1, update)
	assert.Error(t, err)
}

func TestDeleteSuccess(t *testing.T) {
	currentTime := time.Now()
	user := domain.User{
		Id:        1,
		Name:      "test user",
		Email:     "test@user.com",
		Password:  "test-password",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	mocking := domainTest.NewUserRepository(t)
	mocking.On("GetById", 1).Return(user, nil)
	mocking.On("Delete", 1).Return(nil)
	service := UserServiceImpl{
		UserRepository: mocking,
	}

	err := service.Delete(1)
	assert.NoError(t, err)
}

func TestDeleteUserNotFound(t *testing.T) {
	mocking := domainTest.NewUserRepository(t)
	mocking.On("GetById", 1).Return(domain.User{}, errors.New("user is not found"))
	service := UserServiceImpl{
		UserRepository: mocking,
	}

	err := service.Delete(1)
	assert.Error(t, err)
}

func TestDeleteFail(t *testing.T) {
	currentTime := time.Now()
	user := domain.User{
		Id:        1,
		Name:      "test user",
		Email:     "test@user.com",
		Password:  "test-password",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	mocking := domainTest.NewUserRepository(t)
	mocking.On("GetById", 1).Return(user, nil)
	mocking.On("Delete", 1).Return(errors.New("failed to delete user"))
	service := UserServiceImpl{
		UserRepository: mocking,
	}

	err := service.Delete(1)
	assert.Error(t, err)
}
