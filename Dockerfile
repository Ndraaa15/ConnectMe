FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/app/main.go

FROM alpine:latest

RUN apk --update add ca-certificates curl && rm -rf /var/cache/apk/*

WORKDIR /app

RUN mkdir log template

EXPOSE 8080

COPY --from=build /app/main /app/.env ./

COPY --from=build /app/internal/adapter/pkg/template /app/template

CMD ["./main"]