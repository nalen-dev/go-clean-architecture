FROM golang:1.21.9-bullseye AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api-golang ./cmd/main.go

EXPOSE 8080

CMD ["./api-golang"]