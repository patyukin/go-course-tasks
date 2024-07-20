FROM golang:1.22.1-alpine3.19 AS builder

ENV config=docker

WORKDIR /app

COPY .. /app

RUN go mod tidy && \
    go mod download && \
    go get github.com/githubnemo/CompileDaemon && \
    go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o server main.go" --command=./server
