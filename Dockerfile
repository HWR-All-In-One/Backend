FROM alpine:latest

ARG PB_VERSION=0.11.2

# Der Port, welcher angesprochen können soll
ENV PORT=8008

RUN apk add --no-cache \
    unzip \
    ca-certificates

# download and unzip PocketBase
ADD https://github.com/pocketbase/pocketbase/releases/download/v${PB_VERSION}/pocketbase_${PB_VERSION}_linux_amd64.zip /tmp/pb.zip
RUN unzip /tmp/pb.zip -d /pb/

EXPOSE $PORT

# start PocketBase
CMD ["/pb/pocketbase", "serve", "--http=0.0.0.0:8080"]