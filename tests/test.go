package main

import (
	"fmt"
	"os/exec"
)

func init() {
	res, err := exec.Command("docker", "info").Output()
	if err != nil || len(string(res)) == 0 {
		fmt.Println("[JARVIS] seems like docker is not installed")
	}
	res, err = exec.Command("heroku").Output()
	if err != nil || len(string(res)) == 0 {
		fmt.Println("[JARVIS] seems like heroku cli is not installed")
	}
}

type herokuDetails struct {
	emailID, password string
}

func (dep herokuDetails) herokuLogin() bool {

	res, err := exec.Command("heroku", "login").Output()
	if err != nil {
		fmt.Println("heroku login error")
		panic(err)
	}
	fmt.Printf("heroku login \n %s", res)
	return true
}

func herokuContainerPush() bool {
	res, err := exec.Command("heroku", "container:push web").Output()
	if err != nil {
		fmt.Println("heroku container push error")
		panic(err)
	}
	fmt.Printf("heroku container push \n %s", res)
	return true
}

func herokuCreate(expectedName string) string {
	res, err := exec.Command("heroku", "create", expectedName).Output()
	if err != nil {
		fmt.Println("heroku create error")
		panic(err)
	}
	fmt.Printf("heroku create \n %s", res)
	return string(res)
}

func herokuContainerRelease() bool {
	res, err := exec.Command("heroku", "container:release web").Output()
	if err != nil {
		fmt.Println("heroku container release error")
		panic(err)
	}
	fmt.Printf("heroku container release \n %s", res)
	return true
}

func herokuOpen() bool {
	res, err := exec.Command("heroku", "open").Output()
	if err != nil {
		fmt.Println("heroku container release error")
		panic(err)
	}
	fmt.Printf("heroku container release \n %s", res)
	return true
}

// DeploymentFunction facilitates autodeployment functionality of the project using docker with heroku client
func DeploymentFunction() {

	obj := herokuDetails{
		emailID:  "harkishensingh@hotmail.com",
		password: "Bbsr@131",
	}
	obj.herokuLogin()
	herokuCreate("")
	herokuContainerPush()
	herokuContainerRelease()
	herokuOpen()
	fmt.Println("-do-")
}

func main() {
	DeploymentFunction()
}
