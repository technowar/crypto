package utils

import "fmt"

func Usage() {
	fmt.Println("\nCommands available:")
	fmt.Println("  help    \tDisplay help.")
	fmt.Println("  watch   \tWatch top 10 coins.")
	fmt.Println("  coin    \tDisplay specific coin. Defaults to BTC.")
	fmt.Println("\nAll data are from https://coinmarketcap.com/api/")
}
