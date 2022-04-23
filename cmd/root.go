package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/adlio/trello"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	client  *trello.Client
	Top     bool

	rootCmd = &cobra.Command{
		Use:   "trelldo",
		Short: "A Trello CLI",
		Long:  "A simple Trello CLI using the Trello API.",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.trelldo.yaml)")
	rootCmd.PersistentFlags().StringP("board", "b", "", "board name")
	rootCmd.PersistentFlags().StringP("list", "l", "", "list name")
	rootCmd.PersistentFlags().StringP("card", "c", "", "card id")
	viper.BindPFlag("board", rootCmd.PersistentFlags().Lookup("board"))
	viper.BindPFlag("list", rootCmd.PersistentFlags().Lookup("list"))
	viper.BindPFlag("card", rootCmd.PersistentFlags().Lookup("card"))
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".trelldo")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		fmt.Printf("Default board: %v\n", viper.GetString("board"))
		fmt.Printf("Default list: %v\n", viper.GetString("list"))
		fmt.Printf("Default card: %v\n", viper.GetString("card"))
	}
}
func getClient() (client *trello.Client) {
	if client == nil {
		appKey := viper.GetString("appKey")
		token := viper.GetString("token")
		client = trello.NewClient(appKey, token)
	}
	return client
}

func getBoards() (boards []*trello.Board, err error) {
	return getClient().GetMyBoards(trello.Defaults())
}

func getBoard() (board *trello.Board, err error) {
	boardName := viper.GetString("board")
	if boardName == "" {
		return nil, errors.New("Bad board name")
	}

	boards, err := getBoards()
	if err != nil {
		return nil, err
	}
	for _, board := range boards {
		if board.Name == boardName {
			return board, nil
		}
	}
	return nil, errors.New("Board not found")
}

func getLists() (lists []*trello.List, err error) {
	board, err := getBoard()
	if err != nil {
		return nil, err
	}
	return board.GetLists(trello.Defaults())
}

func getList() (list *trello.List, err error) {
	listName := viper.GetString("list")
	if listName == "" {
		return nil, errors.New("Bad list name")
	}

	lists, err := getLists()
	if err != nil {
		return nil, err
	}
	for _, list := range lists {
		if list.Name == listName {
			return list, nil
		}
	}
	return nil, errors.New("List not found")
}

func createList() (list *trello.List, err error) {
	listName := viper.GetString("list")
	if listName == "" {
		return nil, errors.New("Bad list name")
	}
	board, err := getBoard()
	if err != nil {
		return nil, err
	}
	return board.CreateList(listName, trello.Defaults())
}

func getCards() (cards []*trello.Card, err error) {
	list, err := getList()
	if err != nil {
		return nil, err
	}
	return list.GetCards(trello.Defaults())
}

func createCard(name string, desc string) (card *trello.Card, err error) {
	if name == "" {
		return nil, errors.New("Bad card name")
	}
	list, err := getList()
	if err != nil {
		return nil, err
	}
	card = &trello.Card{Name: name, Desc: desc}
	err = list.AddCard(card, trello.Defaults())
	if err != nil {
		return nil, err
	}
	viper.Set("card", card.ID)
	return card, nil
}

func getTopCard() (card *trello.Card, err error) {
	cards, err := getCards()
	if err != nil {
		return nil, err
	}
	card = cards[0]
	viper.Set("card", card.ID)
	return card, nil
}

func getCard() (card *trello.Card, err error) {
	cardId := viper.GetString("card")
	return getClient().GetCard(cardId, trello.Defaults())
}
