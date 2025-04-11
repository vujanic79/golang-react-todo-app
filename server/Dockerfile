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

RUN go build -o /app/bin/golang-react-todo-app

WORKDIR /app

EXPOSE 8000

CMD ["/app/bin/golang-react-todo-app"]

# ------------------------------
# Builder image for production
# ------------------------------
FROM base AS builder

WORKDIR /app/cmd/web

RUN CGO_ENABLED=0 go build -o /app/bin/golang-react-todo-app

# ------------------------------
# Final production image
# ------------------------------
FROM scratch AS production

WORKDIR /prod

COPY --from=builder /app/bin/golang-react-todo-app ./
COPY pkg/db/data/ ./pkg/db/data/

EXPOSE 8000

CMD ["/prod/golang-react-todo-app"]