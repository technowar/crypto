package lib

import (
	"fmt"
	"strings"
	"time"
)

func Crypto(crypto string, coinList Coins, now time.Time) {
	var id string
	detail := []Coin{}

	for _, item := range coinList.Coin {
		if item.Symbol == crypto || strings.ToUpper(item.Id) == crypto {
			id = item.Id
			detail = append(detail, item)
		}
	}

	if len(detail) == 0 {
		fmt.Printf("%v: Unable to locate.\n\n", crypto)
	} else {
		Display(Coins{detail}, now)
		fmt.Printf("URL: https://coinmarketcap.com/currencies/%v/\n", id)
	}
}
