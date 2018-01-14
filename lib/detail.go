package lib

import (
	"fmt"
	"strings"
	"time"
)

func Details(crypto string, coinList Coins, now time.Time) {
	detail := []Coin{}

	switch len(crypto) {
	case 0:
		for i := 0; i < 10; i++ {
			detail = append(detail, coinList.Coin[i])
		}
	default:
		for _, item := range coinList.Coin {
			if item.Symbol == crypto || strings.ToUpper(item.Id) == crypto {
				detail = append(detail, item)
			}
		}
	}

	if len(detail) == 0 {
		fmt.Printf("%v: Unable to locate.\n\n", crypto)
	} else {
		Display(Coins{detail}, now)
	}
}
