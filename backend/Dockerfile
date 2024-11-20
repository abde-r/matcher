FROM golang:1.22.1-alpine AS build

# Set environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app
COPY . .
# Download dependencies
RUN go mod download
RUN go mod tidy
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go get -u github.com/swaggo/gin-swagger
RUN go get -u github.com/swaggo/files
RUN swag init -g cmd/main.go
RUN apk add --no-cache make
RUN make build

FROM alpine:latest
# Copy only the binary from the build stage
COPY --from=build /app/bin/matcherx ./matcherx
COPY . .
CMD ["./matcherx"]


# # Set environment variables
# ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# # Set the working directory inside the container

# # Copy go.mod and go.sum files
# COPY go.mod go.sum ./


# # Copy the source code and Makefile

# # Install make (if needed) and build the Go application using Makefile

# # Stage 2: Create a smaller image with the built binary

# # Set working directory
# WORKDIR /app


# COPY wait-for-it.sh /usr/local/bin/wait-for-it
# RUN chmod +x /usr/local/bin/wait-for-it

# # Ensure binary is executable
# RUN chmod +x ./matcherx

# # Copy the .env file into the final container if it is needed at runtime

# # Expose the port on which the backend will run (adjust if necessary)
# EXPOSE 8000
# # Run the binary directly

