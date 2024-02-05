FROM golang:1.21.5 AS builder

WORKDIR /app
COPY . .
RUN go build -o main ./cmd/etax/main.go
# RUN go get

# FROM alpine:3.18
# WORKDIR /app
# COPY --from=builder /app/main .
# COPY --from=builder /app/config /app/config

# CMD [ "/app/main" ]