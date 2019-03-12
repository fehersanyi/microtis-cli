package cmd

import (
	"github.com/bitrise-io/go-utils/log"
	"github.com/fehersanyi/microtis-cli/stargate"
	"github.com/spf13/cobra"
)

// jumpCmd represents the jump command
var jumpCmd = &cobra.Command{
	Use:   "jump",
	Short: "jump will hehe, jump to a given directory for you",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if err := stargate.Jump(args[0]); err != nil {
				log.Errorf("Tilk, something happened, we can't jump: \n%s", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(jumpCmd)
}
