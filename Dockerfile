# Use the official Go image as a build environment.
FROM golang:1.22-alpine AS builder

# Set the working directory.
WORKDIR /app

# Copy go.mod and go.sum files to the workspace.
COPY go.mod go.sum ./

# Download dependencies.
RUN go mod download

# Copy the entire source code to the container.
COPY . .

# Build the Go app.
RUN go build -o main .

# Use a minimal base image to reduce the container size.
FROM alpine:latest

# Set the working directory in the container.
WORKDIR /root/

# Copy the binary from the build stage.
COPY --from=builder /app/main .

# Copy the data directory to the container.
COPY --from=builder /app/data ./data

# Expose the port the app runs on.
EXPOSE 9700

# Run the binary.
CMD ["./main"]
