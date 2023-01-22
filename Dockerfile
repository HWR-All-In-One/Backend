# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY ./cmd ./cmd
COPY ./internal ./internal

COPY go.mod ./
COPY go.sum ./
RUN go mod download

#COPY ./cmd ./

RUN cd ./cmd && go build -o ../docker-gs-ping

EXPOSE 8090

CMD [ "/docker-gs-ping", "serve" ]