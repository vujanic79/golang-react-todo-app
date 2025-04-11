# Base image
FROM golang:1.24-alpine AS base

# Making image for development
FROM base AS development

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

RUN cd cmd/web && go build -gcflags="all=-N -l" -o bin/golang-react-todo-app

EXPOSE 8000 40000

CMD ["/go/bin/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/build/cmd/web/bin/golang-react-todo-app"]
#CMD ["/build/cmd/web/bin/golang-react-todo-app"]

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