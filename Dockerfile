FROM ubuntu:latest

RUN apt-get update
RUN apt-get dist-upgrade -y
# for speech to text
RUN apt install mplayer -y

# installing dependencies

# # nodejs
RUN apt-get install nodejs -y
RUN apt-get install npm -y
RUN node --version
RUN npm --version

# # python
RUN apt-get install python2.7 -y
RUN apt-get install python -y

# # golang
RUN apt-get update
RUN apt-get install golang-go -y
RUN mkdir -p $HOME/go
RUN echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc
RUN source $HOME/.bashrc
RUN go env GOPATH

# dependencies installation
RUN npm install -g http-server selenium-webdriver mocha mochawesome eslint
RUN go get -v -u -d ./...
RUN python --version
