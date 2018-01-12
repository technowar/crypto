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

	coinList, now := lib.Fetch()
	reader := bufio.NewReader(os.Stdin)

	utils.Clear(runtime.GOOS)

	for {
		fmt.Print("> ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		texts := strings.Split(text, " ")
		newNow := time.Now().Unix()
		after := now.Add(5 * time.Minute).Unix()

		if newNow > after {
			coinList, now = lib.Fetch()
		}

		if texts[0] == "coin" {
			switch len(texts) <= 1 {
			case true:
				lib.Crypto("BTC", coinList, now)
			default:
				lib.Crypto(strings.ToUpper(texts[1]), coinList, now)
			}
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

			lib.Price(strings.ToUpper(from), strings.ToUpper(to))
		} else if text == "watch" {
			lib.Watch(coinList, now)
		} else if text == "clear" || text == "cls" {
			utils.Clear(runtime.GOOS)
		} else if text == "help" {
			utils.Usage()
		} else {
			fmt.Printf("Bad option: %v\n\n", text)
		}
	}
}
