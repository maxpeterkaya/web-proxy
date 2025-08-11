#!/bin/bash

# The workdir from Dockerfile
cd /app

# The & is important after the run command to ensure the command runs in the background
echo "Starting web server..."
npm start &

echo "Starting web-proxy..."
/usr/bin/web-proxy &

wait