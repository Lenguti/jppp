SHELL := bash

PROJECT_NAME  ?= jppp
GO_BINARY_DIR := bin
CGO_ENABLED   := 0
GOOS          ?= linux
GOARCH        ?= arm64
MIGRATIONS_DIR = business/data/migrations

.PHONY : build
build  :
	@CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(GO_BINARY_DIR)/$(PROJECT_NAME) .

.PHONY : run
run    :
	@docker compose up -d --build

.PHONY : stop
stop   :
	@docker compose down

.PHONE : migrate
migrate:
	@cd $(MIGRATIONS_DIR) && goose postgres "host=localhost user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME}" up

.PHONY : test
test   :
	@CGO_ENABLED=$(CGO_ENABLED) go test ./... --cover
