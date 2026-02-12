# ALL YOUR BUILD STEPS HERE
# In your build step add the following (works with node:lts-alpine)
RUN apk add --no-cache curl
# This package allows docker to retrive the web-proxy binary from forgejo

# This tells docker that all your files for your website / web-proxy are in this directory and this should be the default execution directory
WORKDIR /app

# Download the latest linux amd64 version of web-proxy and give the binary execution permissions
RUN curl -L -o /usr/bin/web-proxy https://vc.maxkaya.com/maxpeterkaya/web-proxy/releases/download/latest/web-proxy_linux_amd64 && chmod +x /usr/bin/web-proxy

# Port 3000 is the default port that web-proxy is configured to run on, this can be left alone since you can map a different port on your host to the container
EXPOSE 3000

# Run binary, add any neccessary flags per your need.
# In order to add the authentication feature, add the neccessary environment variables mentioned in the readme.
CMD ["/usr/bin/web-proxy"]