package main

import (
	"fmt"
	"os"

	"github.com/fdesjardins/neb/cmd"
)

func main() {
	if err := cmd.NebCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
