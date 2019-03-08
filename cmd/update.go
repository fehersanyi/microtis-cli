package cmd

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/bitrise-io/go-utils/log"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "this will update to the latest version of microtis",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := clone(); err != nil {
			log.Errorf("%s", err)
		}
		os := getOS()
		if err := access(os); err != nil {
			log.Errorf("failed to grant access %s", err)
		}
		home := home()
		if err := makeDir(home + "/.microtis"); err != nil {
			log.Errorf("failed to create directory: %s", err)
		}
		if err := touchBashProfile(home); err != nil {
			log.Errorf("failed to create bash_profile %s", err)
		}
		if err := move(os, home); err != nil {
			log.Errorf("failed to move binary %s", err)
		}
		err := pathExport(home)
		if err != nil {
			log.Errorf("failed to export PATH %s", err)
		}
		if err := remove(); err != nil {
			log.Errorf("failed to remove tmp dir %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func clone() error {
	log.Infof("Cloning into /tmp/bin")
	_, err := exec.Command("git", "clone", "https://github.com/fehersanyi/bin.git", "/tmp/bin/").Output()
	if err != nil {
		return err
	}
	log.Successf("Cloning succeeded")
	return nil
}

func access(os string) error {
	log.Infof("granting exec right to binary")
	if os == "darwin" {
		_, err := exec.Command("chmod", "777", "/tmp/bin/mac/microtis").Output()
		if err != nil {
			return err
		}
	} else {
		_, err := exec.Command("chmod", "777", "/tmp/bin/unix/microtis").Output()
		if err != nil {
			return err
		}
	}
	return nil
}

func makeDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err := exec.Command("mkdir", path).Output()
		if err != nil {
			return err
		}
	}
	return nil
}

func move(ops, home string) error {
	log.Infof("Overwriting binary")
	if ops == "mac" {
		_, err := exec.Command("cp", "-f", "/tmp/bin/mac/microtis", home+"/.microtis/microtis").Output()
		if err != nil {
			return err
		}
	} else {
		_, err := exec.Command("cp", "/tmp/bin/unix/microtis", home+"/.microtis/microtis").Output()
		if err != nil {
			return err
		}
	}
	return nil
}

func getOS() string {
	if runtime.GOOS == "darwin" {
		return "mac"
	}
	return "unix"
}

func remove() error {
	log.Infof("removing temp dir")
	_, err := exec.Command("rm", "-rf", "/tmp/bin", "-y").Output()
	if err != nil {
		return err
	}
	return nil
}

func home() string {
	home := os.Getenv("HOME")
	return home
}

func touchBashProfile(home string) error {
	if _, err := os.Stat(home + "/.bash_profile"); os.IsNotExist(err) {
		_, err := exec.Command("touch", "~/.bash_profile").Output()
		if err != nil {
			return err
		}
	}
	return nil
}
func pathExport(home string) error {
	log.Infof("Exporting PATH")
	bashProfile := home + "/.bash_profile"

	profile, err := ioutil.ReadFile(bashProfile)
	if err != nil {
		return err
	}
	if strings.Contains(string(profile), "export PATH=~/.mictoris/:$PATH\n") {
		return nil
	}
	overWrite := string(profile) + "export PATH=~/.mictoris/:$PATH\n"
	err = ioutil.WriteFile(bashProfile, []byte(overWrite), 777)
	if err != nil {
		return err
	}
	_, err = exec.Command("bash", "-c", "export", "PATH=~/.mictoris/:$PATH").Output()
	if err != nil {
		return err
	}
	return nil
}
