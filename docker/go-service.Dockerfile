# Use the official lightweight Go image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY backend/go-services/go.mod backend/go-services/go.sum ./

# Download all dependencies
RUN go mod download

# Copy the entire source code to the working directory
COPY backend/go-services/ .

# Build the Go application
RUN go build -o main .

# Expose the port the service will run on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
