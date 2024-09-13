include .env

build:
	go build -o bin/$(APP_NAME) cmd/app/main.go

air: 
	-air -c .air.toml

migrate-up:
	-go run cmd/app/main.go migrate -action up

migrate-down:
	-go run cmd/app/main.go migrate -action down

run:
	-go run cmd/app/main.go

seed:
	-go run cmd/app/main.go seed