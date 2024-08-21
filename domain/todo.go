package domain

import "time"

type Todo struct {
	Id        int64     `json:"id" gorm:"column:id;"`
	Title     string    `json:"title" gorm:"column:title;"`
	Status    string    `json:"status" gorm:"column:status;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

type TodoRepository interface {
	GetAll() ([]Todo, error)
	GetById(int) (Todo, error)
	Update(int, Todo) error
	Delete(int) error
	Create(data Todo) (int64, error)
}
