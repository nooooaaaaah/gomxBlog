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
COPY edgedb.toml /app/edgedb.toml


# Build the Go app
RUN go build -o /bin/blog ./cmd/main.go

# Stage 2: Create the final image
FROM alpine:latest

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy the built binary from the build stage
COPY --from=build /bin/blog /bin/main

COPY --from=build /app/robots.txt ./
COPY --from=build /app/sitemap.xml ./

# Copy static files and templates
COPY --from=build /app/ui/static /ui/static
COPY --from=build /app/ui/html/ /ui/html
RUN rm /ui/static/js/refreshMeDaddy.js

# Expose the application port
EXPOSE 4200

# Command to run the executable
ENTRYPOINT ["/bin/main"]
