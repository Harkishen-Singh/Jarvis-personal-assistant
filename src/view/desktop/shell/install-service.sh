#!/usr/bin/env bash

which git
which go

echo "fetching jarvis-service from https://github.com/Harkishen-Singh/Jarvis-personal-assistant"

git clone https://github.com/Harkishen-Singh/Jarvis-personal-assistant
echo "finished"

cd Jarvis-personal-assistant/service/
echo "installing go dependencies"
go get -u -v -d ./...

echo "preparing to execute service .."
go run main.go &
echo "successfully executed the service"

