FROM itsjwala/locus_runner:latest as locus_runner

FROM python:3-alpine

COPY --from=locus_runner /tmp/locus_runner /tmp

ENTRYPOINT ["/tmp/locus_runner"]
