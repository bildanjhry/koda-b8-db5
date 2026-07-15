FROM golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM postgres:alpine

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=admin

COPY init.sql /docker-entrypoint-initdb.d
COPY --from=build /app/main /app

WORKDIR /var/app
COPY --chmod=755 entrypoint.sh .

CMD ["/var/app/entrypoint.sh"]