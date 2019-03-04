package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "microtis-cli",
	Short: "CLI app for microtis business",
	Long: `
	███╗   ███╗██╗ ██████╗██████╗  ██████╗ ████████╗██╗███████╗
	████╗ ████║██║██╔════╝██╔══██╗██╔═══██╗╚══██╔══╝██║██╔════╝
	██╔████╔██║██║██║     ██████╔╝██║   ██║   ██║   ██║███████╗
	██║╚██╔╝██║██║██║     ██╔══██╗██║   ██║   ██║   ██║╚════██║
	██║ ╚═╝ ██║██║╚██████╗██║  ██║╚██████╔╝   ██║   ██║███████║
	╚═╝     ╚═╝╚═╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝    ╚═╝   ╚═╝╚══════╝
																														 
	
	version: 0.0.1
`,
}

//Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
