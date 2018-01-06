package utils

import "fmt"

func Usage() {
	fmt.Println("\nCommands available:")
	fmt.Println("  help    \tDisplay help.")
	fmt.Println("  watch   \tWatch top 10 coins.")
	fmt.Println("\nAll data are from https://coinmarketcap.com/api/")
}
