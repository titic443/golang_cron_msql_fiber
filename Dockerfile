FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download


RUN go mod download
RUN go build -o main ./cmd/etax/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/config /app/config

CMD [ "/app/main" ]