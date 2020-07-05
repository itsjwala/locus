FROM itsjwala/locus_runner:latest as locus_runner

FROM node:14.5.0-alpine3.10

ENV CODE_DIR=/playground

WORKDIR /playground

COPY --from=locus_runner /tmp/locus_runner /playground

ENTRYPOINT ["/playground/locus_runner"]
