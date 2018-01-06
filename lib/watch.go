package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Coin struct {
	Rank               string
	Name               string
	Symbol             string
	Price_usd          string
	Price_btc          string
	Percent_change_1h  string
	Percent_change_24h string
	Percent_change_7d  string
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
		data = append(data, []string{
			item.Rank,
			item.Symbol,
			item.Name,
			item.Price_usd,
			item.Price_btc,
			item.Percent_change_1h,
			item.Percent_change_24h,
			item.Percent_change_7d,
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	now := time.Now()
	message := fmt.Sprintf("Updated: %v", now.Format(time.RFC1123))

	table.SetHeader([]string{
		"Rank",
		"Symbol",
		"Name",
		"Price [USD]",
		"Price [BTC]",
		"Percent Change [1h]",
		"Percent Change [24h]",
		"Percent Change [7d]",
	})
	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor})
	table.SetCaption(true, message)

	for _, item := range data {
		table.Append(item)
	}

	table.Render()
}
