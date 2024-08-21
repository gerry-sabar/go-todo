package repository

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"go-todo/config"
	"go-todo/domain"

	// "go-todo/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, now(), now())")).
		WithArgs("test", "test@test.com", "test-password").
		WillReturnResult(sqlmock.NewResult(1, 1))

	var repo UserRepository
	var user domain.User
	user.Name = "test"
	user.Email = "test@test.com"
	user.Password = "test-password"
	id, err := repo.Create(user)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserCreateError(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, now(), now())")).
		WithArgs("test", "test@test.com", "test-password").
		WillReturnError(fmt.Errorf("db error"))

	var repo UserRepository
	var user domain.User
	user.Name = "test"
	user.Email = "test@test.com"
	user.Password = "test-password"
	id, err := repo.Create(user)

	assert.Error(t, err)
	assert.Equal(t, int64(0), id)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET name = ?, email = ?, password = ? where id = ?")).
		WithArgs("update user", "test@update.com", "edit-password", int64(1)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var repo UserRepository
	var user domain.User
	user.Name = "update user"
	user.Email = "test@update.com"
	user.Password = "edit-password"

	err = repo.Update(1, user)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserUpdateError(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET name = ?, email = ?, password = ? where id = ?")).
		WithArgs("update user", "test@update.com", "edit-password", int64(1)).
		WillReturnError(fmt.Errorf("db error"))

	var repo UserRepository
	var user domain.User
	user.Name = "update user"
	user.Email = "test@update.com"
	user.Password = "edit-password"

	err = repo.Update(1, user)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE id = ?")).
		WithArgs(int64(1)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var repo UserRepository
	err = repo.Delete(1)
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserDeleteError(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE id = ?")).
		WithArgs(int64(1)).
		WillReturnError(fmt.Errorf("db error"))

	var repo UserRepository
	err = repo.Delete(1)
	assert.Error(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	currentTime := time.Now()
	expected := domain.User{
		Id:        1,
		Name:      "testing",
		Email:     "test@test.com",
		Password:  "test-password",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
		AddRow(1, "testing", "test@test.com", "test-password", currentTime, currentTime)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `name`, `email`, `password`, `created_at`, `updated_at` FROM users WHERE id = ?")).
		WithArgs(1).
		WillReturnRows(rows)

	var repo UserRepository
	user, err := repo.GetById(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}

func TestUserGetByIDError(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	expected := domain.User{}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `name`, `email`, `password`, `created_at`, `updated_at` FROM users WHERE id = ?")).
		WithArgs(1).
		WillReturnError(fmt.Errorf("db error"))

	var repo UserRepository
	user, err := repo.GetById(1)

	assert.Error(t, err)
	assert.Equal(t, expected, user)
}

func TestUserGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	currentTime := time.Now()
	expected := []domain.User{
		{
			Id:        1,
			Name:      "testing",
			Email:     "test@test.com",
			Password:  "test-password",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
		{
			Id:        2,
			Name:      "testing2",
			Email:     "test2@test.com",
			Password:  "test-password",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
		AddRow(1, "testing", "test@test.com", "test-password", currentTime, currentTime).
		AddRow(2, "testing2", "test2@test.com", "test-password", currentTime, currentTime)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `name`, `email`, `password`, `created_at`, `updated_at` FROM users")).
		WillReturnRows(rows)

	var repo UserRepository
	users, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, expected, users)
}

func TestUserGetAllError(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	var expected []domain.User
	mock.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `name`, `email`, `password`, `created_at`, `updated_at` FROM users")).
		WillReturnError(fmt.Errorf("db error"))

	var repo UserRepository
	users, err := repo.GetAll()
	assert.Error(t, err)
	assert.Equal(t, expected, users)
}
