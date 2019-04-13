package herokuhost

import (
	"fmt"
	"os/exec"
)

func checkDockerInstallation() bool {

	res, err := exec.Command("docker info").Output()
	if err != nil || len(string(res)) == 0 {
		fmt.Println("[JARVIS] seems like docker is not installed")
		return false
	}

	return true
}

func checkHerokuCLIInstallation() bool {

	res, err := exec.Command("heroku").Output()
	if err != nil || len(string(res)) == 0 {
		fmt.Println("[JARVIS] seems like heroku is not installed")
		return false
	}

	return true
}

// DeploymentFunction facilitates autodeployment functionality of the project using docker with heroku client
func DeploymentFunction() {

	fmt.Println("-do-")
}