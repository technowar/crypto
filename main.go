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

		switch text {
		case "coins":
			utils.Clear(runtime.GOOS)
			lib.Watch()
		default:
			utils.Clear(runtime.GOOS)
			utils.Usage()
		}
	}
}
