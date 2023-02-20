# Variables
BINARY_NAME=main
DOCKER_REPO=vorkurk/go-voting-app
GO_VERSION=$(shell go version)
UNAME=$(shell uname)
ifeq ($(UNAME), Darwin)
    AVAILABLE_CPU=$(shell sysctl -n hw.logicalcpu)
    AVAILABLE_RAM=$(shell sysctl -n hw.memsize | awk '{print int($$1/1024^3)"GB"}')
else
    AVAILABLE_CPU=$(shell nproc)
    AVAILABLE_RAM=$(shell free -h | awk '/^Mem:/ {print $$2}')
endif

# Commands
deps:
	go mod download

deps-update:
	go get -u ./...
	go mod tidy

build:
	go build -o $(BINARY_NAME) main -v ./cmd/*.go

run:
	go run ./cmd/*.go

docker-build:
	docker build -t $(DOCKER_REPO) .

info:
	@echo "Go version: $(GO_VERSION)"
	@echo "Available CPU: $(AVAILABLE_CPU)"
	@echo "Available RAM: $(AVAILABLE_RAM)"

test:
	go test ./tests/...

help:
	@echo "Available commands:"
	@echo "deps: Fetches the project dependencies."
	@echo "build: Compiles the project and creates an executable binary file."
	@echo "run: Runs the program without building an executable binary file."
	@echo "docker-build: Creates a Docker image with the name mydockerhubusername/myapp."
	@echo "info: Displays information about the system, including the version of Go installed, the available number of CPU cores, and the available RAM."
	@echo "help: Displays this help message."

.PHONY: deps deps-update build run docker-build info help
