// This file is slate due to change of
// aims in the project.
// If future plannings happen to change
// this would be indeed needed.

package herokuhost

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

const (
	// auto-deployment is disabled for now.
	// shouldbe done in future versions if required.
	herokuMailID string = ""
)

func herokuLogin() bool {
	return false
}

type automateDeploymentContainer interface {
	herokuLogin() bool
	herokuContainerPush() bool
	herokuContainerLogin() bool
	herokuCreate() string
	herokuContainerRelease() bool
	herokuOpen() bool
}

type herokuDetails struct {
	emailID, password string
}

func (dep herokuDetails) herokuLogin() bool {

	res, err := exec.Command("heroku login").Output()
	if err != nil {
		fmt.Println("heroku login error")
		panic(err)
	}
	fmt.Printf("heroku login \n %s", res)
	return true
}

func herokuContainerPush() bool {
	res, err := exec.Command("heroku container:push web").Output()
	if err != nil {
		fmt.Println("heroku container push error")
		panic(err)
	}
	fmt.Printf("heroku container push \n %s", res)
	return true
}

func herokuContainerLogin() bool {
	res, err := exec.Command("heroku container:login").Output()
	if err != nil {
		fmt.Println("heroku container login error")
		panic(err)
	}
	fmt.Printf("heroku container login \n %s", res)
	return true
}

func herokuCreate(expectedName string) string {
	res, err := exec.Command("heroku create", expectedName).Output()
	if err != nil {
		fmt.Println("heroku create error")
		panic(err)
	}
	fmt.Printf("heroku create \n %s", res)
	return string(res)
}

func herokuContainerRelease() bool {
	res, err := exec.Command("heroku container:release web").Output()
	if err != nil {
		fmt.Println("heroku container release error")
		panic(err)
	}
	fmt.Printf("heroku container release \n %s", res)
	return true
}

func herokuOpen() bool {
	res, err := exec.Command("heroku open").Output()
	if err != nil {
		fmt.Println("heroku container release error")
		panic(err)
	}
	fmt.Printf("heroku container release \n %s", res)
	return true
}

type automateDeploymentGithub interface {
	herokuGithubSubprocess() string
	herokuGithubLogs() string
}

type herokuGithubCredentials struct {
	email, password, repoName, result string
}

func (cred herokuGithubCredentials) herokuGithubSubprocess() string {
	fmt.Println(cred)
	fmt.Println("upside")

	res, err := exec.Command("node", "subprocesses/deploy_heroku.js", cred.repoName).Output()
	if err != nil {
		fmt.Printf("[JARVIS] error occurred while handling heroku deployment subprocess")
		panic(err)
	}
	cred.result = string(res)
	fmt.Printf("from subprocess")
	fmt.Printf("%s", res)
	return string(res)
}

func (cred herokuGithubCredentials) herokuGithubLogs(strs string) string {

	logs := strs
	logsLen := len(logs)
	subs := "link"
	subslen := len(subs)
	var appLink string
	for i := 0; i < logsLen-subslen; i++ {
		if logs[i:i+subslen] == subs {
			appLink = logs[i+subslen : logsLen]
			fmt.Printf("appLink hosted at %s", appLink)
			return appLink
		}
	}
	return "deployment at heroku containers failed!"
}

type deployResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// DeploymentFunction facilitates autodeployment functionality of the project using docker with heroku client
func DeploymentFunction(repoName string, res http.ResponseWriter) string {

	credentials := herokuGithubCredentials{
		email:    "harkishensingh@hotmail.com",
		password: "",
		repoName: repoName,
		result:   "",
	}
	strs := credentials.herokuGithubSubprocess()
	response := credentials.herokuGithubLogs(strs)
	var resp deployResponse
	if response == "deployment at heroku containers failed!" {
		resp.Status = true
		resp.Message = "deployment at heroku containers failed!"
	} else {
		resp.Status = true
		resp.Message = "Successfully deployed. Link " + response
	}
	unmarshall, _ := json.Marshal(resp)
	res.Write(unmarshall)
	return resp.Message
}
