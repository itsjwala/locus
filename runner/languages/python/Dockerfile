FROM itsjwala/locus_runner:latest as locus_runner

FROM python:3-alpine

ENV CODE_DIR=/playground

WORKDIR /playground

COPY --from=locus_runner /tmp/locus_runner /playground

ENTRYPOINT ["/playground/locus_runner"]
