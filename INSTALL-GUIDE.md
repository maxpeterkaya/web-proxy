# How to install web-proxy

## NPM Package

You can install web-proxy as a wrapper in the form of a npm package! This allows you to dynamically download the latest
version of web-proxy at all times in your build stage without having to tweak your process to manually fetch the latest
version through commands.

### Step 1

Install: ``npm i @maxpeterkaya/web-proxy concurrently``

While running ``npm i``, it will automatically download the latest binary with execution permissions.

Currently you need to run your app and web-proxy through concurrently to ensure both are running concurrently.

### Step 2

Add the following scripts to your package.json:

```json
{
  "scripts": {
    "start:proxy": "./node_modules/.bin/web-proxy",
    "start:web": "YOUR COMMAND TO START YOUR APP",
    "start": "concurrently --prefix none \"npm run start:web\" \"npm run start:proxy\""
  }
}
```

You can view the docs for concurrently [here](https://github.com/open-cli-tools/concurrently/tree/main/docs).

#

In the future web-proxy will be able to start your web app itself without the use of concurrently.

### Step 3

Simply just run ``npm run start`` and enjoy structured logs.

## Container Middleware

Follow the steps [here](https://github.com/maxpeterkaya/web-proxy/tree/main/examples/container-middleware).
