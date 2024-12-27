FROM golang:1.22.3 AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./*.go ./
COPY ./common ./common
COPY ./build ./build
COPY ./internal ./internal
COPY ./handler ./handler

RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY build/app.env build/app.env

EXPOSE 8502

CMD ["./main"]