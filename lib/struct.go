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

type To struct {
	Currency string
	Value    float64
}

type SortTo []To

func (a SortTo) Len() int           { return len(a) }
func (a SortTo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortTo) Less(i, j int) bool { return a[i].Currency < a[j].Currency }
