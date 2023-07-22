#!/usr/bin/env sh

printf "\033[93m→ Destroying the database\033[0m\n"

. configs/.env
CONTAINER_NAME="${DB_NAME}-db"

docker stop "${CONTAINER_NAME}" > /dev/null 2>&1 && docker rm "${CONTAINER_NAME}" > /dev/null 2>&1

printf "\033[92m✔︎ The database %s has been destroyed\033[0m\n" "${CONTAINER_NAME}"