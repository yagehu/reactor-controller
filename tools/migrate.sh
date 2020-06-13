#!/usr/bin/env bash

export POSTGRESQL_URL='postgres://reactor:reactor@localhost:5432/reactor?sslmode=disable'

migrate -database "$POSTGRESQL_URL" -path db/migrations up
