# Use the official Golang image with a lightweight Alpine version
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Set environment variables to load the .env file inside the container
ENV ENV_FILE_PATH=/app/.env

# Copy go.mod and go.sum to the container
# This is to cache dependencies and avoid re-downloading them every build
COPY ./go.mod ./go.sum ./

# Download all Go dependencies
RUN go mod download

# Copy the Go source code into the container
COPY ./ ./

# Copy the .env file into the container (make sure it exists in the root of the project)
COPY ./.env ./.env

# Build the Go application
RUN go build -o main .

# Expose the port on which the service will run
EXPOSE 8080

# Command to run the Go executable
CMD ["./main"]
