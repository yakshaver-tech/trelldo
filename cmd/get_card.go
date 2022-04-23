package cmd

import (
	"fmt"
	"github.com/adlio/trello"
	"github.com/spf13/cobra"
)

func init() {
	getCardCmd.Flags().BoolVarP(&Top, "top", "t", false, "Get card at top of list")
	getCmd.AddCommand(getCardCmd)
}

var getCardCmd = &cobra.Command{
	Use:   "card",
	Short: "Get card",
	Long:  "Get card",
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
		fmt.Printf("%v: %v: %v\n", card.ID, card.Name, card.Desc)
	},
}
