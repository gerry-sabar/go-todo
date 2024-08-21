package service

import (
	"errors"
	domain "go-todo/domain"

	domainTest "go-todo/mock/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
// BDD behavior example
func TestGetAllFail(t *testing.T) {
	mocking := domainTest.NewTodoRepository(t)
	mocking.On("GetAll").Return(nil, errors.New("failed to update todo"))
	service := &srv.TodoServiceImpl{
		TodoRepository: mocking,
	}

	todos, err := service.GetAll()
	assert.Error(t, err)
	assert.Equal(t, todos, []domain.Todo(nil))

}

func TestGetAllSuccess(t *testing.T) {
	expected := []domain.Todo{
		{
			Id:        1,
			Title:     "first todo",
			Status:    "Doing",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:        2,
			Title:     "second todo",
			Status:    "Doing",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mocking := domainTest.NewTodoRepository(t)
	mocking.On("GetAll").Return(expected, nil)
	service := &srv.TodoServiceImpl{
		TodoRepository: mocking,
	}
	result, err := service.GetAll()
	assert.Equal(t, expected, result)
	assert.NoError(t, err)
}

func TestCreateSuccess(t *testing.T) {
	todo := domain.Todo{
		Id:     1,
		Title:  "first todo",
		Status: "Doing",
	}

	mocking := domainTest.NewTodoRepository(t)
	mocking.On("Create", todo).Return(int64(1), nil)
	service := &srv.TodoServiceImpl{
		TodoRepository: mocking,
	}
	result, err := service.Create(todo)
	assert.Equal(t, todo, result)
	assert.NoError(t, err)

}

func TestCreateFail(t *testing.T) {
	mocking := domainTest.NewTodoRepository(t)
	mocking.On("Create", domain.Todo{}).Return(int64(0), errors.New("failed to create todo"))
	service := &srv.TodoServiceImpl{
		TodoRepository: mocking,
	}

	result, err := service.Create(domain.Todo{})
	assert.Equal(t, domain.Todo{}, result)
	assert.Error(t, err)

}

func TestUpdateSuccess(t *testing.T) {
	todo := domain.Todo{
		Id:        1,
		Title:     "first todo",
		Status:    "Doing",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	update := domain.Todo{
		Title:  "edit first todo",
		Status: "Doing",
	}

	mocking := domainTest.NewTodoRepository(t)
	mocking.On("GetById", 1).Return(todo, nil)
	mocking.On("Update", 1, update).Return(nil)
	service := &srv.TodoServiceImpl{
		TodoRepository: mocking,
	}

	err := service.Update(1, update)
	assert.NoError(t, err)
}

func TestUpdateTodoNotFound(t *testing.T) {
	mocking := domainTest.NewTodoRepository(t)
	mocking.On("GetById", 1).Return(domain.Todo{}, errors.New("todo is not found"))
	update := domain.Todo{
		Title:  "edit first todo",
		Status: "Doing",
	}
	service := &srv.TodoServiceImpl{
		TodoRepository: mocking,
	}

	err := service.Update(1, update)
	assert.Error(t, err)
}

func TestUpdateFail(t *testing.T) {
	todo := domain.Todo{
		Id:        1,
		Title:     "first todo",
		Status:    "Doing",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	update := domain.Todo{
		Title:  "edit first todo",
		Status: "Doing",
	}

	mocking := domainTest.NewTodoRepository(t)
	mocking.On("GetById", 1).Return(todo, nil)
	mocking.On("Update", 1, update).Return(errors.New("failed to update todo"))
	service := &srv.TodoServiceImpl{
		TodoRepository: mocking,
	}

	err := service.Update(1, update)
	assert.Error(t, err)
}

func TestDeleteSuccess(t *testing.T) {
	todo := domain.Todo{
		Id:        1,
		Title:     "first todo",
		Status:    "Doing",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mocking := domainTest.NewTodoRepository(t)
	mocking.On("GetById", 1).Return(todo, nil)
	mocking.On("Delete", 1).Return(nil)
	service := &srv.TodoServiceImpl{
		TodoRepository: mocking,
	}

	err := service.Delete(1)
	assert.NoError(t, err)
}

func TestDeleteTodoNotFound(t *testing.T) {
	mocking := domainTest.NewTodoRepository(t)
	mocking.On("GetById", 1).Return(domain.Todo{}, errors.New("todo is not found"))
	service := &srv.TodoServiceImpl{
		TodoRepository: mocking,
	}

	err := service.Delete(1)
	assert.Error(t, err)
}

func TestDeleteFail(t *testing.T) {
	todo := domain.Todo{
		Id:        1,
		Title:     "first todo",
		Status:    "Doing",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mocking := domainTest.NewTodoRepository(t)
	mocking.On("GetById", 1).Return(todo, nil)
	mocking.On("Delete", 1).Return(errors.New("failed to delete todo"))
	service := &srv.TodoServiceImpl{
		TodoRepository: mocking,
	}

	err := service.Delete(1)
	assert.Error(t, err)
}
*/

// Tabel driven test case example
func TestTodoGetAll(t *testing.T) {
	currentTime := time.Now()
	tests := []struct {
		name  string
		todos []domain.Todo
		err   error
	}{
		{
			"success case",
			[]domain.Todo{
				{
					Id:        1,
					Title:     "first todo",
					Status:    "Doing",
					CreatedAt: currentTime,
					UpdatedAt: currentTime,
				},
				{
					Id:        2,
					Title:     "second todo",
					Status:    "Doing",
					CreatedAt: currentTime,
					UpdatedAt: currentTime,
				},
			},
			nil,
		},
		{"failed case", []domain.Todo(nil), errors.New("failed")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocking := domainTest.NewTodoRepository(t)
			mocking.On("GetAll").Return(tt.todos, tt.err)
			service := &TodoServiceImpl{
				TodoRepository: mocking,
			}
			result, err := service.GetAll()
			assert.Equal(t, tt.todos, result)
			if tt.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
