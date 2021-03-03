all: build test

modules:
	go mod tidy

build: modules
	go build cmd/orderservice/main.go

test:
	go test ./...
