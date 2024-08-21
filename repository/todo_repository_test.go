package repository

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"go-todo/config"
	"go-todo/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTodoCreate(t *testing.T) {
	tests := []struct {
		name string
		id   int64
		err  error
	}{
		{
			"success case",
			1,
			nil,
		},
		{"failed case", 0, errors.New("failed")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			config.DB = db
			defer config.DB.Close()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			if tt.name == "success case" {
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO todo_items (title, status, created_at, updated_at) VALUES (?, ?, now(), now())")).
					WithArgs("testing", "Doing").
					WillReturnResult(sqlmock.NewResult(1, 1))
			}

			if tt.name == "failed case" {
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO todo_items (title, status, created_at, updated_at) VALUES (?, ?, now(), now())")).
					WithArgs("testing", "Doing").
					WillReturnError(tt.err)
			}

			var repo TodoRepository
			var todo domain.Todo
			todo.Title = "testing"
			todo.Status = "Doing"
			id, err := repo.Create(todo)
			if tt.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.id, id)
		})
	}
}

func TestTodoUpdate(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			"success case",
			nil,
		},
		{"failed case", errors.New("failed")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			config.DB = db
			defer config.DB.Close()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			if tt.name == "success case" {
				mock.ExpectExec(regexp.QuoteMeta("UPDATE todo_items SET title = ?, status = ? where id = ?")).
					WithArgs("update test", "Finished", int64(1)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			}

			if tt.name == "failed case" {
				mock.ExpectExec(regexp.QuoteMeta("UPDATE todo_items SET title = ?, status = ? where id = ?")).
					WithArgs("update test", "Finished", int64(1)).
					WillReturnError(tt.err)
			}

			var repo TodoRepository
			var todo domain.Todo
			todo.Title = "update test"
			todo.Status = "Finished"
			err = repo.Update(1, todo)
			if tt.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestTodoDelete(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			"success case",
			nil,
		},
		{"failed case", errors.New("failed")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			config.DB = db
			defer config.DB.Close()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			if tt.name == "success case" {
				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM todo_items WHERE id = ?")).
					WithArgs(int64(1)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			}

			if tt.name == "failed case" {
				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM todo_items WHERE id = ?")).
					WithArgs(int64(1)).
					WillReturnError(tt.err)
			}

			var repo TodoRepository
			err = repo.Delete(1)
			if tt.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestTodoGetById(t *testing.T) {
	currentTime := time.Now()
	tests := []struct {
		name string
		todo domain.Todo
		err  error
	}{
		{
			"success case",
			domain.Todo{
				Id:        1,
				Title:     "Testing",
				Status:    "Doing",
				CreatedAt: currentTime,
				UpdatedAt: currentTime,
			},
			nil,
		},
		{"failed case", domain.Todo{}, errors.New("failed")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			config.DB = db
			defer config.DB.Close()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			rows := sqlmock.NewRows([]string{"id", "title", "status", "updated_at", "created_at"}).AddRow(1, "Testing", "Doing", currentTime, currentTime)
			if tt.name == "success case" {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM todo_items WHERE id = ?")).
					WithArgs(1).
					WillReturnRows(rows)
			}

			if tt.name == "failed case" {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM todo_items WHERE id = ?")).
					WithArgs(1).
					WillReturnError(tt.err)
			}

			var repo TodoRepository
			todo, err := repo.GetById(1)
			if tt.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.todo, todo)
		})
	}
}

func TestTodoGetAll(t *testing.T) {
	currentTime := time.Now()
	var todos []domain.Todo
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
					Title:     "Testing",
					Status:    "Doing",
					CreatedAt: currentTime,
					UpdatedAt: currentTime,
				},
				{
					Id:        2,
					Title:     "Testing2",
					Status:    "Doing",
					CreatedAt: currentTime,
					UpdatedAt: currentTime,
				},
			},
			nil,
		},
		{"failed case", todos, errors.New("failed")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			config.DB = db
			defer config.DB.Close()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			if tt.name == "success case" {
				rows := sqlmock.NewRows([]string{"id", "title", "status", "created_at", "updated_at"}).
					AddRow(1, "Testing", "Doing", currentTime, currentTime).
					AddRow(2, "Testing2", "Doing", currentTime, currentTime)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM todo_items")).
					WillReturnRows(rows)
			}

			if tt.name == "failed case" {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM todo_items")).
					WillReturnError(tt.err)
			}

			var repo TodoRepository
			todos, err := repo.GetAll()
			if tt.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.todos, todos)
		})
	}
}
