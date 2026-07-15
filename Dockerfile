FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o contact-list main.go

FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache netcat-openbsd
COPY init.sql /docker-entrypoint-initdb.d
COPY --from=builder /app/contact-list .
COPY --chmod=755 entrypoint.sh .

RUN apk add --no-cache postgresql postgresql-contrib postgresql-client
CMD ["./entrypoint.sh"]