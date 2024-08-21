package domain

import "time"

type User struct {
	Id        int64     `json:"id" gorm:"column:id;"`
	Name      string    `json:"name" gorm:"column:name;"`
	Email     string    `json:"email" gorm:"column:email;"`
	Password  string    `json:"password" gorm:"column:password;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(int) (User, error)
	Update(int, User) error
	Delete(int) error
	Create(user User) (int64, error)
}
