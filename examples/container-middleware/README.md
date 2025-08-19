# Container Middleware

## What is it?

Essentially runs web-proxy alongside your web app in the same container and routes all traffic inside the container
through web-proxy to your web-app.

## Files

### entrypoint.sh

This runs web-proxy and your web-app at the same time, make sure to copy over this file in your root directory of your
app and adjust it to your needs.

### Dockerfile

This has all the necessary steps to add web-proxy as a binary along with in-depth explanations.

### docker-compose.yml

Just a quick template for port-mapping if you're not familiar with it.

> For anything else you may need, refer to the
> root [README](https://github.com/maxpeterkaya/web-proxy/blob/main/README.md).