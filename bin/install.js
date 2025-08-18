#!/usr/bin/env node
const os = require('os');
const path = require('path');
const fs = require('fs');
const {Readable} = require('stream');
const {finished} = require('stream/promises');

const platform = os.platform();
const arch = os.arch();

const platformMap = {
    'linux-x64': 'linux_amd64',
    'linux-arm64': 'linux_arm64',
    'darwin-x64': 'darwin_amd64',
    'darwin-arm64': 'darwin_arm64',
    'win32-x64': 'windows_amd64.exe'
};

const key = `${platform}-${arch}`;
const assetName = platformMap[key];

if (!assetName) {
    console.error(`Unsupported platform: ${key}`);
    process.exit(1);
}

const url = `https://github.com/maxpeterkaya/web-proxy/releases/latest/download/web-proxy_${assetName}`;
const destDir = path.join(__dirname);
const destFile = path.join(destDir, `${platform === "win32" ? "web-proxy.exe" : "web-proxy"}`);

(async () => {
    console.info(`Downloading ${assetName} from ${url}...`);
    fs.mkdirSync(destDir, {recursive: true});

    const res = await fetch(url);
    if (!res.ok) throw new Error(`Failed to fetch ${url}: ${res.statusText}`);

    const fileStream = fs.createWriteStream(destFile, {mode: 0o755});
    await finished(Readable.fromWeb(res.body).pipe(fileStream));

    console.info('web-proxy installed.');
})();