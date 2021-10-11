# :minidisc: Installation Instructions

## Prerequisites

To run this project, the system should have the following prerequisites, if not then click on the link mentioned in the prerequites and install it and then follow the installation steps.

The prerequisites for installing the project :
1. [golang](https://golang.org/dl/ "Install GOLang")
2. [nodejs with npm](https://nodejs.org/en/download/ )

## Installation and executing web-app
1. Clone the repository in your GOPATH

    ```bash
        $ mkdir -p $GOPATH/src/github.com/Harkishen-Singh/
        $ cd $GOPATH/src/github.com/Harkishen-Singh/
        $ git clone https://github.com/Harkishen-Singh/Jarvis-personal-assistant.git
        $ cd Jarvis-personal-assistant
    ```
2. Install Go and npm packages.
    ```
        $ sudo npm install -g http-server mocha mochawesome eslint
    ```
    Important, execute separately: 
    ```
        $ go get -v -u -d ./...
    ```
3. Start the HTTP server. *Do not close this terminal*
    ```bash
        cd view/web && http-server
    ```
4. Start the GoLang server. *Do not close this terminal*
    ```bash
        cd service && go run main.go
    ```
5. [optional but recommended] For listening to the jarvis voice download and install **mplayer** in linux or windows, and set the **path** so that it responds when typed `mplayer` in the terminal or command prompt.

6. On running the service, it will occupy the port `:3000` and the front-end will run at port `:8080`

### for executing tests

7. `eslint view/web/app-jarvis.js` # for linting checks
8. `cd tests && npm install && mocha tests.js`

## Running the Desktop-app

***prerequisite :***
1. ```golang``` with path as global
2. ```node version >= 7.0.0``` & corresponding ```npm```

*copy the following instructions in your terminal :*

1. Move to the desktop app folder
    ```bash
    cd view/desktop
    ```
2. execute the command in the terminal
    ```bash
    npm install
    ```
3. Then run to locally execute jarvis-desktop
    ```bash
     npm start
     ```
