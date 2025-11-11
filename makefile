.PHONY: run migrate

run:
	go run cmd/app/main.go

migrate:
	go run cmd/app/main.go migrate

setup:
	go mod tidy