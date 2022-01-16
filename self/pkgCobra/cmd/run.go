package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)


var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Say hello!",
	Long:  "Address a wonderful greeting to the majestic executioner of this CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run.go cmd.runCmd.Run() +++++++")
	},
}

func init() {
	fmt.Println("run.go cmd.init() ++++++")
	RootCmd.AddCommand(runCmd)

	addAPIFlags(runCmd)
	addDBFlags(runCmd)
	addAddressesFlags(runCmd)

}