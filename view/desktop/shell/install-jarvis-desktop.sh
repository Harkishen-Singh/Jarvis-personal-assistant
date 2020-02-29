#!/usr/bin/env bash

which npm

cd ..
echo "install development dependencies"
npm install --dev

echo "test existing code"
npm test
