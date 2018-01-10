package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)

const tsyms = "BTC,USD,EUR,JPY,PHP"

func Price(from, to string) {
	query := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%v&tsyms=%v", from, tsyms)
	response, err := http.Get(query)

	defer response.Body.Close()

	if err != nil {
		fmt.Println("Internal server error.")
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Internal server error.")
		os.Exit(1)
	}

	var priceMap map[string]float64
	var prices []To

	json.Unmarshal(data, &priceMap)

	for currency, value := range priceMap {
		price := To{currency, value}
		prices = append(prices, price)
	}

	sort.Sort(SortTo(prices))

	prices = append([]To{To{from, 1}}, prices...)

	fmt.Println(prices)
}
