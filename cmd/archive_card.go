package cmd

import (
	"fmt"
	"github.com/adlio/trello"
	"github.com/spf13/cobra"
)

func init() {
	archiveCardCmd.Flags().BoolVarP(&Top, "top", "t", false, "Archive card at top of list")
	archiveCmd.AddCommand(archiveCardCmd)
}

var archiveCardCmd = &cobra.Command{
	Use:   "card",
	Short: "Archive card",
	Long:  "Archive card",
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
		err = card.Archive()
		if err != nil {
			er(err)
		}
		fmt.Printf("Archived card %v: %v: %v\n", card.ID, card.Name, card.Desc)
	},
}
