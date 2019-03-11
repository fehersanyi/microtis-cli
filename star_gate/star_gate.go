package stargate

import (
	"io/ioutil"
	"os/exec"
	"strings"
)

//Jump will cd into a given directory
func Jump(alias string) error {
	path, err := readMap(alias)
	if err != nil {
		return err
	}
	_, err = exec.Command("cd", path).Output()
	if err != nil {
		return err
	}
	return nil
}

//CheckPoint will mark a directory
func CheckPoint(alias string) error {
	if err := createMap(); err != nil {
		return err
	}
	path, err := getLocation()
	if err != nil {
		return err
	}
	if err := recordLocation(alias, path); err != nil {
		return err
	}
	return nil
}

/*
this will make a .map file
that will hold the checkpoints with a
name=pwd fomat
*/
func createMap() error {
	_, err := exec.Command("touch", "~/.microtis/.map").Output()
	if err != nil {
		return err
	}
	return err
}

/*
this shall read the .map file and return a path
based on the input
*/
func readMap(alias string) (string, error) {
	mapFile, err := ioutil.ReadFile("~/.microtis/.map")
	path := ""
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(mapFile), "\n")
	if len(lines) > 0 {
		for _, line := range lines {
			if strings.Contains(line, alias) {
				p := strings.Split(line, "=")
				path = p[1]
			}
		}
	}
	return path, err
}

/*
this will write the given path to the .bash_profile
it is important to log out that after this you need to
source ~/.bash_profile !!!!
*/
func recordLocation(alias, dir string) error {
	data := []byte(alias + "=" + dir + "\n")
	if err := ioutil.WriteFile("~/.microtis/.map", data, 777); err != nil {
		return err
	}
	return nil
}

// thisl will get the current location
func getLocation() (string, error) {
	path, err := exec.Command("pwd").Output()
	if err != nil {
		return string(path), err
	}
	return string(path), nil
}
