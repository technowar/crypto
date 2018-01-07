package lib

type Coin struct {
	Id                 string
	Rank               string
	Name               string
	Symbol             string
	Price_usd          string
	Market_cap_usd     string
	Available_supply   string
	Price_btc          string
	Percent_change_1h  string
	Percent_change_24h string
}

type Coins struct {
	Coin []Coin
}
