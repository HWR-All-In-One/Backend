FROM golang:1.19.5

ENV PORT 8090 

#RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY /cmd/main.go .
COPY /internal ./internal

RUN go install
RUN go build -o /godocker .

EXPOSE $PORT

# Nicht als Array, weil so einfacher der Port als ENV genutzt werden kann
CMD /godocker serve --http=0.0.0.0:$PORT
