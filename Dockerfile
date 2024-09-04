# syntax=docker/dockerfile:1
FROM golang:1.23 AS builder

# Set destination for COPY
WORKDIR /opt/app

COPY ./backend/ .

# Install swag
RUN export PATH=$(go env GOPATH)/bin:$PATH
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy the source code.
RUN swag init -g ./cmd/main.go

# Download Go modules
RUN go mod download

# Build the Go application
RUN go build -o app ./cmd/...

ENTRYPOINT [ "/opt/app/app" ]
