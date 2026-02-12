FROM node:lts-alpine AS builder

# Ensure all needed dependencies are installed with this setup
RUN apk add --no-cache curl git npm
WORKDIR /build

# Copy files and install node modules
COPY package.json /build
RUN npm i
COPY . /build

# Build the frontend
RUN npm run build

# Download the latest linux amd64 binary
RUN curl -L -o /usr/bin/web-proxy https://vc.maxkaya.com/maxpeterkaya/web-proxy/releases/download/latest/web-proxy_linux_amd64 && chmod +x /usr/bin/web-proxy

FROM scratch

# Copy all built static files and binary
COPY --from=builder /build/dist /app
COPY --from=builder /usr/bin/web-proxy /usr/bin/web-proxy

# Ensure docker knows that port 3000 is exposed by default
EXPOSE 3000

# Run binary in static mode and point to directory with static files
CMD ["/usr/bin/web-proxy", "-static", "-static-dir", "/app"]