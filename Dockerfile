FROM debian:trixie AS certs

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

FROM scratch

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ARG TARGETPLATFORM

ENTRYPOINT ["/usr/bin/web-proxy"]

COPY $TARGETPLATFORM/web-proxy /usr/bin/