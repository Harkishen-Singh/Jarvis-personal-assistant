# Setup

To run this project, the system should have the following prerequisites, if not then click on the link mentioned in the prerequites and install it and then follow the installation steps.

## Prerequisites

The prerequisites for installing the project :
1. [golang](https://golang.org/dl/ "Install GOLang")
2. [nodejs with npm](https://nodejs.org/en/download/ )


## Installation
1. `sudo npm install -g http-server selenium-webdriver mocha mochawesome eslint`
2. `go get -u github.com/op/go-logging`
3. `go get -u github.com/hegedustibor/htgo-tts`
4. `go get -u github.com/Harkishen-Singh/Jarvis-personal-assistant/service/controllers`
5. `go get -u github.com/Harkishen-Singh/Jarvis-personal-assistant/service/utils`
6.  Execute the command `npm start` in the view folder
7.  Execute the command `go run *.go` in the service folder
### for executing tests
8. `eslint view/app-jarvis.js` # for linting checks
9. `cd tests && npm install && mocha tests.js`
