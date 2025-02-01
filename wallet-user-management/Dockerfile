# Use an official Golang image as the base
FROM golang:1.23.5-alpine

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o wallet-app .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./wallet-app"]
