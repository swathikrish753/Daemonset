# Use the official Go image as a base
FROM golang:1.20-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o prometheus-metrics .

# Use a smaller image for the final build
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the executable built in the previous stage
COPY --from=build /app/prometheus-metrics .

# Expose the port on which the Go application listens
EXPOSE 8080

# Command to run the Go application
CMD ["./prometheus-metrics"]
