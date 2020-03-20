// This file is slate due to change of
// aims in the project.
// If future plannings happen to change
// this would be indeed needed.

package herokuhost

import (
	"fmt"
	"net/http"
	"os/exec"
)

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

	return resp.Message
}
