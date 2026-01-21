FROM golang:1.24 AS builder

# define the build arguments
ARG GIT_VERSION
ARG GIT_COMMIT
ARG GIT_DATE

ENV GIT_VERSION=${GIT_VERSION}
ENV GIT_COMMIT=${GIT_COMMIT}
ENV GIT_DATE=${GIT_DATE}

WORKDIR /app

# build everything
COPY . .
RUN CGO_ENABLED=0 go build -o web-proxy -ldflags="-X 'main.version=${GIT_VERSION}' -X 'main.commit=${GIT_COMMIT}' -X 'main.date=${GIT_DATE}'" .

FROM scratch

# Copy binary
COPY --from=builder /app/web-proxy /usr/bin/

EXPOSE 3000

ENTRYPOINT ["/usr/bin/web-proxy"]