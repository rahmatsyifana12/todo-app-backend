# Use the official Golang image to build the Go app
FROM golang:1.20-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache for dependencies
COPY go.mod go.sum ./

# Download all Go modules
RUN go mod download

# Copy the source code, including the module.go file and other .go files
COPY ./src ./src

# Build the Go app, specifying the main package where both main.go and module.go are located
RUN go build -o main ./src

# Start a new stage from a smaller image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the pre-built Go binary from the previous stage
COPY --from=build /app/main .

# Copy the .env file
COPY .env .env

# Expose port from the .env file
ARG PORT

# Expose the dynamic port (it will be assigned in runtime from .env file)
EXPOSE ${PORT}

# Set environment variables from the .env file
ENV $(cat .env | xargs)

# Command to run the executable
CMD ["./main"]
