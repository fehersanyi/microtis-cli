package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/bitrise-io/go-utils/log"
	"github.com/spf13/cobra"
)

// aliasCmd represents the alias command
var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "with this you can set a shorter alias for microtis",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if err := setAlias(args[0]); err != nil {
				log.Errorf("%s", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(aliasCmd)
}

func setAlias(alias string) error {
	home := os.Getenv("HOME")
	bashProfile := home + "/.bash_profile"
	profile := home + "/.zprofile"
	setFile(bashProfile, alias)
	sourcing(bashProfile)
	setFile(profile, alias)
	sourcing(profile)
	return nil
}

func setFile(file, alias string) error {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	set := fmt.Sprintf("alias %s=/usr/local/bin/\"microtis\"\n", alias)
	d := string(s) + set
	ioutil.WriteFile(file, []byte(d), 777)
	return nil
}

func sourcing(file string) error {
	_, err := exec.Command("source", file).Output()
	if err != nil {
		return err
	}
	return nil
}
