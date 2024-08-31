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
RUN go build -o main .

# Stage 2: Create a lightweight image to run the Go application
FROM alpine:latest

# Install ca-certificates in the container
RUN apk --update add ca-certificates && rm -rf /var/cache/apk/*

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=build /app/main .

# Expose the port that the Go app listens on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
