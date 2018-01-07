package lib

import (
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

	Display(Coins{detail}, now)
}
