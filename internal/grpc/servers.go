package grpc

import (
	"context"
	"github.com/babenkoivan/todo/internal/todos"
	"github.com/babenkoivan/todo/internal/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

var (
	errInvalidCredentials = status.Errorf(codes.NotFound, "incorrect username/password")
	errInvalidInput       = status.Errorf(codes.InvalidArgument, "invalid or duplicated input")
)

type userServiceServer struct {
}

func (u *userServiceServer) Login(ctx context.Context, request *LoginRequest) (*TokenResponse, error) {
	user := users.NewUser(request.Username, request.Password)
	token, err := user.Login()

	if err != nil {
		return nil, errInvalidCredentials
	}

	return newTokenResponse(token), nil
}

func (u *userServiceServer) Create(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	user := users.NewUser(request.Username, request.Password)
	err := user.Create()

	if err != nil {
		return nil, errInvalidInput
	}

	return newUserResponse(user), err
}

func (u *userServiceServer) mustEmbedUnimplementedUserServiceServer() {
}

type todoServiceServer struct {
}

func (t *todoServiceServer) Create(ctx context.Context, request *TodoRequest) (*TodoResponse, error) {
	user := getUser(ctx)
	todo := todos.NewTodo(request.Task, user)
	err := todo.Create()

	if err != nil {
		return nil, errInvalidInput
	}

	return newTodoResponse(todo), nil
}

func (t *todoServiceServer) List(ctx context.Context, empty *emptypb.Empty) (*TodoListResponse, error) {
	user := getUser(ctx)
	allTodos := todos.GetAll(user)

	return newTodoListResponse(allTodos), nil
}

func (t *todoServiceServer) mustEmbedUnimplementedTodoServiceServer() {
}

func StartServer(network, address string) {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(authenticate))
	RegisterUserServiceServer(s, &userServiceServer{})
	RegisterTodoServiceServer(s, &todoServiceServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
