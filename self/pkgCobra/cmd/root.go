/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	//"os"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var log = logrus.WithField("module", "main")


// rootCmd represents the base command when called without any subcommands
var (
	configPath    string
	RootCmd = &cobra.Command{
		Use:   "pkgCobra",
		Short:	"Run the pkgCobra test",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("root.go  cmd.rootCmd.PersistentPreRun() ++++++")

			err := viper.BindPFlags(cmd.Flags())
			if err != nil {
				log.Fatal(err)
			}
		},
		
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) { 
			fmt.Println("root.go  cmd.rootCmd.Run() ++++++")
			cmd.HelpFunc()(cmd, args)
		},
	}
)

func init() {
	fmt.Println("root.go cmd.init()  ++++++")
	cobra.OnInitialize(func() {
		fmt.Println("root.go cmd.init() cobra.OnInitialize() ++++++")
	})

	RootCmd.PersistentFlags().StringVar(&configPath, "config","", "/path/to/config.yml")

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


