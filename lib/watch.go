package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
)

type Coin struct {
	Rank               string
	Name               string
	Symbol             string
	Price_usd          string
	Market_cap_usd     string
	Available_supply   string
	Price_btc          string
	Percent_change_1h  string
	Percent_change_24h string
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
		capFloat, _ := strconv.ParseFloat(item.Market_cap_usd, 64)
		cap := fmt.Sprintf("$%v", humanize.Commaf(capFloat))
		priceFloat, _ := strconv.ParseFloat(item.Price_usd, 64)
		price := fmt.Sprintf("$%v", humanize.Commaf(priceFloat))
		priceBtc := fmt.Sprintf("%v BTC", item.Price_btc)
		supplyFloat, _ := strconv.ParseFloat(item.Available_supply, 64)
		supply := fmt.Sprintf("%v %v", humanize.Commaf(supplyFloat), item.Symbol)

		data = append(data, []string{
			item.Rank,
			item.Symbol,
			item.Name,
			cap,
			price,
			priceBtc,
			supply,
			item.Percent_change_1h,
			item.Percent_change_24h,
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	now := time.Now()
	message := fmt.Sprintf("Updated: %v", now.Format(time.RFC1123))

	table.SetHeader([]string{
		"Rank",
		"Symbol",
		"Name",
		"Market Cap",
		"Price",
		"Price",
		"Circulating Supply",
		"Change [1h]",
		"Change [24h]",
	})
	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
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
