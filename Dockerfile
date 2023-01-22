FROM alpine:latest

# Der Port, welcher angesprochen können soll
ENV PORT=8090

COPY /bin/linux/filename /

EXPOSE $PORT

# start PocketBase
CMD ["filename", "serve"]