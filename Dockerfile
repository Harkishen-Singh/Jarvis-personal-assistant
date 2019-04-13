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
RUN go env GOPATH

# dependencies installation
RUN npm install -g http-server selenium-webdriver mocha mochawesome eslint
RUN python --version

# install git
RUN apt install git -y
RUN git --version

RUN go get -u github.com/Harkishen-Singh/Jarvis-personal-assistant/service

RUN cd root/go/src/github.com/Harkishen-Singh/Jarvis-personal-assistant/service && go get -u -v -d ./...
WORKDIR /root/go/src/github.com/Harkishen-Singh/Jarvis-personal-assistant/service
CMD go run maintut.go