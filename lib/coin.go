package lib

import (
	"fmt"
	"strings"
	"time"
)

func Crypto(crypto string, coinList Coins, now time.Time) {
	detail := []Coin{}

	for _, item := range coinList.Coin {
		if item.Symbol == crypto || strings.ToUpper(item.Id) == crypto {
			detail = append(detail, item)
		}
	}

	if len(detail) == 0 {
		fmt.Printf("%v: Unable to locate.\n\n", crypto)
	} else {
		Display(Coins{detail}, now)
	}
}
