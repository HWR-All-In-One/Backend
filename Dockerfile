FROM golang:1.19.5

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY /cmd/main.go .
COPY /internal ./internal

RUN go mod download
RUN go build -o /godocker .

EXPOSE 8090

CMD ["/godocker", "serve"]
