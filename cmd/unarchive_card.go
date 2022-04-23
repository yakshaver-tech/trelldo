package cmd

import (
	"fmt"
	"github.com/adlio/trello"
	"github.com/spf13/cobra"
)

func init() {
	unarchiveCardCmd.Flags().BoolVarP(&Top, "top", "t", false, "Unarchive card at top of list")
	unarchiveCmd.AddCommand(unarchiveCardCmd)
}

var unarchiveCardCmd = &cobra.Command{
	Use:   "card",
	Short: "Unarchive card",
	Long:  "Unarchive card",
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
		err = card.Unarchive()
		if err != nil {
			er(err)
		}
		fmt.Printf("Unarchived card %v: %v: %v\n", card.ID, card.Name, card.Desc)
	},
}
