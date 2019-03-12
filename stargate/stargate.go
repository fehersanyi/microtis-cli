package stargate

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bitrise-io/go-utils/log"
)

//Jump will cd into a given directory
func Jump(alias string) error {
	path, err := readMap(alias)
	fmt.Print("hetes égzár kódolva")
	if err != nil {
		log.Errorf("[X]")
		return err
	}
	log.Successf("[✓]")
	_, err = exec.Command("cd", path).Output()
	if err != nil {
		return err
	}
	fmt.Print("sikeresen megérkezett: ")
	log.Infof("%s", path)
	return nil
}

//CheckPoint will mark a directory
func CheckPoint(alias string) error {
	err := createMap()
	fmt.Print("mapfile found: ")
	if err != nil {
		log.Errorf("[X]")
		return err
	}
	log.Successf("[✓]")

	path, err := getLocation()
	if err != nil {
		return err
	}
	if err := recordLocation(alias, path); err != nil {
		return err
	}
	fmt.Print("saved location: ")
	log.Infof("%s", strings.TrimSpace(path))
	fmt.Print("you can reach here by running:")
	log.Successf(" jump %s", alias)
	return nil
}

/*
this will make a .map file
that will hold the checkpoints with a
name=pwd fomat
*/
func createMap() error {
	home := getHome()
	mapfile := ".microtis/mapfile"
	f, err := os.Create(filepath.Join(home, mapfile))
	if err != nil {
		return err
	}
	defer f.Close()
	return err
}

/*
this shall read the .map file and return a path
based on the input
*/
func readMap(alias string) (string, error) {
	home := getHome()
	mapfile := ".microtis/mapfile"
	mapFile, err := ioutil.ReadFile(filepath.Join(home, mapfile))
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
	home := getHome()
	mapfile := ".microtis/mapfile"
	data := []byte(alias + "=" + dir + "\n")
	if err := ioutil.WriteFile(filepath.Join(home, mapfile), data, 0777); err != nil {
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

func getHome() string {
	home := os.Getenv("HOME")
	return home
}
