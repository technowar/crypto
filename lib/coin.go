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
		} else {
			fmt.Printf("Unable to find %v in the database.\n\n", crypto)

			return
		}
	}

	Display(Coins{detail}, now)
}
