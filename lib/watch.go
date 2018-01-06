package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
)

type Coin struct {
	Name   string
	Symbol string
	Rank   string
}

type Coins struct {
	Coin []Coin
}

func Watch() {
	response, err := http.Get("https://api.coinmarketcap.com/v1/ticker/?limit=10")

	defer response.Body.Close()

	if err != nil {
		fmt.Println("Internal server error.")
		os.Exit(1)
	}

	coin := new([]Coin)
	err = json.NewDecoder(response.Body).Decode(coin)

	if err != nil {
		fmt.Println("Internal server error.")
		os.Exit(1)
	}

	coins := Coins{*coin}
	data := [][]string{}

	for _, item := range coins.Coin {
		data = append(data, []string{item.Name, item.Symbol, item.Rank})
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Name", "Symbol", "Rank"})

	for _, item := range data {
		table.Append(item)
	}

	table.Render()
}
