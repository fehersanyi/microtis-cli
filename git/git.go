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

//Kiscica ...
func Kiscica(args []string) error {
	_, err := exec.Command("git", "add", ".").Output()
	if err != nil {
		return err
	}
	status, err := exec.Command("git", "status").Output()
	if err != nil {
		return err
	}
	statusLines := strings.Split(string(status), "\n")
	var addedFiles []string
	var modifiedFiles []string
	for _, line := range statusLines {
		if strings.Contains(line, "new file") {
			s := strings.Split(line, " ")
			addedFiles = append(addedFiles, s[len(s)-1])
		}
		if strings.Contains(line, "modified") {
			s := strings.Split(line, " ")
			modifiedFiles = append(modifiedFiles, s[len(s)-1])
		}
	}

	//commiting changes
	cm := strings.Join(args, " ")
	message := []string{"commit", "-m"}
	message = append(message, cm)
	log.Warnf("%s", message)
	kiscica, err := exec.Command("git", message...).Output()
	if err != nil {
		log.Printf("itt")
		log.Printf("%s", kiscica)
		return err
	}

	fmt.Print("The following files were added: ")
	log.Warnf("%s", addedFiles)
	fmt.Print("You changed these files: ")
	log.Warnf("%s", modifiedFiles)
	fmt.Print("You committed them with the following message: ")
	log.Successf("%s", message[1:])
	return nil
}

//Cica ...
func Cica(args ...string) error {
	return nil
}
