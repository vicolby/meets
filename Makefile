build:
	@go build -o bin/meets cmd/main.go

run: build
	@bin/meets

test:
	@go test -v ./...
