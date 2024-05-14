# Use official Golang image as the base image
FROM golang:1.21 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code from the current directory to the
# working directory inside the container
COPY ../.. .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Expose the port the service listens on
EXPOSE 80

# Command to run the executable

CMD [ "/app/app"]
