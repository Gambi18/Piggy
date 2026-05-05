.PHONY: run-app build test clean migrate-up migrate-down sqlc docker-up docker-down

# Application commands
run-app:
	go run ./cmd/main.go

build:
	go build -o bin/piggy ./cmd/main.go

test:
	go test ./...

clean:
	rm -rf bin/

# Database commands
migrate-up:
	migrate -path ./internal/db/migrations -database "postgres://admin:2323@localhost:5433/piggy?sslmode=disable" up

migrate-down:
	migrate -path ./internal/db/migrations -database "postgres://admin:2323@localhost:5433/piggy?sslmode=disable" down

migrate-force:
	migrate -path ./internal/db/migrations -database "postgres://admin:2323@localhost:5433/piggy?sslmode=disable" force $(version)

sqlc:
	sqlc generate

# Docker commands
docker-up:
	docker-compose up -d db

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

# Development setup
dev-setup: docker-up
	sleep 5
	make migrate-up
	echo "Development environment is ready!"

# Production commands
prod-build:
	CGO_ENABLED=0 GOOS=linux go build -o bin/piggy-linux ./cmd/main.go

run-prod:
	./bin/piggy-linux