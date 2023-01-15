# Ubuntu, damit zusätzliche Software (hier: gcc) installiert werden kann
#FROM ubuntu:22.04

# Das GO-Image
FROM golang:1.19

# Der Port, welcher angesprochen können soll
ENV PORT=8008

# Ein Verzeichnis für die App erstellen
RUN mkdir /app
 
# Alle Dateien vom derzeitigen Verzeichnis ins Docker-Verzeichnis kopieren
COPY /cmd /app
COPY go.mod /app

# working directory setzen
WORKDIR /app

# GCC installieren, damit dieses währen des build-Prozesses ausgeführt werden kann
#RUN apt-get update && \
#    apt-get install -y build-essential && \
#    apt-get clean

#RUN go install github.com/swaggo/swag/cmd/swag@latest

#RUN swag init -g restfulApiService.go

# Baut eine ausführbare Datei namens server im aktuellen Ordner
RUN go build -o server . 

# Dokumentation des Ports, welcher angesprochen wird
EXPOSE $PORT

# server.exe ausführen
CMD [ "/app/server" ]