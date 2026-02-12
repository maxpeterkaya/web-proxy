# Authentication

## What is it?

With authentication inside web-proxy you're able to lock down your downstream proxy to pass authorization before being
able to view and interact with your proxy.

## Needed environment variables
You will need to provide these variables into the container in order to properly use the authentication feature.
```dotenv
AUTH_TYPE=basic # none, basic
AUTH_USER=admin # default username
AUTH_PASS=      # insert SHA256 hashed password here
```

By default, if no password is provided, i.e. empty, then a cryptographical password will be created to ensure zero access.

## Files

### Dockerfile

This has all the necessary steps to add web-proxy as a binary along with in-depth explanations.

> For anything else you may need, refer to the
> root [README](https://vc.maxkaya.com/maxpeterkaya/web-proxy/src/branch/main/README.md).