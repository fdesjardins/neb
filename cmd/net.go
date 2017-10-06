package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	NebCmd.AddCommand(netCmd)
}

var netCmd = &cobra.Command{
	Use:   "net",
	Short: "Network your nebs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[neb net]")
	},
}
