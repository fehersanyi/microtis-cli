package stargate

import "os/exec"

//Jump will cd into a given directory
func Jump() {}

//CheckPoint will mark a directory
func CheckPoint() {}

/*
this will make a .map file
that will hold the checkpoints with a
name=pwd fomat
*/
func createMap() {}

/*
this shall read the .map file and return a path
based on the input
*/
func readMap() {}

/*
this will write the given path to the .bash_profile
it is important to log out that after this you need to
source ~/.bash_profile !!!!
*/
func recordLocation() {}

// thisl will get the current location
func getLocation() (string, error) {
	path, err := exec.Command("pwd").Output()
	if err != nil {
		return string(path), err
	}
	return string(path), nil
}
