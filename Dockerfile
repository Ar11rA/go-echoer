# Start from the official Go base image
FROM golang:1.22-alpine as build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o server .

# Start a new stage from scratch
FROM alpine:latest  

# Copy the pre-built binary file from the previous stage
COPY --from=build /app/server /app/server

# Expose the port on which the Echo server will run
EXPOSE 7001

# Run the binary
ENTRYPOINT ["/app/server"]