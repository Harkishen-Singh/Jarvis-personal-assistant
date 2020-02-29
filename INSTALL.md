# Setup

To run this project, the system should have the following prerequisites, if not then click on the link mentioned in the prerequites and install it and then follow the installation steps.

## Prerequisites

The prerequisites for installing the project :
1. [golang](https://golang.org/dl/ "Install GOLang")
2. [nodejs with npm](https://nodejs.org/en/download/ )
3. [python v2](https://www.python.org/)

## Installation

1. `sudo npm install -g http-server selenium-webdriver mocha mochawesome eslint`
2. `npm install chromedriver`
3. `go get -v -u -d ./...`
4. Install Python 2.*.* . To verify if its installed, `python --version`.
5. Execute the command `pip install -r requirements.txt` in the subprocesses folder. Use `sudo` if required!
6.  Execute the command `http-server` in the view folder. *Do not close this terminal*.
7.  Execute the command `go run maintut.go` in the service folder. *Do not close this terminal*.
8. [optional but recommended] For listening to the jarvis voice download and install **mplayer** in linux or windows, and set the **path** so that it responds when typed `mplayer` in the terminal or command prompt.

### for executing tests

8. `eslint view/app-jarvis.js` # for linting checks
9. `cd tests && npm install && mocha tests.js`
