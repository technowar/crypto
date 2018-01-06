package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Crypto(crypto string) {
	response, err := http.Get("https://api.coinmarketcap.com/v1/ticker?limit=1000")

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
	detail := []Coin{}

	for _, item := range coins.Coin {
		if item.Symbol == crypto || strings.ToUpper(item.Id) == crypto {
			detail = append(detail, item)
		}
	}

	Display(Coins{detail})
}
