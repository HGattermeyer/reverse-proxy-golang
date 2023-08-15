# Use an official Go image as the base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and install Go module dependencies
RUN go mod download

# Copy the rest of the project files to the container
COPY . .

EXPOSE 8080

# Build the Go application
RUN go build -o proxy-golang

# Define the entry point command
CMD ["./proxy-golang"]
