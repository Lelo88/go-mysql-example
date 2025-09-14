# Makefile for go-mysql-example

# Run tests
test:
	go test ./...

# Build the project
build:
	go build -o go-mysql-example

# Run the project
run:
	go run main.go

# Clean build artifacts
clean:
	rm -f go-mysql-example