package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	NebCmd.AddCommand(storeCmd)
}

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Put stuff in your nebs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[neb contents]")
	},
}
