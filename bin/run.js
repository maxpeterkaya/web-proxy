#!/usr/bin/env node
const {spawnSync} = require('child_process');
const path = require('path');
const binaryName = process.platform === 'win32' ? 'web-proxy.exe' : 'web-proxy';
const binaryPath = path.join(__dirname, binaryName);

spawnSync(binaryPath, {stdio: 'inherit'});