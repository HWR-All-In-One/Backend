# Backend von der HWR-App

Für die Projektstruktur siehe: https://github.com/golang-standards/project-layout


docker build -t hwraio/backend:local .
docker run -p 80:80 -it hwraio/backend:local
docker run -it hwraio/backend:local
docker run hwraio/backend:local