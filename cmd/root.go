package cmd

import (
	"fmt"
	"os"

	"github.com/bitrise-io/go-utils/log"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "microtis-cli",
	Short: "CLI app for microtis business",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		version := "0.1.0"
		fmt.Println(`
	███╗   ███╗██╗ ██████╗██████╗  ██████╗ ████████╗██╗███████╗
	████╗ ████║██║██╔════╝██╔══██╗██╔═══██╗╚══██╔══╝██║██╔════╝
	██╔████╔██║██║██║     ██████╔╝██║   ██║   ██║   ██║███████╗
	██║╚██╔╝██║██║██║     ██╔══██╗██║   ██║   ██║   ██║╚════██║
	██║ ╚═╝ ██║██║╚██████╗██║  ██║╚██████╔╝   ██║   ██║███████║
	╚═╝     ╚═╝╚═╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝    ╚═╝   ╚═╝╚══════╝`)
		fmt.Println()
		log.Successf("Version: %s", version)
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println()
		fmt.Print("microtis ")
		log.Infof("[command] [argument]")
		fmt.Println()
		fmt.Println("Available commands:")
		fmt.Println()
		fmt.Println("		holvagyok		will print out the current repository, and branch")
		fmt.Println("		kiscica			will add changes, and commit, commit message will be the argument list you provide")
		fmt.Println("		cica			will kiscica and push to the current branch, commit message will be the argument list you provide")
		fmt.Println("		alias			the binary name is long, with this you can add an alias to make your work easier")
		fmt.Println("		update			will update to the latest version of the app")
		fmt.Println()
		fmt.Print("this cli app was created by ")
		log.Infof("Sandor Feher")
		fmt.Println("If you meet him, please consider buying him a beer. ;)")
	},
}

//Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
