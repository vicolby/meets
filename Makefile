run: build
	@./bin/events

build:
	@go build -tags dev -o bin/events main.go

pup:
	@docker-compose up -d

pdown:
	@docker-compose down

reset:
	@go run cmd/reset/main.go

down: ## Database migration down
	@go run cmd/migrate/main.go down

up: ## Database migration up
	@go run cmd/migrate/main.go up

migration: ## Migrations against the database
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

