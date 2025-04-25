build:
	@go build -o bin/jotjournal cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/jotjournal
	
migration:
	@migrate create -ext sql -dir db/migrate/migrations -seq $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run db/migrate/main.go up

migrate-down:
	@go run db/migrate/main.go down

generate-task:
	@go run db/seed/main.go

swag:
	@swag init -g cmd/main.go
