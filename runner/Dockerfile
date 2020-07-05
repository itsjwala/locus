FROM golang:1.14-alpine

ENV GOOS=linux

WORKDIR /src

COPY . .

RUN cd /src  \
    && go build -o locus_runner \
    && mv locus_runner /tmp

ENTRYPOINT  ["/tmp/locus_runner"]
