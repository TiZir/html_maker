#!/bin/sh
#entrypoint.sh
wait-for "${PG_HOST}:${PG_PORT}" -- "$@"
go build -o main ./cmd/main.go
./main