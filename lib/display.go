package lib

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	"github.com/technowar/crypto/utils"
)

func Display(coins Coins, now time.Time) {
	data := [][]string{}

	for _, item := range coins.Coin {
		capFloat, _ := strconv.ParseFloat(item.Market_cap_usd, 64)
		cap := fmt.Sprintf("$%v", humanize.Commaf(capFloat))
		priceFloat, _ := strconv.ParseFloat(item.Price_usd, 64)
		price := fmt.Sprintf("$%v", humanize.Commaf(priceFloat))
		priceBtc := fmt.Sprintf("฿%v", item.Price_btc)
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

	utils.Clear(runtime.GOOS)
	table.Render()
}
