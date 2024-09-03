package server

import (
	"encoding/json"
	"errors"
	"go-todo/domain"
	service "go-todo/service/mocks"

	// "go-todo/mock/service"
	"io"
	"net/http"
	"net/http/httptest"
	"time"

	"testing"

	"github.com/go-playground/assert"
)

func TestTodoGetAllSuccess(t *testing.T) {
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

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	mBody, _ := json.Marshal(expected)
	result := "{\"data\":" + string(mBody) + "}"
	assert.Equal(t, string(body), result)
}

func TestTodoGetAllFailed(t *testing.T) {
	mocking := service.NewTodoService(t)
	var expected []domain.Todo
	RouterService := RouterService{
		TodoService: mocking,
	}

	mocking.On("GetAll").Return(expected, errors.New("error from GetAll"))
	router := SetupRouter(RouterService)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/todos", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	result := "{\"error\":\"error from GetAll\"}"
	assert.Equal(t, string(body), result)
}
