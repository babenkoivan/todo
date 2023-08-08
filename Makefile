.PHONY: start-db stop-db destroy-db start-rest-server start-graphql-server start-grpc-server protoc help

.DEFAULT_GOAL := help
SHELL := /bin/bash

start-db: ## Start the database container
	@sh scripts/start-db.sh
stop-db: ## Stop the database container
	@sh scripts/stop-db.sh
destroy-db: ## Destroy the database container
	@sh scripts/destroy-db.sh
start-rest-server: ## Start REST API server
	@sh scripts/start-rest-server.sh
start-graphql-server: ## Start GraphQL API server
	@sh scripts/start-graphql-server.sh
start-grpc-server: ## Start gRPC server
	@sh scripts/start-grpc-server.sh
protoc: ## Generate code based on .proto file definition
	@sh scripts/protoc.sh

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'