#!/usr/bin/env sh

printf "\033[93m→ Starting the database\033[0m\n"

. configs/.env
CONTAINER_NAME="${DB_NAME}-db"

docker start "${CONTAINER_NAME}" > /dev/null 2>&1 || docker run -d \
  --name "${CONTAINER_NAME}" \
  --platform linux/x86_64 \
  -p "${DB_PORT}":3306 \
  -e MYSQL_RANDOM_ROOT_PASSWORD=yes \
  -e MYSQL_DATABASE="${DB_NAME}" \
  -e MYSQL_USER="${DB_USER}" \
  -e MYSQL_PASSWORD="${DB_PASSWORD}" \
  -v "$(pwd)/init/db":/docker-entrypoint-initdb.d \
  mysql:latest > /dev/null 2>&1

printf "\033[92m✔︎ The database %s has been started\033[0m\n" "${CONTAINER_NAME}"