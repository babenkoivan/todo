#!/usr/bin/env sh

printf "\033[93m→ Starting GraphQL API server\033[0m\n"

set -a && source configs/.env && go run cmd/graphql.go