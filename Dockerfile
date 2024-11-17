FROM alpine:latest
RUN apk update && apk upgrade && \
    apk add --no-cache ca-certificates
COPY dist/birdcli-linux /bin/birdcli
ENTRYPOINT ["/bin/birdcli"]
