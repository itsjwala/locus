FROM golang:1.14-alpine

ENV GOOS=linux

COPY locus_runner.go /tmp

RUN cd /tmp \ 
    && go build -o runner locus_runner.go

ENTRYPOINT  ["/tmp/runner"]
