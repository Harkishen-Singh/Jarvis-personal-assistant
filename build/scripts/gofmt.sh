#!/bin/bash

echo 'running go fmt on all packages...'
invalidFiles=$(gofmt -l . 2>&1)
if [ "$invalidFiles" ]; then
  echo "These files did not pass the 'go fmt' check, please run 'go fmt' on them:"
  echo $invalidFiles
  exit 1
fi
