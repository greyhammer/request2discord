# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o /request2discord

EXPOSE 8080

CMD [ "/request2discord" ]