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

func fileChecker(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func setAlias(alias string) error {
	home := os.Getenv("HOME")
	bashProfile := home + "/.bash_profile"
	profile := home + "/.zprofile"
	if err := fileChecker(bashProfile); err != nil {
		fmt.Println("filecheck bash profile")
		return err
	}
	if err := fileChecker(profile); err != nil {
		fmt.Println("filecheck profile")
		return err
	}
	if err := setFile(bashProfile, alias); err != nil {
		fmt.Println("file write bash profile")
		return err
	}
	// if err := sourcing(bashProfile); err != nil {
	// 	fmt.Println("source bash profile")
	// 	return err
	// }
	if err := setFile(profile, alias); err != nil {
		fmt.Println("file write profile")
		return err
	}
	// if err := sourcing(profile); err != nil {
	// 	fmt.Println("source profile")
	// 	return err
	// }
	return nil
}

func setFile(file, alias string) error {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		if err.Error() == fmt.Sprintf(`open %s: no such file or directory`, file) {
			exec.Command("touch", file)
			if err := setFile(file, alias); err != nil {
				return err
			}
		}
	}
	set := fmt.Sprintf("alias %s=~/.microtis/microtis\n", alias)
	d := string(s) + set
	if err := ioutil.WriteFile(file, []byte(d), 777); err != nil {
		return err
	}
	return nil
}

func sourcing(file string) error {
	_, err := exec.Command("bash", "-c", "source", file).Output()
	if err != nil {
		return err
	}
	return nil
}
