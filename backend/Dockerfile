# Stage 1: Build the Go binary
FROM golang:1.22.1-alpine AS build

# Set environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Copy the .env file to the container
COPY .env .env

# Build the Go application using the Makefile
RUN apk add --no-cache make
RUN make build

# Stage 2: Create a smaller image with the built binary
FROM alpine:latest

# Set working directory
WORKDIR /app

COPY . .

# Copy the binary from the build stage
COPY --from=build /app/bin/matcherx ./bin/matcherx

# Copy the .env file into the final container
COPY .env .env

# Expose the port on which the backend will run (adjust if necessary)
EXPOSE 8080

# Load the environment variables from .env file and run the binary
CMD ["sh", "-c", "./bin/matcherx"]
