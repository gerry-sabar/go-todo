package server

import (
	"go-todo/domain"
	"go-todo/mock/service"
	"net/http"
	"net/http/httptest"
	"time"

	"testing"

	"github.com/go-playground/assert"
)

func TestGetTodo(t *testing.T) {
	mocking := service.NewTodoService(t)
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
	RouterService := RouterService{
		TodoService: mocking,
	}

	mocking.On("GetAll").Return(expected, nil)
	router := SetupRouter(RouterService)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/todos", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
