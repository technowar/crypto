package main

import (
	"./lib"
	"./utils"
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
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
