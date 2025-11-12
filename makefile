.PHONY: run setup help migrate

# Load .env if it exists (silent if missing)
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
	@echo "  make migrate new NAME=xxx  - Create a new SQL migration (e.g., make migrate new NAME=create_users_table)"
	@echo "  make migrate up          - Apply all pending migrations"
	@echo "  make migrate down        - Rollback the last migration"
	@echo "  make migrate reset       - Rollback all migrations, then re-apply them"
	@echo ""
	@echo "💡 Tip: Ensure 'migrate' CLI is installed (https://github.com/golang-migrate/migrate)"

# --- Migration Subcommands ---
migrate:  ## Delegate to subcommands (for help clarity)

migrate-new:
ifndef NAME
	$(error Please provide NAME, e.g., make migrate new NAME=create_users_table)
endif
	migrate create -ext sql -dir migrations -seq $(NAME)

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" -verbose down

migrate-reset:
	@echo "🔄 Resetting database migrations..."
	make migrate-down
	make migrate-up

# --- Alias-style routing (migrate new → migrate-new) ---
migrate new: migrate-new
migrate up: migrate-up
migrate down: migrate-down
migrate reset: migrate-reset

# Prevent direct calls to internal targets
migrate-new migrate-up migrate-down migrate-reset:
	@true