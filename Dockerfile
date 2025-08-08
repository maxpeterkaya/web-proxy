FROM scratch

COPY web-proxy /usr/bin/web-proxy

ENTRYPOINT ["/usr/bin/web-proxy"]