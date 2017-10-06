package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NebCmd comment
var NebCmd = &cobra.Command{
	Use:   "neb",
	Short: "nebulous",
	Long:  "cloud-agnostic infrastructure cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Neb")
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	viper.SetDefault("license", "MIT")
}

func initConfig() {
	viper.SetConfigFile("config")
}

// Execute comment
func Execute() {
	NebCmd.Execute()
}
