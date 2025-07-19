FROM golang:1.24 AS builder

WORKDIR /app
COPY . .
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /walet-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /walet-service .
COPY --from=builder /app/.env ./
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./walet-service"]