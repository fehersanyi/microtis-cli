package cmd

import (
	"github.com/bitrise-io/go-utils/log"
	"github.com/fehersanyi/microtis-cli/git"
	"github.com/spf13/cobra"
)

// holvagyokCmd represents the holvagyok command
var holvagyokCmd = &cobra.Command{
	Use:   "holvagyok",
	Short: "prints out wheretheHELLyouare",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := git.Holvagyok(); err != nil {
			log.Errorf("Nem tudom \n%s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(holvagyokCmd)
}
