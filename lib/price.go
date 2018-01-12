package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Price(from, to string) {
	query := fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%v&tsyms=%v", from, to)
	response, err := http.Get(query)

	defer response.Body.Close()

	if err != nil {
		fmt.Println("Internal server error.")
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Internal server error.")
		os.Exit(1)
	}

	var respMap = map[string]map[string]map[string]To{}

	json.Unmarshal(data, &respMap)

	lastUpdate := respMap["RAW"][from][to].Lastupdate
	now := time.Unix(lastUpdate, 0)

	Compare(respMap["RAW"][from], now)
}
