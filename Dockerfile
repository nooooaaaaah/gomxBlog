# Stage 1: Build the Go app
FROM golang:1.22-alpine AS build

# Install necessary build tools
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy the Go modules manifest
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o /bin/main ./cmd/main.go

# Stage 2: Create the final image
FROM alpine:latest

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy the built binary from the build stage
COPY --from=build /bin/main /bin/main

# Copy static files and templates
COPY --from=build /app/ui/static /ui/static
COPY --from=build /app/ui/html/ /ui/html

# Expose the application port
EXPOSE 80

# Command to run the executable
ENTRYPOINT ["/bin/main"]
