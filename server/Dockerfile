# ------------------------------
# Base image
# ------------------------------
FROM golang:1.24-alpine AS base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ------------------------------
# Development image
# ------------------------------
FROM base AS development

WORKDIR /app/cmd/web

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go build -gcflags="all=-N -l" -o /app/bin/golang-react-todo-app

WORKDIR /app

EXPOSE 8000 40000