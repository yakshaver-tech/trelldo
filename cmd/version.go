package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of trelldo",
	Long:  "Print the version number of trelldo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trelldo v0.0.1")
	},
}
