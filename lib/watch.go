package lib

func Watch(coinList Coins) {
	detail := []Coin{}

	for i := 0; i < 10; i++ {
		detail = append(detail, coinList.Coin[i])
	}

	Display(Coins{detail})
}
