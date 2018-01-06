package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/technowar/crypto/lib"
	"github.com/technowar/crypto/utils"
)

func main() {
	utils.Clear(runtime.GOOS)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		texts := strings.Split(text, " ")

		if texts[0] == "coin" {
			utils.Clear(runtime.GOOS)

			switch len(texts) <= 1 {
			case true:
				lib.Crypto("BTC")
			default:
				lib.Crypto(strings.ToUpper(texts[1]))
			}
		} else if text == "help" {
			utils.Usage()
		} else if text == "watch" {
			utils.Clear(runtime.GOOS)
			lib.Watch()
		} else {
			fmt.Printf("Bad option: %v\n", text)
			utils.Usage()
		}
	}
}
