FROM golang:1.22.2-alpine

WORKDIR /app

COPY . /app
RUN go mod tidy && go build cmd/main.go



