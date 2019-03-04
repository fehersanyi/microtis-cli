package git

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/bitrise-io/go-utils/log"
)

//Holvagyok ...
func Holvagyok() error {
	//fetching repository
	var repository string
	config, err := ioutil.ReadFile("./.git/config")
	if err != nil {
		return err
	}
	configFile := strings.Split(string(config), "\n")
	for _, line := range configFile {
		if strings.Contains(line, "url") {
			l := strings.Split(line, " ")
			repository = l[len(l)-1]
		}
	}

	//fetching branch
	var branch string
	cmd, err := exec.Command("git", "branch").Output()
	if err != nil {
		return err
	}
	branches := strings.Split(string(cmd), "\n")
	for _, line := range branches {
		if strings.Contains(line, "*") {
			l := strings.Split(line, " ")
			branch = l[len(l)-1]
		}
	}
	//fetch local path
	path, err := exec.Command("pwd").Output()
	if err != nil {
		return err
	}
	simplePath := strings.Split(string(path), "\n")
	fmt.Print("locally: ")
	log.Infof(simplePath[0])
	fmt.Print("repository: ")
	log.Infof(repository)
	fmt.Print("branch: ")
	log.Infof(branch)
	return nil
}
