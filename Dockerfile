# syntax=docker/dockerfile:1
FROM golang:1.22.5
ENV CGO_ENABLED=1

WORKDIR /app

COPY go.mod go.sum ./
RUN apt-get update && apt-get install -y netcat-openbsd
RUN go mod download

COPY . .

COPY config/wait-mysql.sh /wait-mysql.sh
RUN chmod +x /wait-mysql.sh

RUN CGO_ENABLED=1 GOOS=linux go build -o /app/contacts-api

EXPOSE 8080

CMD ["/wait-mysql.sh", "/app/contacts-api"]
