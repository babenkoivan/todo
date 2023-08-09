# Todo

This is a simple todo application written in Go that provides 3 different interfaces to manage your todo list: REST API, GraphQL and gRPC.

There are 4 different actions available to clients:

* create new user
* login to the application
* create new todo
* get all todos

The last two actions are user-specific and require authentication, which is based on JWT tokens.

## Configuration

The application relies on the configuration file, which must be located in the `configs` folder. 
To create one, just copy `configs/.env.dist` to `configs/.env`.
You can adjust the database credentials and the application port, or leave the recommended settings. 

## Setup

Before you start making queries, you need to start a database and an API server of your choice:

```bash
# start the database
make start-db

# start REST API server
make start-rest-server

# start GraphQL API server
make start-graphql-server

# start gRPC server
make start-grpc-server
```

## Usage

You can use `Postman` to perform all types of queries: just import desired API definition from `api` folder and create a new collection out of it.

The default base URL for REST API and GraphQL is `localhost:8080/api`, and `localhost:8080` for gRPC. The authentication token received via `login` action can be provided in the `Authorization` header.

Below you can find the demonstration of the application using three different interfaces. 

### REST API

<img alt="REST API" src="assets/rest.gif" width="830"/>

### GraphQL

<img alt="GraphQL" src="assets/graphql.gif" width="830"/>

### gRPC

<img alt="gRPC" src="assets/grpc.gif" width="830"/>