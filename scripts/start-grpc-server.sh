#!/usr/bin/env sh

printf "\033[93m→ Starting gRPC server\033[0m\n"

set -a && source configs/.env && go run cmd/grpc.go