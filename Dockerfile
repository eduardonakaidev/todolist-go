FROM golang:1.23 as builder

RUN apt-get update && apt-get install -y gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

ENV CGO_ENABLED=1 GOOS=linux
## external dependy run v
RUN go build -ldflags '-linkmode external -extldflags "-static -fpic"' -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
RUN chmod +x main
EXPOSE 8080

CMD ["./main"]
