include .env

.PHONY: build
build:
	go build -o bin/$(APP_NAME) cmd/app/main.go

.PHONY: air
air: 
	air -c .air.toml

.PHONY: migrate-up
migrate-up: build
	./bin/$(APP_NAME) migrate -action up

.PHONY: migrate-down
migrate-down: build
	./bin/$(APP_NAME) migrate -action down

.PHONY: run
run: build
	./bin/$(APP_NAME)

.PHONY: test
seed: build
	./bin/$(APP_NAME) seed