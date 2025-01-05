include .env

run:
	@make build
	./currency-conversion-api

build:
	@go build -v -o currency-conversion-api ./cmd/main.go