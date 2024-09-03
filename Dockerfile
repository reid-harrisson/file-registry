# syntax=docker/dockerfile:1
FROM golang:1.23 AS builder

# Set destination for COPY
WORKDIR /opt/app

COPY . .

WORKDIR /opt/app/backend

COPY .env .

# Install swag
RUN export PATH=$(go env GOPATH)/bin:$PATH && \
    go install github.com/swaggo/swag/cmd/swag@latest

# Copy the source code.
RUN swag init -g ./cmd/main.go

# Download Go modules
RUN go mod download

# Build the Go application
RUN go build -o app ./cmd/...

# Stage 2: Add Node.js and npm
FROM node:16 AS node_builder

# Set working directory for Node.js
WORKDIR /opt/app/smart_contract

# Copy the smart contract files
COPY --from=builder /opt/app/smart_contract .

# Install npm packages
RUN npm install

# Final stage: Build the final image
FROM golang:1.23

# Set destination for COPY
WORKDIR /opt/app

# Copy the built Go application from the builder stage
COPY --from=builder /opt/app/backend/app ./backend/

# Copy the smart contract artifacts
COPY --from=node_builder /opt/app/smart_contract .

WORKDIR /opt/app/backend

ENTRYPOINT [ "/opt/app/backend/app" ]
