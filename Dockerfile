

# Switch to an Alpine-based image
FROM golang:alpine

# Install build-base and sqlite
RUN apk --no-cache add build-base sqlite

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=1 GOOS=linux go build -o /authservice .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/authservice"]
