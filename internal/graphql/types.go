package graphql

import (
	"github.com/graphql-go/graphql"
	"time"
)

var dateTime = graphql.NewScalar(
	graphql.ScalarConfig{
		Name: "DateTime",
		Serialize: func(value interface{}) interface{} {
			return value.(time.Time).String()
		},
	},
)

var user = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: dateTime,
			},
		},
	},
)

var todo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"task": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: dateTime,
			},
			"user": &graphql.Field{
				Type: user,
			},
		},
	},
)

var token = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Token",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var mutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"login": &graphql.Field{
				Type: token,
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: login,
			},
			"createUser": &graphql.Field{
				Type: user,
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: createUser,
			},
			"createTodo": &graphql.Field{
				Type: todo,
				Args: graphql.FieldConfigArgument{
					"task": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: createTodo,
			},
		},
	},
)

var query = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"todos": &graphql.Field{
				Type:    graphql.NewList(todo),
				Resolve: getTodos,
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	},
)
