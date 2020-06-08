FROM itsjwala/locus_runner:latest as locus_runner

FROM python:3-alpine

COPY --from=locus_runner /tmp/runner /tmp

ENTRYPOINT ["/tmp/runner"]
