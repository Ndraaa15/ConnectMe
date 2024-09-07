# Stage 1: Build the Go application
FROM golang:1.22-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main cmd/app/main.go

# Stage 2: Create a lightweight image to run the Go application
FROM alpine:latest

# Install ca-certificates and curl in the container
RUN apk --update add ca-certificates curl && rm -rf /var/cache/apk/*

# Set the working directory inside the container
WORKDIR /app

# Create directory log for saving log
RUN mkdir log

# Expose port 8080
EXPOSE 8080

# Copy the compiled binary from the build stage
COPY --from=build /app/main /app/.env ./

# Command to run the application
CMD ["./main"]