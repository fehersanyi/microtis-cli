package cmd

import (
	"github.com/bitrise-io/go-utils/log"
	"github.com/fehersanyi/microtis-cli/stargate"
	"github.com/spf13/cobra"
)

// checkinCmd represents the checkin command
var checkinCmd = &cobra.Command{
	Use:   "checkin",
	Short: "checkin will mark down your current position",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if err := stargate.CheckPoint(args[0]); err != nil {
				log.Errorf("You can't mark this place\n%s", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkinCmd)
}
