package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	createCmd.AddCommand(createListCmd)
}

var createListCmd = &cobra.Command{
	Use:   "list",
	Short: "Create list",
	Long:  "Create list",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			viper.Set("list", args[0])
		}
		list, err := createList()
		if err != nil {
			er(err)
		}
		fmt.Printf("Created list %v\n", list.Name)
	},
}
