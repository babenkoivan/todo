package grpc

import (
	"github.com/babenkoivan/todo/internal/todos"
	"github.com/babenkoivan/todo/internal/users"
)

func newTokenResponse(token string) *TokenResponse {
	return &TokenResponse{
		Token: token,
	}
}

func newUserResponse(user *users.User) *UserResponse {
	return &UserResponse{
		Id:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.String(),
	}
}

func newTodoResponse(todo *todos.Todo) *TodoResponse {
	return &TodoResponse{
		Id:        todo.ID,
		Task:      todo.Task,
		CreatedAt: todo.CreatedAt.String(),
		User:      newUserResponse(todo.User),
	}
}

func newTodoListResponse(allTodos []*todos.Todo) *TodoListResponse {
	response := &TodoListResponse{}

	for _, todo := range allTodos {
		response.Todos = append(response.Todos, newTodoResponse(todo))
	}

	return response
}
