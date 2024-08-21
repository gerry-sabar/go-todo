package repository

import (
	"go-todo/config"
	"go-todo/domain"
)

type TodoRepository struct {
	todoItems []domain.Todo
	todoItem  domain.Todo
}

func NewTodoRepository() *TodoRepository {
	var domain []domain.Todo
	return &TodoRepository{todoItems: domain}
}

func (r *TodoRepository) Create(todo domain.Todo) (int64, error) {
	query := `INSERT INTO todo_items (title, status, created_at, updated_at) VALUES (?, ?, now(), now())`
	row, err := config.DB.Exec(query, todo.Title, todo.Status)
	if err != nil {
		return 0, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *TodoRepository) Update(id int, todo domain.Todo) error {
	_, err := config.DB.Exec("UPDATE todo_items SET title = ?, status = ? where id = ?", todo.Title, todo.Status, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM todo_items WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) GetById(id int) (domain.Todo, error) {
	err := config.DB.QueryRow("SELECT * FROM todo_items WHERE id = ?", id).Scan(&r.todoItem.Id, &r.todoItem.Title, &r.todoItem.Status, &r.todoItem.CreatedAt, &r.todoItem.UpdatedAt)
	if err != nil {
		return r.todoItem, err
	}

	return r.todoItem, nil
}

func (r *TodoRepository) GetAll() ([]domain.Todo, error) {
	var items []domain.Todo
	rows, err := config.DB.Query("SELECT * FROM todo_items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item domain.Todo
		if err := rows.Scan(&item.Id, &item.Title, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
