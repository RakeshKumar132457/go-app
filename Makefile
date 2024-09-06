# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=go-app
BINARY_UNIX=$(BINARY_NAME)_unix
MAIN_PATH=./cmd/go-app
BIN_FOLDER=bin

all: test build

build:
	mkdir -p $(BIN_FOLDER)
	$(GOBUILD) -o $(BIN_FOLDER)/$(BINARY_NAME) -v $(MAIN_PATH)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BIN_FOLDER)

run: build
	./$(BIN_FOLDER)/$(BINARY_NAME)

deps:
	$(GOGET) -v -t -d ./...

# Migrations
migrate-up:
	migrate -path ./migrations -database "postgresql://username:password@localhost:5432/database_name?sslmode=disable" up

migrate-down:
	migrate -path ./migrations -database "postgresql://username:password@localhost:5432/database_name?sslmode=disable" down

# Cross compilation
build-linux:
	mkdir -p $(BIN_FOLDER)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BIN_FOLDER)/$(BINARY_UNIX) -v $(MAIN_PATH)

# Docker
docker-build:
	docker build -t $(BINARY_NAME):latest .

# Linting
lint:
	golangci-lint run

# Hot Reloading with nodemon
dev:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run $(MAIN_PATH)

.PHONY: all build test clean run deps migrate-up migrate-down build-linux docker-build lint dev
