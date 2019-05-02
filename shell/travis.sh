#!/usr/bin/env bash
which eslint
eslint app/*.js app/scripts/*.js

echo "building Jarvis-Desktop package"

npm run build
echo "successfully build Jarvis-Desktop"

echo "cleaning ... "
rm -R Jarvis-Desktop-linux-x64
