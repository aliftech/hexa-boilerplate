.PHONY: run setup help migrate

# Load .env if it exists
-include .env

# --- Main Targets ---
run:
	go run cmd/app/main.go

setup:
	go mod tidy

help:
	@echo "🚀 Available commands:"
	@echo ""
	@echo "  make run                 - Run the application"
	@echo "  make setup               - Install Go dependencies"
	@echo ""
	@echo "  make migrate-new NAME=xxx  - Create new migration (e.g., make migrate-new NAME=create_users_table)"
	@echo "  make migrate-up            - Apply all pending migrations"
	@echo "  make migrate-down          - Rollback last migration"
	@echo "  make migrate-reset         - Reset: down then up"
	@echo ""
	@echo "💡 Tip: Install 'migrate' CLI: go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"

# --- Migrations ---
migrate-new:
ifndef NAME
	$(error Please specify NAME, e.g., make migrate-new NAME=create_users_table)
endif
	migrate create -ext sql -dir migrations -seq $(NAME)

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" -verbose down

migrate-reset:
	@echo "🔄 Resetting migrations..."
	make migrate-down
	make migrate-up