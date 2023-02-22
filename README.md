# Backend von der HWR-All-In-One-App

## Projektstruktur
Für die Projektstruktur siehe: https://github.com/golang-standards/project-layout

## Branchstruktur
### main
**Der main-Branch ist stable.** Somit wird nur der dev-Branch in den main-Branch gemerged, wenn der dev-Branch fehlerfrei funktioniert. Es wird ausschließlich aus dem dev-Branch in den main-Branch gemerged.

### dev
In dem dev-Branch werden die Änderungen gesammelt und, wenn diese stabil laufen, in den main-Branch gemerged. **Somit ist der dev-Branch nicht stable.** Es wird aus dem dev-Branch heraus in main gemerged, somit kommen Änderungen in main nur aus dev.

### Weitere Branchebezeichnungen
* *feature*: Hier wird ein feature implementiert (Nummer von JIRA)
* *subfeature*: Hier wird ein subfeature implementiert (Nummer von JIRA)
* *bug*: Hier wird ein Bug behoben (Nummer von JIRA)

## Docker
### Docker lokal selbst bauen und verwenden
**Docker build:** `docker build -t hwraio/backend:local .`  
**Docker run:** `docker run -p 8090:8090 -it hwraio/backend:local`  
Nun kann beispielweise das Admin-Interface über die URL [http://0.0.0.0:8090/_/](http://0.0.0.0:8090/_/) erreicht werden.  

### Auf dem Server zentral gebauten Docker verwenden
Wenn auf den dev oder main Branch des Backend Repositories gepusht wird, wird ein neues Docker Image erstellt und mit zwei verschiedenen Tags versehen (s. u.).

#### GHCR login
`docker login https://ghcr.io -u USERNAME -p TOKEN`  

#### Docker Image nutzen
**Docker run (Schema):** `docker run --name NAME -p PORT-NACH-EXTERN:8090 -it ghcr.io/hwr-all-in-one/backend:TAG`  
**Docker run (Beispiel):** `docker run --name hwraio-backend -p 8090:8090 -it ghcr.io/hwr-all-in-one/backend:latest`  

#### Verfügbare Tags
* Tags für Images von main (diese sind stable)
  * latest (neuste Version von dem `main`)
  * Versionstag (z. B. `0.1.0-build20`)
* Tags für Images von dev (diese sind noch in Entwicklung, sodass diese fehlerhaft sein können / nicht stabil sind)
  * dev-latest (neuste Version von dem `dev` Branch)
  * Versionstag (z. B. `0.1.0-build7-dev1`)