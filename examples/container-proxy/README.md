# Container Proxy

## What is it?

Run the web-proxy container in the same stack as your desired service that should be proxied for logging and/or authentication purposes.

## Container Routing

### Portainer

When you run your stack via portainer, the name of your stack is automatically prepended to every service name.
Say you have the following and your example service runs on port 3000:
```yaml
services:
  proxy:
    image: vc.maxkaya.com/maxpeterkaya/web-proxy:latest
  example:
    image: example:latest
```
Your ``example`` service would be accessible by the full URL of ``STACKNAME_example:3000``. In order to properly route your downstream proxy, make sure you use ``STACKNAME_example`` for the ``PROXY_HOST`` variable and ``3000`` for the ``PROXY_PORT`` variable.

### docker-compose

Say you have the same or a similar compose config as the above, the URL to access your service from inside the web-proxy container would be ``example:3000``. 

## Files

### docker-compose.yml

Rough structure for running web-proxy and your other service(s).

> For anything else you may need, refer to the
> root [README](https://vc.maxkaya.com/maxpeterkaya/web-proxy/src/branch/main/README.md).