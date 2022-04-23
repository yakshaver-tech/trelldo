package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	createCardCmd.Flags().BoolVarP(&Top, "top", "t", false, "Add card to top of list")
	createCmd.AddCommand(createCardCmd)
}

var createCardCmd = &cobra.Command{
	Use:   "card",
	Short: "Create card",
	Long:  "Create card",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			args = append(args, "")
		}
		card, err := createCard(args[0], args[1])
		if err != nil {
			er(err)
		}
		fmt.Printf("Created card %v: %v: %v\n", card.ID, card.Name, card.Desc)
		if Top {
			err = card.MoveToTopOfList()
			if err != nil {
				er(err)
			}
			fmt.Println("Moved to top of list")
		} else {
			err = card.MoveToBottomOfList()
			if err != nil {
				er(err)
			}
			fmt.Println("Moved to bottom of list")
		}
	},
}
