# Makefile

APP_NAME=server
MAIN_PATH=./main.go
BINARY_DIR=./bin

.DEFAULT_GOAL := help

## dev: Hot reload with Air (daily use ke liye ye sabse zyada chalega)
dev:
	air

## run: Run app directly, no hot reload
run:
	go run $(MAIN_PATH)

## build: Build binary
build:
	go build -o $(BINARY_DIR)/$(APP_NAME) $(MAIN_PATH)

## tidy: go mod tidy
tidy:
	go mod tidy

## fmt: Format code
fmt:
	go fmt ./...

## vet: Static check for bugs
vet:
	go vet ./...

## test: Run all tests
test:
	go test ./... -v

## clean: Remove build + tmp files
clean:
	rm -rf $(BINARY_DIR) tmp

## docker-up: Start postgres container
docker-up:
	docker compose up -d

## docker-down: Stop containers (data safe rahega)
docker-down:
	docker compose down

## docker-logs: Watch postgres logs
docker-logs:
	docker compose logs -f postgres

## docker-reset: Stop + delete volumes (DB data wipe ho jayega)
docker-reset:
	docker compose down -v

## help: List all commands
help:
	@echo "Available commands:"
	@grep -E '^## ' Makefile | sed 's/## /  /'

.PHONY: dev run build tidy fmt vet test clean docker-up docker-down docker-logs docker-reset help