# Setup

To run this project, the system should have the following prerequisites, if not then click on the link mentioned in the prerequites and install it and then follow the installation steps.

## Prerequisites

The prerequisites for installing the project :
1. [golang](https://golang.org/dl/ "Install GOLang")
2. [nodejs with npm](https://nodejs.org/en/download/ )
3. [python v2](https://www.python.org/)

## Installation

1. Run the following commands: `sudo npm install -g http-server selenium-webdriver mocha mochawesome eslint`   
2. `npm install chromedriver` to install the dependencies.
3. Make a default repository for cloning the project. This should be strictly inside the GOPATH. 
4. Navigate to the directory via `cd $GOPATH/src`
5. Clone the Repo via `git clone https://github.com/Harkishen-Singh/Jarvis-personal-assistant.git`
6. Navigate to the cloned Repository via `cd Jarvis-personal-assistant`
7. Install all the GO dependencies via `go get -v -u -d ./...` 
8. Install Python 2.*.* . To verify if its installed, `python --version`.
9. Execute the command `pip install -r requirements.txt` in the ./service/subprocesses folder. Use `sudo` if required!
10.  Execute the command `http-server` in the view folder. *Do not close this terminal*.
11.  Execute the command `go run maintut.go` in the service folder. *Do not close this terminal*.
12. [optional but recommended] For listening to the jarvis voice download and install **mplayer** in linux or windows, and set the **path** so that it responds when typed `mplayer` in the terminal or command prompt.
13. On Successfully executing the above commands Angular server will start at PORT: 8081 and GOLang Server on port: 3000. Open `http://localhost:8081` to get started.

### for executing tests

14. `eslint view/app-jarvis.js` # for linting checks
15. `cd tests && npm install && mocha tests.js`
