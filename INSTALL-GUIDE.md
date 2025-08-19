# How to install web-proxy

## NPM Package

You can install web-proxy as a wrapper in the form of a npm package! This allows you to dynamically download the latest
version of web-proxy at all times in your build stage without having to tweak your process to manually fetch the latest
version through commands.

### Step 1

Install: ``npm i @maxpeterkaya/web-proxy``

While running ``npm i``, it will automatically download the latest binary with execution permissions.


### Step 2

> You can run your web app either through web-proxy or concurrently. To run your app through web-proxy, pass the **-app
** argument.

Add the following scripts to your package.json to run your app through web-proxy:

```json
{
  "scripts": {
    "start:web": "YOUR COMMAND TO START YOUR APP",
    "start": "npx web-proxy -app"
  }
}
```

Add the following scripts to your package.json to run your app and web-proxy concurrently:

Install: ``npm i concurrently``

```json
{
  "scripts": {
    "start:proxy": "npx web-proxy",
    "start:web": "YOUR COMMAND TO START YOUR APP",
    "start": "concurrently --prefix none \"npm run start:web\" \"npm run start:proxy\""
  }
}
```

> You can view the docs for concurrently [here](https://github.com/open-cli-tools/concurrently/tree/main/docs).

In the future web-proxy will be able to start your web app itself without the use of concurrently.

### Step 3

Simply just run ``npm run start`` and enjoy structured logs.

## Container Middleware

Follow the steps [here](https://github.com/maxpeterkaya/web-proxy/tree/main/examples/container-middleware).
