FROM fedora:latest

RUN dnf install golang -y
RUN mkdir -p $HOME/go && echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc && source $HOME/.bashrc
RUN go env GOPATH

RUN dnf install nodejs -y

RUN go get -u github.com/Harkishen-Singh/Jarvis-personal-assistant/service
RUN cd root/go/src/github.com/Harkishen-Singh/Jarvis-personal-assistant/service && go get -u -v -d ./...
WORKDIR /root/go/src/github.com/Harkishen-Singh/Jarvis-personal-assistant/service/
EXPOSE 3000
CMD go run maintut.go
