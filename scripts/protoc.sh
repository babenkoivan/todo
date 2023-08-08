#!/usr/bin/env sh

printf "\033[93m→ Generating code\033[0m\n"

protoc --proto_path=api/grpc \
  --go_out=internal/grpc \
  --go_opt=paths=source_relative \
  --go-grpc_out=internal/grpc \
  --go-grpc_opt=paths=source_relative \
  index.proto

printf "\033[92m✔︎ The code has been generated\033[0m\n"
