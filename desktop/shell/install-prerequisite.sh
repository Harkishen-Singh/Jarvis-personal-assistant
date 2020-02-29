#!/usr/bin/env bash
# run this file under sudo

echo "installing prerequisites"
platform = ./platform.sh

if [[ ${platform} == *_RedHat_* ]]; then
    # tested
    dnf install golang nodejs -y
    mkdir -p $HOME/go
    echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc
    source $HOME/.bashrc
    go env GOPATH
elif [[ ${platform} == *_Debian_* ]]; then
    # untested
    wget https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.4.2.linux-amd64.tar.gz
    rm go1.4.2.linux-amd64.tar.gz
    mkdir $HOME/go
    echo 'export GOPATH=$HOME/go'
    echo 'PATH=$PATH:/usr/local/go/bin:$GOPATH/bin'
    source $HOME/.bashrc

    apt-get install nodejs-legacy npm -y

# provide support for other platforms as well
fi
echo "complete!"
