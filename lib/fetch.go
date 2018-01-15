package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func FetchCoins() (Coins, time.Time) {
	data, err := fetch("https://api.coinmarketcap.com/v1/ticker/?limit=1000")
	if err != nil {
		fmt.Println("Internal server error.")
		os.Exit(1)
	}

	var coin = new([]Coin)

	json.Unmarshal(data, &coin)

	coins := Coins{*coin}
	now := time.Now()

	return coins, now
}

func FetchPrice(from, to string) {
	query := fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%v&tsyms=%v", from, to)
	data, err := fetch(query)
	if err != nil {
		fmt.Println("Internal server error.")
		os.Exit(1)
	}

	var respMap = map[string]map[string]map[string]To{}

	json.Unmarshal(data, &respMap)

	lastUpdate := respMap["RAW"][from][to].Lastupdate
	now := time.Unix(lastUpdate, 0)
	resp := respMap["RAW"][from]

	if len(resp) == 0 {
		fmt.Printf("%v/%v: Unable to locate.\n\n", from, to)
	} else {
		Compare(resp, now)
	}
}

func fetch(urlPath string) ([]byte, error) {
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
