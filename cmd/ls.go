package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	NebCmd.AddCommand(lsCmd)
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List your nebs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("neb neb neb")
	},
}
