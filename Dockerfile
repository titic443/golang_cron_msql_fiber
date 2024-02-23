FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o main ./cmd/etax/main.go
# RUN go get

FROM alpine:3.18
RUN apk update
RUN apk --no-cache add  telnet tzdata
WORKDIR /app
COPY --from=builder /app/config.yaml .
COPY --from=builder /app/main .
EXPOSE 8888

CMD [ "/app/main" ]