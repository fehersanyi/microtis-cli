package cmd

import (
	"github.com/bitrise-io/go-utils/log"
	"github.com/fehersanyi/microtis-cli/git"
	"github.com/spf13/cobra"
)

// kiscicaCmd represents the kiscica command
var kiscicaCmd = &cobra.Command{
	Use:   "kiscica",
	Short: "will add + commit",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := git.Kiscica(args); err != nil {
			log.Errorf("Nincs kiscica: \n%s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(kiscicaCmd)
}
