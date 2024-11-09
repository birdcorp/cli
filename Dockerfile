FROM alpine
RUN apk update && apk upgrade && \
  apk add --no-cache ca-certificates
COPY birdcli /bin/birdcli
ENTRYPOINT ["/bin/birdcli"]
