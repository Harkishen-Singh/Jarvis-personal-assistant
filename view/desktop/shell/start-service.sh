#!/usr/bin/env bash

if [ -d "Jarvis-personal-assistant" ]; then
    echo "folder exists! skipping service installation steps"
else
    echo "Jarvis-personal-assistant directory does not exists! Downloading ..."
    ./install-service.sh
    echo "installation of Jarvis-personal-assistant service successful"
fi

cd Jarvis-personal-assistant/service/
echo "preparing to execute service .."
go run main.go &
echo "successfully executed the service"
