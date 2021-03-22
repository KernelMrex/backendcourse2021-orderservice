all: build test run

modules:
	go mod tidy

build: modules
	go build -o bin/orderservice ./cmd/orderservice/.

run:
	docker build -t orderservice .
	docker run -d -p8000:8000 orderservice

test:
	go test ./...
