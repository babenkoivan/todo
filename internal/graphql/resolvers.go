package graphql

import (
	"errors"
	"github.com/babenkoivan/todo/internal/todos"
	"github.com/babenkoivan/todo/internal/users"
	"github.com/graphql-go/graphql"
)

var (
	errInvalidCredentials = errors.New("invalid credentials")
	errUnauthorized       = errors.New("unauthorized")
)

func login(params graphql.ResolveParams) (interface{}, error) {
	user := users.NewUser(params.Args["username"].(string), params.Args["password"].(string))
	token, err := user.Login()

	if err != nil {
		return nil, errInvalidCredentials
	}

	return map[string]string{"token": token}, nil
}

func createUser(params graphql.ResolveParams) (interface{}, error) {
	user := users.NewUser(params.Args["username"].(string), params.Args["password"].(string))
	user.Create()

	return user, nil
}

func createTodo(params graphql.ResolveParams) (interface{}, error) {
	user, ok := getUser(params.Context)

	if !ok {
		return nil, errUnauthorized
	}

	todo := todos.NewTodo(params.Args["task"].(string), user)
	todo.Create()

	return todo, nil
}

func getTodos(params graphql.ResolveParams) (interface{}, error) {
	user, ok := getUser(params.Context)

	if !ok {
		return nil, errUnauthorized
	}

	return todos.GetAll(user), nil
}
