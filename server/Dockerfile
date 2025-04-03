FROM golang:1.24-alpine AS base

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cd cmd/web && go build -o bin/golang-react-todo-app

EXPOSE 8000

CMD ["/build/cmd/web/bin/golang-react-todo-app"]