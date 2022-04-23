package cmd

import (
	"fmt"
	"github.com/adlio/trello"
	"github.com/spf13/cobra"
)

func init() {
	deleteCardCmd.Flags().BoolVarP(&Top, "top", "t", false, "Delete card at top of list")
	deleteCmd.AddCommand(deleteCardCmd)
}

var deleteCardCmd = &cobra.Command{
	Use:   "card",
	Short: "Delete card",
	Long:  "Delete card",
	Run: func(cmd *cobra.Command, args []string) {
		var card *trello.Card
		var err error
		if Top {
			card, err = getTopCard()
		} else {
			card, err = getCard()
		}
		if err != nil {
			er(err)
		}
		err = card.Delete()
		if err != nil {
			er(err)
		}
		fmt.Printf("Deleted card %v: %v: %v\n", card.ID, card.Name, card.Desc)
	},
}
