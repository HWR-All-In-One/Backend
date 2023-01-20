FROM alpine:latest

# Der Port, welcher angesprochen können soll
ENV PORT=8090

COPY /bin/main.exe /

EXPOSE $PORT

# start PocketBase
CMD ["main.exe", "serve", "--http=0.0.0.0:8090"]