package lib

import "time"

func Watch(coinList Coins, now time.Time) {
	detail := []Coin{}

	for i := 0; i < 10; i++ {
		detail = append(detail, coinList.Coin[i])
	}

	Display(Coins{detail}, now)
}
