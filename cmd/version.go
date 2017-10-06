package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	NebCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "v0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.1")
	},
}
