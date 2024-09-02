# syntax=docker/dockerfile:1
FROM golang:1.23

# Set destination for COPY
WORKDIR /opt/app

COPY . .
RUN cd backend
RUN export PATH=$(go env GOPATH)/bin:$PATH
RUN go install github.com/swaggo/swag/cmd/swag@latest


# Copy the source code.

RUN swag init -g ./cmd/main.go

# Download Go modules
RUN go mod download

RUN go build -o app ./cmd/...


ENTRYPOINT [ "/opt/app/app" ]