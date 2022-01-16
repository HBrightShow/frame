package cmd

import(
	"fmt"
	"io/ioutil"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/lacasian/gogo/confgen"
)

var (
	ignore = []string{"verbose", "v","vv", "version", "config", "help", "connection-string", "with-defaults" }
)

var(
	generateConfigCmd = &cobra.Command{
		Use:	"generate-config",
		Short:	"generate a sample config file",
		Long:	"generates a sample config file named config-generated.yml",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmd .generateConfigCmd.PersistentPreRun() ++++++")
			err := viper.BindPFlags(cmd.Flags())
			if err != nil {
				log.Fatal(err)
			}
			if !viper.GetBool("with-defaults") {
				RootCmd.PersistentPreRun(cmd, args)
			}
		
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmd.generateConfigCmd.Run() ++++++")
			log.Println("start config")
			c := viper.AllSettings()
			fmt.Println("viper length: ", len(c))
			for key,value := range c {
				fmt.Printf("%s=>%s\n", key, value)
			}

			ba, err := confgen.Viper(c, cmd, ignore) 
			if err != nil {
				log.Fatal(err)
			}

			err = ioutil.WriteFile("config-generated.yal",ba, 0644)
			if err != nil {
				log.Fatal(err)
			}

			log.Info("done writing config")
			fmt.Println("done writing config")
		},
	}
)

func init() {
	fmt.Println("generate-config init() ++++++")
	RootCmd.AddCommand(generateConfigCmd)
	addDBFlags(generateConfigCmd)
	addAPIFlags(generateConfigCmd)
	addMetricsFlags(generateConfigCmd)
	addAddressesFlags(generateConfigCmd)
}