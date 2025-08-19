# Web-proxy

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/web-proxy)](https://artifacthub.io/packages/search?repo=web-proxy)

## Introduction

This project aims to provide a better logging solution to web applications that lack structured logs.
These applications all mostly lack a uniform logger that is customizable as well as even non-existent in production
builds.
Some of these frameworks may include:

- NextJS
- PrismaORM

## Getting Started

### NodeJS

Install: ``npm i @maxpeterkaya/web-proxy``

Add the following scripts to your ``package.json``:

```json
{
  "scripts": {
    "start": "npx web-proxy -app",
    // Passing the -app is very important to tell web-proxy to run the web app for you.
    "start:web": "YOUR COMMAND TO START THE WEB APP"
    // This could be next start, remix-serve build, ng serve, node dist/main, etc.
  }
}
```

After doing so, make sure web-proxy is properly configured

## [Installation Guide](https://github.com/maxpeterkaya/web-proxy/blob/main/INSTALL-GUIDE.md)

## Examples

- [Container middleware](https://github.com/maxpeterkaya/web-proxy/tree/main/examples/container-middleware)

## Roadmap

These bullet points are in no particular order of importance or difficulty.

- Install methods
- Install guide
- Usage guide
- Templates
- Screenshots
- Benchmarks
- Configurable logs