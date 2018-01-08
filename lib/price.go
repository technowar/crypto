package lib

import (
	"fmt"
	"net/http"
	"os"
)

func Price(from, to string) {
	query := fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%v&tsyms=%v", from, to)
	response, err := http.Get(query)

	defer response.Body.Close()

	if err != nil {
		fmt.Println("Internal server error.")
		os.Exit(1)
	}

	fmt.Println(response)
}
