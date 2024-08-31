include .env

build:
	go build -o bin/$(APP_NAME) cmd/app/main.go

run: 
	air -c .air.toml