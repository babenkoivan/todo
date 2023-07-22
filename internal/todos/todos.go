package todos

import (
	"github.com/babenkoivan/todo/internal/db"
	"github.com/babenkoivan/todo/internal/users"
	"time"
)

type Todo struct {
	ID        int64     `json:"id"`
	Task      string    `json:"task"`
	CreatedAt time.Time `json:"createdAt"`
	User      *users.User
}

func (t *Todo) Create() {
	now := time.Now().UTC()
	res, _ := db.Connection.Exec("INSERT INTO todos (user_id, task, created_at) VALUES (?, ?, ?)", t.User.ID, t.Task, now)
	id, _ := res.LastInsertId()

	t.ID = id
	t.CreatedAt = now
}

func GetAll(user *users.User) []Todo {
	todos := make([]Todo, 0)

	res, _ := db.Connection.Query("SELECT id, task, created_at FROM todos WHERE user_id = ?", user.ID)
	defer res.Close()

	for res.Next() {
		todo := Todo{User: user}
		res.Scan(&todo.ID, &todo.Task, &todo.CreatedAt)
		todos = append(todos, todo)
	}

	return todos
}
