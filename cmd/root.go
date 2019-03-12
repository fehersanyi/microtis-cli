package cmd

import (
	"fmt"
	"os"

	"github.com/bitrise-io/go-utils/log"
	"github.com/spf13/cobra"
)

var cfgFile string

var version = "0.2.2"

var rootCmd = &cobra.Command{
	Use:   "microtis-cli",
	Short: "CLI app for microtis business",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
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
		fmt.Println()
		fmt.Println("		checkin			will save your current location with the given arg as an alias")
		fmt.Println("		jump			will jump into the saved directory you specify by the argument")
		fmt.Println()
		fmt.Println("		help			will print a short explanation on usage")
		fmt.Println("		version			will plain print the current version")
		fmt.Println()
		fmt.Println("		update			will update to the latest version of the app")
		fmt.Println()
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
