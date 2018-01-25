package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/technowar/crypto/lib"
	"github.com/technowar/crypto/utils"
)

func main() {
	utils.Clear(runtime.GOOS)
	fmt.Println("Fetching data...")

	coinList, now := lib.FetchCoins()
	reader := bufio.NewReader(os.Stdin)
	tickerDone := make(chan struct{})

	utils.Clear(runtime.GOOS)

	for {
		fmt.Print("> ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		texts := strings.Split(text, " ")

		select {
		case tickerDone <- struct{}{}:
		default:
		}

		if texts[0] == "coin" {
			coin := "BTC"

			if len(texts) >= 2 {
				coin = strings.ToUpper(texts[1])
			}

			lib.CoinDetails(coin, coinList, now)

			go timeTick("coin", coin, "", 5, tickerDone)
		} else if texts[0] == "price" {
			from := "BTC"
			to := "BTC"

			if len(texts) == 2 {
				from = texts[1]
			}

			if len(texts) >= 3 {
				from = texts[1]
				to = texts[2]
			}

			lib.FetchPrice(strings.ToUpper(from), strings.ToUpper(to))

			go timeTick("price", strings.ToUpper(from), strings.ToUpper(to), 1, tickerDone)
		} else if text == "watch" {
			lib.CoinDetails("", coinList, now)

			go timeTick("watch", "", "", 5, tickerDone)
		} else if text == "clear" || text == "cls" {
			utils.Clear(runtime.GOOS)
		} else if text == "help" {
			utils.Usage()
		} else {
			fmt.Printf("Bad option: %v\n\n", text)
		}
	}
}

func timeTick(format, from, to string, tick int, done <-chan struct{}) {
	tickChan := time.NewTicker(time.Minute * time.Duration(tick))

	for {
		select {
		case <-tickChan.C:
			switch format {
			case "watch":
				coinList, now := lib.FetchCoins()

				lib.CoinDetails("", coinList, now)
			case "coin":
				coinList, now := lib.FetchCoins()

				lib.CoinDetails(from, coinList, now)
			case "price":
				lib.FetchPrice(from, to)
			}
		case <-done:
			return
		}
	}
}
