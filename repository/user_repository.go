package repository

import (
	"go-todo/config"
	"go-todo/domain"
)

type UserRepository struct {
	users []domain.User
	user  domain.User
}

func NewUserRepository() *UserRepository {
	var domain []domain.User
	return &UserRepository{users: domain}
}

func (r *UserRepository) Create(user domain.User) (int64, error) {
	query := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, now(), now())`
	row, err := config.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *UserRepository) Update(id int, user domain.User) error {
	_, err := config.DB.Exec("UPDATE users SET name = ?, email = ?, password = ? where id = ?", user.Name, user.Email, user.Password, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetById(id int) (domain.User, error) {
	err := config.DB.QueryRow("SELECT `id`, `name`, `email`, `password`, `created_at`, `updated_at` FROM users WHERE id = ?", id).Scan(&r.user.Id, &r.user.Name, &r.user.Email, &r.user.Password, &r.user.CreatedAt, &r.user.UpdatedAt)
	if err != nil {
		return r.user, err
	}

	return r.user, nil
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	rows, err := config.DB.Query("SELECT `id`, `name`, `email`, `password`, `created_at`, `updated_at` FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil

}
