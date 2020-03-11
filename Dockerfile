FROM alpine:latest

COPY /dockerhub-statsd /

ENTRYPOINT ["/bin/sh", "-c", "/dockerhub-statsd $0 $@"]
