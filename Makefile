all: go-build go-test build run

modules:
	go mod tidy

go-build: modules
	go build -o bin/orderservice ./cmd/orderservice/.

build:
	docker-compose build

run:
	docker-compose up

go-test:
	go test ./...
