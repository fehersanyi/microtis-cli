package cmd

import (
	"github.com/bitrise-io/go-utils/log"
	"github.com/fehersanyi/microtis-cli/git"
	"github.com/spf13/cobra"
)

// cicaCmd represents the cica command
var cicaCmd = &cobra.Command{
	Use:   "cica",
	Short: "cica is a grown ass kiscica and will push your code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := git.Cica(args); err != nil {
			log.Errorf("No cica: \n%s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cicaCmd)
}
