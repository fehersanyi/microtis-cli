package cmd

import (
	"fmt"
	"strings"

	"github.com/bitrise-io/go-utils/log"
	"github.com/fehersanyi/microtis-cli/stargate"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "will print every custom setup",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		gates, err := stargate.ListMap()
		if err != nil {
			log.Errorf(err.Error())
		}
		fmt.Println()
		log.Infof("Your list of stargates: ")
		fmt.Println()
		if len(gates) > 0 {
			for i := 0; i < len(gates); i++ {
				gate := strings.Split(gates[i], "=")
				fmt.Printf("jump %s will port you to: ", gate[0])
				log.Infof("%s", gate[1])
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
