package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// This is overwritten at compile time with build flags with the current tag
// See build step in Makefile to get a sense of what happens
var version = "v0.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("notask version %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
