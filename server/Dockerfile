# Base image
FROM golang:1.24-alpine AS base

# Making image for development
FROM base AS development

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cd cmd/web && go build -o bin/golang-react-todo-app

EXPOSE 8000

CMD ["/build/cmd/web/bin/golang-react-todo-app"]

# Builder for production
FROM base AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cd cmd/web && CGO_ENABLED=0 go build -o bin/golang-react-todo-app

# Making small image for production
FROM scratch AS production

WORKDIR /prod

COPY --from=builder /build/cmd/web/bin/golang-react-todo-app ./

COPY pkg/db/data/ ./pkg/db/data/

EXPOSE 8000

CMD ["/prod/golang-react-todo-app"]